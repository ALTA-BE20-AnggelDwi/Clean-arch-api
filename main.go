package main

import (
	"clean-arch/app/configs"
	"clean-arch/app/databases"
	"clean-arch/app/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	dbSql := databases.InitDBMysql(cfg)

	e := echo.New()
	routers.InitRouter(dbSql, e)
	//start server and port
	e.Logger.Fatal(e.Start(":8080"))
}
