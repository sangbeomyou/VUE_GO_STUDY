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

// 댓글 불러오기
func GetCommentHandler(c echo.Context) error {
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
	BbsIdStr := c.QueryParam("bbsId")

	// page와 limit을 정수로 변환합니다. 오류가 있을 경우 기본값을 설정할 수 있습니다.
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1 // 기본 페이지 값 설정
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10 // 기본 limit 값 설정
	}
	BbsId, err := strconv.Atoi(BbsIdStr)
	if err != nil || limit < 1 {
		return c.JSON(http.StatusOK, map[string]string{
			"success": "N",
			"error":   "Failed to bind request body: " + err.Error(),
		})
	}

	var commentList []models.TN_COMMENT
	var count int64
	offset := (page - 1) * limit // offset 계산

	// 먼저, 전체 게시글의 수를 가져옵니다.
	if result := client.Model(&models.TN_COMMENT{}).Where("delete_yn = ? AND bbs_id = ?", "N", BbsId).Count(&count); result.Error != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": "N",
			"error":   "Failed to count TN_COMMENT: " + result.Error.Error(),
		})
	}

	// 게시글을 limit와 offset을 이용하여 불러옵니다.
	if result := client.Where("delete_yn = ? AND bbs_id = ?", "N", BbsId).Select("idx, bbs_id, user_id, user_name, content, reg_date").
		Offset(offset).Limit(limit).Find(&commentList); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"success": "N",
				"error":   "No records found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": "N",
			"error":   "Failed to query TN_COMMENT: " + result.Error.Error(),
		})
	}

	// 정상적으로 처리되었을 경우, 데이터와 함께 전체 페이지 수와 현재 페이지 번호를 클라이언트에 반환합니다.
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": "Y",
		"result":  commentList,
		"total":   count,
		"page":    page,
		"pages":   (int(count) + limit - 1) / limit, // 전체 페이지 수 계산
	})
}

// 댓글 작성 부분
func PostCommentWriteHandler(c echo.Context) error {
	client, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": "N",
			"error":   "Failed to connect to database: " + err.Error(),
		})
	}
	var comment models.TN_COMMENT

	// 요청 본문에서 정보 가져오기
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": "N",
			"error":   "파라미터 오류: " + err.Error(),
		})
	}
	// 나머지 설정
	comment.RegDate = time.Now().Format("2006-01-02 15:04:05")
	comment.DeleteYn = "N"

	// 게시물 추가
	if err := client.Create(&comment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusOK, map[string]string{
				"success": "N",
				"error":   "Failed to insert into TN_COMMENT: " + err.Error(),
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": "Y",
		"result":  comment, //생성한 객체 보내기
	})
}

func PostCommentEditHandler(c echo.Context) error {
	client, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to connect to database: " + err.Error(),
		})
	}

	var comment models.TN_COMMENT_API

	// 요청 본문에서 정보 가져오기
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"error": "Failed to bind request body: " + err.Error(),
		})
	}
	// 나머지 설정
	comment.UpdDate = time.Now().Format("2006-01-02 15:04:05")

	// 게시물 업데이트
	if err := client.Model(&models.TN_COMMENT{}).Where("idx = ?", comment.Idx).Updates(models.TN_COMMENT{
		Content: comment.Content,
		UpdDate: comment.UpdDate,
	}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Failed to update TN_COMMENT: Record not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update TN_COMMENT: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"success": "Y",
	})
}

func PostCommentDeleteHandler(c echo.Context) error {
	client, err := db.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"success": "N",
			"error":   "Failed to connect to database: " + err.Error(),
		})
	}

	var req struct {
		Idx string `json:"Idx"`
	}

	// 요청 본문에서 정보 가져오기
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"success": "N",
			"error":   "Failed to bind request body: " + err.Error(),
		})
	}

	// idx로 delete_yn 필드를 'Y'로 업데이트
	if err := client.Model(&models.TN_COMMENT{}).Where("idx = ?", req.Idx).Update("delete_yn", "Y").Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"success": "N",
				"error":   "Failed to mark TN_COMMENT as deleted: " + err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"success": "Y",
	})
}
