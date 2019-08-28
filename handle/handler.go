package handle

import (
    "doubleJWT/pkg/errno"
    "github.com/gin-gonic/gin"
    "net/http"
)

type Result struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

func SendResult(c *gin.Context, err error, data interface{}) {
    code, message := errno.DecodeErr(err)

    // always return http.StatusOK
    c.JSON(http.StatusOK, Result{
        Code:    code,
        Message: message,
        Data:    data,
    })
}

