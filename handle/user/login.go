package user

import (
    "doubleJWT/handle"
    "doubleJWT/models"
    "doubleJWT/pkg/token"
    "doubleJWT/pkg/errno"
    "github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
    var req LoginReq
    if err := c.BindJSON(&req); err != nil {
        handle.SendResult(c, errno.ErrBind, nil)
        return
    }

    user, err := models.GetTBUser(map[string]interface{}{"username":req.Username})
    if err != nil {
        handle.SendResult(c, errno.ErrUserNotFound, nil)
        return
    }

    if err = user.Compare(req.Password); err != nil {
        handle.SendResult(c, errno.ErrPassword, nil)
        return
    }

    shortToken, err := token.SignToken(token.Context{ID:user.ID, Username:user.UserName})
    if err != nil {
        handle.SendResult(c, errno.ErrToken, nil)
        return
    }

    RefreshToken, err := token.SignRefreshToken(token.Context{ID:user.ID, Username:user.UserName})
    if err != nil {
        handle.SendResult(c, errno.ErrToken, nil)
        return
    }

    resp := LoginRes{
        Token:          shortToken,
        RefreshToken:   RefreshToken,
    }

    handle.SendResult(c, nil, resp)
}
