package routers

import (
	"gin-demo/controllers"
	"gin-demo/middlewares/jwt"
	"gin-demo/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", controllers.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", controllers.GetTags)
		//新建标签
		apiv1.POST("/tags", controllers.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", controllers.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", controllers.DeleteTag)

		article := apiv1.Group("/articles")
		article.Use(jwt.JWT())
		{
			//获取文章列表
			article.GET("", controllers.GetArticles)
			//获取指定文章
			article.GET("/:id", controllers.GetArticle)
			//新建文章
			article.POST("", controllers.AddArticle)
			//更新指定文章
			article.PUT("/:id", controllers.EditArticle)
			//删除指定文章
			article.DELETE("/:id", controllers.DeleteArticle)
		}
	}

	return r
}
