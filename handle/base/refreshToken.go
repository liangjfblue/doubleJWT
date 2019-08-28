package base

import (
    "doubleJWT/handle"
    "doubleJWT/models"
    "doubleJWT/pkg/errno"
    "doubleJWT/pkg/token"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func RefreshToken(c *gin.Context) {
    t, exists := c.Get("token")
    if !exists {
        handle.SendResult(c, errno.ErrToken, nil)
        return
    }

    tCtx := t.(token.Context)
    logrus.Info(tCtx)
    user, err := models.GetTBUser(map[string]interface{}{"id":tCtx.ID})
    if err != nil {
        handle.SendResult(c, errno.ErrUserNotFound, nil)
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

    resp := RefreshTokenRes{
        Token:          shortToken,
        RefreshToken:   RefreshToken,
    }

    handle.SendResult(c, nil, resp)

}
