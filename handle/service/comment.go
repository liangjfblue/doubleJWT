package service

import (
    "doubleJWT/handle"
    "doubleJWT/models"
    "doubleJWT/pkg/errno"
    "github.com/gin-gonic/gin"
)

func Comment(c *gin.Context) {
    var (
        err error
        req CommentReq
    )
    if err = c.BindJSON(&req); err != nil {
        handle.SendResult(c, errno.ErrBind, nil)
        return
    }

    if _, err = models.GetTBUser(map[string]interface{}{"id": req.FromUid}); err != nil {
        handle.SendResult(c, errno.ErrUserNotFound, nil)
        return
    }

    comment := models.TBComment{
        TopicId:    req.TopicId,
        TopicType:  req.TopicType,
        Content:    req.Content,
        FromUid:    req.FromUid,
    }

    if err = comment.Validate(); err != nil {
        handle.SendResult(c, errno.ErrValidation, nil)
        return
    }

    if err = comment.AddTBComment(); err != nil {
        handle.SendResult(c, errno.ErrDatabase, nil)
        return
    }

    handle.SendResult(c, nil, nil)
}
