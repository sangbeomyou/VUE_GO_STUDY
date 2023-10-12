package handlers

import (
	"goBack/db"
	"goBack/models"
	"net/http"
	"time"

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

func GetBbsInfoHandler(c echo.Context) error {
	client, err := db.ConnectDB() // Gorm으로 데이터베이스 연결을 설정한다고 가정합니다.
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to connect to database: " + err.Error(),
		})
	}

	// URL에서 idx 파라미터를 가져옵니다.
	idx := c.Param("idx")

	var bbs models.TN_BBS
	if err := client.Where("delete_yn = ? AND idx = ?", "N", idx).Select("idx, user_id, user_name, title, content, reg_date").First(&bbs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "No records found for idx: " + idx,
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to query TN_BBS: " + err.Error(),
		})
	}

	c.Logger().Info(bbs)
	return c.JSON(http.StatusOK, models.Response{
		Success: "Y",
		Result:  bbs, // 단일 게시물만 반환
	})
}

func PostBbsWriteHandler(c echo.Context) error {
	client, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to connect to database: " + err.Error(),
		})
	}

	var bbs models.TN_BBS

	// 요청 본문에서 정보 가져오기
	if err := c.Bind(&bbs); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to bind request body: " + err.Error(),
		})
	}
	// 나머지 설정
	bbs.RegDate = time.Now().Format("2006-01-02 15:04:05")
	bbs.DeleteYn = "N"

	// 게시물 추가
	if err := client.Create(&bbs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to insert into TN_BBS: " + err.Error(),
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Success": "Y",
	})
}

func PostBbsDeleteHandler(c echo.Context) error {
	client, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to connect to database: " + err.Error(),
		})
	}

	var req struct {
		Idx string `json:"Idx"`
	}

	// 요청 본문에서 정보 가져오기
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to bind request body: " + err.Error(),
		})
	}

	// idx로 delete_yn 필드를 'Y'로 업데이트
	if err := client.Model(&models.TN_BBS{}).Where("idx = ?", req.Idx).Update("delete_yn", "Y").Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to mark TN_BBS as deleted: " + err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Success": "Y",
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
