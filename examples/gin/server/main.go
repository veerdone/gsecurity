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
	"github.com/gin-gonic/gin"
	"github.com/veerdone/gsecurity"
	"github.com/veerdone/gsecurity/adaptor/ginadaptor"
	"log"
	"strconv"
)

func main() {
	g := gin.Default()

	g.POST("/login", func(c *gin.Context) {
		queryId := c.Query("id")
		id, _ := strconv.ParseInt(queryId, 10, 64)
		gsecurity.LoginAndSet(id, ginadaptor.New(c))
		c.JSON(200, gin.H{"msg": "ok"})
	})

	g.GET("/isLogin", func(c *gin.Context) {
		isLogin := gsecurity.IsLogin(ginadaptor.New(c))

		c.JSON(200, gin.H{"isLogin": isLogin})
	})

	err := g.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
