package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/CodeSpaceiiii/dalutest-20220202/client"
	"github.com/CodeSpaceiiii/dalutest-20220202/client/WebsocketAwapDemoApi"
	"github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
	credential "github.com/aliyun/credentials-go/credentials"
)

// GeneralWebSocketHandler 实现 GeneralWebSocketHandler 接口，用于处理 WebSocket General 协议消息
type GeneralWebSocketHandler struct {
	dara.AbstractGeneralWebSocketHandler
	ConnectedCalled      bool
	MessageReceivedCount int
	ErrorCount           int
	ClosedCalled         bool
	LastTextMessage      *dara.GeneralMessage
	LastBinaryMessage    []byte
}

func (h *GeneralWebSocketHandler) AfterConnectionEstablished(session *dara.WebSocketSessionInfo) error {
	h.ConnectedCalled = true
	fmt.Printf("[Handler] Connection established. Session ID: %s, Remote: %s\n", session.SessionID, session.RemoteAddr)
	return nil
}

func (h *GeneralWebSocketHandler) HandleGeneralMessage(session *dara.WebSocketSessionInfo, message *dara.GeneralMessage) error {
	h.MessageReceivedCount++
	if message.Format == dara.GeneralMessageFormatText {
		h.LastTextMessage = message
		fmt.Printf("[Handler] Received text message. Body: %v\n", message.Body)
	} else if message.Format == dara.GeneralMessageFormatBinary {
		h.LastBinaryMessage = message.Body.([]byte)
		fmt.Printf("[Handler] Received binary message. Body: %v\n", message.Body)
	}

	// 尝试解析 body 为具体的类型
	if bodyBytes, ok := message.Body.([]byte); ok {
		fmt.Printf("[Handler] Body as bytes: %s\n", string(bodyBytes))
	} else if bodyStr, ok := message.Body.(string); ok {
		fmt.Printf("[Handler] Body as string: %s\n", bodyStr)
	} else {
		bodyJSON, _ := json.Marshal(message.Body)
		fmt.Printf("[Handler] Body as JSON: %s\n", string(bodyJSON))
	}

	return nil
}

func (h *GeneralWebSocketHandler) HandleRawMessage(session *dara.WebSocketSessionInfo, message *dara.WebSocketMessage) error {
	// 这个方法通常不会被调用，因为 GeneralWebSocketHandler 会优先使用 HandleGeneralMessage
	// 如果用户未自定义HandleGeneralMessage, 这里会返回原始messsage 供用户处理
	fmt.Printf("[Handler] HandleRawMessage called. Type: %d, Size: %d bytes, message: %v\n", message.Type, len(message.Payload), message)
	return nil
}

func (h *GeneralWebSocketHandler) HandleError(session *dara.WebSocketSessionInfo, err error) error {
	h.ErrorCount++
	fmt.Printf("[Handler] Error occurred: %v\n", err)
	return nil
}

func (h *GeneralWebSocketHandler) AfterConnectionClosed(session *dara.WebSocketSessionInfo, code int, reason string) error {
	h.ClosedCalled = true
	fmt.Printf("[Handler] Connection closed. Code: %d, Reason: %s\n", code, reason)
	return nil
}

func (h *GeneralWebSocketHandler) SupportsPartialMessages() bool {
	return false
}

// awap

// AwapWebSocketHandler 实现 AWAP WebSocket Handler，用于处理 WebSocket AWAP 协议消息
type AwapWebSocketHandler struct {
	dara.AbstractAwapWebSocketHandler
	ConnectedCalled           bool
	MessageReceivedCount      int
	ErrorCount                int
	ClosedCalled              bool
	LastDownstreamTextEvent   *WebsocketAwapDemoApi.DownstreamTextEvent
	LastDownstreamBinaryEvent []byte
	LastMessageReceiveEvent   bool
}

func (h *AwapWebSocketHandler) AfterConnectionEstablished(session *dara.WebSocketSessionInfo) error {
	h.ConnectedCalled = true
	fmt.Printf("[AWAP Handler] Connection established. Session ID: %s, Remote: %s\n", session.SessionID, session.RemoteAddr)
	return nil
}

func (h *AwapWebSocketHandler) HandleAwapMessage(session *dara.WebSocketSessionInfo, message *dara.AwapMessage) error {
	h.MessageReceivedCount++

	// 从 message 中提取 messageType 和 data
	messageType := message.Type
	data := message.Payload

	fmt.Printf("[AWAP Handler] Received AWAP message. Type: %s\n", messageType)

	// 根据消息类型处理不同的数据
	switch messageType {
	case WebsocketAwapDemoApi.DownstreamTextEvent_MessageType:
		if eventData, ok := data.(*WebsocketAwapDemoApi.DownstreamTextEvent); ok {
			h.LastDownstreamTextEvent = eventData
			eventJSON, _ := json.Marshal(eventData)
			fmt.Printf("[AWAP Handler] DownstreamTextEvent data: %s\n", string(eventJSON))
		} else {
			// 尝试从 map 转换
			if dataMap, ok := data.(map[string]interface{}); ok {
				eventJSON, _ := json.Marshal(dataMap)
				fmt.Printf("[AWAP Handler] DownstreamTextEvent data (as map): %s\n", string(eventJSON))
			} else {
				fmt.Printf("[AWAP Handler] DownstreamTextEvent data type: %T, value: %v\n", data, data)
			}
		}

	case WebsocketAwapDemoApi.DownstreamBinaryEvent_MessageType:
		if binaryData, ok := data.([]byte); ok {
			h.LastDownstreamBinaryEvent = binaryData
			fmt.Printf("[AWAP Handler] DownstreamBinaryEvent. Size: %d bytes\n", len(binaryData))
		} else {
			fmt.Printf("[AWAP Handler] DownstreamBinaryEvent data type: %T\n", data)
		}

	case WebsocketAwapDemoApi.ReceiveEvent_MessageType:
		h.LastMessageReceiveEvent = true
		fmt.Printf("[AWAP Handler] MessageReceiveEvent received\n")
		dataJSON, _ := json.Marshal(data)
		fmt.Printf("data: %v\n", string(dataJSON))

	default:
		dataJSON, _ := json.Marshal(data)
		fmt.Printf("[AWAP Handler] Unknown message type: %s, data: %s\n", messageType, string(dataJSON))
	}

	return nil
}

func (h *AwapWebSocketHandler) HandleRawMessage(session *dara.WebSocketSessionInfo, message *dara.WebSocketMessage) error {
	// HandleAwapMessage, 这里会返回原始messsage 供用户处理
	fmt.Printf("[AWAP Handler] HandleRawMessage called. Type: %d, Size: %d bytes, message: %v\n", message.Type, len(message.Payload), message)
	return nil
}

func (h *AwapWebSocketHandler) HandleError(session *dara.WebSocketSessionInfo, err error) error {
	h.ErrorCount++
	fmt.Printf("[AWAP Handler] Error occurred: %v\n", err)
	return nil
}

func (h *AwapWebSocketHandler) AfterConnectionClosed(session *dara.WebSocketSessionInfo, code int, reason string) error {
	h.ClosedCalled = true
	fmt.Printf("[AWAP Handler] Connection closed. Code: %d, Reason: %s\n", code, reason)
	return nil
}

func (h *AwapWebSocketHandler) SupportsPartialMessages() bool {
	return false
}

func main() {
	// 初始化客户端配置
	credentials, _err := credential.NewCredential(nil)
	if _err != nil {
		log.Fatalf("Failed to create credential: %v", _err)
	}
	config := &utils.Config{
		Credential: credentials,
		Endpoint:   dara.String("dalutest-pre.aliyuncs.com"),
	}

	// 创建客户端
	apiClient, err := client.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// 创建 WebSocket handler
	handler := &GeneralWebSocketHandler{}

	// 创建请求
	request := &client.WebsocketGeneralDemoApiRequest{
		Auth:       dara.String("test-auth"),
		RespBody:   dara.String("test-response-body"),
		RespStatus: dara.String("200"),
		SleepMs:    dara.String("0"),
	}

	// 创建运行时选项，设置 WebSocketHandler
	runtime := &dara.RuntimeOptions{
		WebSocketHandler:         handler,
		ConnectTimeout:           dara.Int(5000),
		ReadTimeout:              dara.Int(30000),
		WebSocketPingInterval:    dara.Int(0),      // 禁用 ping，用于测试
		WebSocketEnableReconnect: dara.Bool(false), // 禁用自动重连，用于测试
	}

	// 创建 headers
	headers := make(map[string]*string)

	// 调用 WebSocket API
	fmt.Println("Calling WebsocketGeneralDemoApiWithOptions...")
	response, err := apiClient.WebsocketGeneralDemoApiWithOptions(request, headers, runtime)
	if err != nil {
		log.Fatalf("Failed to call WebsocketGeneralDemoApiWithOptions: %v", err)
	}

	// 检查响应
	if response == nil {
		log.Fatal("Response is nil")
	}

	if response.WebSocketClient == nil {
		log.Fatal("WebSocketClient is nil")
	}

	wsClient := response.WebSocketClient

	// 检查连接状态
	if !wsClient.IsConnected() {
		log.Fatal("WebSocket client is not connected")
	}

	fmt.Println("WebSocket connection established successfully!")

	// 获取会话信息
	session := wsClient.GetSessionInfo()
	if session != nil {
		fmt.Printf("Session ID: %s\n", session.SessionID)
		fmt.Printf("Connected at: %s\n", session.ConnectedAt.Format(time.RFC3339))
		fmt.Printf("Remote address: %s\n", session.RemoteAddr)
		fmt.Printf("Local address: %s\n", session.LocalAddr)
	}

	// 等待一段时间以接收消息
	fmt.Println("Waiting for messages (1 seconds)...")
	time.Sleep(1 * time.Second)

	// 打印 handler 统计信息
	fmt.Printf("\n=== Handler Statistics ===\n")
	fmt.Printf("Connected called: %v\n", handler.ConnectedCalled)
	fmt.Printf("Messages received: %d\n", handler.MessageReceivedCount)
	fmt.Printf("Errors occurred: %d\n", handler.ErrorCount)
	fmt.Printf("Closed called: %v\n", handler.ClosedCalled)

	// 如果需要，可以发送一些测试消息
	if handler.ConnectedCalled {
		// 发送文本消息示例
		testMessage := &dara.GeneralMessage{
			Body: map[string]interface{}{
				"message":   "Hello from client",
				"timestamp": time.Now().Unix(),
			},
		}

		messageJSON, err := testMessage.ToJSON()
		if err != nil {
			log.Printf("Failed to marshal test message: %v", err)
		} else {
			if err := wsClient.SendGeneralTextMessage(string(messageJSON)); err != nil {
				log.Printf("Failed to send text message: %v", err)
			} else {
				fmt.Println("Sent test text message")
			}
		}

		// 再等待一段时间以接收响应
		time.Sleep(5 * time.Second)
	}

	// 关闭连接
	fmt.Println("Closing WebSocket connection...")
	if err := wsClient.Close(); err != nil {
		log.Printf("Error closing WebSocket: %v", err)
	}

	// 等待连接关闭完成
	time.Sleep(1 * time.Second)

	// 最终统计
	fmt.Printf("\n=== Final Statistics ===\n")
	fmt.Printf("Connected called: %v\n", handler.ConnectedCalled)
	fmt.Printf("Messages received: %d\n", handler.MessageReceivedCount)
	fmt.Printf("Errors occurred: %d\n", handler.ErrorCount)
	fmt.Printf("Closed called: %v\n", handler.ClosedCalled)

	fmt.Println("general Test completed!")
	testAwap(apiClient)
}

func testAwap(apiClient *client.Client) {
	// awap start
	fmt.Println("awap Test started!")
	// 创建 AWAP WebSocket handler
	handler := &AwapWebSocketHandler{}

	// 创建请求
	request := &client.WebsocketAwapDemoApiRequest{
		Auth:       dara.String("test-auth"),
		RespBody:   dara.String("test-response-body"),
		RespStatus: dara.String("200"),
		SleepMs:    dara.String("0"),
	}

	// 创建运行时选项，设置 WebSocketHandler
	runtime := &dara.RuntimeOptions{
		WebSocketHandler:         handler,
		ConnectTimeout:           dara.Int(5000),
		ReadTimeout:              dara.Int(30000),
		WebSocketPingInterval:    dara.Int(0),      // 禁用 ping，用于测试
		WebSocketEnableReconnect: dara.Bool(false), // 禁用自动重连，用于测试
	}

	// 创建 headers
	headers := make(map[string]*string)

	// 调用 WebSocket API
	fmt.Println("Calling WebsocketAwapDemoApiWithOptions...")
	response, err := apiClient.WebsocketAwapDemoApiWithOptions(request, headers, runtime)
	if err != nil {
		log.Fatalf("Failed to call WebsocketAwapDemoApiWithOptions: %v", err)
	}

	// 检查响应
	if response == nil {
		log.Fatal("Response is nil")
	}

	if response.WebSocketClient == nil {
		log.Fatal("WebSocketClient is nil")
	}

	wsClient := response.WebSocketClient

	// 检查连接状态
	if !wsClient.IsConnected() {
		log.Fatal("WebSocket client is not connected")
	}

	fmt.Println("WebSocket connection established successfully!")

	// 获取会话信息
	session := wsClient.GetSessionInfo()
	if session != nil {
		fmt.Printf("Session ID: %s\n", session.SessionID)
		fmt.Printf("Connected at: %s\n", session.ConnectedAt.Format(time.RFC3339))
		fmt.Printf("Remote address: %s\n", session.RemoteAddr)
		fmt.Printf("Local address: %s\n", session.LocalAddr)
	}

	// 等待一段时间以接收消息
	fmt.Println("Waiting for messages (1 second)...")
	time.Sleep(1 * time.Second)

	// 打印 handler 统计信息
	fmt.Printf("\n=== Handler Statistics ===\n")
	fmt.Printf("Connected called: %v\n", handler.ConnectedCalled)
	fmt.Printf("Messages received: %d\n", handler.MessageReceivedCount)
	fmt.Printf("Errors occurred: %d\n", handler.ErrorCount)
	fmt.Printf("Closed called: %v\n", handler.ClosedCalled)

	// 如果需要，可以发送一些测试消息
	if handler.ConnectedCalled {
		// 发送 UpstreamTextEvent 消息示例
		testEvent := &WebsocketAwapDemoApi.UpstreamTextEvent{
			Name: dara.String("test-event"),
			Object: &struct {
				StrField  *string   `json:"strField,omitempty" xml:"strField,omitempty"`
				IntField  *int32    `json:"intField,omitempty" xml:"intField,omitempty"`
				InnerList []*string `json:"innerList,omitempty" xml:"innerList,omitempty"`
			}{
				StrField:  dara.String("test-string"),
				IntField:  dara.Int32(123),
				InnerList: []*string{dara.String("item1"), dara.String("item2")},
			},
			List: []*struct {
				Key2     *bool   `json:"boolField,omitempty" xml:"boolField,omitempty"`
				StrField *string `json:"strField,omitempty" xml:"strField,omitempty"`
			}{
				{
					Key2:     dara.Bool(true),
					StrField: dara.String("list-item-1"),
				},
			},
			Map: map[string]interface{}{
				"key1": "value1",
				"key2": 42,
			},
		}

		if err := wsClient.SendRawAwapTextMessage(WebsocketAwapDemoApi.UpstreamTextEvent_MessageType, testEvent); err != nil {
			log.Printf("Failed to send AWAP message: %v", err)
		} else {
			fmt.Println("Sent test UpstreamTextEvent message")
		}
		fmt.Println("Note: AWAP message sending needs to be implemented based on actual API")
	}

	// 再等待一段时间以接收响应
	time.Sleep(5 * time.Second)

	// 关闭连接
	fmt.Println("Closing WebSocket connection...")
	if err := wsClient.Close(); err != nil {
		log.Printf("Error closing WebSocket: %v", err)
	}

	// 等待连接关闭完成
	time.Sleep(1 * time.Second)

	// 最终统计
	fmt.Printf("\n=== Final Statistics ===\n")
	fmt.Printf("Connected called: %v\n", handler.ConnectedCalled)
	fmt.Printf("Messages received: %d\n", handler.MessageReceivedCount)
	fmt.Printf("Errors occurred: %d\n", handler.ErrorCount)
	fmt.Printf("Closed called: %v\n", handler.ClosedCalled)

	if handler.LastDownstreamTextEvent != nil {
		eventJSON, _ := json.Marshal(handler.LastDownstreamTextEvent)
		fmt.Printf("Last DownstreamTextEvent: %s\n", string(eventJSON))
	}

	fmt.Println("AWAP test completed!")
}
