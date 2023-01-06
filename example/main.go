/*
 * @Date: 2022-10-29 10:41:08
 * @LastEditTime: 2023-01-06 10:06:28
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package main

import (
	"bot-go-sdk/openapi"
	"bot-go-sdk/subscription"
	"fmt"
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

//发送文本消息
//token来自于云湖官网控制台
func SendTextMessage(recvId string, recvType string, text string) {
	openApi := openapi.NewOpenApi("token")
	textMessage := openapi.TextMessage{
		RecvId:   recvId,
		RecvType: recvType,
		Text:     text,
	}
	openApi.SendTextMessage(textMessage)
}
