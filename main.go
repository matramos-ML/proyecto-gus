package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/practicas/service/model"
)

var usuarios []model.Usuario

func main() {

	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {

		c.JSON(http.StatusOK, usuarios)
	})

	router.POST("/post", func(c *gin.Context) {

		mail := c.PostForm("mail")
		nombre := c.PostForm("nombre")

		if ok := model.Contains(usuarios, mail); ok == false {
			usuarios = append(usuarios, model.Usuario{Mail: mail, Nombre: nombre})
		}
	})

	router.Run(":80")
}
