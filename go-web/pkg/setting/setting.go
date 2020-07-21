package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var (
	RunMode      string
	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	PageSize     int
	JwtSecret    string
)

func init()  {
	viper.SetConfigName("config")
	viper.AddConfigPath("conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	LoadApp()
	LoadBase()
	LoadServer()
}

func LoadBase()  {
	RunMode = viper.GetString("runMode")
}
func LoadServer()  {
	HttpPort = viper.GetInt("server.port")
	ReadTimeOut = viper.GetDuration("readTimeOut")*time.Second
	WriteTimeOut = viper.GetDuration("writeTimeOut")*time.Second
}

func LoadApp()  {
	JwtSecret = viper.GetString("secret")
	PageSize = viper.GetInt("pageSize")
}