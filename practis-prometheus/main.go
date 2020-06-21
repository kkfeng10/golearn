package main

// 尝试使用Prometheus接口

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// func main() {
// 	http.Handle("/metrics", promhttp.Handler())
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// 	ph := make(chan<- int) // 发送int类型的数据到channel
// 	ch := make(<-chan int) // 从channel中接收int类型的数据
// 	v := <-ch
// }

func main() {
	server := http.NewServeMux()
	server.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9001", server)
}
