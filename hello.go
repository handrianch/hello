package hello

import (
	"rsc.io/quote/v3"
)

// Hello printed
func Hello() string {
	return quote.HelloV3()
}

// Proverb method
func Proverb() string {
	return quote.Concurrency()
}
