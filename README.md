# gofetch

A simple HTTP package for Go.

## Usage

```go
package main

func main(){
    // create client
    client := gofetch.New()
    // with url
    client.With(gofetch.WithURL("http://something.com/get"))
    // with query params
    client.With(gofetch.WithParams(gofetch.Params{
        "userid": 1
    }))
    // with method
    client.With(gofetch.WithMethod("GET"))
    // with header
    client.With(gofetch.WithHeader("Content-Type", "application/json"))
    // with headers
    client.With(gofetch.WithHeaders(gofetch.Headers{
        "Authorization": "Bearer token",
        "Content-Type": "application/json",
        // ...
    }))
    // with form
    client.With(gofetch.WithForm(gofetch.Params{
        "userid": 1
    }))
    // with body
    client.With(gofetch.WithBody(`user=1`))
    // with json
    client.With(gofetch.WithJSON(struct{User int `json:"user"`}{
        User: 1,
    }))
    // with timeout
    client.With(gofetch.WithTimeout(5 * time.Second))

    // fetch result
    ret, err := client.Fetch()
}
```

### Simple GET

```go
package main

func main(){
    client := gofetch.New()
    ret, err := gofetch.GET("http://something.com/get", gofetch.Params{
        "userid": 1
    });
}
```

### Simple POST

```go
package main

func main(){
    client := gofetch.New()
    ret, err := gofetch.POST("http://something.com/post", gofetch.Params{
        "userid": 1
    });
}
```

### POST Json

```go
package main

type User struct {
    User int `json:"user"`
}

func main(){
    client := gofetch.New()
    ret, err := gofetch.POSTJson("http://something.com/post", User{User: 1});
}
```


## With

You can also custom the `with option` by using `DoFunc` type.

```go
package gofetch

type DoFunc func(*Request)
```

### Example

use function to set cookie

```go
package main

func main(){
    client := gofetch.New()
    client.With(WithCookie("cookie string"))
    client.GET("http://something.com/get", nil)
}

func WithCookie(cookie string) gofetch.DoFunc {
    return func(req *gofech.Request) {
        req.Headers["Cookie"] = cookie
    }
}
```

another method for set cookie

```go
package main

func main(){
    client := gofetch.New()
    client.With(func(req *gofetch.Request) {
        req.Headers["Cookie"] = "cookie string"
    })
    client.GET("http://something.com/get", nil)
}
```
