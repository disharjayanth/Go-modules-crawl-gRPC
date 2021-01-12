package hello

import (
	"rsc.io/quote/v3"
)

// Hello func returns "Hello, world." string
func Hello() string {
	return quote.HelloV3()
}

// Proverb func returns "Concurrency is not parallelism." string
func Proverb() string {
	return quote.Concurrency()
}
