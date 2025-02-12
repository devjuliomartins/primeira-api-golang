package main

//importando a biblioteca do gin
import "github.com/gin-gonic/gin"

func main() {
	// Criando a router
	r := gin.Default()

	// Define uma rota HTTP GET para o caminho /
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// Inicia o servidor na porta 8080
	r.Run(":8080")
}
