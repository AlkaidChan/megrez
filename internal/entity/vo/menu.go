package vo

import "github.com/megrez/internal/entity/po"

type Menu struct {
	Name string
	Slug string
}

func GetMenuFromPO(po po.Menu) *Menu {
	menu := &Menu{
		Name: po.Name,
		Slug: po.Slug,
	}
	return menu
}