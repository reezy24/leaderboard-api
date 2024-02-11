package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/reezy24/leaderboard-api/db"
)

func setupRouter(database db.DB) *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Create leaderboard.
	r.POST("/leaderboards", func(ctx *gin.Context) {
		var body db.Leaderboard

		if err := ctx.BindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "invalid request body")
			return
		}

		if body.Name == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "name is required")
			return
		}

		if len(body.FieldNames) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "at least one field name is required")
			return
		}

		leaderboard := database.CreateLeaderboard(body.Name, body.FieldNames, body.Entries)

		ctx.JSON(http.StatusCreated, leaderboard)
	})

	// Read leaderboard.
	r.GET("/leaderboards/:id", func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")

		parsedId, err := uuid.Parse(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
			return
		}

		leaderboard := database.ReadLeaderboard(parsedId)

		if leaderboard == nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "leaderboard not found")
			return
		}

		ctx.JSON(http.StatusOK, leaderboard)
	})

	// Create leaderboard entry.
	r.POST("/leaderboards/:id/entries", func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")

		parsedId, err := uuid.Parse(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
			return
		}

		var body db.Entry

		if err := ctx.BindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "invalid request body")
			return
		}

		if body.Name == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "name is required")
			return
		}

		entry, err := database.CreateEntry(parsedId, body.Name, body.FieldNamesToValues)
		if err != nil {
			code := http.StatusInternalServerError
			msg := "failed to create entry: %s"

			switch err {
			case db.ErrLeaderboardInvalidFieldName:
				code = http.StatusBadRequest
				msg = fmt.Sprintf(msg, "invalid field name")
			case db.ErrLeaderboardNotFound:
				code = http.StatusNotFound
				msg = fmt.Sprintf(msg, "leaderboard not found")
			default:
				msg = fmt.Sprintf(msg, err)
			}

			ctx.AbortWithStatusJSON(code, msg)
			return
		}

		ctx.JSON(http.StatusOK, entry)
	})

	// Update leaderboard entries.
	r.PATCH("/leaderboards/:lid/entries/:eid", func(ctx *gin.Context) {
		lid := ctx.Params.ByName("lid")

		leaderboardId, err := uuid.Parse(lid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
			return
		}

		eid := ctx.Params.ByName("eid")

		entryId, err := uuid.Parse(eid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
			return
		}

		var body db.Entry

		if err := ctx.BindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "invalid request body")
			return
		}

		if len(body.FieldNamesToValues) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "specify at least one field to update")
			return
		}
	
		entry, err := database.UpdateEntry(leaderboardId, entryId, body.FieldNamesToValues)
		if err != nil {
			code := http.StatusInternalServerError
			msg := "failed to create entry: %s"

			switch err {
			case db.ErrLeaderboardInvalidFieldName:
				code = http.StatusBadRequest
				msg = fmt.Sprintf(msg, "invalid field name")
			case db.ErrLeaderboardNotFound:
				code = http.StatusNotFound
				msg = fmt.Sprintf(msg, "leaderboard not found")
			case db.ErrEntryNotFound:
				code = http.StatusNotFound
				msg = fmt.Sprintf(msg, "entry not found")
			default:
				msg = fmt.Sprintf(msg, err)
			}

			ctx.AbortWithStatusJSON(code, msg)
			return

		}

		ctx.JSON(http.StatusOK, entry)
	})

	return r
}

func main() {
	database := db.NewMemDB()
	r := setupRouter(database)

	r.Run("localhost:8080")
}
