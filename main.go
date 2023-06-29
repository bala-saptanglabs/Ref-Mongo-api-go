package main

/*
go get url -> download url file and add it in go. mod file , you can import and use it
go mod tidy -> download the files you mentioned in import and add it to go.mod file

to kill a process
-----------------
lsof -i tcp:port_number
kill -9 p_id


in vscode to use different git accounts -> profile -> github profile (log in /out based on need)

*/

import (
	"log"

	"github.com/bala-saptang/fib-go-mongo-akilsharma/config"
	"github.com/joho/godotenv" 


	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/bala-saptang/fib-go-mongo-akilsharma/routes"
)

func setUpRoutes(app *fiber.App){
	app.Get("/",func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"messaege":"you are at endpoint",
		})
	})

	//api grou
	api := app.Group("/api")
	api.Get("/",func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"messaege":"you are at api endpoint",
		})
	})

	routes.TodoRoute(api.Group("/todos"))
	
}

func main(){

	app:= fiber.New()

	// logs the http requests
	app.Use(logger.New())	

	err:= godotenv.Load();

	if err!=nil{
		log.Fatal("error loading env file")
	}

	config.ConnectDB()


	

	setUpRoutes(app)

	err = app.Listen(":8000")

	if err!=nil{
		panic(err)
	}
}