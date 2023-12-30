package models

import (
    "apiGo/db"
)

func get(id) (todo Todo, err error){
    conn, err := db.OpenConn()
    if err != nil{
        return
    }

    defer conn.Close()

    row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)

    err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

    return
}