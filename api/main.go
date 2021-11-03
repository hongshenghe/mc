package main

import (
	_ "api/boot"
	_ "api/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
