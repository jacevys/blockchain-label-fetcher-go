package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestGetLabel(t *testing.T) {
	// 創建一個新的Gin引擎
	r := gin.Default()

	// 註冊路由
	r.GET("/v1/get_label", GetLabel)

	// 創建一個測試請求
	req, err := http.NewRequest("GET", "/v1/get_label?chain_name=ethereum&address=0x1234&source_list_code=code123&search_flag=true&quick_mode=false", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 創建一個HTTPRecorder來捕捉響應
	w := httptest.NewRecorder()

	// 執行請求
	r.ServeHTTP(w, req)

	// 斷言響應狀態碼是200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// 斷言響應內容包含 "status": true
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(response)
	// 驗證返回的結果
	assert.Equal(t, true, response["status"])
}
