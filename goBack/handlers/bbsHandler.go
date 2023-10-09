package handlers

import (
	"goBack/db"
	"goBack/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetBbsHandler(c echo.Context) error {
	client, err := db.ConnectDB() // Gorm으로 데이터베이스 연결을 설정한다고 가정합니다.
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to connect to database: " + err.Error(),
		})
	}

	var bbsList []models.TN_BBS
	if err := client.Where("delete_yn = ?", "N").Select("idx, user_id, user_name, title, content, reg_date").Find(&bbsList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "No records found",
			})
		}
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

// package handlers

// import (
// 	"context"
// 	"goBack/db"
// 	"goBack/models"

// 	"goBack/ent/tn_bbs" // Import the bbs entity
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func GetBbsHandler(c echo.Context) error {
// 	client, err := db.ConnectDB()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Failed to connect to database: " + err.Error(),
// 		})
// 	}
// 	defer client.Close()

// 	bbsList, err := client.TN_BBS.
// 		Query().
// 		Where(tn_bbs.DelectYnEQ("N")).
// 		Select(
// 			tn_bbs.FieldUserID,
// 			tn_bbs.FieldUserName,
// 			tn_bbs.FieldTitle,
// 			tn_bbs.FieldContent,
// 			tn_bbs.FieldRegDate,
// 		).
// 		All(context.Background())
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Failed to query TN_BBS: " + err.Error(),
// 		})
// 	}

// 	c.Logger().Info(bbsList)
// 	return c.JSON(http.StatusOK, models.Response{
// 		Success: "Y",
// 		Result:  bbsList,
// 	})
// }
