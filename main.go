package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/amirulazreen/CHIP/internal"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	connect := os.Getenv("PG")

	db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	r.Use(internal.CorsMiddleware())
	r.POST("/save", internal.PostForm(db))
	r.Run(":8080")
}
