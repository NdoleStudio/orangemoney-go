# orangemoney-go

[![Build](https://github.com/NdoleStudio/orangemoney-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/orangemoney-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/orangemoney-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/orangemoney-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/orangemoney-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/orangemoney-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/orangemoney-go)](https://goreportcard.com/report/github.com/NdoleStudio/orangemoney-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/orangemoney-go)](https://github.com/NdoleStudio/orangemoney-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/orangemoney-go?color=brightgreen)](https://github.com/NdoleStudio/orangemoney-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/orangemoney-go)](https://pkg.go.dev/github.com/NdoleStudio/orangemoney-go)


This package provides a generic `go` client template for the Orange Money API

## Installation

`orangemoney-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/orangemoney-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/orangemoney-go"
```


## Implemented

- **Token**
  - `POST {baseURL}/token`: Get Access Token
- **Merchant Payments**
  - `POST {baseURL}/omcoreapis/1.0.2/mp/init`: Generate pay token.
  - `GET {baseURL}/omcoreapis/1.0.2/mp/push/{payToken}`: Push payment notification.
  - `GET {baseURL}/omcoreapis/1.0.2/mp/paymentstatus/{payToken}`: Get the status of a transaction
  - `GET {baseURL}/omcoreapis/1.0.1/mp/pay`: Execute and initiated transaction

## Usage

### Initializing the Client

An instance of the client can be created using `New()`.

```go
package main

import (
    "github.com/NdoleStudio/orangemoney-go"
)

func main() {
    client := orangemoney.New(
        orangemoney.WithUsername(""),
        orangemoney.WithPassword(""),
        orangemoney.WithAuthToken(""),
    )
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
status, response, err := client.MerchantPayment.Init(context.Background())
if err != nil {
    //handle error
}
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
