package handlers

import (
	"goBack/db"
	"goBack/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetBbsHandler(c echo.Context) error {
	client, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": "N",
			"error":   "Failed to connect to database: " + err.Error(),
		})
	}

	// 쿼리 파라미터에서 page와 limit 값을 가져옵니다.
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	// page와 limit을 정수로 변환합니다. 오류가 있을 경우 기본값을 설정할 수 있습니다.
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1 // 기본 페이지 값 설정
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10 // 기본 limit 값 설정
	}

	var bbsList []models.TN_BBS
	var count int64
	offset := (page - 1) * limit // offset 계산

	// 먼저, 전체 게시글의 수를 가져옵니다.
	if result := client.Model(&models.TN_BBS{}).Where("delete_yn = ?", "N").Count(&count); result.Error != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": "N",
			"error":   "Failed to count TN_BBS: " + result.Error.Error(),
		})
	}

	// 게시글을 limit와 offset을 이용하여 불러옵니다.
	if result := client.Where("delete_yn = ?", "N").Select("idx, user_id, user_name, title, content, reg_date").
		Offset(offset).Limit(limit).Find(&bbsList); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"success": "N",
				"error":   "No records found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": "N",
			"error":   "Failed to query TN_BBS: " + result.Error.Error(),
		})
	}

	// 정상적으로 처리되었을 경우, 데이터와 함께 전체 페이지 수와 현재 페이지 번호를 클라이언트에 반환합니다.
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": "Y",
		"result":  bbsList,
		"total":   count,
		"page":    page,
		"pages":   (int(count) + limit - 1) / limit, // 전체 페이지 수 계산
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

func PostBbsEditHandler(c echo.Context) error {
	client, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to connect to database: " + err.Error(),
		})
	}

	var bbs models.TN_BBS_API

	// 요청 본문에서 정보 가져오기
	if err := c.Bind(&bbs); err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"error": "Failed to bind request body: " + err.Error(),
		})
	}
	// 나머지 설정
	bbs.UpdDate = time.Now().Format("2006-01-02 15:04:05")

	// 게시물 업데이트
	if err := client.Model(&models.TN_BBS{}).Where("idx = ?", bbs.Idx).Updates(models.TN_BBS{
		Title:   bbs.Title,
		Content: bbs.Content,
		UpdDate: bbs.UpdDate,
	}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Failed to update TN_BBS: Record not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update TN_BBS: " + err.Error(),
		})
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
