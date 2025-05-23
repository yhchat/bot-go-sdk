/*
 * @Date: 2022-10-29 10:41:08
 * @LastEditTime: 2023-03-21 21:16:38
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package main

import (
	"fmt"
	"time"

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

// 接收普通消息
func onMessageNormal(event subscription.MessageEvent) {
	fmt.Println(event)
}

// 接收指令消息
func onMessageInstruction(event subscription.MessageEvent) {
	fmt.Println(event)
}

// 接收群新成员加入事件
func onGroupJoin(event subscription.GroupJoinEvent) {
	fmt.Println(event)
}

// 接收机器人被关注事件
func onBotFollowed(event subscription.BotFollowedEvent) {
	fmt.Println(event)
}

/* 示例方法
 * 发送文本消息
 * token来自于云湖官网控制台
 */
func SendTextMessage(recvId string, recvType string, text string) (openapi.BasicResponse, error) {
	openApi := openapi.NewOpenApi("token")
	textMessage := openapi.TextMessage{
		RecvId:   recvId,
		RecvType: recvType,
		Text:     text,
	}
	return openApi.SendTextMessage(textMessage)
}

/* 示例方法
 * 发送流式消息
 * token来自于云湖官网控制台
 */
func SendStreamMessage(recvId string, recvType string) {

	openApi := openapi.NewOpenApi("token")
	writer, _ := openApi.StreamMessageWriter(recvId, recvType, "markdown")

	// 发送消息数据
	str := "## 测试流式消息"
	for _, r := range str {
		writer.Write([]byte(string(r)))
		time.Sleep(50 * time.Millisecond)
	}

	// 关闭流式消息发送连接
	writer.Close()

	// 获取并打印服务端响应
	resp, _ := writer.GetResponse()
	fmt.Println("Server response:", string(resp))
}

/* 示例方法
 * 发送文本消息
 * token来自于云湖官网控制台
 * buttons消息结构如下：(参考代码文件：openapi/openapi_test.go)
 * 1、[]openapi.Button{}
 * 2、[][]openapi.Button{}
 */
func SendTextMessage1(recvId string, recvType string, text string, buttons interface{}) (openapi.BasicResponse, error) {
	openApi := openapi.NewOpenApi("token")
	textMessage := openapi.TextMessage{
		RecvId:   recvId,
		RecvType: recvType,
		Text:     text,
		Buttons:  buttons,
	}
	return openApi.SendTextMessage(textMessage)
}

/* 示例方法
 * 编辑文本消息
 * msgId字段是来着发送消息时返回的消息ID
 */
func EditTextMessage(msgId string, recvId string, recvType string, text string) (openapi.BasicResponse, error) {
	openApi := openapi.NewOpenApi("token")
	newTextMessage := openapi.EditTextMessage{
		MsgId:    msgId,
		RecvId:   recvId,
		RecvType: recvType,
		Text:     text,
		Buttons:  nil,
	}
	return openApi.EditTextMessage(newTextMessage)
}

/* 示例方法
 * 撤回消息
 * msgId字段是来着发送消息时返回的消息ID
 */
func RecallMessage(msgId string, chatId string, chatType string) (openapi.BasicResponse, error) {
	openApi := openapi.NewOpenApi("token")
	recallMessage := openapi.RecallMessageRequest{
		MsgId:    msgId,
		ChatId:   chatId,
		ChatType: chatType,
	}
	return openApi.RecallMessage(recallMessage)
}
