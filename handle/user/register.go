package user

import (
    "doubleJWT/handle"
    "doubleJWT/models"
    "doubleJWT/pkg/errno"
    "doubleJWT/pkg/uuid"
    "github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
    var (
        err error
        req RegisterReq
    )
    if err = c.BindJSON(&req); err != nil {
        handle.SendResult(c, errno.ErrBind, nil)
        return
    }

    if _, err = models.GetTBUser(map[string]interface{}{"username": req.Username}); err == nil {
        handle.SendResult(c, errno.ErrUserHadRegister, nil)
        return
    }

    user := models.TBUser{
        UUID:       uuid.NewUUID(),
        UserName:   req.Username,
        PassWord:   req.Password,
    }

    if err = user.Encrypt(); err != nil {
        handle.SendResult(c, errno.ErrPassword, nil)
        return
    }

    if err = user.AddTBUser(); err != nil {
        handle.SendResult(c, errno.ErrDatabase, nil)
        return
    }

    res := RegisterRes{
        UUID:user.UUID,
    }

    handle.SendResult(c, nil, res)
}
