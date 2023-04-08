package subscription

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yhchat/bot-go-sdk/utils"

	"github.com/gin-gonic/gin"
)

type Subscription struct {
	Port                 int
	OnGroupJoin          func(GroupJoinEvent)
	OnGroupLeave         func(GroupLeaveEvent)
	OnMessageNormal      func(MessageEvent)
	OnMessageInstruction func(MessageEvent)
	OnBotFollowed        func(BotFollowedEvent)
	OnButtonReportInline func(ButtonReportInlineEvent)
}

func NewSubscription(port int) *Subscription {
	return &Subscription{Port: port}
}

func (s *Subscription) Start() {
	router := gin.Default()
	router.POST("/sub", func(c *gin.Context) {
		var sr SubScriptionResp
		if err := c.BindJSON(&sr); err != nil {
			return
		}
		s.Parse(sr)

	})

	//测试使用
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func (s *Subscription) Parse(sr SubScriptionResp) {
	log.Println(sr)
	header := sr.Header
	event := sr.Event

	if header.EventType == "group.join" {
		groupJoinEvent := GroupJoinEvent{
			AvatarUrl: event["avatarUrl"].(string),
			ChatId:    event["chatId"].(string),
			ChatType:  event["chatType"].(string),
			Nickname:  event["nickname"].(string),
			Time:      utils.InterfaceToInt64(event["time"]),
			UserId:    event["userId"].(string),
		}
		if s.OnGroupJoin != nil {
			s.OnGroupJoin(groupJoinEvent)
		}
		return
	}

	if header.EventType == "group.leave" {
		groupLeaveEvent := GroupLeaveEvent{
			AvatarUrl: event["avatarUrl"].(string),
			ChatId:    event["chatId"].(string),
			ChatType:  event["chatType"].(string),
			Nickname:  event["nickname"].(string),
			Time:      utils.InterfaceToInt64(event["time"]),
			UserId:    event["userId"].(string),
		}
		if s.OnGroupLeave != nil {
			s.OnGroupLeave(groupLeaveEvent)
		}
		return
	}

	/**
	 * @desc: 普通消息
	 * event消息内容如下
	 */
	if header.EventType == "message.receive.normal" {
		chat := event["chat"].(map[string]interface{})
		sender := event["sender"].(map[string]interface{})
		message := event["message"].(map[string]interface{})
		content, ok := message["content"].(map[string]interface{})
		if !ok {
			log.Println("普通消息，content字段值不合法。")
			return
		}
		messageEvent := MessageEvent{
			Chat: MessageEventChat{
				ChatId:   utils.InterfaceToString(chat["chatId"]),
				ChatType: utils.InterfaceToString(chat["chatType"]),
			},
			Sender: MessageEventSender{
				SenderId:        utils.InterfaceToString(sender["senderId"]),
				SenderType:      utils.InterfaceToString(sender["senderType"]),
				SenderUserLevel: utils.InterfaceToString(sender["senderUserLevel"]),
				SenderNickname:  utils.InterfaceToString(sender["senderNickname"]),
			},
			Message: MessageEventMessage{
				MsgId:           utils.InterfaceToString(message["msgId"]),
				ParentId:        utils.InterfaceToString(message["parentId"]),
				ContentType:     utils.InterfaceToString(message["contentType"]),
				Content:         content,
				InstructionId:   utils.InterfaceToInt64(message["instructionId"]),
				InstructionName: utils.InterfaceToString(message["instructionName"]),
			},
		}
		if s.OnMessageNormal != nil {
			s.OnMessageNormal(messageEvent)
		}
		return
	}

	/**
	 * @desc: 指令消息
	 * event消息内容如下
	 */
	if header.EventType == "message.receive.instruction" {
		chat := event["chat"].(map[string]interface{})
		sender := event["sender"].(map[string]interface{})
		message := event["message"].(map[string]interface{})
		content, ok := message["content"].(map[string]interface{})
		if !ok {
			log.Println("指令消息，content字段值不合法。")
			return
		}
		messageEvent := MessageEvent{
			Chat: MessageEventChat{
				ChatId:   utils.InterfaceToString(chat["chatId"]),
				ChatType: utils.InterfaceToString(chat["chatType"]),
			},
			Sender: MessageEventSender{
				SenderId:        utils.InterfaceToString(sender["senderId"]),
				SenderType:      utils.InterfaceToString(sender["senderType"]),
				SenderUserLevel: utils.InterfaceToString(sender["senderUserLevel"]),
				SenderNickname:  utils.InterfaceToString(sender["senderNickname"]),
			},
			Message: MessageEventMessage{
				MsgId:           utils.InterfaceToString(message["msgId"]),
				ParentId:        utils.InterfaceToString(message["parentId"]),
				ContentType:     utils.InterfaceToString(message["contentType"]),
				Content:         content,
				InstructionId:   utils.InterfaceToInt64(message["instructionId"]),
				InstructionName: utils.InterfaceToString(message["instructionName"]),
			},
		}
		if s.OnMessageInstruction != nil {
			s.OnMessageInstruction(messageEvent)
		}

		return
	}

	/**
	 * @desc: 关注机器人事件
	 * event消息内容如下
	 * map[avatarUrl:xxxx chatId:xxxx chatType:bot nickname:xxxx time:1658835054923 userId:123456]
	 */
	if header.EventType == "bot.followed" {
		botFollowedEvent := BotFollowedEvent{
			AvatarUrl: utils.InterfaceToString(event["avatarUrl"]),
			ChatId:    utils.InterfaceToString(event["chatId"]),
			ChatType:  utils.InterfaceToString(event["chatType"]),
			Nickname:  utils.InterfaceToString(event["nickname"]),
			Time:      utils.InterfaceToInt64(event["time"]),
			UserId:    utils.InterfaceToString(event["userId"]),
		}
		if s.OnBotFollowed != nil {
			s.OnBotFollowed(botFollowedEvent)
		}
		return
	}

	/**
	 * @desc: 取消关注机器人
	 * event消息内容如下
	 * map[avatarUrl:xxxxx chatId:xxxxx chatType:bot nickname:xxxxx  time:1658835054923 userId:123456]
	 */
	if header.EventType == "bot.unfollowed" {
		fmt.Println(event)
	}

	/**
	 * @desc: 消息下按钮点击回调事件
	 * event消息内容如下
	 * map[msgId:xxx recvId:xxx recvType:bot senderId:xxx time:1.679899979517e+12 value:xxx]}
	 */
	if header.EventType == "button.report.inline" {
		buttonReportInlineEvent := ButtonReportInlineEvent{
			MsgId:    utils.InterfaceToString(event["msgId"]),
			RecvId:   utils.InterfaceToString(event["recvId"]),
			RecvType: utils.InterfaceToString(event["recvType"]),
			SenderId: utils.InterfaceToString(event["senderId"]),
			Value:    utils.InterfaceToString(event["value"]),
			Time:     utils.InterfaceToInt64(event["time"]),
		}
		if s.OnButtonReportInline != nil {
			s.OnButtonReportInline(buttonReportInlineEvent)
		}
		return
	}
}
