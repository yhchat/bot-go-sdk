package subscription

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/assert/v2"
)

// 测试解析普通消息
func TestParseMessageNormal(t *testing.T) {
	jsonStr := `{
		"version":"1.0",
		"header":{
			"eventId":"ebaf7a6e2eca4c508c439cd213981f76",
			"eventType":"message.receive.normal",
			"eventTime":1666962601669
		},
		"event":{
			"sender":{
				"senderId":"7058262",
				"senderType":"user",
				"senderUserLevel":"member"
			},
			"chat":{
				"chatId":"84660887",
				"chatType":"bot"
			},
			"message":{
				"msgId":"706f3a24f7ea476196caafa84d2ba34d",
				"parentId":"",
				"sendTime":1666962601658,
				"chatId":"84660887",
				"chatType":"bot",
				"contentType":"text",
				"content":{
					"text":"这里是消息text"
				},
				"instructionId":0,
				"instructionName":"",
				"commandId":1,
				"commandName":"commandName"
			}
		}
	}`
	var sr SubScriptionResp
	if err := json.Unmarshal([]byte(jsonStr), &sr); err != nil {
		panic(err)
	}
	subscription := NewSubscription(0)
	subscription.OnMessageNormal = func(event MessageEvent) {
		content := event.Message.Content
		t.Log(content)
		text := content["text"]
		t.Log(text)

		assert.Equal(t, text, "这里是消息text")
	}
	subscription.Parse(sr)
}
