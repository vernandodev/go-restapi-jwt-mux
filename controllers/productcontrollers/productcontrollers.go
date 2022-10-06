package productcontrollers

import (
	"net/http"

	"github.com/vernandodev/go-restapi-jwt-mux/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// membuat data statis berbentuk slice dan akan dikirimkan dalam bentuk json
	data := []map[string]interface{}{
		{
			"id":           1,
			"nama_product": "kemeja",
			"stok":         1000,
		},
		{
			"id":           2,
			"nama_product": "celana",
			"stok":         1000,
		},
		{
			"id":           3,
			"nama_product": "jaket",
			"stok":         200,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}
