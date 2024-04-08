package server

import (
	"weixin_backend/server/middlewares"
	"weixin_backend/server/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	//一些基础配置
	config := cors.DefaultConfig()
	config.ExposeHeaders = []string{"Authorization"}
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	api := r.Group("api")
	api.Use(gin.Recovery())
	{
		//默认接口，用于与微信对接
		api.GET("ping", service.WX_authorization)
		api.POST("ping", service.Reply)
		trueApi := api.Group("")
		trueApi.Use(middlewares.IPAuthMiddleware())
		{
			user := trueApi.Group("user")
			{
				// get /api/user | 用于获取用户信息
				user.GET("", service.HandlerNoBind(&service.UserInfo{}))
				// post /api/user | 用于更新用户信息
				user.PUT("", service.HandlerBind(&service.UpdateUserInfo{}))
			}

			salary := trueApi.Group("salary")
			{
				// get /api/salary?page_size=?&page=?&city=?&company=? | 用于检索信息
				salary.GET("", service.HandlerBindQuery(&service.GetSalary{}))
				// get /api/salary/getByUserId?page_size=?&page=?&user_id=? | 用于检索信息
				salary.GET("getByUserId", service.HandlerBindQuery(&service.GetSalaryByUserId{}))
				// get /api/salary/getById?id=? | 用于检索信息
				salary.GET("getById", service.HandlerBindQuery(&service.GetSalaryById{}))
				// post /api/salary/create | 用于创建信息
				salary.POST("create", service.HandlerBind(&service.CreateSalary{}))
				// post /api/salary/creates | 用于创建信息
				salary.POST("creates", service.HandlerBind(&service.CreateSalaries{}))
				// PUT /api/salary | 修改信息
				salary.PUT("", service.HandlerBind(&service.UpdateSalary{}))
				// DELETE /api/salary?id=? | 删除信息
				salary.DELETE("", service.HandlerBindQuery(&service.DeleteSalary{}))
			}
		}
	}
	return r
}
