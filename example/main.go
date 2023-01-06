/*
 * @Date: 2022-10-29 10:41:08
 * @LastEditTime: 2023-01-06 15:53:53
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package main

import (
	"fmt"

	"github.com/yhchat/bot-go-sdk/openapi"
	"github.com/yhchat/bot-go-sdk/subscription"
)

func main() {
	port := 8805
	subscription := subscription.NewSubscription(port)
	subscription.OnMessageNormal = onMessageNormal
	subscription.OnMessageInstruction = onMessageInstruction
	subscription.OnGroupJoin = onGroupJoin
	subscription.OnBotFollowed = onBotFollowed
	subscription.Start()
}

//接收普通消息
func onMessageNormal(event subscription.MessageEvent) {
	fmt.Println(event)
}

//接收指令消息
func onMessageInstruction(event subscription.MessageEvent) {
	fmt.Println(event)
}

//接收群新成员加入事件
func onGroupJoin(event subscription.GroupJoinEvent) {
	fmt.Println(event)
}

//接收机器人被关注事件
func onBotFollowed(event subscription.BotFollowedEvent) {
	fmt.Println(event)
}

/* 示例方法
 * 发送文本消息
 * token来自于云湖官网控制台
 */
func SendTextMessage(recvId string, recvType string, text string) {
	openApi := openapi.NewOpenApi("token")
	textMessage := openapi.TextMessage{
		RecvId:   recvId,
		RecvType: recvType,
		Text:     text,
	}
	openApi.SendTextMessage(textMessage)
}

/* 示例方法
 * 发送文本消息
 * token来自于云湖官网控制台
 * buttons消息结构如下：(参考代码文件：openapi/openapi_test.go)
 * 1、[]openapi.Button{}
 * 2、[][]openapi.Button{}
 */
func SendTextMessage1(recvId string, recvType string, text string, buttons interface{}) {
	openApi := openapi.NewOpenApi("token")
	textMessage := openapi.TextMessage{
		RecvId:   recvId,
		RecvType: recvType,
		Text:     text,
		Buttons:  buttons,
	}
	openApi.SendTextMessage(textMessage)
}
