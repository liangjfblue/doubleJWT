package config

import (
    "github.com/fsnotify/fsnotify"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
    "strings"
)

// ServerConfig server config
type ServerConfig struct {
    HTTPConf	*HTTPConfig
    MysqlConf	*MysqlConfig
    RedisConf	*RedisConfig
}

// HTTPConfig http config
type HTTPConfig struct {
    RunMode		string
    Addr		string
    Name		string
}

// MysqlConfig mysql config
type MysqlConfig struct {
    Addr 		string
    Db			string
    User 		string
    Password	string
}

// RedisConfig redis config
type RedisConfig struct {
    Host 			string
    Port			string
    ClusterHost 	string
    IsCluster		bool
}

// NewServerConf 实例化server config
func NewServerConf() *ServerConfig {
    return &ServerConfig{
        HTTPConf: &HTTPConfig{
            RunMode:	viper.GetString("runmode"),
            Addr:		viper.GetString("addr"),
            Name:		viper.GetString("name"),
        },
        MysqlConf: &MysqlConfig{
            Addr:		viper.GetString("mysql.addr"),
            Db:			viper.GetString("mysql.db"),
            User:		viper.GetString("mysql.user"),
            Password:	viper.GetString("mysql.password"),
        },
        RedisConf: &RedisConfig{
            Host:			viper.GetString("redis.host"),
            Port:			viper.GetString("redis.port"),
        },
    }
}

// init 初始化config
func init() {
    if err := initConfig(); err != nil {
        panic(err)
    }
    initLog()
    watchConfig()
}

// initConfig 初始化config
func initConfig() error {
    viper.AddConfigPath(".")
    viper.SetConfigName("config")

    viper.SetConfigType("yaml")
    viper.AutomaticEnv()
    viper.SetEnvPrefix("DOUBLEJWT")
    replacer := strings.NewReplacer(".", "_")
    viper.SetEnvKeyReplacer(replacer)
    if err := viper.ReadInConfig(); err != nil {
        return err
    }

    return nil
}

// watchConfig 监听config文件
func watchConfig() {
    viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
        logrus.Info("Config file changed: %s", e.Name)
    })
}

// initLog 初始化log
func initLog() {
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetOutput(os.Stdout)
    logrus.SetLevel(logrus.Level(viper.GetInt("log.level")))
    logrus.SetReportCaller(viper.GetBool("log.reportCaller"))
}
