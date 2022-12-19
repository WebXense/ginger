# ginger

Use [gin](https://github.com/gin-gonic/gin) in easier way.

## Installation

```bash
go get -u github.com/WebXense/ginger
```

## Initialization

This package will initialize gin engine at startup. You can call the engine directly by `ginger.Engine`.

```go
engine := ginger.Engine
```

## Routing

This package did not provide methods for routing. You should setup the routes by calling the methods of engine. Please refer to [gin](https://github.com/gin-gonic/gin).

## Request

This package provides some methods to get the request data easily.

### uri parameter

Setup the `uri` tag in the struct to get the uri parameter.

```go
type RequestObject struct {
    ID int `uri:"id"`
}

obj := ginger.GetRequest[RequestObject](ctx)
```

### query parameter

Setup the `form` tag in the struct to get the query parameter.

```go
type RequestObject struct {
    ID int `form:"id"`
}

obj := ginger.GetRequest[RequestObject](ctx)
```

### body

Setup the `json` tag in the struct to get the body.

```go
type RequestObject struct {
    ID int `json:"id"`
}

obj := ginger.GetRequest[RequestObject](ctx)
```

### pagination

Get the `page` and `size` from query parameter. This method will return a `Pagination` struct from [sql]("https://github.com/WebXense/sql").

```go
// url should be like this: /api/v1/users?page=1&size=10
pagination := ginger.GetPagination(ctx)
```

### sort

Get the `sortBy` and `asc` from query parameter. This method will return a `Sort` struct from [sql]("https://github.com/WebXense/sql").

```go
// url should be like this: /api/v1/users?sortBy=id&asc=true
sort := ginger.GetSort(ctx)
```

## Response

This package provides a standard restful response.

### success

```go
ginger.OK(ctx, data, pagination)
```

### error

```go
ginger.ERR(ctx, errCode, errMsg, data)
```

## CORS

```go
cors.Setup(ginger.Engine, cors.Options{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
    ExposeHeaders: []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge: 12 * time.Hour,
})
```

## Rate Limit

Setup the global rate limit.

```go
rateLimit.GeneralRateLimit(ginger.Engine, time.Hour * 24, 1000) // 1000 requests per day
```

Get the rate limit middleware.

```go
middleware := rateLimit.RateLimit(time.Hour * 24, 1000) // 1000 requests per day
```
