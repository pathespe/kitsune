package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pathespe/kitsune/pkg/handlers"
	"github.com/piquette/finance-go/equity"
	"github.com/piquette/finance-go/etf"
	"github.com/piquette/finance-go/quote"
)

func GetStockQuotes(symbols []string) {
	for _, symbol := range symbols {

		q, err := quote.Get(symbol)
		if err != nil {
			panic(err)
		}

		eq, err := equity.Get(symbol)
		if err != nil {
			panic(err)
		}

		et, err := etf.Get(symbol)
		if err != nil {
			panic(err)
		}

		//  YTDReturn                    float64
		//  TrailingThreeMonthReturns    float64
		//  TrailingThreeMonthNavReturns float64

		fmt.Println(q.Ask, eq.TrailingAnnualDividendYield, et.YTDReturn)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", handlers.Health).Methods(http.MethodGet)
	router.HandleFunc("/stonks", handlers.Stonks).Methods(http.MethodGet)
	http.ListenAndServe(":3000", router)
}
