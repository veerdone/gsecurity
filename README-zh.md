# gsecurity

------------
[English](README.md) | 中文

参考 [sa-token](https://github.com/dromara/Sa-Token) 的 Golang 实现

```shell
go get github.com/veerdone/gsecurity
```
要实现登录，只需要一行代码
```go
// 返回一个token, 将 token 设置到 cookie 或者响应头或响应体中
toekn := gsecurity.Login(10010)

// 或者使用 LoginAndSet, 登录后将 token 设置到 cookie 中
// 第二个方法参数需要实现 adaptor.Adaptor
gsecurity.LoginAndSet(10010, standardadaptor.New(request))
```
## 文档

-----------
### 示例
- [standard(标准库)](examples/standard)
- [gin](examples/gin)