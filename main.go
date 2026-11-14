package main

import (
	_ "msys_payment_app_gateway/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"

	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {
	sqlConn, err := beego.AppConfig.String("sqlconn")
	if err != nil {
		logs.Error("%s", err)
	}
	// sqlConn2, err2 := beego.AppConfig.String("sqlconn2")
	// if err2 != nil {
	// 	logs.Error("%s", err)
	// }
	orm.RegisterDataBase("default", "mysql", sqlConn)
	// orm.RegisterDataBase("db2", "mysql", sqlConn2)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	logs.SetLogger(logs.AdapterFile, `{"filename":"../logs/msys-app-api-gateway.log"}`)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"https://schools.msysconsultgh.com", "https://app.msysconsultgh.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With", "X-Application-Code"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.Run()
}
