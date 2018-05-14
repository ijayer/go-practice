package apitest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	// 创建一个 HTTP Request, 发送请求
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Errorf("request failed: %v\n", err)
	}

	// 创建一个 HTTP ResponseWriter, 记录服务器响应的结果
	resp := httptest.NewRecorder()

	// 测试 Handler
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(resp, req)

	// 验证响应结果
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handlers returnned wrong status code: got %v, want %v", status, http.StatusOK)
	}
	expected := `{"alive":"true"}`
	if resp.Body.String() != expected {
		t.Errorf("handlers returned unexpected body: got %v, want %v", resp.Body.String(), expected)
	}
}

func TestHealthCheckHandler2(t *testing.T) {
	reqData := struct {
		Info string `json:"info"`
	}{Info: "yours"}
	reqBody, _ := json.Marshal(reqData)
	t.Log("input:", string(reqBody))

	req := httptest.NewRequest(
		http.MethodPost,
		"/health-check",
		bytes.NewReader(reqBody),
	)
	req.Header.Set("user_id", "abc123")

	resp := httptest.NewRecorder()
	HealthCheckHandler(resp, req)

	result := resp.Result()
	body, _ := ioutil.ReadAll(result.Body)
	t.Log(string(body))

	if result.StatusCode != http.StatusOK {
		t.Errorf("handlers returned unexpected body: got %v, want %v",
			resp.Body.String(), result.StatusCode)
	}
}
