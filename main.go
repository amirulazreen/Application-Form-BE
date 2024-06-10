package main

import (
	"database/sql"
	"log"

	"github.com/amirulazreen/CHIP/internal"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	connect := "user=postgres.xqrxxtdmtwhykrktnrii password=chipdb123123!! host=aws-0-ap-southeast-1.pooler.supabase.com port=6543 dbname=postgres ?sslmode=disable"

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
