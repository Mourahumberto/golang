package routes

import (
	"projeto-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.ExibeUmAluno)
	r.DELETE("/alunos/:id", controllers.DeletaUmAluno)
	r.PATCH("/alunos/:id", controllers.EditaUmAluno)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.Run()
}
