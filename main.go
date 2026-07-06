package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/sijms/go-ora/v2"
)

func main() {
	connString := "oracle://system:Oracle123@localhost:1521/FREEPDB1"

	db, err := sql.Open("oracle", connString)
	if err != nil {
		log.Fatal("Error abriendo conexión:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error conectando a Oracle:", err)
	}

	log.Println("Conexión exitosa a Oracle Database")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "API TransLogix funcionando",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"api":      "ok",
			"database": "oracle conectado",
		})
	})

	r.Run(":8080")
}
