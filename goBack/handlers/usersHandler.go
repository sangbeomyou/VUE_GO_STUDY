package handlers

import (
	"context"
	"goBack/db"
	"goBack/utils" //
	"os"

	"goBack/ent/tn_user" // Import the bbs entity
	"goBack/models"
	"log"
	"net/http"

	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// 로그인
func GetUsersLoginHandler(c echo.Context) error {

	request := new(models.LoginRequest)
	// 리퀘스트 값이 없을때
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
		Select(
			tn_user.FieldUserID,
			tn_user.FieldUserName,
		).
		All(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to query TN_BBS: " + err.Error(),
		})
	}

	//user 빈값 체크
	if len(user) == 0 {
		return c.JSON(http.StatusOK, map[string]string{
			"success": "N",
			"error":   "No user found",
		})
	}

	// JWT 토큰 생성

	jwtKey := os.Getenv("SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userid"] = user[0].UserID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //토큰 완료 시간 설정

	t, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create the token: " + err.Error(),
		})
	}

	c.Logger().Info(user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": "Y",
		"token":   t,
		"result":  user,
	})
}
