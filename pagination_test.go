package goclient

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestParsePagination(t *testing.T) {
	tests := []struct {
		Pagination Pagination
		Next       Cursor
		HasMore    bool
	}{
		{
			Pagination: Pagination{
				Cursor:     Cursor{Page: 2, PerPage: 5, NoFilter: 0},
				ItemCount:  5,
				TotalCount: 13,
				TotalPages: 3,
				Links: Links{
					First:   "https://api.flow.swiss/v4/compute/instances?page=1&per_page=5",
					Last:    "https://api.flow.swiss/v4/compute/instances?page=3&per_page=5",
					Current: "https://api.flow.swiss/v4/compute/instances?page=2&per_page=5",
					Prev:    "https://api.flow.swiss/v4/compute/instances?page=1&per_page=5",
					Next:    "https://api.flow.swiss/v4/compute/instances?page=3&per_page=5",
				},
			},
			Next: Cursor{
				Page:     3,
				PerPage:  5,
				NoFilter: 0,
			},
			HasMore: true,
		},
		{
			Pagination: Pagination{
				Cursor:     Cursor{Page: 1, PerPage: 5, NoFilter: 0},
				ItemCount:  5,
				TotalCount: 13,
				TotalPages: 3,
				Links: Links{
					First:   "https://api.flow.swiss/v4/compute/instances?page=1&per_page=5",
					Last:    "https://api.flow.swiss/v4/compute/instances?page=3&per_page=5",
					Current: "https://api.flow.swiss/v4/compute/instances?page=1&per_page=5",
					Prev:    "",
					Next:    "https://api.flow.swiss/v4/compute/instances?page=2&per_page=5",
				},
			},
			Next: Cursor{
				Page:     2,
				PerPage:  5,
				NoFilter: 0,
			},
			HasMore: true,
		},
		{
			Pagination: Pagination{
				Cursor:     Cursor{Page: 3, PerPage: 5, NoFilter: 0},
				ItemCount:  3,
				TotalCount: 13,
				TotalPages: 3,
				Links: Links{
					First:   "https://api.flow.swiss/v4/compute/instances?page=1&per_page=5",
					Last:    "https://api.flow.swiss/v4/compute/instances?page=3&per_page=5",
					Current: "https://api.flow.swiss/v4/compute/instances?page=3&per_page=5",
					Prev:    "https://api.flow.swiss/v4/compute/instances?page=2&per_page=5",
					Next:    "",
				},
			},
			HasMore: false,
		},
	}

	for idx, expectation := range tests {
		t.Run(fmt.Sprintf("#%d", idx), func(t *testing.T) {
			expectedPagination := expectation.Pagination

			var links []string
			buildLink := func(name, link string) string {
				return fmt.Sprintf(`<%s>; rel="%s"`, link, name)
			}
			if len(expectedPagination.Links.First) > 0 {
				links = append(links, buildLink("first", expectedPagination.Links.First))
			}
			if len(expectedPagination.Links.Last) > 0 {
				links = append(links, buildLink("last", expectedPagination.Links.Last))
			}
			if len(expectedPagination.Links.Current) > 0 {
				links = append(links, buildLink("self", expectedPagination.Links.Current))
			}
			if len(expectedPagination.Links.Next) > 0 {
				links = append(links, buildLink("next", expectedPagination.Links.Next))
			}
			if len(expectedPagination.Links.Prev) > 0 {
				links = append(links, buildLink("prev", expectedPagination.Links.Prev))
			}

			header := http.Header{}
			header.Set("X-Pagination-Current-Page", fmt.Sprint(expectedPagination.Page))
			header.Set("X-Pagination-Limit", fmt.Sprint(expectedPagination.PerPage))
			header.Set("X-Pagination-Count", fmt.Sprint(expectedPagination.ItemCount))
			header.Set("X-Pagination-Total-Count", fmt.Sprint(expectedPagination.TotalCount))
			header.Set("X-Pagination-Total-Pages", fmt.Sprint(expectedPagination.TotalPages))
			header.Set("Link", strings.Join(links, ", "))

			pagination := ParsePagination(header)

			if !reflect.DeepEqual(expectedPagination, pagination) {
				t.Errorf("expected pagination to equal %v, got %v", expectedPagination, pagination)
			}

			if expectation.HasMore {
				if !pagination.HasMore() {
					t.Errorf("expected pagination to have more pages")
				}

				next := pagination.Next()
				if !reflect.DeepEqual(expectation.Next, next) {
					t.Errorf("expected next cursor to equal %v, got %v", expectation.Next, next)
				}
			} else if pagination.HasMore() {
				t.Errorf("expected pagination not to have more pages")
			}
		})
	}
}
