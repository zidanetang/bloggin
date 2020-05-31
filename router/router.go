package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zidanetang/devopsgin/handler"
	"github.com/zidanetang/devopsgin/handler/system"
)

func InitRouter(content *gin.Context) {

	r := gin.New()

	//加载中间建

	r.GET("", handler.Index)
	r.Static("/static", "./static")
	r.GET("/healthcheck", handler.Healthcheck)

	systemMessage := r.Group("/system")
	{
		systemMessage.GET("/info", sys.Info)
	}

	r.POST("/login")
	r.GET("/refresh_token")
	r.GET("/logout")

	apiv1 := r.Group("/api/v1")
	{
		user := apiv1.Group("/usrs")
		{
			user.GET("/current")
		}

		group := apiv1.Group("/groups")
		{
			group.GET("/current")
		}
		role := apiv1.Group("/roles")
		{
			role.GET("/current")
		}
		shift := apiv1.Group("/shifts")
		{
			shift.GET("/binds")
		}
		oncall := apiv1.Group("/oncalls")
		{
			oncall.GET("/current")
		}

	}
}
