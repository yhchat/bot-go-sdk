/*
 * @Date: 2022-07-25 08:29:52
 * @LastEditTime: 2023-03-21 21:01:01
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package openapi

import (
	"encoding/json"
	"fmt"

	"github.com/yhchat/bot-go-sdk/utils"
)

/**
 * @description: 单条，发送文本消息
 * @param {TextMessage} message
 */
func (o *OpenApi) SendTextMessage(message TextMessage) (BasicResponse, error) {
	contentType := "text"
	content := map[string]interface{}{"text": message.Text, "buttons": message.Buttons}
	return o.SendMessage(message.RecvId, message.RecvType, contentType, content)
}

/**
 * @description: 单条，发送markdown消息
 * @param {MarkdownMessage} message
 */
func (o *OpenApi) SendMarkdownMessage(message MarkdownMessage) (BasicResponse, error) {
	contentType := "markdown"
	content := map[string]interface{}{"text": message.Text}
	return o.SendMessage(message.RecvId, message.RecvType, contentType, content)
}

/**
 * @description: 单条，发送单条消息
 */
func (o *OpenApi) SendMessage(recvId string, recvType string, contentType string, content map[string]interface{}) (BasicResponse, error) {
	var smr = SendMessageRequest{
		RecvId:      recvId,
		RecvType:    recvType,
		ContentType: contentType,
		Content:     content,
	}
	data := utils.InterfaceToJsonBytes(smr)
	url := fmt.Sprintf("%s/bot/send?token=%s", API_BASE_URL, o.Token)
	resp, err := utils.HttpPost(url, data)
	var basicResp BasicResponse
	json.Unmarshal(resp.Body(), &basicResp)
	return basicResp, err
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
	url := fmt.Sprintf("%s/bot/batch_send?token=%s", API_BASE_URL, o.Token)
	utils.HttpPost(url, data)
}

/**
 * @description: 单条，编辑文本消息
 * @param {EditTextMessage} newMessage
 */
func (o *OpenApi) EditTextMessage(newMessage EditTextMessage) (BasicResponse, error) {
	contentType := "text"
	content := map[string]interface{}{"text": newMessage.Text, "buttons": newMessage.Buttons}
	return o.EditMessage(newMessage.RecvId, newMessage.RecvType, newMessage.MsgId, contentType, content)
}

/**
 * @description: 单条，编辑单条消息
 */
func (o *OpenApi) EditMessage(recvId string, recvType string, msgId string, contentType string, content map[string]interface{}) (BasicResponse, error) {
	var smr = EditMessageRequest{
		MsgId:       msgId,
		RecvId:      recvId,
		RecvType:    recvType,
		Content:     content,
		ContentType: contentType,
	}
	data := utils.InterfaceToJsonBytes(smr)
	url := fmt.Sprintf("%s/bot/edit?token=%s", API_BASE_URL, o.Token)
	resp, err := utils.HttpPost(url, data)
	var basicResp BasicResponse
	json.Unmarshal(resp.Body(), &basicResp)
	return basicResp, err
}
