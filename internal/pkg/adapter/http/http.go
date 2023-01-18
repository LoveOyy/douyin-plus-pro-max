package http

import "github.com/gin-gonic/gin"

var ginEngine *gin.Engine
var dyHTTPRouter *gin.RouterGroup

func Init() {
	ginEngine = gin.Default()
	dyHTTPRouter = ginEngine.Group("/douyin")
}
func Run() {
	ginEngine.Run(":8081")
}
