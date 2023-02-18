package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Nombres y contenido del struct debe ser publico si no no se hace el binding
type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/api")
	{

		api.POST("/auth/login", func(c *gin.Context) {
			u := Login{}
			//Con el binding guardamos el json de la petición en u que es de tipo login
			if err := c.BindJSON(&u); err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			fmt.Println(u)
			c.JSON(http.StatusAccepted, &u)
		})

		api.POST("/auth/register", func(c *gin.Context) {
			u := Register{}
			//Con el binding guardamos el json de la petición en u que es de tipo login
			if err := c.BindJSON(&u); err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			fmt.Println(u)
			c.JSON(http.StatusAccepted, &u)
		})
	}

	// Start and run the server
	router.Run(":3001")
}
