package model

import "gorm.io/gorm"

type Category struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Slug        string `gorm:"type:varchar(255);uniqueIndex" json:"slug"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Status      int    `gorm:"type:int(11)" json:"status"`
}

// CreateCategory handle create category
func CreateCategory(tx *gorm.DB, category *Category) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Create(category)
	return result.Error
}

// UpdateCategoryByID update article by id and data
func UpdateCategoryByID(tx *gorm.DB, id uint, category *Category) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Model(&category).Where("id= ？", id).Updates(&category)
	return result.Error
}

// DeleteCategoryByID delete article by id
func DeleteCategoryByID(tx *gorm.DB, id uint) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Delete(&Category{}, id)
	return result.Error
}

// DeleteArticleCategoriesByArticleID delete article categories by articleID
func DeleteArticleCategoriesByArticleID(tx *gorm.DB, aid uint) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Delete(&ArticleCategory{}, "article_id = ?", aid)
	return result.Error
}

// GetCategoryByID return category by id
func GetCategoryByID(id uint) (Category, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	category := Category{}
	result := db.First(&category, id)
	return category, result.Error
}

// GetCategoryBySlug return category by slug
func GetCategoryBySlug(slug string) (Category, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	category := Category{}
	result := db.First(&category, "`slug` = ?", slug)
	return category, result.Error
}

// ListAllCategories return all categories
func ListAllCategories() ([]Category, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var categories []Category
	result := db.Find(&categories)
	return categories, result.Error
}

// ListCategoriesByPage return categories by page
func ListCategoriesByPage(pageNum int, pageSize int) ([]Category, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var categories []Category
	result := db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&categories)
	return categories, result.Error
}

// ListCategoriesByArticleID return categories by articleID
func ListCategoriesByArticleID(aid uint) ([]Category, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var categories []Category
	result := db.Where("id in (?)", db.Table("article_categories").Select("category_id").Where("article_id = ?", aid)).Find(&categories)
	return categories, result.Error
}

// CountCategories return count of categories
func CountCategories() (int64, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var count int64
	result := db.Model(&Category{}).Count(&count)
	return count, result.Error
}
