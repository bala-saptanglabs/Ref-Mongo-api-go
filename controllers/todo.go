package controllers

import (
	"os"
	// "strconv"
	"time"

	"github.com/bala-saptang/fib-go-mongo-akilsharma/config"
	"github.com/bala-saptang/fib-go-mongo-akilsharma/model"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/keyauth"
	"go.mongodb.org/mongo-driver/bson"           // new
	"go.mongodb.org/mongo-driver/bson/primitive" // new
	"go.mongodb.org/mongo-driver/mongo"          // new
)



func GetTodos(c *fiber.Ctx) error{
	todo_collection := config.MI.Db.Collection(os.Getenv("TODO_COLLECTION"))

	query := bson.D{{}}

	cursos,err := todo_collection.Find(c.Context(), query)

	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"message":"something went wrong",
			"error":err.Error(),
		})
	}

	// creating a slice
	var todo_list []model.Task = make([]model.Task,0)

	err = cursos.All(c.Context(), &todo_list)

	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"message":"something went wrong",
			"error":err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(todo_list)
}

func CreateTodo(c *fiber.Ctx) error{
	
	
	data := new(model.Task)
	errr := c.BodyParser(&data)

	if errr!=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"message":"cannot parse json",
		})
	}

	todo_collection := config.MI.Db.Collection(os.Getenv("TODO_COLLECTION"))


	data.Id= nil
	f:= false
	data.Completed = &f
	data.CreatedAt = time.Now()
	data.UpdateAt = time.Now()

	result,err := todo_collection.InsertOne(c.Context(), data)

	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"message":"cannot insert data",
			"error" : err.Error(),
		})
	}

	// now get the same data
	query := bson.D{{Key:"_id",Value: result.InsertedID}}

	var todo model.Task

	err = todo_collection.FindOne(c.Context(),query).Decode(&todo)

	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"message":"cannot fetch data",
			"error" : err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"success": true,
		"data":todo,
	})
}

func GetTodoById(c *fiber.Ctx) error{
	param_id := c.Params("id") // string format

	id,err := primitive.ObjectIDFromHex(param_id)

	if err!=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "success":  false,
            "message": "Cannot parse Id",
        })
	}

	// get the exact task using id
	todo_collection := config.MI.Db.Collection(os.Getenv("TODO_COLLECTION"))

	query:= bson.D{{Key: "_id",Value: id}}

	var todo model.Task

	err = todo_collection.FindOne(c.Context(),query).Decode(&todo)

	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"message":"cannot fetch data",
			"error" : err.Error(),
		})
	}

 	
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"success":  true,
		"data": todo,
	})
}

func UpdateTodo(c *fiber.Ctx) error{

	param_id  :=  c.Params("id")

	id,err  := primitive.ObjectIDFromHex(param_id)

	if err!=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "success":  false,
            "message": "Cannot parse Id",
        })
	}

	data := new(model.Task)
	err = c.BodyParser(&data)

	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "success":  false,
            "message": "Cannot parse JSON",
        })
	}

	// find and update
	todo_collection := config.MI.Db.Collection(os.Getenv("TODO_COLLECTION"))

	query := bson.D{{Key: "_id",Value: id}}

	var update_doc bson.D

	if data.Title!=nil{
		update_doc = append(update_doc, bson.E{Key: "title",Value: data.Title})
	}
	if data.Completed!=nil{
		update_doc = append(update_doc, bson.E{Key: "completed",Value: data.Completed})
	}
	update_doc = append(update_doc, bson.E{Key: "updatedAt",Value: time.Now()})

	update := bson.D{{Key: "$set", Value: update_doc}}

	err = todo_collection.FindOneAndUpdate(c.Context(),query, update).Err()


	if err!=nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success":  false,
			"message": "Not found",
		})
	}

	var updated_todo model.Task

	 todo_collection.FindOne(c.Context(),query).Decode(&updated_todo)

	

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"success":  true,
		"data":updated_todo,
	})

}

func DeleteTodo(c *fiber.Ctx) error{
	// get params id and delete
	param_id := c.Params("id")

	id, err := primitive.ObjectIDFromHex(param_id)

	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "success":  false,
            "message": "Cannot parse Id",
        })
	}

	// delete id if found

	todo_collection := config.MI.Db.Collection(os.Getenv("TODO_COLLECTION"))

	query := bson.D{{Key: "_id",Value: id}}

	err = todo_collection.FindOneAndDelete(c.Context(), query).Err()


	if err!=nil{
		if err == mongo.ErrNoDocuments{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success":  false,
				"message": "Not found to delete",
				"error":err,
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":  false,
			"message": "cant  delete",
			"error":err,
		})
	}


	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"success":  true,
		"Message": "successfully deleted",
	})

}