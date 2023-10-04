package middleware

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	jwtKey := os.Getenv("SECRET_KEY")

	return func(c echo.Context) error {
		//헤더에서 토큰값 가지고 오기
		tokenString := c.Request().Header.Get("Authorization")

		// 토큰이 없는 경우 에러 응답 로그인 화면으로 보내야함
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "authorization is required")
		}

		// JWT 토큰을 파싱하고 검증
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		// 토큰이 유효한 경우, "userid" 클레임 값을 컨텍스트에 저장
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userid", claims["userid"])
			return next(c)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}
}
