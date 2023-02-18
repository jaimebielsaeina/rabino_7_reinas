package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Nombres y contenido del struct debe ser publico si no no se hace el binding **Mejor moverlo a otro archivo**
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
		//Procesa una petición de login
		api.POST("/auth/login", postLogin)

		//Procesa una petición de registro
		api.POST("/auth/register", postRegister)

		//Ejemplo de paso de parametros por url
		api.GET("/prueba/:param", getParam)

		//Ejemplo para devolver más de un dato
		api.GET("/prueba/names", getNames)

		//Ejemplo para devolver structs de datos
		api.GET("/prueba/users", getUsers)
	}

	// Start and run the server
	router.Run(":3001")
}

func postLogin(c *gin.Context) {

	u := Login{}
	//Con el binding guardamos el json de la petición en u que es de tipo login
	if err := c.BindJSON(&u); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(u)
	c.JSON(http.StatusAccepted, &u)

}

func postRegister(c *gin.Context) {

	u := Register{}
	//Con el binding guardamos el json de la petición en u que es de tipo login
	if err := c.BindJSON(&u); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(u)
	c.JSON(http.StatusAccepted, &u)

}

func getParam(c *gin.Context) {
	param := c.Param("param")
	c.JSON(http.StatusOK, gin.H{
		"name": param,
	})

}

func getNames(c *gin.Context) {
	var param = []string{"Adolfo", "Adrián", "Agustín", "Aitor", "Aitor-tilla"}
	c.JSON(http.StatusOK, gin.H{
		"name": param,
	})

}

func getUsers(c *gin.Context) {
	var users = []Login{{"a@gmail.com", "1234"}, {"b@gmail.com", "5678"}, {"c@gmail.com", "91011"}, {"d@gmail.com", "1234"}}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
