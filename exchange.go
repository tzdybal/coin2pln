package main

import (
	"sort"
)

type exchange struct {
	name   string
	url    string
	fees   []feeSchedule
	quotes []quote
}

func newExchange(name, url string, quotes []quote) *exchange {
	ex := &exchange{name: name, url: url, quotes: quotes}

	sort.Sort(sortableQuotes(ex.quotes))

	return ex
}

func (ex *exchange) sourceCurrencies() (currencies []string) {
	// quotes are sorted by currency pair
	last := ""
	for _, q := range ex.quotes {
		if q.pair.src != last {
			currencies = append(currencies, q.pair.src)
		}
		last = q.pair.src
	}
	return
}

func (ex *exchange) destinationCurrencies() []string {
	// quotes are sorted by first currency, so we need to sort&unique results
	var currencies []string
	for _, q := range ex.quotes {
		currencies = append(currencies, q.pair.dst)
	}
	sort.Strings(currencies)
	return uniqueStrings(currencies)
}
