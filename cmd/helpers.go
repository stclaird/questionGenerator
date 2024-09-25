package main

import (
	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
)

type Config struct{
	port string
}

func GetConfig() Config {

	viper.SetConfigFile(".env")
    viper.ReadInConfig()

    port := viper.Get("PORT")
    var portString string

    if port == nil {
        portString = ":5001"
    } else {
        portString = port.(string)
    }

	return Config{
		port: portString,
	}
}

func CORSConfig() cors.Config {
    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{"http://localhost:5001"}
    corsConfig.AllowCredentials = true
    corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
    corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
    return corsConfig
}