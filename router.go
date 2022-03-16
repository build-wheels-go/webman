package main

import (
	"webman/framework/gin"
	"webman/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	//core.Get("foo", FooControllerHandler)
	// 静态路由+HTTP方法匹配
	core.GET("/user/login", middleware.Cost(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Cost(), middleware.Test2())
		// 动态路由
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Use(middleware.Test2())
			subjectInnerApi.GET("/name", SubjectNameController)
		}

	}

}
