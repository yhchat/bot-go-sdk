/*
 * @Date: 2022-07-25 08:14:34
 * @LastEditTime: 2023-03-30 12:10:17
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package openapi

import (
	"io"
	"net/http"
)

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
 * 用于撤回消息
 * msgId 为消息ID
 * chatId 为消息的用户ID或者群ID
 * chatType 取值group、user
 */
type RecallMessageRequest struct {
	MsgId    string `json:"msgId"`
	ChatId   string `json:"chatId"`
	ChatType string `json:"chatType"`
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

/**
 * 机器人看板展示结构体
 */
type BotBoardRequest struct {
	MemberId    string `json:"memberId"`
	ChatId      string `json:"chatId"`
	ChatType    string `json:"chatType"`
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
	ExpireTime  uint64 `json:"expireTime"`
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

// 流式消息 StreamWriter 封装结构
type StreamWriter struct {
	pr         *io.PipeReader
	pw         *io.PipeWriter
	client     *http.Client
	req        *http.Request
	responseCh chan []byte // 用于接收响应
}
