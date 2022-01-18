package site

import (
	"log"
	"strconv"

	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
)

func RouteLink(g *gin.Engine) {
	g.GET("/links", listLinks)
	g.GET("/links/comment-page/:pageNum", listLinks)
	g.POST("/admin/link", createLink)
}

func listLinks(c *gin.Context) {
	var pageNum, pageSize int
	var err error
	if c.Param("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(c.Param("pageNum"))
		if err != nil {
			log.Println("incorrect param pageNum, err:", err)
			// TODO: 应该是 4XX?
			c.Redirect(500, "/error")
		}
	}
	if c.Param("pageSize") == "" {
		pageSize = 8
	} else {
		pageSize, err = strconv.Atoi(c.Param("pageSize"))
		if err != nil {
			log.Println("incorrect param pageSize, err:", err)
			c.Redirect(500, "/error")
		}
	}

	linkPOs, err := model.ListAllLinks()
	links := []*vo.Link{}
	for _, linkPO := range linkPOs {
		links = append(links, vo.GetLinkFromPO(linkPO))
	}

	commentPOs, err := model.ListRootCommentsByPageID(2, pageNum, pageSize)
	if err != nil {
		c.Redirect(500, "/error")
	}
	var comments []*vo.Comment
	for _, commentPO := range commentPOs {
		comment, err := vo.GetCommentFromPO(commentPO)
		if err != nil {
			c.Redirect(500, "/error")
		}
		comments = append(comments, comment)
	}

	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(500, "/error")
	}

	page := struct {
		ID          uint
		Name        string
		Slug        string
		CommentsNum int64
		Visits      int64
	}{
		ID:     2,
		Name:   "友链",
		Slug:   "links",
		Visits: 10086,
	}
	commentsNum, err := model.CountRootCommentsByPageID(page.ID)
	if err != nil {
		c.Redirect(500, "/error")
	}
	page.CommentsNum = commentsNum
	pagination := vo.CalculatePagination(pageNum, pageSize, int(commentsNum))
	c.HTML(200, "links.html", pongo2.Context{"page": page, "pagination": pagination, "links": links, "comments": comments, "global": globalOption})
}

func createLink(c *gin.Context) {
	name := c.PostForm("name")
	addr := c.PostForm("addr")
	logo := c.PostForm("logo")

	if name == "" || addr == "" || logo == "" {
		c.JSON(500, "invalid parameter")
		return
	}
	link := &model.Link{
		Name: name,
		Addr: addr,
		Logo: logo,
	}
	err := model.CreateLink(link)
	if err != nil {
		c.JSON(500, "failed to create link")
		return
	}
	c.JSON(200, "success")
}