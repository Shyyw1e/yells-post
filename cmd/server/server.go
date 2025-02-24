package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"yells-post/graph"
	"yells-post/internal/postgres"
	"yells-post/internal/usecase"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func isPortAvailable(port string) bool {
    ln, err := net.Listen("tcp", ":"+port)
    if err != nil {
        return false
    }
    ln.Close()
    return true
}

func main() {
    dbHost := os.Getenv("DB_HOST")
    if dbHost == "" {
        dbHost = "db"
    }

    dbPort := os.Getenv("DB_PORT")
    if dbPort == "" {
        dbPort = "5432"
    }

    dbUser := os.Getenv("DB_USER")
    if dbUser == "" {
        dbUser = "postgres"
    }

    dbPassword := os.Getenv("DB_PASSWORD")
    if dbPassword == "" {
        dbPassword = "Shyywie_8169"
    }

    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "postgres"
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    portNum, err := strconv.Atoi(port)
    if err != nil || portNum < 1 || portNum > 65535 {
        slog.Error("неверный порт: ", port, "используется порт по умолчанию", defaultPort)
        portNum = 8080
        port = defaultPort
    }

    if !isPortAvailable(fmt.Sprintf("%d", portNum)) {
        slog.Error("Порт ", port, " занят, попробуйте другой порт")
        return
    }

    db, err := postgres.NewDB(dbHost, dbPort, dbUser, dbPassword, dbName)
    if err != nil {
        slog.Error("Ошибка подключения к DB: ", err)
        return
    }
    defer db.Close()


    pgRepo := postgres.NewRepo(db)
    postUsecase := usecase.NewPostUsecase(pgRepo)
    commentUsecase := usecase.NewCommentUsecase(pgRepo)

    resolver := &graph.Resolver{
        PostUsecase:    postUsecase,
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

    

    log.Printf("connect to http://localhost:%d/ for GraphQL playground", portNum)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        slog.Error("Ошибка при запуске сервера: ", err)
    }
    
}