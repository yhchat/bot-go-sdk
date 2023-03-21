/*
 * @Date: 2023-01-06 14:16:19
 * @LastEditTime: 2023-03-21 21:13:39
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package openapi

import (
	"testing"
	"time"
)

//测试竖向一维的buttons
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

//测试横竖向二维的buttons
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
