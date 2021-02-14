package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
)

func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	gin.Recovery()

	// Set monitor
	getMonitor().Use(r)

	// Set pprof
	pprof.Register(r, "/debug/pprof")

	// Set zip
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	return r
}

func getMonitor() *ginmetrics.Monitor {
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/debug/metrics")
	m.SetSlowTime(5)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	return m
}
