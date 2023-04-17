/*
 * @Date: 2023-04-17 13:42:10
 * @LastEditTime: 2023-04-17 13:57:43
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yhchat/bot-go-sdk/subscription"
)

func main() {
	port := 8805
	s := subscription.NewSubscription(port)
	s.OnMessageNormal = onMessageNormal
	s.Router = CustomRouter(s)
	s.Start()
}

func CustomRouter(s *subscription.Subscription) *gin.Engine {
	router := gin.Default()
	router.POST("/custom", func(c *gin.Context) {
		var sr subscription.SubScriptionResp
		if err := c.BindJSON(&sr); err != nil {
			return
		}
		s.Parse(sr)
	})

	return router
}

//接收普通消息
func onMessageNormal(event subscription.MessageEvent) {
	fmt.Println(event)
}
