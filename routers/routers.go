package routers

import (
	"gin-demo/controllers"
	"gin-demo/middlewares/jwt"
	"gin-demo/pkg/export"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	_ "gin-demo/docs"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", controllers.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.POST("/upload", controllers.UploadImage)

	c := r.Group("/api/v1")
	c.Use(jwt.JWT())
	{
		//获取标签列表
		c.GET("/tags", controllers.GetTags)
		//新建标签
		c.POST("/tags", controllers.AddTag)
		//更新指定标签
		c.PUT("/tags/:id", controllers.EditTag)
		//删除指定标签
		c.DELETE("/tags/:id", controllers.DeleteTag)
		//导出标签
		r.POST("/tags/export", controllers.ExportTag)
		//导入标签
		r.POST("/tags/import", controllers.ImportTag)

		//获取文章列表
		c.GET("/articles", controllers.GetArticles)
		//获取指定文章
		c.GET("/articles/:id", controllers.GetArticle)
		//新建文章
		c.POST("/articles", controllers.AddArticle)
		//更新指定文章
		c.PUT("/articles/:id", controllers.EditArticle)
		//删除指定文章
		c.DELETE("/articles/:id", controllers.DeleteArticle)
		//生成文章海报
		c.POST("/articles/poster/generate", controllers.GenerateArticlePoster)
	}

	return r
}
