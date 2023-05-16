# gsecurity

------------
[English](README.md) | 中文

参考 [sa-token](https://github.com/dromara/Sa-Token) 的 Golang 实现

- [x] 登陆认证
- [x] 踢人下线
- [x] 账号封禁
- [x] 会话查询
- [ ] 权限认证

```shell
go get github.com/veerdone/gsecurity
```

## 文档


### 登陆
```go
// 使用账号登陆, 返回 token, 可以将 token 设置到 Cookie 或 Response Header
token := gsecurity.Login(10010)

// 或者使用 LoginAndSet, 登陆并设置 token 到 Cookie
// 第二个参数需要实现 gsecurity.Adaptor
http.HandleFunc("/set/session", func(w http.ResponseWriter, req *http.Request) {
    gsecurity.LoginAndSet(10010, standardadaptor.New(req, w))
})
```

### Account Banned
```go
// 封禁账号并指定封禁时间
gsecurity.Disable(10010, 3600)

// 阶梯封禁, 封禁账号指定封禁级别和时间
gsecurity.DisableWithLevel(10010, 1, 3600)

// 分类封禁, 封禁账号指定封禁级别、时间和服务
gsecurity.DisableWithLevelAndService(10010, 1, 3600, "comment")

// 检查账号是否被封禁
disabled := gsecurity.IsDisable(10010)
// or
gsecurity.IsDisableWithLevel(10010, 1)
gsecurity.IsDisableWithLevelAndService(10010, 1, "comment")

// 检查账号是否被封禁, 如果被封禁了会返回 error
disbleErr := gsecurity.CheckDisable(10010)
// or
gsecurity.CheckDisableWithLevel(10010, 1)
gsecurity.CheckDisableWithLevelAndService(10010, 1, "comment")
```

### 会话查询
```go
http.HandleFunc("/set/session", func (w http.ResponseWriter, req *http.Request) {
// 如果没有登陆, 会返回 nil
    session := gsecurity.Sessions(standardadaptor.New(req, w))
    if session != nil {
        session.Set("key", "value")
        session.Get("key")
    }
})

```


-----------
### 示例
- [standard(标准库)](examples/standard)
- [gin](examples/gin)