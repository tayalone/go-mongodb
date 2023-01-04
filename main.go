package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	MongoClient "github.com/tayalone/go-mongodb/mongo"
	"github.com/tayalone/go-mongodb/todo"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("Let's Try")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoClient.Init()

	client, errMongoClient := MongoClient.GetClient()

	if errMongoClient != nil {
		log.Fatalln("errMongoClient", errMongoClient.Error())
	}

	r := gin.Default()

	r.GET("/todos", func(ctx *gin.Context) {
		var todos []todo.Domain

		coll := client.GetCollection("todos")

		cursor, err := coll.Find(ctx, bson.M{})
		if err != nil {
			panic(err)
		}
		if err = cursor.All(ctx, &todos); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"todos":   todos,
		})

	})

	r.GET("/pinga", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// r.Run(":3000") // listen and serve on 0.0.0.0:8080

	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	// // run srv in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// // create channel of os signal for waiting signal
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// // close `quit` channel when app're closed
	// defer func() {
	// 	close(quit)
	// }()
	s := <-quit
	log.Println("signal is: ", s)
	log.Println("Shutting down app...")

	// // The context is used to inform the server it has 5 seconds to finish
	// // the request it is currently handling
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// // create context -> waiting every process in server is done
	ctx := context.Background()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("App forced to shutdown:", err)
	}
	client.Deconnect()

	log.Println("App exiting")

}
