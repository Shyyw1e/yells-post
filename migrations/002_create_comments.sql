CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL REFERENCES posts(id),
    text TEXT NOT NULL,
    author TEXT NOT NULL,
    parent_id INTEGER -- если комментарий является ответом на другой комментарий
);
