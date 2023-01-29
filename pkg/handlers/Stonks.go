package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pathespe/kitsune/pkg/stonks"
)

func Stonks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	quotes, _ := stonks.GetStockQuotes([]string{"CLNE.AX", "FMG.AX", "STW.AX", "DHHF.AX", "FANG.AX", "HACK.AX", "HPI.AX", "IOO.AX", "IVV.AX", "MVR.AX", "TPG.AX", "TUA.AX", "VESG.AX", "VGAD.AX"})

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}
