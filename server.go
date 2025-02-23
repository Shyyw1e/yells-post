package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"yells-post/graph"
	//"yells-post/internal/inmemory"
	"yells-post/internal/usecase"
	"yells-post/internal/postgres"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := postgres.NewDB("localhost", "5432", "postgres", "Shyywie_8169", "postgres")
	if err != nil {
		slog.Error("Ошибка подключения к DB: ", err)
		return
	}

	pgRepo := postgres.NewRepo(db)
	
	postUsecase := usecase.NewPostUsecase(pgRepo)
	commentUsecase := usecase.NewCommentUsecase(pgRepo)

	resolver := &graph.Resolver{
		PostUsecase: postUsecase,
		CommentUsecase: commentUsecase,
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


// netstat -ano | findstr :8080  список процессов
// taskkill /PID <номер_PID> /F закрытие процесса (для того чтобы localhost освобождать, можно просто менять порт локалхоста)
