package goweb

import (
	"io"
	"net/http"
)

func main() {
	//将请求路径和逻辑函数绑定
	http.HandleFunc("/user/login",
		func(writer http.ResponseWriter, req *http.Request) {
			io.WriteString(writer, "你好")
		})

	//启动服务器 http
	http.ListenAndServe(":8080", nil)
}