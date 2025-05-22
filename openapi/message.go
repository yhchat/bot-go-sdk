/*
 * @Date: 2022-07-25 08:29:52
 * @LastEditTime: 2023-03-30 12:10:08
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package openapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
	content := map[string]interface{}{"text": message.Text, "buttons": message.Buttons}
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
 * @description: 流式消息
 */
func (o *OpenApi) StreamMessageWriter(recvId string, recvType string, contentType string) (*StreamWriter, error) {
	pr, pw := io.Pipe()
	url := fmt.Sprintf("%s/bot/send-stream?token=%s&recvId=%s&recvType=%s&contentType=%s",
		API_BASE_URL, o.Token, recvId, recvType, contentType)

	req, err := http.NewRequest("POST", url, pr)
	if err != nil {
		return nil, err
	}

	respCh := make(chan []byte, 1)

	client := &http.Client{}
	go func() {
		defer close(respCh)

		resp, err := client.Do(req)
		if err != nil {
			pw.CloseWithError(err)
			respCh <- []byte(err.Error())
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		respCh <- body
	}()

	return &StreamWriter{
		pr:         pr,
		pw:         pw,
		client:     client,
		req:        req,
		responseCh: respCh,
	}, nil
}

// GetResponse 获取服务端响应（阻塞直到收到响应或出错）
func (w *StreamWriter) GetResponse() ([]byte, error) {
	resp := <-w.responseCh
	if len(resp) == 0 {
		return nil, io.ErrUnexpectedEOF
	}
	return resp, nil
}

// Write 写入数据（线程安全）
func (w *StreamWriter) Write(data []byte) error {
	_, err := w.pw.Write(data)
	return err
}

// Close 关闭写入器
func (w *StreamWriter) Close() error {
	return w.pw.Close()
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
	content := map[string]interface{}{"text": message.Text, "buttons": message.Buttons}
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
 * @description: 撤回消息
 * @param {RecallMessageRequest} message
 */
func (o *OpenApi) RecallMessage(message RecallMessageRequest) (BasicResponse, error) {
	data := utils.InterfaceToJsonBytes(message)
	url := fmt.Sprintf("%s/bot/recall?token=%s", API_BASE_URL, o.Token)
	resp, err := utils.HttpPost(url, data)
	var basicResp BasicResponse
	json.Unmarshal(resp.Body(), &basicResp)
	return basicResp, err
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

/**
 * @description: 机器人看板设置接口
 * 机器人看板类型contentType取值: text、markdown、html
 * expireTime: 看板过期时间，11位时间戳。比如过期时间为10分钟，则expireTime为当前时间戳+600秒（time.Now().Unix() + 600）。
 */
func (o *OpenApi) SetBotBoard(chatId string, chatType string, memberId string, contentType string, content string, expireTime uint64) (BasicResponse, error) {
	var smr = BotBoardRequest{
		ChatId:      chatId,
		ChatType:    chatType,
		MemberId:    memberId,
		Content:     content,
		ContentType: contentType,
		ExpireTime:  expireTime,
	}
	data := utils.InterfaceToJsonBytes(smr)
	url := fmt.Sprintf("%s/bot/board?token=%s", API_BASE_URL, o.Token)
	resp, err := utils.HttpPost(url, data)
	var basicResp BasicResponse
	json.Unmarshal(resp.Body(), &basicResp)
	return basicResp, err
}

/**
 * @description: 机器人看板批量设置接口
 * 机器人看板类型contentType取值: text、markdown、html
 * expireTime: 看板过期时间，11位时间戳。比如过期时间为10分钟，则expireTime为当前时间戳+600秒（time.Now().Unix() + 600）。
 */
func (o *OpenApi) SetBotBoardAll(contentType string, content string, expireTime uint64) (BasicResponse, error) {
	var smr = BotBoardRequest{
		Content:     content,
		ContentType: contentType,
		ExpireTime:  expireTime,
	}
	data := utils.InterfaceToJsonBytes(smr)
	url := fmt.Sprintf("%s/bot/board-all?token=%s", API_BASE_URL, o.Token)
	resp, err := utils.HttpPost(url, data)
	var basicResp BasicResponse
	json.Unmarshal(resp.Body(), &basicResp)
	return basicResp, err
}

/**
 * @description: 机器人看板取消接口
 */
func (o *OpenApi) DismissBotBoard(chatId string, chatType string, memberId string) (BasicResponse, error) {
	var smr = BotBoardRequest{
		ChatId:   chatId,
		ChatType: chatType,
		MemberId: memberId,
	}
	data := utils.InterfaceToJsonBytes(smr)
	url := fmt.Sprintf("%s/bot/board-dismiss?token=%s", API_BASE_URL, o.Token)
	resp, err := utils.HttpPost(url, data)
	var basicResp BasicResponse
	json.Unmarshal(resp.Body(), &basicResp)
	return basicResp, err
}

/**
 * @description: 机器人看板批量取消接口
 */
func (o *OpenApi) DismissBotBoardAll() (BasicResponse, error) {
	url := fmt.Sprintf("%s/bot/board-all-dismiss?token=%s", API_BASE_URL, o.Token)
	resp, err := utils.HttpPost(url, []byte{})
	var basicResp BasicResponse
	json.Unmarshal(resp.Body(), &basicResp)
	return basicResp, err
}
