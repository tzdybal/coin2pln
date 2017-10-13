package main

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestQuoteComaparator(t *testing.T) {
	q1 := quote{pair{"ETC", "LTC"}, decimal.New(12345, 2)}
	q2 := quote{pair{"ETC", "BTC"}, decimal.New(12345, 2)}
	q3 := quote{pair{"BTC", "LTC"}, decimal.New(12345, 2)}
	q4 := quote{pair{"BTC", "LTC"}, decimal.New(23456, 2)}
	q5 := quote{pair{"BTC", "LTC"}, decimal.New(23456, 2)}

	assert.True(t, q1.compare(q2) > 0)
	assert.True(t, q2.compare(q1) < 0)

	assert.True(t, q1.compare(q3) > 0)
	assert.True(t, q3.compare(q1) < 0)

	assert.True(t, q3.compare(q4) < 0)

	assert.Equal(t, q4.compare(q5), 0)
}

func TestQuoteSorting(t *testing.T) {
	var quotes []quote

	currs := []string{"ETC", "ETH", "BTC", "LTC", "PLN"}

	for i, c1 := range currs {
		for j, c2 := range currs {
			if i != j {
				for k := 0; k < 5; k++ {
					quotes = append(quotes, quote{pair{c1, c2}, decimal.New(rand.Int63(), 2)})
				}
			}
		}
	}

	sort.Sort(sortableQuotes(quotes))

	for i := 1; i < len(quotes); i++ {
		// check with comparator
		assert.True(t, quotes[i-1].compare(quotes[i]) < 0)
	}
}
