package main

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func sampleQuotes() []quote {
	return []quote{
		{pair{"ETC", "LTC"}, decimal.New(12345, 2)},
		{pair{"ETH", "ETC"}, decimal.New(54321, 2)},
		{pair{"ETC", "BTC"}, decimal.New(23456, 2)},
		{pair{"BTC", "ETC"}, decimal.New(54321, 2)},
	}
}

func TestNewExchange(t *testing.T) {
	quotes := sampleQuotes()
	ex := newExchange("test", "http://test.test", quotes)

	assert.NotNil(t, ex)
	assert.Equal(t, len(quotes), len(ex.quotes))
	assert.Nil(t, ex.fees)
}

func TestSourceCurrencies(t *testing.T) {
	quotes := sampleQuotes()
	ex := newExchange("test", "http://test.test", quotes)

	assert.Equal(t, []string{"BTC", "ETC", "ETH"}, ex.sourceCurrencies())
}

func TestDestinationCurrencies(t *testing.T) {
	quotes := sampleQuotes()
	ex := newExchange("test", "http://test.test", quotes)

	assert.Equal(t, []string{"BTC", "ETC", "LTC"}, ex.destinationCurrencies())
}
