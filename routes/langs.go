package routes

import (
	"encoding/json"
	"main/util"
	"main/util/types"
	"net/http"
)

func Langs(w http.ResponseWriter, r *http.Request, app *util.App) {
	util.SetHeaders(w, r)

	keys := make([]string, 0, len(app.Reasons))

	for k := range app.Reasons {
		keys = append(keys, k)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(types.StandardResponse[[]string]{
		Data:    keys,
		Message: nil,
		Error:   nil,
		Success: true,
	})
}
