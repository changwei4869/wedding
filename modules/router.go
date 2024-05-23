package modules

import (
	_ "github.com/changwei4869/wedding/docs"
	"github.com/changwei4869/wedding/middleware"
	"github.com/changwei4869/wedding/modules/admin"
	"github.com/changwei4869/wedding/modules/file"
	"github.com/changwei4869/wedding/modules/health"
	"github.com/changwei4869/wedding/modules/permission"
	"github.com/changwei4869/wedding/modules/role"
	"github.com/changwei4869/wedding/modules/site"
	"github.com/changwei4869/wedding/modules/tag"
	"github.com/changwei4869/wedding/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api", middleware.Cors())
	{
		// health
		api.POST("/health", health.HealthCheck)
		// tag
		api.GET("/tag/:id", tag.GetTagById)
		api.POST("/tag", tag.AddTag)
		api.DELETE("/tag/:id", tag.DeleteTag)
		api.PUT("/tag", tag.EditTag)
		// file
		api.POST("/file/upload", file.UploadFile)
		//role
		api.GET("/role", role.ListRole)
		api.POST("/role", role.AddRole)
		api.DELETE("/role/:id", role.DeleteRole)
		api.PUT("/role", role.EditRole)
		// admin
		api.GET("/admin", admin.ListAdmin)
		api.POST("/admin", admin.AddAdmin)
		api.DELETE("/admin/:id", admin.DeleteAdmin)
		api.PUT("/admin", admin.EditAdmin)
		//permission
		api.GET("/permissions", permission.ListPermission)
		// site
		api.GET("/site/:id", site.GetSiteById)
		api.POST("/site", site.AddSite)
		api.DELETE("/site/:id", site.DeleteSite)
		api.PUT("/site", site.EditSite)
	}

	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	_ = r.Run(utils.HttpPort)

}
