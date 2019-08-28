package models

import "gopkg.in/go-playground/validator.v9"

// TBUser 评论表
type TBComment struct {
    BaseModel
    TopicId     uint        `gorm:"column:topic_id;not null;"       json:"content"    description:"主题id" binding:"required" `
    TopicType   string      `gorm:"column:topic_type;not null;"     json:"content"    description:"主题类型" binding:"required" `
    Content     string      `gorm:"column:content;not null;"        json:"content"    description:"评论内容" binding:"required" `
    FromUid     uint        `gorm:"column:from_uid;not null;"       json:"content"    description:"评论用户id" binding:"required" `
}

// TableName 表名
func (t *TBComment) TableName() string {
    return "tb_comment"
}

// AddTBComment 插入
func (t *TBComment) AddTBComment() error {
    return MysqlPool.Create(t).Error
}

// GetTBComment 查找
func GetTBComment(query map[string]interface{}) (*TBComment, error) {
    var comment TBComment
    err := MysqlPool.Where(query).First(&comment).Error
    return &comment, err
}

// GetAllTBComments 获取所有记录
func GetAllTBComments(query map[string]interface{}, offset int, limit int) ([]TBComment, error) {
    if offset < 0 {
        offset = 0
    }
    if limit <= 0 {
        limit = 1
    }
    var comments  []TBComment
    err := MysqlPool.Where(query).Offset(offset).Limit(limit).Find(&comments).Error
    return comments, err
}

func (t *TBComment) Validate() error {
    validate := validator.New()
    return validate.Struct(t)
}