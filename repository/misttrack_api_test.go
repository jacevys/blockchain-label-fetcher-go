package repository

import (
	"fmt"
	"testing"
)

func TestGetLabelMist(t *testing.T) {
	// // Mock API Server
	// server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// 檢查 HTTP 方法
	// 	assert.Equal(t, "GET", r.Method)

	// 	// 讀取和檢查請求的 Body
	// 	body, err := io.ReadAll(r.Body)
	// 	assert.NoError(t, err)
	// 	defer r.Body.Close()

	// 	var payload map[string]interface{}
	// 	err = json.Unmarshal(body, &payload)
	// 	assert.NoError(t, err)

	// 	// 驗證 payload 中的參數
	// 	assert.Equal(t, "BTC", payload["coin"])
	// 	assert.Equal(t, "dummy_address", payload["address"])
	// 	assert.Equal(t, "test_api_key", payload["api_key"])

	// 	// 模擬服務器返回成功響應
	// 	w.Header().Set("Content-Type", "application/json")
	// 	response := map[string]interface{}{
	// 		"success": true,
	// 		"data": map[string]interface{}{
	// 			"label_type": "exchange",
	// 			"label_list": []interface{}{"deposit", "withdraw"},
	// 		},
	// 	}
	// 	json.NewEncoder(w).Encode(response)
	// }))
	// defer server.Close()

	// 初始化 MisttrackApi
	misttrack := NewMisttrackApi("")
	chainName := "bitcoin"
	address := "bc1qgdjqv0av3q56jvd82tkdjpy7gdp9ut8tlqmgrpmv24sq90ecnvqqjwvw97"

	// 測試數據
	label, err := misttrack.GetLabel(chainName, address)
	// 驗證結果
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}
	fmt.Println(label)
}
