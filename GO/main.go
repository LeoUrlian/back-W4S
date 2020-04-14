package main

import (
	"github.com/gin-gonic/gin"
	"w4s/DB"
	"w4s/controllers"
	"w4s/middleware"
)

//
/*type Authc struct {
	Email string `json:"email" binding:required `
	Token string `json:"token" binding:required `
}*/

func main() {
	//creating connection with database
	r := gin.Default()     //starting the gin. //Iniciando o gin
	db := DB.SetupModels() //Connection database //Conexão banco de dados
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	authorized := r.Group("/v2")
	r.POST("/login", controllers.Login)
	//Cria usuario
	r.POST("/user", controllers.CreateUser)
	authorized.Use(middleware.AuthRequired)
	{
		authorized.GET("/searchall", controllers.FindUser)
		authorized.GET("/search", controllers.FindUserByNick)
	}

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") // listando e escutando no localhost:8080
	if err != nil {
		panic("NOT POSSIBLE RUN")
	}
}