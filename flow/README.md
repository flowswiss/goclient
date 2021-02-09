# GoClient (deprecated)

***Deprecation Notice:*** this is the old api integration which is deprecated
since February 2021 and will be removed in a later version. Do not use this for
new integrations.

## Example
```go
package main

import (
  "context"
  "fmt"

  "github.com/flowswiss/goclient/flow"
)

func main() {
  client := flow.NewClientWithToken("your-application-token")

  servers, _, err := client.Server.List(context.Background(), flow.PaginationOptions{
    Page:    1,
    PerPage: 5,
  })

  if err != nil {
    fmt.Println("error listing servers: ", err)
  }

  for _, server := range servers {
    fmt.Println("found server with id ", server.Id)
  }
}
```
