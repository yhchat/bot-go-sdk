# 云湖官方Golang SDK
## 使用
```
go install github.com/yhchat/bot-go-sdk
```

## 使用示例
example目录中包括两个示例项目

|项目|说明|
|--|--|
|basic|建议参考，支持消息发送和接收|
|custom_router|使用自定义路由的示例代码项目|

### basic示例
默认端口号为8805，可以随意修改

将示例代码中 SendTextMessage 等方法中的token替换为云湖控制台提供的Token字符串，即可发送消息到对应机器人。

启动后，可以将 http://ip:8805/sub 填入到云湖控制台，并打开云湖控制台的事件订阅按钮，程序即可收到云湖消息。



## 官方文档
[https://www.yhchat.com/document/1-3](https://www.yhchat.com/document/1-3)