package main

import (
    "doubleJWT/config"
    "doubleJWT/models"
    "doubleJWT/router"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    g := gin.New()
    gin.SetMode(viper.GetString("runmode"))

    serverConf := config.NewServerConf()

    router.Router(g)

    models.InitMysql(serverConf.MysqlConf)
    defer models.Close()

    go func() {
        logrus.Infof("server start addr:%s  name:%s", serverConf.HTTPConf.Addr, serverConf.HTTPConf.Name)
        if err := http.ListenAndServe(serverConf.HTTPConf.Addr, g); err != nil {
            logrus.Error("server start error : ", err)
            panic(err)
        }
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
    for {
        s := <-c
        logrus.Infof("server get a signal %s", s.String())
        switch s {
        case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
            logrus.Info("server exit")
            return
        case syscall.SIGHUP:
        default:
            return
        }
    }
}
