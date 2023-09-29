package handlers

import (
	"context"
	"goBack/db"
	"goBack/utils" //

	"goBack/ent/tn_bbs"  // Import the bbs entity
	"goBack/ent/tn_user" // Import the bbs entity
	"goBack/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersHandler(c echo.Context) error {
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
		All(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to query TN_BBS: " + err.Error(),
		})
	}

	c.Logger().Info(bbsList)
	return c.JSON(http.StatusOK, bbsList)
}

func GetUsersLoginHandler(c echo.Context) error {
	// 리퀘스트 값이 없을때
	request := new(models.LoginRequest)
	if err := c.Bind(request); err != nil {
		log.Println("Failed to bind request:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "bad request"})
	}
	// db연결
	client, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to connect to database: " + err.Error(),
		})
	}
	defer client.Close()

	user, err := client.TN_USER.
		Query().
		Where(
			tn_user.DelectYnEQ("N"),
			tn_user.UserIDEQ(request.UserID),
			tn_user.PasswordEQ(utils.HashPassword(request.Password)),
		).
		All(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to query TN_BBS: " + err.Error(),
		})
	}

	c.Logger().Info(user)
	return c.JSON(http.StatusOK, user)
}
