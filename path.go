package goclient

import (
	"fmt"
	"path"
)

func Join(parts ...interface{}) (p string) {
	for _, part := range parts {
		p = path.Join(p, fmt.Sprint(part))
	}
	return
}
