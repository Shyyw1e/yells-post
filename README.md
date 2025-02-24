# yells-post


Реализация с помощью GraphQL, запросы  
query {
  posts(page: 1, pageSize: 10) {
    id
    title
    content
    allowComments
    comments(page: 1, pageSize: 5) {
      id
      text
      author
      parentID
      replies(page: 1, pageSize: 3) {
        id
        text
      }
    }
  }
}

– Запрашивает список постов с указанной пагинацией.
– Для каждого поста запрашивает поля id, title, content, allowComments.
– Для каждого поста также запрашивает список комментариев (с пагинацией) и для каждого комментария – вложенные ответы (если они есть).


mutation {
  createPost(title: "New Post", content: "This is a new post created via mutation", allowComments: true) {
    id
    title
    content
    allowComments
  }
}

– Создает новый пост, сервер возвращает созданны объект с сгенерированным id


mutation {
  createComment(postID: "1", parentID: null, text: "This is a new comment") {
    id
    text
    author
    parentID
  }
}

– Вызывает мутацию createComment для поста с ID "1".
– В данном примере parentID передается как null (если комментарий не является ответом на другой комментарий).
– Сервер должен вернуть созданный комментарий.

subscription {
  commentAdded(postID: "1") {
    id
    text
    author
    parentID
  }
}


– Подписывается на событие commentAdded для поста с ID "1".
– При появлении нового комментария сервер должен отправить данные по подписке.


// netstat -ano | findstr :8080  список процессов
// taskkill /PID <номер_PID> /F закрытие процесса (для того чтобы localhost освобождать, можно просто менять порт локалхоста)
