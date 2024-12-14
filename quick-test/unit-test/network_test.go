package main

import (
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!"))
}

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Failed", err)
	}
}

// 需要测试某个 API 接口的 handler 能够正常工作，例如 helloHandler
func TestCon(t *testing.T) {
	// 尽量不对 http 和 net 库使用 mock，这样可以覆盖较为真实的场景

	ln, err := net.Listen("tcp", "127.0.0.1:0") // 监听一个未被占用的端口，并返回 Listener
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler) // 注册一个请求处理器
	go http.Serve(ln, nil) // 启动 http 服务

	resp, err := http.Get("http://" + ln.Addr().String() + "/hello") // 发起一个 Get 请求，检查返回值是否正确
	handleError(t, err)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "Hello World!" {
		t.Fatal("expected", "Hello World!", "but got", string(body))
	}
}

// 针对 http 开发的场景，使用标准库 net/http/httptest 进行测试更为高效
func TestHttp(t *testing.T) {
	request := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	helloHandler(w, request) // 使用 httptest 模拟请求对象(req)和响应对象(w)，达到了相同的目的
	bytes, _ := io.ReadAll(w.Result().Body)
	if string(bytes) != "Hello World!" {
		t.Fatal("expected", "Hello World!", "but got", string(bytes))
	}
}