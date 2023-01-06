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

//GroupJoin
type GroupJoinEvent struct {
	AvatarUrl string `json:"avatarUrl"`
	ChatId    string `json:"chatId"`
	ChatType  string `json:"chatType"`
	Nickname  string `json:"nickname"`
	Time      int64  `json:"time"`
	UserId    string `json:"userId"`
}

//GroupLeave
type GroupLeaveEvent struct {
	AvatarUrl string `json:"avatarUrl"`
	ChatId    string `json:"chatId"`
	ChatType  string `json:"chatType"`
	Nickname  string `json:"nickname"`
	Time      int64  `json:"time"`
	UserId    string `json:"userId"`
}

//BotFollowed
type BotFollowedEvent struct {
	AvatarUrl string `json:"avatarUrl"`
	ChatId    string `json:"chatId"`
	ChatType  string `json:"chatType"`
	Nickname  string `json:"nickname"`
	Time      int64  `json:"time"`
	UserId    string `json:"userId"`
}

//Message
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
	MsgId           string                 `json:"msgId"`
	ParentId        string                 `json:"parentId"`
	ContentType     string                 `json:"contentType"`
	Content         map[string]interface{} `json:"content"`
	InstructionId   int64                  `json:"instructionId"`
	InstructionName string                 `json:"instructionName"`
}

type MessageEventSender struct {
	SenderId        string `json:"senderId"`
	SenderType      string `json:"senderType"`
	SenderUserLevel string `json:"senderUserLevel"`
}
