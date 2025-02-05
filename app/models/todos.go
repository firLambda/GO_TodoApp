package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content, 
		user_id, 
		created_at) values (?, ?, ?)`
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `SELECT * FROM todos where id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	return todo, err
}

func GetAllTodos() (todos []Todo, err error) {
	cmd := `SELECT * FROM todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}

	return todos, err
}

func (u *User) GetAllTodosByUser() (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at from todos where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}

	return todos, err
}

func (t *Todo) UpdateTodo() (err error) {
	cmd := `UPDATE todos set content = ?, user_id = ? where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() (err error) {
	cmd := `DELETE from todos WHERE id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
