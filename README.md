# gsecurity

------------
English | [中文](README-zh.md)

Golang implementation imitating [sa-token](https://github.com/dromara/Sa-Token)


```shell
go get github.com/veerdone/gsecurity
```

Implement login with just one line of code
```go
// Login return a token, then set token to cookie or response header or response body
token := gsecurity.Login(10010)

// or use LoginAndSet,login then set token to the cookie
// the second method parameter needs to be implemented Adaptor
gsecurity.LoginAndSet(10010, standardadaptor.New(request))
```
## Document

-----------
### Example
- [standard](examples/standard)
- [gin](examples/gin)