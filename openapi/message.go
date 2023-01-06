/*
 * @Date: 2022-07-25 08:29:52
 * @LastEditTime: 2023-01-06 11:01:15
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package openapi

import (
	"fmt"

	"github.com/yhchat/bot-go-sdk/utils"
)

/**
 * @description: 单条，发送文本消息
 * @param {TextMessage} message
 */
func (o *OpenApi) SendTextMessage(message TextMessage) {
	contentType := "text"
	content := map[string]interface{}{"text": message.Text, "buttons": message.Buttons}
	o.SendMessage(message.RecvId, message.RecvType, contentType, content)
}

/**
 * @description: 单条，发送markdown消息
 * @param {MarkdownMessage} message
 */
func (o *OpenApi) SendMarkdownMessage(message MarkdownMessage) {
	contentType := "markdown"
	content := map[string]interface{}{"text": message.Text}
	o.SendMessage(message.RecvId, message.RecvType, contentType, content)
}

/**
 * @description: 单条，发送单条消息
 */
func (o *OpenApi) SendMessage(recvId string, recvType string, contentType string, content map[string]interface{}) {
	var smr = SendMessageRequest{
		RecvId:      recvId,
		RecvType:    recvType,
		ContentType: contentType,
		Content:     content,
	}
	data := utils.InterfaceToJsonBytes(smr)
	url := fmt.Sprintf("https://chat-go.jwzhd.com/open-apis/v1/bot/send?token=%s", o.Token)
	utils.HttpPost(url, data)
}

/**
 * @description: 批量，发送文本消息
 * @param {BatchTextMessage} message
 */
func (o *OpenApi) BatchSendTextMessage(message BatchTextMessage) {
	contentType := "text"
	content := map[string]interface{}{"text": message.Text, "buttons": message.Buttons}
	o.BatchSendMessage(message.RecvIds, message.RecvType, contentType, content)
}

/**
 * @description: 批量，发送markdown消息
 * @param {BatchMarkdownMessage} message
 */
func (o *OpenApi) BatchSendMarkdownMessage(message BatchMarkdownMessage) {
	contentType := "markdown"
	content := map[string]interface{}{"text": message.Text}
	o.BatchSendMessage(message.RecvIds, message.RecvType, contentType, content)
}

/**
 * @description: 批量，批量发送消息
 */
func (o *OpenApi) BatchSendMessage(recvIds []string, recvType string, contentType string, content map[string]interface{}) {
	var smr = BatchSendMessageRequest{
		RecvIds:     recvIds,
		RecvType:    recvType,
		ContentType: contentType,
		Content:     content,
	}
	data := utils.InterfaceToJsonBytes(smr)
	url := fmt.Sprintf("https://chat-go.jwzhd.com/open-apis/v1/bot/batch_send?token=%s", o.Token)
	utils.HttpPost(url, data)
}
