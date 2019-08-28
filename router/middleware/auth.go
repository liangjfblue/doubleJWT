package middleware

import (
    "doubleJWT/handle"
    "doubleJWT/pkg/errno"
    "doubleJWT/pkg/token"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        var (
            err error
            t *token.Context
        )

        logrus.Info(c.Request.URL.Path)
        if c.Request.URL.Path != "/v1/user/register" && c.Request.URL.Path != "/v1/user/login" {
            if c.Request.URL.Path != "/v1/base/refreshtoken" {
                if t, err = token.ParseRequest(c, true); err != nil {
                    logrus.Error(err)
                    handle.SendResult(c, errno.ErrToken, nil)
                    c.Abort()
                    return
                }
            } else {
                if t, err = token.ParseRequest(c, false); err != nil {
                    logrus.Error(err)
                    handle.SendResult(c, errno.ErrToken, nil)
                    c.Abort()
                    return
                }
            }
        }

        c.Set("token", *t)
        c.Next()
    }
}
