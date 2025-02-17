package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func defaultResult() map[string]interface{} {
	return map[string]interface{}{
		"status":  true,
		"message": "Request processing time out",
		"data": map[string]interface{}{
			"blockchain_security": []map[string]interface{}{
				{"type": "Unknown", "name": "Unknown", "labels": []string{}},
			},
			"chainanalysis": []map[string]interface{}{
				{"type": "Unknown", "name": "Unknown", "labels": []string{}},
			},
			"qlue": []map[string]interface{}{
				{"type": "Unknown", "name": "Unknown", "labels": []string{}},
			},
			"smart_contract": false,
			"black_list":     false,
			"timeout":        true,
		},
	}
}

func GetLabel(c *gin.Context) {
	// chainName := c.Query("chain_name")
	// address := c.Query("address")
	// sourceListCode := c.Query("source_list_code")
	// searchFlag := c.Query("search_flag")
	// quickMode := c.Query("quick_mode")

	// c.JSON(http.StatusOK, defaultResult())
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "test",
	})
}

// ok
