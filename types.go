package main

import (
	"strings"

	"github.com/shopspring/decimal"
)

type pair struct {
	src string
	dst string
}

type quote struct {
	pair  pair
	price decimal.Decimal
}

func (q *quote) compare(other quote) int {
	if q.pair.src == other.pair.src {
		if q.pair.dst == other.pair.dst {
			return q.price.Cmp(other.price)
		}
		return strings.Compare(q.pair.dst, other.pair.dst)
	}
	return strings.Compare(q.pair.src, other.pair.src)
}

type sortableQuotes []quote

// sort.Interface
func (q sortableQuotes) Len() int {
	return len(q)
}

func (q sortableQuotes) Less(i int, j int) bool {
	return q[i].compare(q[j]) < 0
}

func (q sortableQuotes) Swap(i int, j int) {
	q[i], q[j] = q[j], q[i]
}

type fee struct {
	// upper bound limit for particular fee
	volume decimal.Decimal
	maker  decimal.Decimal
	taker  decimal.Decimal
}

type feeSchedule struct {
	pair pair
	fees []fee
}
