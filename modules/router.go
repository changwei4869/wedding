package modules

import (
	_ "github.com/changwei4869/wedding/docs"
	"github.com/changwei4869/wedding/middleware"
	"github.com/changwei4869/wedding/modules/admin"
	"github.com/changwei4869/wedding/modules/file"
	"github.com/changwei4869/wedding/modules/health"
	"github.com/changwei4869/wedding/modules/permission"
	qrcode "github.com/changwei4869/wedding/modules/qr_code"
	"github.com/changwei4869/wedding/modules/registration"
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
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// health
	r.GET("/api/health", health.HealthCheck)
	r.POST("/api/admin/login", middleware.Cors(), admin.AdminLogin)

	api := r.Group("/api", middleware.Logger(), middleware.JwtToken(), middleware.Cors())
	{
		// tag
		api.GET("/tag/:id", tag.GetTagById)
		api.GET("/tag", tag.GetAllTags)
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
		api.GET("/site", site.GetAllSites)
		//registration
		api.GET("/registration", registration.GetRegistration)
		api.GET("/registrations", registration.GetAllRegistrations)
		api.DELETE("/registration/:id", registration.DeleteRegistration)
		api.PUT("/registration", registration.EditRegistration)
		api.POST("/registration/checkregistration", registration.CheckRegistration)
		api.POST("/registration", registration.InitRegistration)
		api.DELETE("/registration", registration.DelBatchRegistrations)
		// qrcode
		api.GET("/qrcodes", qrcode.GetAllQrCode)
		api.POST("/qrcode", qrcode.AddQrCode)
		api.PUT("/qrcode", qrcode.EditQrCode)
	}
	_ = r.Run(utils.HttpPort)

}
