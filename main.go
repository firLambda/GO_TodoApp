package main

import (
	"fmt"
	"golang_todoApp/todo_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(2)
	// fmt.Println(u)
	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(2)
	// fmt.Println(u.Name, u.Email)
	// users, _ := models.GetAllUsers()
	// fmt.Println(users[0])

	// u.DeleteUser()
	// u, _ = models.GetUser(0)
	// fmt.Println(u)

	// fmt.Println(models.GetAllTodos())
	// user.CreateTodo("First Todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")

	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetAllTodosByUser()

	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(9)

	u, _ := models.GetUser(t.UserID)
	fmt.Println(u.GetAllTodosByUser())

	t, _ = models.GetTodo(9)
	t.DeleteTodo()
	fmt.Println(models.GetAllTodos())

}
