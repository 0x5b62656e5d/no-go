package routes

import (
	"encoding/json"
	"main/util"
	"main/util/types"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	util.SetHeaders(w, r)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(types.StandardResponse[string]{
		Data:    time.Now().UTC().Format(time.RFC3339Nano),
		Message: util.StringPtr("Healthy"),
		Error:   nil,
		Success: true,
	})
}
