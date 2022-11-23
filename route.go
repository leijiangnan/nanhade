package main

import (
	"github.com/leijiangnan/nanhade/framework"
	"github.com/leijiangnan/nanhade/framework/middleware"
)

func registerRouter(core *framework.Core) {
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 注意路径前带 /
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test3())
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", middleware.Test2(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}
