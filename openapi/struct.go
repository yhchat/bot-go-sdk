/*
 * @Date: 2022-07-25 08:14:34
 * @LastEditTime: 2023-03-30 12:10:17
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package openapi

/**
 * 单条，发送消息结构体
 */
type SendMessageRequest struct {
	RecvId      string                 `json:"recvId"`
	RecvType    string                 `json:"recvType"`
	ContentType string                 `json:"contentType"`
	Content     map[string]interface{} `json:"content"`
}

/**
 * 批量，发送消息结构体
 */
type BatchSendMessageRequest struct {
	RecvIds     []string               `json:"recvIds"`
	RecvType    string                 `json:"recvType"`
	ContentType string                 `json:"contentType"`
	Content     map[string]interface{} `json:"content"`
}

/**
 * 单条，用于发送文本消息
 * RecvId 为接收消息的用户ID或者群ID
 * RecvType 取值group、user
 * Buttons 取值[]Button或[][]Button
 */
type TextMessage struct {
	RecvId   string      `json:"recvId"`
	RecvType string      `json:"recvType"`
	Text     string      `json:"text"`
	Buttons  interface{} `json:"buttons"`
}

/**
 * 批量，用于发送文本消息
 * RecvId 为接收消息的用户ID或者群ID
 * RecvType 取值group、user
 * Buttons 取值[]Button或[][]Button
 */
type BatchTextMessage struct {
	RecvIds  []string    `json:"recvIds"`
	RecvType string      `json:"recvType"`
	Text     string      `json:"text"`
	Buttons  interface{} `json:"buttons"`
}

/**
 * 单条，用于发送Markdown消息
 * RecvId 为接收消息的用户ID或者群ID
 * RecvType 取值group、user
 */
type MarkdownMessage struct {
	RecvId   string      `json:"recvId"`
	RecvType string      `json:"recvType"`
	Text     string      `json:"text"`
	Buttons  interface{} `json:"buttons"`
}

/**
 * 批量，用于发送Markdown消息
 * RecvId 为接收消息的用户ID或者群ID
 * RecvType 取值group、user
 */
type BatchMarkdownMessage struct {
	RecvIds  []string    `json:"recvIds"`
	RecvType string      `json:"recvType"`
	Text     string      `json:"text"`
	Buttons  interface{} `json:"buttons"`
}

/**
 * text         按钮上的文字
 * actionType   1: 跳转URL; 2: 复制; 3: 点击汇报
 * url          当actionType为1时使用
 * value        当actionType为2时，该值会复制到剪贴板; 当actionType为3时，该值会发送给订阅端
 */
type Button struct {
	Text       string `json:"text"`
	ActionType int    `json:"actionType"`
	Url        string `json:"url"`
	Value      string `json:"value"`
}

/**
 * 单条，用于编辑文本消息
 * RecvId 为消息的用户ID或者群ID
 * RecvType 取值group、user
 */
type EditTextMessage struct {
	MsgId    string      `json:"msgId"`
	RecvId   string      `json:"recvId"`
	RecvType string      `json:"recvType"`
	Text     string      `json:"text"`
	Buttons  interface{} `json:"buttons"`
}

/**
 * 单条，编辑消息结构体
 */
type EditMessageRequest struct {
	MsgId       string                 `json:"msgId"`
	RecvId      string                 `json:"recvId"`
	RecvType    string                 `json:"recvType"`
	ContentType string                 `json:"contentType"`
	Content     map[string]interface{} `json:"content"`
}

/////////////////////////////////响应内容///////////////////////////////////
/**
 * 发送消息返回消息对象
 */
type BasicResponse struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
