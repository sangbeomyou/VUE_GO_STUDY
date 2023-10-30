package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

// FirebaseAuthentication은 Firebase로부터 발급받은 토큰을 검증하는 미들웨어 함수를 반환합니다.
func FirebaseAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// 이경로는 인증 필요 안함
			if strings.HasPrefix(c.Request().URL.Path, "/public") {
				return next(c)
			}
			// 쿠키에서 "authToken" 이름의 값을 읽어옵니다.
			cookie, err := c.Cookie("authToken")
			if err != nil || cookie.Value == "" {
				log.Printf("missing token in cookie")
				// 쿠키없으면 로그인
				return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/login")
			}

			// 쿠키의 값을 ID 토큰으로 사용합니다.
			idToken := cookie.Value

			ctx := context.Background()

			opt := option.WithCredentialsFile("./vuefrontConfig.json")

			app, err := firebase.NewApp(ctx, nil, opt)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, "error initializing firebase app")
			}

			// Firebase Auth 클라이언트를 생성합니다.
			client, err := app.Auth(ctx)
			if err != nil {
				log.Printf("Error creating Firebase Auth client: %v", err) // 로깅
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"error":  "error creating auth client",
					"detail": err.Error(),
				})
			}

			// ID 토큰을 검증합니다.
			tokenInfo, err := client.VerifyIDToken(ctx, idToken)
			if err != nil {
				//인증토큰 아니라면
				log.Printf("Error verifying ID token: %v", err)
				// authToken 쿠키 삭제
				c.SetCookie(&http.Cookie{
					Name:     "authToken",
					Value:    "",
					HttpOnly: true,
					Expires:  time.Unix(0, 0),
					Path:     "/",
					SameSite: http.SameSiteStrictMode,
				})
				//로그인 페이지로
				return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/login")
			}

			// 검증된 토큰의 사용자 UID를 컨텍스트에 저장합니다.
			// 이 후의 핸들러에서 UID를 사용하여 사용자에게 관련된 작업을 수행할 수 있습니다.
			c.Set("uid", tokenInfo.UID)
			// Firebase Auth 클라이언트를 컨텍스트에 저장
			c.Set("firebase_auth_client", client)
			return next(c)
		}
	}
}
