package main

import (
	_ "auto-work-helper/boot"
	_ "auto-work-helper/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
