# statuspage-go ![GitHub Actions](https://github.com/nagelflorian/statuspage-go/actions/workflows/main.yml/badge.svg)

This repository contains a Go client library for accessing the [Statuspage REST API v1](https://developer.statuspage.io).

## Install the package

Install `statuspage-go` via `go-get`:

```bash
go get github.com/nagelflorian/statuspage-go
```

## Getting Started

Before you can initialize an instance you'll have to [obtain an API key](https://developer.statuspage.io/#section/Authentication) from the Statuspage account view.

```go
package main

import "github.com/nagelflorian/statuspage-go"

func main() {
  client := statuspage.New("YOUR_API_KEY", nil)

  // Use the client.

  // Get the page profile for a given page id
  page, err := client.Page.GetPage(context.TODO(), "YOUR_PAGE_ID")
}
```

## API Documentation

The official Statuspage API documentation can be found here: [developer.statuspage.io](https://developer.statuspage.io).

## License

This library is distributed under the MIT-style license found in the LICENSE file.
