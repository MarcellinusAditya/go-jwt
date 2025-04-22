package productcontroller

import (
	"net/http"

	"github.com/MarcellinusAditya/go-jwt/helper"
)

func Index(w http.ResponseWriter, r* http.Request){
	data :=[]map[string]interface{}{
		{
			"id": 1,
			"nama_product": "kemeja",
			"stok": 1000,
		},
		{
			"id": 2,
			"nama_product": "Celana",
			"stok": 1000,
		},
		{
			"id": 3,
			"nama_product": "Dasi",
			"stok": 500,
		},
	}

	helper.ResponseJson(w, http.StatusOK, data)
}