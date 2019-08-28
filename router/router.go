package router

import (
    "doubleJWT/handle/base"
    "doubleJWT/handle/service"
    "doubleJWT/handle/user"
    "doubleJWT/router/middleware"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Router(g *gin.Engine) {
    g.Use(gin.Recovery())

    g.NoRoute(func(c *gin.Context) {
        c.String(http.StatusNotFound, "The incorrect API route.")
    })

    u := g.Group("/v1/user")
    u.Use()
    {
        u.POST("/register", user.Register)
        u.POST("/login", user.Login)
    }

    b := g.Group("/v1/base")
    b.Use(middleware.AuthMiddleware())
    {
        b.POST("/refreshtoken", base.RefreshToken)
    }

    s := g.Group("/v1/service")
    s.Use(middleware.AuthMiddleware())
    {
        s.POST("/comment", service.Comment)
    }
}
