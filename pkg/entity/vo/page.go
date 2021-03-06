package vo

import "github.com/megrez/pkg/model"

type Page struct {
	ID          uint
	Name        string
	Slug        string
	Cover       string
	Password    string
	Private     bool
	Visits      int64
	Likes       int64
	CommentsNum int64
	Status      int
}

func GetPageFromPO(page model.Page) *Page {
	return &Page{
		Name:     page.Name,
		Slug:     page.Slug,
		Cover:    page.Cover,
		Password: page.Password,
		Private:  page.Private,
		Visits:   page.Visits,
		Likes:    page.Likes,
		Status:   page.Status,
	}
}
