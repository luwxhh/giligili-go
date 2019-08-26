package server

import (
	"giligili/api"
	"giligili/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors()) //跨域
	// 获取用户身份
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		//视频详情
		v1.GET("video/:id", api.ShowVideo)
		//查询视频列表
		v1.GET("videos", api.ListVideo)
		// 排行榜
		v1.GET("rank/daily", api.DailyRank)
		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		v1.Use(middleware.AuthRequired())
		{
			// User Routing
			v1.GET("user/me", api.UserMe)
			v1.DELETE("user/logout", api.UserLogout)
			//用户投稿
			v1.POST("videos", api.CreateVideo)
			//更新视频
			v1.PUT("video/:id", api.UpdateVideo)
			//删除食品
			v1.DELETE("video/:id", api.DeleteVideo)
			//头像上传
			v1.POST("upload/token", api.UploadToken)
			//视频上传
			v1.POST("upload/tack", api.UploadTack)
		}

		//

	}
	return r
}