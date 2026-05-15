package profiler

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	group := r.Group("/debug/pprof")
	{
		group.GET("/", gin.WrapF(pprof.Index))
		group.GET("/heap", gin.WrapF(pprof.Handler("heap").ServeHTTP))
		group.GET("/goroutine", gin.WrapF(pprof.Handler("goroutine").ServeHTTP))
		group.GET("/profile", gin.WrapF(pprof.Profile))
	}
}