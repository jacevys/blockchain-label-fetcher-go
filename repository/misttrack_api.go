package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type MisttrackApi struct {
	ApiKey       string
	DefaultRsult []map[string]interface{}
	Mapping      map[string]string
	Client       *http.Client
}

func NewMisttrackApi(apiKey string) *MisttrackApi {
	return &MisttrackApi{
		ApiKey: apiKey,
		DefaultRsult: []map[string]interface{}{
			{"type": "Unknown", "name": "Unknown", "labels": []interface{}{}},
		},
		Mapping: map[string]string{
			"bitcoin":   "BTC",
			"ethereum":  "ETH",
			"tron":      "TRX",
			"omni":      "BTC",
			"avalanche": "AVAX-Avalanche",
			"bunance":   "BNB",
			"polygon":   "POL-Polygon",
		},
		Client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (m *MisttrackApi) GetLabel(chainName string, address string) (map[string]interface{}, error) {
	url := fmt.Sprintf(
		"https://openapi.misttrack.io/v1/address_labels?coin=%s&address=%s&api_key=%s",
		m.Mapping[chainName],
		url.QueryEscape(address),
		m.ApiKey,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := m.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-OK status: %d, body: %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	fmt.Printf("API Response: %+v\n", result)

	if success, ok := result["success"].(bool); ok && success {
		data, ok := result["data"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid data format in API response")
		}

		labelType := data["label_type"].(string)
		labelList := data["label_list"].([]interface{})

		label := map[string]interface{}{
			"type":   labelType,
			"name":   labelList[0],
			"labels": labelList,
		}
		return label, nil
	}

	return nil, fmt.Errorf("API error: %v", result["msg"])
}

func contains(slice []interface{}, item string) bool {
	for _, v := range slice {
		if str, ok := v.(string); ok && str == item {
			return true
		}
	}
	return false
}
