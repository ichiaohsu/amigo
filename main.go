package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mirror-media/amigo/config"
	"github.com/mirror-media/amigo/pkg/member"
)

type RouterHandler interface {
	SetRoutes(router *gin.Engine)
}

func setRoutes(router *gin.Engine) {
	for _, h := range []RouterHandler{
		&member.Router,
	} {
		h.SetRoutes(router)
	}
}

func main() {

	var configFile string
	flag.StringVar(&configFile, "path", "", "Configuration file path.")
	flag.Parse()

	conf, err := config.LoadConfig(configFile)
	if err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}
	fmt.Println(conf.MySQL)

	// router := gin.New()
	// router.Use(gin.Recovery())

	// setRoutes(router)

}
