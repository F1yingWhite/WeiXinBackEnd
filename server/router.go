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
			// get /api/search?company=xxx&city=xxx | 用于检索信息
			trueApi.GET("search", service.HandlerBindQuery(&service.Search{}))
			user := trueApi.Group("user")
			{
				// get /api/user/get?user_id=xxx | 用于获取用户信息
				user.GET("get", service.HandlerNoBind(&service.UserInfo{}))
				// post /api/user/update | 用于更新用户信息
				user.POST("update", service.HandlerBind(&service.UpdateUserInfo{}))
			}
		}
	}
	return r
}
