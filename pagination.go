package goclient

import (
	"net/http"
	"regexp"
	"strconv"
)

type Cursor struct {
	// Number of the current page
	Page int `url:"page,omitempty"`
	// Maximum number of items per page
	PerPage int `url:"per_page,omitempty"`
	// If set to 1, disables pagination entirely so that all entities will be returned
	NoFilter int `url:"no_filter,omitempty"`
}

func (c Cursor) Next() Cursor {
	c.Page += 1
	return c
}

// Provides direct links to the individual pages
type Links struct {
	// URL to the first page
	First string
	// URL to the last page
	Last string
	// URL to the current page
	Current string
	// URL to the previous page
	Prev string
	// URL to the next page
	Next string
}

type Pagination struct {
	// Cursor of the current request
	Cursor
	// Number of items on the current page
	ItemCount int
	// Total number of items
	TotalCount int
	// Total number of pages
	TotalPages int
	// Direct links to the individual pages
	Links Links
}

func (p Pagination) HasMore() bool {
	return p.Page < p.TotalPages
}

// Parses the pagination headers from the given http response.
func ParsePagination(header http.Header) Pagination {
	return Pagination{
		Cursor: Cursor{
			Page:     intOrZero(header.Get("X-Pagination-Current-Page")),
			PerPage:  intOrZero(header.Get("X-Pagination-Limit")),
			NoFilter: 0,
		},
		ItemCount:  intOrZero(header.Get("X-Pagination-Count")),
		TotalCount: intOrZero(header.Get("X-Pagination-Total-Count")),
		TotalPages: intOrZero(header.Get("X-Pagination-Total-Pages")),
		Links:      parseLinks(header.Get("Link")),
	}
}

func parseLinks(header string) Links {
	res := Links{}
	regex := regexp.MustCompile("<([^>]+)>; rel=\"(\\w+)\"(?:,\\s?)?")

	links := regex.FindAllStringSubmatch(header, 5)
	for _, link := range links {
		switch link[2] {
		case "first":
			res.First = link[1]
		case "last":
			res.Last = link[1]
		case "self":
			res.Current = link[1]
		case "next":
			res.Next = link[1]
		case "prev":
			res.Prev = link[1]
		default:
			continue
		}
	}

	return res
}

func intOrZero(val string) int {
	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0
	}

	return int(i)
}
