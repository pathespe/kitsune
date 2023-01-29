package stonks

import (
	"fmt"

	"github.com/pathespe/kitsune/pkg/models"
	"github.com/piquette/finance-go/equity"
	"github.com/piquette/finance-go/etf"
	"github.com/piquette/finance-go/quote"
)

func GetStockQuotes(symbols []string) ([]models.Quote, error) {
	var output = []models.Quote{}
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

    var newq = models.Quote{}
    newq.Ask = q.Ask
    newq.Name = symbol
    newq.DividendYield = eq.TrailingAnnualDividendYield
    newq.YTDReturn = et.YTDReturn

    output = append(output, newq)
	}
	return output, nil
}
