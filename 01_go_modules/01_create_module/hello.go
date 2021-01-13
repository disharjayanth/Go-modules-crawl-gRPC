package hello

import (
	"github.com/disharjayanth/quotes"
	quoteV2 "rsc.io/quote/v2"
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

func worldTraverlers() string {
	return quote.GlassV3()
}

func glassV3() string {
	return quote.GlassV3()
}

func optV2() string {
	return quoteV2.OptV2()
}

func quotesFromGithubRep() []string {
	return quotes.Favs()
}
