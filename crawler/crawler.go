package crawler

import (
	"context"
	"sync"
)

type Crawler interface {
	// Find product information from the website
	Crawl(ctx context.Context, page int, newProducts chan *Product, wgJob *sync.WaitGroup)
	FindMaxPage(ctx context.Context, totalWebProduct int) (int, error)
	GetQuerySrc() *Query
}

type Product struct {
	Word       string
	ProductID  string
	Name       string
	Price      int
	ImageURL   string
	ProductURL string
	Website    string
}

type Query struct {
	Web     Web
	Keyword string
}

type Web string

const (
	Pchome Web = "pchome"
	Momo   Web = "momo"
)

func newQuery(web Web, keyWord string) *Query {
	return &Query{Web: web, Keyword: keyWord}
}
