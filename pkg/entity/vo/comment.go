package vo

import (
	"time"

	"github.com/megrez/pkg/model"
)

type Role int

const (
	RoleAdmin Role = iota
	RoleGuest
	RoleFriend
)

func (role Role) String() string {
	switch role {
	case RoleAdmin:
		return "admin"
	case RoleGuest:
		return "guest"
	case RoleFriend:
		return "friend"
	default:
		return "unknown"
	}
}

type Comment struct {
	ID          uint
	ParentID    uint
	ArticleID   uint
	PageID      uint
	Content     string
	Status      int
	HasChild    bool
	SubComments []*SubComment
	Author      string
	Avatar      string
	Role        string
	Mail        string
	Site        string
	Agent       string
	CreateTime  time.Time
}

type SubComment struct {
	ID         uint
	ParentID   uint
	ArticleID  uint
	PageID     uint
	Content    string
	Status     int
	Author     string
	Avatar     string
	Role       string
	Mail       string
	Site       string
	Agent      string
	CreateTime time.Time
}

func GetCommentFromPO(comment model.Comment) (*Comment, error) {
	commentVO := &Comment{}
	commentVO.ID = comment.ID
	commentVO.ParentID = comment.ParentID
	commentVO.ArticleID = comment.ArticleID
	commentVO.Content = comment.Content
	commentVO.Status = comment.Status
	commentVO.CreateTime = comment.CreateTime
	commentVO.Author = comment.Author
	commentVO.Role = Role(comment.Role).String()
	commentVO.Mail = comment.Email
	commentVO.Site = comment.Site
	// TODO: 设置头像
	commentVO.Avatar = "https://cdn.rawchen.com/logo/alkaidchen.jpg"
	commentVO.Agent = comment.Agent
	subCommentPOs, err := model.ListCommentsByRootID(commentVO.ID)
	if err != nil {
		return commentVO, err
	}
	var subComments []*SubComment
	for _, subCommentPO := range subCommentPOs {
		subComment, err := GetSubCommentFromPO(subCommentPO)
		if err != nil {
			return commentVO, err
		}
		subComments = append(subComments, subComment)
	}
	commentVO.SubComments = subComments
	if len(subComments) == 0 {
		commentVO.HasChild = false
	} else {
		commentVO.HasChild = true
	}
	return commentVO, nil
}

func GetSubCommentFromPO(comment model.Comment) (*SubComment, error) {
	subComment := &SubComment{}
	subComment.ID = comment.ID
	subComment.ParentID = comment.ParentID
	subComment.ArticleID = comment.ArticleID
	subComment.Content = comment.Content
	subComment.Status = comment.Status
	subComment.CreateTime = comment.CreateTime
	subComment.Author = comment.Author
	subComment.Role = Role(comment.Role).String()
	subComment.Mail = comment.Email
	subComment.Site = comment.Site
	// TODO: 设置头像
	subComment.Avatar = "https://cdn.rawchen.com/logo/alkaidchen.jpg"
	// TODO: 计算Agent的浏览器和内核
	subComment.Agent = comment.Agent
	return subComment, nil
}
