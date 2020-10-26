package main

import (
	"main/context"
	"main/tool"
)

func main()  {
	app := context.New()
	defer app.Recover()
	defer app.Close()

	//支付连接生成工具
	tool.NewUrl(app)
}