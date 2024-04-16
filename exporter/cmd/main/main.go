package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

// 定义指标
var (
	opsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "my_service_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func init() {
	// 在Prometheus中注册指标
	prometheus.MustRegister(opsProcessed)
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc() // 模拟指标的变化
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	// 记录指标
	recordMetrics()

	// 创建新的HTTP服务器
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil) // 在9090端口上暴露/metrics端点
}
