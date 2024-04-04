package server

import (
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
		//默认接口，用于
		api.GET("ping", service.WX_authorization)
		api.POST("ping", service.Reply)
		trueApi := api.Group("")
		{
			user := trueApi.Group("user")
			{
				// get /api/user/get?user_id=xxx | 用于获取用户信息
				user.GET("get", service.HandlerNoBind(&service.UserInfo{}))
				// post /api/user/update | 用于更新用户信息
				user.POST("update", service.HandlerBind(&service.UpdateUserInfo{}))
			}

			salary := trueApi.Group("salary")
			{
				// get /api/get?pagesize=?&page=? | 用于检索信息
				salary.GET("get", service.HandlerBindQuery(&service.GetSalary{}))
				// get /api/getByCompany?pagesize=?&page=?&company=? | 用于检索信息
				salary.GET("getByCompany", service.HandlerBindQuery(&service.GetSalaryByCompany{}))
				// get /api/getByCity?pagesize=?&page=?&city=? | 用于检索信息
				salary.GET("getByCity", service.HandlerBindQuery(&service.GetSalaryByCity{}))
				// get /api/getByCompanyAndCity?pagesize=?&page=?&company=?&city=? | 用于检索信息
				salary.GET("getByCompanyAndCity", service.HandlerBindQuery(&service.GetSalariesByCompanyAndCity{}))
				// get /api/getByUserId?pagesize=?&page=?&user_id=? | 用于检索信息
				salary.GET("getByUserId", service.HandlerBindQuery(&service.GetSalaryByUserId{}))
				// get /api/getById?pagesize=?&page=?&id=? | 用于检索信息
				salary.GET("getById", service.HandlerBindQuery(&service.GetSalaryById{}))
				// post /api/salary/create | 用于创建信息
				salary.POST("create", service.HandlerBind(&service.CreateSalary{}))
			}
		}
	}
	return r
}
