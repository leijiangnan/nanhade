package main

import "github.com/leijiangnan/nanhade/framework"

func registerRouter(core *framework.Core) {
	core.Get("/user/login", UserLoginController)

	// 注意路径前带 /
	subjectApi := core.Group("/subject")
	{
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}
