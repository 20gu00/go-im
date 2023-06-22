package main

import (
	"io"
	"net/http"
)

func main() {
	//将请求路径和逻辑函数绑定
	http.HandleFunc("/user/login",
		func(writer http.ResponseWriter, req *http.Request) {
			//获取参数,这里是表单form
			req.ParseForm()
			//返回数据,默认是text/html,想返回json就要设置header  如果是Gin框架可以直接c.JSON
			writer.Header().Set("Content-type", "application/json")
			//写回header头响应
			writer.WriteHeader(http.StatusOK)
			//返回的json字符串
			//json.Marshal(struct)将结构体转换成json字符串
			//str:=`{"code":"1","msg":"2"}`
			//omitempty设置jso标签,可以忽略空,也就是null的数据不会出错
			//writer.Write([]byte(str))
			io.WriteString(writer, "你好")
		})

	//默认80端口
	//curl http://localhost:8080/user/login

	//启动服务器 http
	http.ListenAndServe(":8080", nil)
}
