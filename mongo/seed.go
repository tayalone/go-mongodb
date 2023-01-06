package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tayalone/go-mongodb/todo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func seed(db *mongo.Database) {
	/* Seed Initial Data */
	coll := db.Collection("todos")

	_, err := coll.DeleteMany(
		context.TODO(),
		bson.D{
			{},
		},
	)
	if err != nil {
		log.Fatalln("Error tuncate `todos` Collection", err.Error())
	}

	current := time.Now()

	newTodos := []interface{}{
		todo.Domain{
			Task:      "task 3",
			Completed: false,
			CreatedAt: current,
			UpdatedAt: current,
		},
		todo.Domain{
			Task:      "task 4",
			Completed: false,
			CreatedAt: current,
			UpdatedAt: current,
		},
	}
	fmt.Println("todos", newTodos)
	results, err := coll.InsertMany(context.TODO(), newTodos)
	if err != nil {
		log.Fatalln("inset new todos error")
	}

	// fmt.Println(results)

	for _, id := range results.InsertedIDs {
		fmt.Printf("\t%s\n", id)
	}
	// coll.UpdateByID()
}
