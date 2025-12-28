package routes

import (
	"encoding/json"
	"main/util"
	"main/util/types"
	"math/rand/v2"
	"net/http"
)

func No(w http.ResponseWriter, r *http.Request, app *util.App) {
	util.SetHeaders(w, r)

	query := r.URL.Query().Get("lang")

	if query == "" {
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(types.StandardResponse[types.NoResponse]{
			Data: types.NoResponse{
				Lang:   "en",
				Reason: app.Reasons["en"][rand.IntN(len(app.Reasons["en"]))],
			},
			Message: nil,
			Error:   nil,
			Success: true,
		})

		return
	}

	if value, ok := app.Reasons[query]; ok {
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(types.StandardResponse[types.NoResponse]{
			Data: types.NoResponse{
				Lang:   query,
				Reason: value[rand.IntN(len(value))],
			},
			Message: nil,
			Error:   nil,
			Success: true,
		})

		return
	} else {
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(types.StandardResponse[types.NoResponse]{
			Data: types.NoResponse{
				Lang:   "en",
				Reason: app.Reasons["en"][rand.IntN(len(app.Reasons["en"]))],
			},
			Message: util.StringPtr("Queried language not found, returning English"),
			Error:   nil,
			Success: true,
		})

		return
	}
}
