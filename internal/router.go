package internal

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostForm(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj Form

		if err := c.ShouldBindJSON(&obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), 500*time.Millisecond)
		defer cancel()

		var nameExists bool
		stmt := `SELECT EXISTS (SELECT 1 FROM application WHERE name = $1)`
		err := db.QueryRowContext(ctx, stmt, obj.Name).Scan(&nameExists)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
			return
		}

		if nameExists {
			c.JSON(http.StatusConflict, gin.H{"error": "Organization of the same name already exists in the database"})
			return
		}

		insertStmt := `INSERT INTO application (name, type, bank, opsyears, ssm, prevgateway, prodtype, storetype, inventory, reference, socmedia, litigation, score)
                       VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
		_, err = db.ExecContext(ctx, insertStmt, obj.Name, obj.Type, obj.Bank, obj.OpsYears, obj.SSM, obj.PrevGateway, obj.ProdType, obj.StoreType, obj.Inventory, obj.Reference, obj.SocMedia, obj.Litigation, obj.Score)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert record"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Application submitted"})
	}
}