/*
 * @Date: 2023-01-06 10:45:46
 * @LastEditTime: 2023-04-17 13:27:00
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package subscription

type SubScriptionResp struct {
	Version string                 `json:"version"`
	Header  SubScriptionRespHeader `json:"header"`
	Event   map[string]interface{} `json:"event"`
}
type SubScriptionRespHeader struct {
	EventId   string `json:"eventId"`
	EventTime int64  `json:"eventTime"`
	EventType string `json:"eventType"`
}

// GroupJoin
type GroupJoinEvent struct {
	AvatarUrl string `json:"avatarUrl"`
	ChatId    string `json:"chatId"`
	ChatType  string `json:"chatType"`
	Nickname  string `json:"nickname"`
	Time      int64  `json:"time"`
	UserId    string `json:"userId"`
}

// GroupLeave
type GroupLeaveEvent struct {
	AvatarUrl string `json:"avatarUrl"`
	ChatId    string `json:"chatId"`
	ChatType  string `json:"chatType"`
	Nickname  string `json:"nickname"`
	Time      int64  `json:"time"`
	UserId    string `json:"userId"`
}

// BotFollowed
type BotFollowedEvent struct {
	AvatarUrl string `json:"avatarUrl"`
	ChatId    string `json:"chatId"`
	ChatType  string `json:"chatType"`
	Nickname  string `json:"nickname"`
	Time      int64  `json:"time"`
	UserId    string `json:"userId"`
}

// BotUnfollowed
type BotUnfollowedEvent struct {
	AvatarUrl string `json:"avatarUrl"`
	ChatId    string `json:"chatId"`
	ChatType  string `json:"chatType"`
	Nickname  string `json:"nickname"`
	Time      int64  `json:"time"`
	UserId    string `json:"userId"`
}

// 按钮事件汇报
type ButtonReportInlineEvent struct {
	MsgId    string `json:"msgId"`
	RecvId   string `json:"recvId"`
	RecvType string `json:"recvType"`
	SenderId string `json:"senderId"`
	Value    string `json:"value"`
	Time     int64  `json:"time"`
}

// Message
type MessageEvent struct {
	Chat    MessageEventChat    `json:"chat"`
	Sender  MessageEventSender  `json:"sender"`
	Message MessageEventMessage `json:"message"`
}
type MessageEventChat struct {
	ChatId   string `json:"chatId"`
	ChatType string `json:"chatType"`
}

type MessageEventMessage struct {
	MsgId       string                 `json:"msgId"`
	ParentId    string                 `json:"parentId"`
	ContentType string                 `json:"contentType"`
	Content     map[string]interface{} `json:"content"`
	// Deprecated: 这个变量将在之后的版本移除，请用 CommandId 替代.
	InstructionId int64 `json:"instructionId"`
	// Deprecated: 这个变量将在之后的版本移除，请用 CommandName 替代.
	InstructionName string `json:"instructionName"`
	CommandId       int64  `json:"commandId"`
	CommandName     string `json:"commandName"`
}

type MessageEventSender struct {
	SenderId        string `json:"senderId"`
	SenderType      string `json:"senderType"`
	SenderUserLevel string `json:"senderUserLevel"`
	SenderNickname  string `json:"senderNickname"`
}
