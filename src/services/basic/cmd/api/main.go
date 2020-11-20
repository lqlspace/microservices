package main

import (
	"flag"
	"fmt"

	"github.com/lqlspace/microservices/src/core/conf"
	"github.com/lqlspace/microservices/src/engine"
	"github.com/lqlspace/microservices/src/services/basic/cmd/api/config"
	"github.com/lqlspace/microservices/src/services/basic/cmd/api/svc"
	"github.com/lqlspace/microservices/src/services/basic/handler"
)

var configFile = flag.String("f", "etc/basic-api.json", " the config file")
func main() {
	flag.Parse()

	var c config.Config

	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(&c)

	e, _ := engine.NewEngine(c)
	handler.RegisterHandlers(e, ctx)

	fmt.Printf("basi-api server startup at %s:%d\n", c.Host, c.Port)
	e.Start()
}


