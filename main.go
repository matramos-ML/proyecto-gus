package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/practicas/service/model"
)

var usuarios = make(map[string]model.Usuario)

func main() {

	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		fmt.Println("mapa de usuarios... ", usuarios)
		c.JSON(http.StatusOK, usuarios)
	})

	router.POST("/post", func(c *gin.Context) {

		mail := c.PostForm("mail")
		nombre := c.PostForm("nombre")

		if mail == "" || nombre == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Need more properties",
			})
			return
		}
		if ok := model.Contains(usuarios, mail); ok == false {
			usuarios[mail] = model.Usuario{
				Mail:   mail,
				Nombre: nombre,
			}
			c.JSON(http.StatusOK, gin.H{
				"msg": "Usuario creado",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Mail ya existe",
		})
		return
	})

	router.Run(":80")
}
