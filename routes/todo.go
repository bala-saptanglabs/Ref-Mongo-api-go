package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/bala-saptang/fib-go-mongo-akilsharma/controllers"
)

func TodoRoute(router fiber.Router){
	router.Get("",controllers.GetTodos)
	router.Put("/:id",controllers.UpdateTodo)
	router.Post("/",controllers.CreateTodo)
	router.Delete("/:id",controllers.DeleteTodo)
	router.Get("/:id",controllers.GetTodoById)
}

