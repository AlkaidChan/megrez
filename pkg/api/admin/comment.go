package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
	"net/http"
	"strconv"
)

func ListComments(c *gin.Context) {
	var pageNum, pageSize int
	var err error
	if c.Query("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(c.Query("pageNum"))
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
			return
		}
	}
	if c.Query("pageSize") == "" {
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(c.Query("pageSize"))
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
			return
		}
	}
	comments, err := model.ListAllComments(pageNum, pageSize)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	var commentDTOs []dto.CommentListDTO
	for _, comment := range comments {
		commentDTO := dto.CommentListDTO{}
		err := commentDTO.LoadFromModel(comment)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		commentDTOs = append(commentDTOs, commentDTO)
	}
	total, err := model.CountAllArticles()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	pagination := dto.Pagination{
		List:     commentDTOs,
		Current:  pageNum,
		PageSize: pageSize,
		Total:    total,
	}
	c.JSON(http.StatusOK, errmsg.Success(pagination))
}
