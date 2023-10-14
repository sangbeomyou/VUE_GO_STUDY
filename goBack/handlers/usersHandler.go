package handlers

import (
	"fmt"
	"time"

	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)

// 로그인
// func GetUsersLoginHandler(c echo.Context) error {

// 	request := new(models.LoginRequest)
// 	// 리퀘스트 값이 없을때
// 	if err := c.Bind(request); err != nil {
// 		log.Println("Failed to bind request:", err)
// 		return c.JSON(http.StatusBadRequest, map[string]string{"status": "bad request"})
// 	}
// 	// db연결
// 	client, err := db.ConnectDB()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Failed to connect to database: " + err.Error(),
// 		})
// 	}
// 	defer client.Close()

// 	user, err := client.TN_USER.
// 		Query().
// 		Where(
// 			tn_user.DelectYnEQ("N"),
// 			tn_user.UserIDEQ(request.UserID),
// 			tn_user.PasswordEQ(utils.HashPassword(request.Password)),
// 		).
// 		Select(
// 			tn_user.FieldUserID,
// 			tn_user.FieldUserName,
// 		).
// 		All(context.Background())

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Failed to query TN_BBS: " + err.Error(),
// 		})
// 	}

// 	//user 빈값 체크
// 	if len(user) == 0 {
// 		return c.JSON(http.StatusOK, map[string]string{
// 			"success": "N",
// 			"error":   "No user found",
// 		})
// 	}

// 	// JWT 토큰 생성

// 	jwtKey := os.Getenv("SECRET_KEY")
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["userid"] = user[0].UserID
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //토큰 완료 시간 설정

// 	t, err := token.SignedString([]byte(jwtKey))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Failed to create the token: " + err.Error(),
// 		})
// 	}

// 	c.Logger().Info(user)
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"success": "Y",
// 		"token":   t,
// 		"result":  user,
// 	})
// }

func SetTokenHandler(c echo.Context) error {
	var body map[string]string
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid data",
		})
	}

	token := body["token"]
	if token == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Token is missing",
		})
	}

	// HTTPOnly 쿠키로 토큰 설정
	c.SetCookie(&http.Cookie{
		Name:     "authToken",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,        // CSRF 방지
		Expires:  time.Now().Add(24 * time.Hour), // 24시간 후 만료
		Path:     "/",                            // 모든 경로에서 쿠키가 유효하게 설정
	})

	return c.JSON(http.StatusOK, map[string]string{
		"message": token,
	})
}

func GetInitUserInfoHandler(c echo.Context) error {
	// 컨텍스트에서 Firebase Auth 클라이언트 가져오기
	client, ok := c.Get("firebase_auth_client").(*auth.Client)
	if !ok {
		return c.JSON(http.StatusBadRequest, "Failed to retrieve Firebase Auth client")
	}

	// 컨텍스트에서 uid 가져오기
	uid, ok := c.Get("uid").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, "Invalid UID")
	}

	// UID를 사용하여 Firebase에서 사용자 정보 가져오기
	userInfo, err := client.GetUser(c.Request().Context(), uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error fetching user info for UID: %s", uid))
	}

	// 사용자 정보를 응답으로 반환
	return c.JSON(http.StatusOK, userInfo)
}

func GetLogoutHandler(c echo.Context) error {
	// 쿠키 만료 시간을 현재 시간보다 이전으로 설정
	c.SetCookie(&http.Cookie{
		Name:     "authToken",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
		Path:     "/",
	})

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Logged out successfully",
	})
}
