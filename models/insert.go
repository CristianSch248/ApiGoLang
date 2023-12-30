package models

import (
    "apiGo/db"
)

func Insert(todo Todo) (id int64, err error){
    conn, err := bd.OpenConn()
    if err != nil {
        return
    }

    defer conn.Close()

    sql := `INSER INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

    err = conn.QueryRow(sql, todo.title, todo.description, todo.done).Scan(&id)

    return
}