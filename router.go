package main

import (
	"webman/framework"
	"webman/framework/middleware"
)

func registerRouter(core *framework.Core) {
	//core.Get("foo", FooControllerHandler)
	// 静态路由+HTTP方法匹配
	core.Get("/user/login", middleware.Recovery(), middleware.Cost(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Cost(), middleware.Test2())
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Use(middleware.Test2())
			subjectInnerApi.Get("/name", SubjectNameController)
		}

	}

}
