/*
 * @Date: 2023-01-06 14:16:19
 * @LastEditTime: 2023-03-30 12:10:38
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package openapi

import (
	"testing"
	"time"
)

// 测试竖向一维的buttons
func TestTextMessage1(t *testing.T) {
	openApi := NewOpenApi("token")
	buttons := []Button{
		{
			Text:       "复制内容1",
			ActionType: 2,
			Value:      "复制内容1",
		},
		{
			Text:       "复制内容2",
			ActionType: 2,
			Value:      "复制内容2",
		},
	}
	textMessage := TextMessage{
		RecvId:   "7058262",
		RecvType: "user",
		Text:     "test",
		Buttons:  buttons,
	}
	openApi.SendTextMessage(textMessage)
}

// 测试横竖向二维的buttons, 两横两竖
func TestTextMessage2(t *testing.T) {
	openApi := NewOpenApi("token")
	buttons := [][]Button{{
		{
			Text:       "第一行复制1",
			ActionType: 2,
			Value:      "第一行复制1",
		},
		{
			Text:       "第一行复制2",
			ActionType: 2,
			Value:      "第一行复制2",
		},
	}, {
		{
			Text:       "第二行复制1",
			ActionType: 2,
			Value:      "第二行复制1",
		},
		{
			Text:       "第二行复制2",
			ActionType: 2,
			Value:      "第二行复制2",
		},
	}}
	textMessage := TextMessage{
		RecvId:   "7058262",
		RecvType: "user",
		Text:     "test",
		Buttons:  buttons,
	}
	openApi.SendTextMessage(textMessage)
}

// 测试编辑消息
func TestEditMessage(t *testing.T) {
	openApi := NewOpenApi("token")
	textMessage := TextMessage{
		RecvId:   "7058262",
		RecvType: "user",
		Text:     "你好！",
	}
	basicResp, _ := openApi.SendTextMessage(textMessage)
	msgId := basicResp.Data.(map[string]interface{})["messageInfo"].(map[string]interface{})["msgId"].(string)
	time.Sleep(time.Second * 10)
	editTextMessage := EditTextMessage{
		MsgId:    msgId,
		RecvId:   "7058262",
		RecvType: "user",
		Text:     "你好吗？",
	}
	openApi.EditTextMessage(editTextMessage)
}

// 测试markdown消息带button
func TestMarkdownMessage1(t *testing.T) {
	openApi := NewOpenApi("token")
	buttons := []Button{
		{
			Text:       "复制内容1",
			ActionType: 2,
			Value:      "复制内容1",
		},
	}
	markdownMessage := MarkdownMessage{
		RecvId:   "7058262",
		RecvType: "user",
		Text:     "# test",
		Buttons:  buttons,
	}
	openApi.SendMarkdownMessage(markdownMessage)
}

// 测试机器人看板设置接口
func TestSetBotBoard(t *testing.T) {
	openApi := NewOpenApi("token")
	chatId := "7058262"
	chatType := "user"
	memberId := ""
	content := `
	<div  style="background-color:#eff3fc;border-radius: 10px;padding:10px; text-align:center; margin:10px"  >
	<a href="https://www.yhchat.com" target="_blank">https://www.yhchat.com</a>
 <h3><strong>🎉欢迎使用XXX机器人🎉</strong></h3>
  <h2>您的机器人总使用时长：100小时</h2>
   </div>
	`

	openApi.SetBotBoard(chatId, chatType, memberId, "html", content, uint64(time.Now().Unix()+600))
}

// 测试取消机器人看板接口
func TestDismissBotBoard(t *testing.T) {
	openApi := NewOpenApi("token")
	chatId := "7058262"
	chatType := "user"
	memberId := ""
	openApi.DismissBotBoard(chatId, chatType, memberId)
}
