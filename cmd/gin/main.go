package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/matheusvcouto/crud-go/handlers"
	"github.com/matheusvcouto/crud-go/repository"
)

func LoggerIp(c *gin.Context) {
	ip := c.ClientIP()
	log.Printf("[gin middleware] IP do cliente: %s", ip)
	c.Next() // continua para o próximo handler
}

func main() {
	usersRepository := repository.NewUserRepository()
	h := handlers.NewHandler(usersRepository)

	r := gin.Default() // ja usa o gin.Recovery(), gin.Logger()
	// r.Use(LoggerIp, gin.Recovery(), gin.Logger()) // => gin.New()
	r.Use(LoggerIp)

	r.GET("/", h.HelloWordGin)
	r.GET("/hello/:name", h.HelloNameGin)
	r.GET("/stream", h.StreamPhraseGin)

	users := r.Group("/users")
	{
		users.GET("/", h.GetUsersGin)
		users.POST("/", h.CreateUserGin)
	}

	wrapf := r.Group("/w")
	{
		wrapf.GET("/stream", gin.WrapF(h.StreamPhrase))
		wrapf.GET("/", gin.WrapF(h.HelloWord))
		users := wrapf.Group("/users")
		{
			users.GET("/", gin.WrapF(h.GetUsers))
			users.POST("/", gin.WrapF(h.CreateUser))
		}
	}

	r.Run(":8000") // listen and serve on 0.0.0.0:8000
}
