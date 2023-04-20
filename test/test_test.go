package test_test

import (
	"testing"

	"github.com/slashid/httpcache"
	"github.com/slashid/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
