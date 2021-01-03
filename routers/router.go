package routers

import (
	"ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"ginblog/middleware"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())

	// router := r.Group("api/v1")
	{
		/* // User模块的路由接口
				// 添加用户
				router.POST("user/add", v1.AddUser)
				// 修改用户信息
				router.PUT("user/:id", v1.EditUser)
				// 查询单个用户
				router.GET("user/:id", v1.GetUserInfo)
				// 删除用户
				router.DELETE("user/:id", v1.DeleteUser)

				// Category模块的路由接口
				// 添加分类
				router.POST("category/add", v1.AddCate)
				// 获取分类列表
				router.GET("category", v1.GetCate)
				// 根据分类id获取某分类
				router.GET("category/:id", v1.GetCateInfo)
				// 修改分类信息
				router.PUT("category/:id", v1.EditCate)
				// 删除分类
				router.DELETE("category/:id", v1.DeleteCate)

				// Article模块的路由接口
		       // 添加文章
		       router.POST("article/add", v1.AddArticle)
		       // 修改文章
		       router.PUT("article/:id", v1.EditArticle)
		       // 删除文章
		       router.DELETE("article/:id", v1.DeleteArticle)

		       // 获取栏目文章列表
		       router.GET("artcate/:id", v1.GetArticle)
		       // 根据文章id获取某文章
			   router.GET("article/:id", v1.GetArticleInfo) */
			//   使用jwt认证
		// 需要使用权限认证的路由
		// 底下14个对应上面14个API
		auth := r.Group("api/v1")
		auth.Use(middleware.JwtToken()) // 使用JwtToken中间件
		{
			
			auth.POST("upload", v1.UpLoad)

			auth.PUT("user/:id", v1.EditUser)
			auth.DELETE("user/:id", v1.DeleteUser)

			auth.POST("category/add", v1.AddCategory)
			auth.PUT("category/:id", v1.EditCate)
			auth.DELETE("category/:id", v1.DeleteCate)

			auth.POST("article/add", v1.AddArticle)
			auth.PUT("article/:id", v1.EditArticle)
			auth.DELETE("article/:id", v1.DeleteArticle)
		}

		// 不需要使用权限认证的路由
		router := r.Group("api/v1")
		{
			router.POST("user/add", v1.AddUser)
			router.GET("user/:id", v1.GetUserInfo)
			router.GET("category", v1.GetCate)
			router.GET("category/:id", v1.GetCateInfo)
			router.GET("article", v1.GetArticle)
			router.GET("article/info/:id", v1.GetArticleInfo)
			// 登陆接口
			router.POST("login", v1.Login)
		}


	}

	_=r.Run(utils.HttpPort)
}
