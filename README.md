# gsecurity

------------
English | [中文](README-zh.md)

Golang implementation imitating [sa-token](https://github.com/dromara/Sa-Token)

- [x] Login authentication
- [x] Kick people offline
- [x] Account banned
- [x] Session query
- [x] Authority certification

```shell
go get github.com/veerdone/gsecurity
```
## Document

### Login
```go
// Login return a token, then set token to cookie or response header or response body
token := gsecurity.Login(10010)

// or use LoginAndSet,login then set token to the cookie
// the second method parameter needs to be implemented Adaptor
http.HandleFunc("/set/session", func(w http.ResponseWriter, req *http.Request) {
    gsecurity.LoginAndSet(10010, standardadaptor.New(req, w))
})
```

### Account Banned
```go
// disable with userId and disable expire time
gsecurity.Disable(10010, 3600)

// disable with userId, disable level, disable expire time
gsecurity.DisableWithLevel(10010, 1, 3600)

// disable with userId, disable level, disable expire time and service
gsecurity.DisableWithLevelAndService(10010, 1, 3600, "comment")

// has been disabled with userId
disabled := gsecurity.IsDisable(10010)
// or
gsecurity.IsDisableWithLevel(10010, 1)
gsecurity.IsDisableWithLevelAndService(10010, 1, "comment")

// has been disabled with userId, return error
disbleErr := gsecurity.CheckDisable(10010)
// or
gsecurity.CheckDisableWithLevel(10010, 1)
gsecurity.CheckDisableWithLevelAndService(10010, 1, "comment")
```

### Session query
```go
http.HandleFunc("/set/session", func (w http.ResponseWriter, req *http.Request) {
// if not login will return nil
    session := gsecurity.Sessions(standardadaptor.New(req, w))
    if session != nil {
        session.Set("key", "value")
        session.Get("key")
    }
})

```

-----------
### Example
- [standard](examples/standard)
- [gin](examples/gin)