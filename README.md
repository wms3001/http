# golang http 请求封装
## 简介
实现http请求
## 使用
```
go get github.com/wms3001/http
```
## 实例
1. get
```go
m := map[string]string{
    "name":    "backy",
    "species": "dog",
}
a := map[string]string{
    "Accept": "application/json",
    //"Accept": "text/plain",
}
hp := Http{
    "https://www.wishpost.cn/api/v3/oauth/access_token", //url
    "", // username
    "", //password
    m,  //header
    a,  //accept
    "", //body
    10, //timeout
    "", //type
    []byte{}, //file
}
resp := hp.Get()
```
2. post
```go
m := map[string]string{
    "name":    "backy",
    "species": "dog",
}
a := map[string]string{
    "Accept": "application/json",
    //"Accept": "text/plain",
}
hp := Http{
    "https://www.wishpost.cn/api/v3/oauth/access_token",
    "",
    "",
    m,
    a,
    "",
    10,
    "",
    []byte{},
}
resp := hp.Post()
```
3. delete
```go
m := map[string]string{
    "name":    "backy",
    "species": "dog",
}
a := map[string]string{
    "Accept": "application/json",
    //"Accept": "text/plain",
}
hp := Http{
    "https://www.wishpost.cn/api/v3/oauth/access_token",
    "",
    "",
    m,
    a,
    "",
    10,
    "",
    []byte{},
}
resp := hp.Delete()
```
4. head
```go
m := map[string]string{
    "name":    "backy",
    "species": "dog",
}
a := map[string]string{
    "Accept": "application/json",
    //"Accept": "text/plain",
}
hp := Http{
    "https://www.wishpost.cn/api/v3/oauth/access_token",
    "",
    "",
    m,
    a,
    "",
    10,
    "",
    []byte{},
}
resp := hp.Head()
```
5. options
```go
m := map[string]string{
    "name":    "backy",
    "species": "dog",
}
a := map[string]string{
    "Accept": "application/json",
    //"Accept": "text/plain",
}
hp := Http{
    "https://www.wishpost.cn/api/v3/oauth/access_token",
    "",
    "",
    m,
    a,
    "",
    10,
    "",
    []byte{},
}
resp := hp.Options()
```
6. put
```go
m := map[string]string{
    "name":    "backy",
    "species": "dog",
}
a := map[string]string{
    "Accept": "application/json",
    //"Accept": "text/plain",
}
hp := Http{
    "https://www.wishpost.cn/api/v3/oauth/access_token",
    "",
    "",
    m,
    a,
    "",
    10,
    "",
    []byte{},
}
resp := hp.Put()
```
7. trace
```go
m := map[string]string{
    "name":    "backy",
    "species": "dog",
}
a := map[string]string{
    "Accept": "application/json",
    //"Accept": "text/plain",
}
hp := Http{
    "https://www.wishpost.cn/api/v3/oauth/access_token",
    "",
    "",
    m,
    a,
    "",
    10,
    "",
    []byte{},
}
resp := hp.Trace()
```