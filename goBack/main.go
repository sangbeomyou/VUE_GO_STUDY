package main

import (
	"goBack/handlers" //  함수가 있는 패키지 임포트
	customMiddleware "goBack/middleware"
	"log"

	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	// 미들웨어 설정 파이어베이스 토큰 확인용
	e.Use(customMiddleware.FirebaseAuthentication())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// e.POST("/user/login", handlers.GetUsersLoginHandler)
	bbsGroup := e.Group("/bbs")
	bbsGroup.GET("/bbsInfo/:idx", handlers.GetBbsInfoHandler)
	bbsGroup.GET("/bbsList", handlers.GetBbsHandler)
	bbsGroup.POST("/write", handlers.PostBbsWriteHandler)
	bbsGroup.POST("/delete", handlers.PostBbsDeleteHandler)
	bbsGroup.POST("/edit", handlers.PostBbsEditHandler)

	commentGroup := e.Group("/comment")
	commentGroup.POST("/write", handlers.PostCommentWriteHandler)
	commentGroup.GET("/list", handlers.GetCommentHandler)
	commentGroup.POST("/edit", handlers.PostCommentEditHandler)
	commentGroup.POST("/delete", handlers.PostCommentDeleteHandler)

	e.GET("/InitUser", handlers.GetInitUserInfoHandler)

	//위 경로 미들웨어 파이어 베이스 토큰 미확인
	publicGroup := e.Group("/public")
	publicGroup.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/login")
	})
	publicGroup.POST("/setToken", handlers.SetTokenHandler) //토큰 쿠키에 세팅
	publicGroup.GET("/logout", handlers.GetLogoutHandler)   //로그아웃 쿠키 지우기

	e.Logger.Fatal(e.Start(":8080"))
}
