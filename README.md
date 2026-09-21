# GoClient
This is the official API client, written in Go, for accessing the
[Flow Swiss](https://flow.swiss/) and [Cloudbit](https://www.cloudbit.ch/)
APIs. Our API Documentation can be found at [my.flow.swiss/#/doc](https://my.flow.swiss/#/doc)
or [my.cloudbit.ch/#/doc](https://my.cloudbit.ch/#/doc).

## Installation
```
go get github.com/flowswiss/goclient/v2
```

## Example
```go
package main

import (
	"context"
	"fmt"

	"github.com/flowswiss/goclient/v2"
	"github.com/flowswiss/goclient/v2/core"
)

func main() {
	client := goclient.WithToken("your-application-token")

	res, err := client.Compute.Server.List(context.Background(), core.Cursor{
		Page:    1,
		PerPage: 5,
	})

	if err != nil {
		fmt.Println("error listing servers: ", err)
	}

	for _, server := range res.Items {
		fmt.Println("found server with id ", server.ID)
	}
}
```

## License
MIT License
