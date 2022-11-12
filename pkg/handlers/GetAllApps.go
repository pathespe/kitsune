package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pathespe/kitsune/pkg/mocks"
)

func GetAllApps(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Apps)
}
