package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/reezy24/leaderboard-api/db"
)


func setupRouter(database *db.DB) *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := db[user]
	// 	if ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	// 	}
	// })

	// Create leaderboard.
	r.POST("/leaderboard", func(ctx *gin.Context) {
		var body db.Leaderboard		

		if err := ctx.BindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "invalid request body")
			return
		}

		if body.Name == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "name is required")
			return
		}

		leaderboard := database.CreateLeaderboard(body.Name, body.Columns)

		ctx.JSON(http.StatusCreated, leaderboard)
	})

	// Read leaderboard.
	r.GET("/leaderboard/:id", func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")

		parsedID, err := uuid.Parse(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
			return
		}

		leaderboard := database.ReadLeaderboard(parsedID)

		if leaderboard == nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "leaderboard not found")
			return
		}

		ctx.JSON(http.StatusOK, leaderboard)
	})

	return r
}

func main() {
	database := db.NewLeaderboardDB()
	r := setupRouter(database)

	r.Run("localhost:8080")
}
