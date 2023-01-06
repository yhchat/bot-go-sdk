/*
 * @Date: 2023-01-06 14:16:19
 * @LastEditTime: 2023-01-06 15:52:39
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package openapi

import (
	"testing"
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
