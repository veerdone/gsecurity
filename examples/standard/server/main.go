/*
 * Copyright 2023 veerdone
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/veerdone/gsecurity"
	"github.com/veerdone/gsecurity/adaptor/standardadaptor"
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

	http.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		queryId := req.URL.Query().Get("id")
		id, _ := strconv.ParseInt(queryId, 10, 64)
		token := gsecurity.LoginAndSet(id, standardadaptor.New(req, w))
		fmt.Fprintf(w, "login success, token is: %s", token)
	})

	http.HandleFunc("/isLogin", func(w http.ResponseWriter, req *http.Request) {
		isLogin := gsecurity.IsLogin(standardadaptor.New(req, w))
		fmt.Fprintf(w, "isLogin: %t", isLogin)
	})

	http.HandleFunc("/set/session", func(w http.ResponseWriter, req *http.Request) {
		session := gsecurity.Sessions(standardadaptor.New(req, w))
		session.Set("key", "value")
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/get/session", func(w http.ResponseWriter, req *http.Request) {
		session := gsecurity.Sessions(standardadaptor.New(req, w))
		v, b := session.Get("key")
		if b {
			fmt.Fprintf(w, "value is %v", v)
		} else {
			fmt.Fprintln(w, "not value exist")
		}
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, req *http.Request) {
		gsecurity.Logout(standardadaptor.New(req, w))
		fmt.Fprintln(w, "logout success")
	})

	http.HandleFunc("/disable", func(w http.ResponseWriter, req *http.Request) {
		gsecurity.Disable(gsecurity.GetLoginId(standardadaptor.New(req, w)), 3600)
		fmt.Fprintln(w, "disable success")
	})

	http.HandleFunc("/isDisable", func(w http.ResponseWriter, req *http.Request) {
		isDisable := gsecurity.IsDisable(gsecurity.GetLoginId(standardadaptor.New(req, w)))
		fmt.Fprintf(w, "is disable: %t", isDisable)
	})

	http.HandleFunc("/getDisableTime", func(w http.ResponseWriter, req *http.Request) {
		time := gsecurity.DisableExTime(gsecurity.GetLoginId(standardadaptor.New(req, w)))
		fmt.Fprintf(w, "disable expire time: %d", time)
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
