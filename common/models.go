package common

import "github.com/flowswiss/goclient/v2/core"

type List[T any] struct {
	Items      []T
	Pagination core.Pagination
}
