package models

import (
    "golang.org/x/crypto/bcrypt"
)

// TBUser 用户表
type TBUser struct {
    BaseModel
    UUID        string  `gorm:"column:uuid;not null;"       json:"uuid"         description:"uuid" `
    UserName    string  `gorm:"column:username;not null;"   json:"username"     description:"账号" validate:"min=1,max=50" `
    PassWord    string  `gorm:"column:password;not null;"   json:"password"     description:"密码" validate:"min=1,max=50" `
}

// TableName 表名
func (t *TBUser) TableName() string {
    return "tb_user"
}

// AddTBUser 插入记录
func (t *TBUser) AddTBUser() error {
    return MysqlPool.Create(t).Error
}

// GetTBUser 查找
func GetTBUser(query map[string]interface{}) (*TBUser, error) {
    var user TBUser
    err := MysqlPool.Where(query).First(&user).Error
    return &user, err
}

// DeleteTBUser 删除记录
func DeleteTBUser(uuid string) error {
    user := TBUser{
        UUID:uuid,
    }
    return MysqlPool.Delete(&user).Error
}

// GetAllTBUsers 获取所有记录
func GetAllTBUsers(query map[string]interface{}, offset int, limit int) ([]TBUser, error) {
    if offset < 0 {
        offset = 0
    }
    if limit <= 0 {
        limit = 1
    }
    var users  []TBUser
    err := MysqlPool.Where(query).Offset(offset).Limit(limit).Find(&users).Error
    return users, err
}

// UpdateTBUser 更新记录
func (t *TBUser) UpdateTBUser() error {
    return MysqlPool.Save(t).Error
}

// Compare 比较密码
func (u *TBUser) Compare(pwd string) (err error) {
    err = bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(pwd))
    return
}

// Encrypt 密码加密
func (u *TBUser) Encrypt() (err error) {
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(u.PassWord), bcrypt.DefaultCost)
    u.PassWord = string(hashedBytes)
    return
}
