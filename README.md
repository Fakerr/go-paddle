# go-paddle #

[![Test Status](https://github.com/Fakerr/go-paddle/actions/workflows/test.yml/badge.svg)](https://github.com/Fakerr/go-paddle/actions/workflows/test.yml)

go-paddle is a Go client library for accessing the [Paddle API](https://developer.paddle.com/api-reference/intro).

## Installation ##


```bash
go get github.com/Fakerr/go-paddle
```


Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/Fakerr/go-paddle"
```

and run `go get` without paramters.

## Usage ##
The package paddle comes with two different clients. A client for the Product, Subscription, Alert APIs that will require
a vendor_id and a vendor_auth_code and a client for the Checkout API.


The services of a client divide the API into logical chunks and correspond to
the structure of the Paddle API documentation at
https://developer.paddle.com/api-reference/.

### Product API, Subscription API, Alert API ###

```go
import "github.com/Fakerr/go-paddle/paddle"
```

Construct a new Paddle client, then use the various services on the client to access different parts of the Paddle API.
The client always requires a vendor_id and a vendor_auth_code arguments that you can get from the Paddle dashboard. For example:

```go
client := paddle.NewClient(vendorId, vendorAuthCode, nil)

// List all users subscribed to any of your subscription plans
users, _, err := client.Users.List(context.Background(), nil)
```

Some API methods have optional parameters that can be passed. For example:

```go
client := paddle.NewClient(vendorId, vendorAuthCode, nil)

// List all users subscribed to any of your subscription plans
opt := &UsersOptions{SubscriptionID: "1"}
users, _, err := client.Users.List(context.Background(), opt)
```

NOTE: Using the [context](https://godoc.org/context) package, one can easily
pass cancelation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then `context.Background()`
can be used as a starting point.

For more sample code snippets, head over to the
[example](https://github.com/fakerr/go-paddle/tree/master/example) directory.

### Checkout API ###

```go
import "github.com/Fakerr/go-paddle/paddle"
```

Construct a new Paddle checkout client, then use the various services on the client to access different parts of the Paddle API.

```go
client := paddle.NewCheckoutClient(nil)

// Retrieve prices for one or multiple products or plans
options := &PricesOptions{CustomerCountry: "tn"}
prices, _, err := client.Prices.Get(context.Background(), "1", options)

```
### Sandbox environment ###
If you want to send requests against a sandbox environment, the package paddle provides two specific clients for that purpose:

```
client := paddle.NewSandboxClient(sandboxVendorId, sandboxVendorAuthCode, nil)
```
or to access the checkout API 

```
client := paddle.NewSandboxCheckoutClient(nil)
```

### Pagination ###

Some requests for resource collections (users, webhooks, etc.)
support pagination. Pagination options are described in the
`paddle.ListOptions` struct and passed to the list methods directly or as an
embedded type of a more specific list options struct (for example
`paddle.UsersOptions`).

```go
client := paddle.NewClient(vendorId, vendorAuthCode, nil)

// List all users subscribed to any of your subscription plans
opt := &UsersOptions{
	SubscriptionID: "1",
	ListOptions: ListOptions{Page: 2},
}
users, _, err := client.Users.List(context.Background(), opt)
```
## Contributing ##
Pull requests are welcome, along with any feedback or ideas. The calling pattern is pretty well established, so adding new methods is relatively
straightforward.

## License ##

This library is distributed under the MIT-style license found in the [LICENSE](./LICENSE)
file.
