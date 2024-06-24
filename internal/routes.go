package internal

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostForm(db *sql.DB,Errors DBError) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
        defer cancel()

		var form Form
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		exist, err := form.Check(ctx, db,query)
		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": Errors.QueryFail})
			return 
		}

		if exist {
			c.JSON(http.StatusConflict, gin.H{"error": Errors.NameConflict})
			return
		} 

		if err := form.Insert(ctx, db, query); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": Errors.InsertFail})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Application submitted"})
	}
}