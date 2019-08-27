package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/practicas/service/model"
)

var usuarios []model.Usuario

func main() {

	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		fmt.Println(usuarios)
		c.JSON(http.StatusOK, usuarios)
	})

	router.POST("/post", func(c *gin.Context) {

		mail := c.PostForm("mail")
		nombre := c.PostForm("nombre")

		if ok := model.Contains(usuarios, mail); ok == false {
			usuarios = append(usuarios, model.Usuario{Mail: mail, Nombre: nombre})
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Mail ya existe",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "Usuario creado",
		})

	})

	router.Run(":80")
}
