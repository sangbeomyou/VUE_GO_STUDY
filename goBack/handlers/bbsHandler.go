package handlers

import (
	"context"
	"goBack/db"
	"goBack/models"

	"goBack/ent/tn_bbs" // Import the bbs entity
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBbsHandler(c echo.Context) error {
	client, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to connect to database: " + err.Error(),
		})
	}
	defer client.Close()

	bbsList, err := client.TN_BBS.
		Query().
		Where(tn_bbs.DelectYnEQ("N")).
		Select(
			tn_bbs.FieldUserID,
			tn_bbs.FieldUserName,
			tn_bbs.FieldTitle,
			tn_bbs.FieldContent,
			tn_bbs.FieldRegDate,
		).
		All(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to query TN_BBS: " + err.Error(),
		})
	}

	c.Logger().Info(bbsList)
	return c.JSON(http.StatusOK, models.Response{
		Success: "Y",
		Result:  bbsList,
	})
}
