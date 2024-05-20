package main

import (
	"github.com/changwei4869/wedding/modules"
	"github.com/changwei4869/wedding/modules/db"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   zang
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8088
// @BasePath  /api

// SwaggerUI: http://localhost:8088/swagger/index.html
func main() {
	// 引用数据库
	db.InitDb()
	// 引入路由组件
	modules.InitRouter()

}
