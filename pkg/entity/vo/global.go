package vo

import (
	"fmt"
	"strconv"

	"github.com/megrez/pkg/dao"
	"github.com/megrez/pkg/entity/po"
)

type Global struct {
	BlogURL         string
	BlogTitle       string
	BlogDescription string
	BlogCover       string
	IPCRecord       string
	BlogAge         float64
	Github          string
	QQ              string
	Email           string
	Telegram        string
	Menus           []*Menu
	Categories      []*BriefCategory
	LatestArticles  []*LatestArticl
	LatestComments  []*LatestComment
}

type LatestArticl struct {
	Title string
	URL   string
}

type LatestComment struct {
	Content string
	URL     string
}

func GetLatestArticleFromPO(article *po.Article) (*LatestArticl, error) {
	latestArticle := &LatestArticl{
		Title: article.Title,
	}
	// dao, err := dao.GetDAO()
	// if err != nil {
	// 	return latestArticle, err
	// }
	// blogURL, err := dao.GetOptionByKey(OptionKeyBlogURL)
	// if err != nil {
	// 	return latestArticle, err
	// }
	latestArticle.URL = "/article/" + strconv.Itoa(int(article.ID))
	return latestArticle, nil
}

func GetLatestCommentFromPO(comment *po.Comment) (*LatestComment, error) {
	dao, err := dao.GetDAO()
	if err != nil {
		return nil, err
	}
	latestComment := &LatestComment{
		Content: comment.Content,
	}
	indexPageSizeStr, err := dao.GetOptionByKey(OptionComentsPageSize)
	if err != nil {
		return nil, err
	}
	indexPageSize, err := strconv.Atoi(indexPageSizeStr)
	if err != nil {
		return nil, err
	}
	var rootComments []po.Comment
	// comment for article
	if comment.Type == 1 {
		rootComments, err = dao.ListRootCommentsByArticleID(comment.ArticleID, 0, 0)
		if err != nil {
			return nil, err
		}
		var index int
		for i, rootComment := range rootComments {
			if rootComment.ID == comment.ID || rootComment.ID == comment.RootID {
				index = i + 1
				break
			}
		}
		pagination := (index + indexPageSize - 1) / indexPageSize
		url := fmt.Sprintf("/article/%d/comment-page/%d#comment-%d", comment.ArticleID, pagination, comment.ID)
		latestComment.URL = url
	} else if comment.Type == 2 {
		// comment for page
		rootComments, err = dao.ListRootCommentsByPageID(comment.PageID, 0, 0)
		if err != nil {
			return nil, err
		}
		var index int
		for i, rootComment := range rootComments {
			if rootComment.ID == comment.ID || rootComment.ID == comment.RootID {
				index = i + 1
				break
			}
		}
		pagination := (index + indexPageSize - 1) / indexPageSize
		page, err := dao.GetPageByID(comment.PageID)
		if err != nil {
			return nil, err
		}
		url := fmt.Sprintf("/%s/comment-page/%d#comment-%d", page.Slug, pagination, comment.ID)
		latestComment.URL = url
	}

	return latestComment, nil
}

func GetGlobalOption() (Global, error) {
	global := Global{}
	dao, err := dao.GetDAO()
	if err != nil {
		return global, err
	}
	blogTitle, err := dao.GetOptionByKey(OptionKeyBlogTitle)
	if err == nil {
		global.BlogTitle = blogTitle
	}
	blogURL, err := dao.GetOptionByKey(OptionKeyBlogURL)
	if err == nil {
		global.BlogURL = blogURL
	}
	blogDescription, err := dao.GetOptionByKey(OptionKeyBlogDescription)
	if err == nil {
		global.BlogDescription = blogDescription
	}
	blogCover, err := dao.GetOptionByKey(OptionKeyBlogCover)
	if err == nil {
		global.BlogCover = blogCover
	}
	ipcRecord, err := dao.GetOptionByKey(OptionKeyIPCRecord)
	if err == nil {
		global.IPCRecord = ipcRecord
	}

	// list menus
	menuPOs, err := dao.ListAllMenus()
	if err != nil {
		return global, err
	}
	menus := []*Menu{}
	for _, mensPO := range menuPOs {
		menus = append(menus, GetMenuFromPO(mensPO))
	}
	global.Menus = menus

	// list category
	categoryPOs, err := dao.ListAllCategories()
	if err != nil {
		return global, err
	}
	categories := []*BriefCategory{}
	for _, categoryPO := range categoryPOs {
		categories = append(categories, GetBriefCategoryFromPO(categoryPO))
	}
	global.Categories = categories

	// latest articles
	articlePOs, err := dao.ListLatestArticles()
	if err == nil {
		latestArticles := []*LatestArticl{}
		for _, articlePO := range articlePOs {
			latestArticle, err := GetLatestArticleFromPO(&articlePO)
			if err == nil {
				latestArticles = append(latestArticles, latestArticle)
			}
		}
		global.LatestArticles = latestArticles
	}

	// latest comments
	commentPOs, err := dao.ListLatestComments()
	if err == nil {
		latestComments := []*LatestComment{}
		for _, commentPO := range commentPOs {
			latestComment, err := GetLatestCommentFromPO(&commentPO)
			if err == nil {
				latestComments = append(latestComments, latestComment)
			}
		}
		global.LatestComments = latestComments
	}
	return global, nil
}