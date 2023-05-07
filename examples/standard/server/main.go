package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/veerdone/gsecurity"
	"net/http"
	"strconv"
)

func main() {
	// use redis store
	redisStore := gsecurity.NewRedisStore(redis.NewClient(&redis.Options{Addr: "localhost:6379"}))
	// new logic
	logic := gsecurity.NewLogic(gsecurity.DefaultConfig, redisStore)
	// set default security
	gsecurity.SetDefaultSecurity(logic)

	http.HandleFunc("/login", func(writer http.ResponseWriter, req *http.Request) {
		queryId := req.URL.Query().Get("id")
		id, _ := strconv.ParseInt(queryId, 10, 64)
		token := gsecurity.Login(id)
		writer.Header().Set(gsecurity.DefaultConfig.TokenName, token)
		_, err := fmt.Fprintln(writer, "ok")
		if err != nil {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/isLogin", func(writer http.ResponseWriter, req *http.Request) {
		isLogin := gsecurity.IsLogin(gsecurity.Standard(req))
		_, err := fmt.Fprintln(writer, "登录成功:", isLogin)
		if err != nil {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/set/session", func(w http.ResponseWriter, req *http.Request) {
		session := gsecurity.Sessions(gsecurity.Standard(req))
		session.Set("key", "value")
	})

	http.HandleFunc("/get/session", func(w http.ResponseWriter, req *http.Request) {
		session := gsecurity.Sessions(gsecurity.Standard(req))
		v, b := session.Get("key")
		fmt.Println(v, b)
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, req *http.Request) {
		gsecurity.Logout(gsecurity.Standard(req))
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
