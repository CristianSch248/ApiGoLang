package models

import (
	"apiGo/db"
)

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConn()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			// o deial seria logar onde deu erro, mas por enquanto vai para o próximo
			continue
		}

		todos = append(todos, todo)
	}

	return
}
