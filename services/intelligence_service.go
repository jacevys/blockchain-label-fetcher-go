package services

import (
	"addressintelligence/repository"
	"context"
)

type IntelligenceService struct {
	ChainList     []string
	DefaultResult map[string]interface{}
	MistApi       *repository.MisttrackApi
}

type GetLabelCib struct{}

func NewIntelligenceService() *IntelligenceService {
	return &IntelligenceService{
		ChainList: []string{"bitcoin", "ethereum", "tron", "omni", "avalanche", "binance", "polygon"},
		DefaultResult: map[string]interface{}{
			"status":  true,
			"message": "chain name not supported",
			"data": map[string]interface{}{
				"blockchain_security": []map[string]interface{}{
					{"type": "Unknown", "name": "Unknown", "labels": []interface{}{}},
				},
				"chainanalysis": []map[string]interface{}{
					{"type": "Unknown", "name": "Unknown"},
				},
				"qlue": []map[string]interface{}{
					{"type": "Unknown", "name": "Unknown"},
				},
				"smart_contract": false,
				"black_list":     false,
			},
		},
	}
}

func (service *IntelligenceService) ProcessRequest(
	ctx context.Context,
	chainName string,
	address string,
	sourceListCode string,
	searchFlag bool,
	quickMode bool,
	affiliation string,
) (map[string]interface{}, error) {
	if !contains(service.ChainList, chainName) {
		return service.DefaultResult, nil
	}

	labelDict := service.MistApi
	return labelDict, nil
}

func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
