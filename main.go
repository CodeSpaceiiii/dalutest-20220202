package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/CodeSpaceiiii/dalutest-20220202/client"
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

func (h *GeneralWebSocketHandler) HandleGeneralTextMessage(session *dara.WebSocketSessionInfo, message *dara.GeneralMessage) error {
	h.MessageReceivedCount++
	h.LastTextMessage = message
	fmt.Printf("[Handler] Received text message. Body: %v\n", message.Body)

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

func (h *GeneralWebSocketHandler) HandleGeneralBinaryMessage(session *dara.WebSocketSessionInfo, data []byte) error {
	h.MessageReceivedCount++
	h.LastBinaryMessage = data
	fmt.Printf("[Handler] Received binary message. Size: %d bytes\n", len(data))
	return nil
}

func (h *GeneralWebSocketHandler) HandleGeneralIncomingMessage(session *dara.WebSocketSessionInfo, message *dara.GeneralIncomingMessage) error {
	fmt.Printf("[Handler] HandleGeneralIncomingMessage called. IsBinary: %v\n", message.IsBinary)
	return nil
}

func (h *GeneralWebSocketHandler) HandleRawMessage(session *dara.WebSocketSessionInfo, message *dara.WebSocketMessage) error {
	// 这个方法通常不会被调用，因为 GeneralWebSocketHandler 会优先使用 HandleGeneralTextMessage/HandleGeneralBinaryMessage
	// 但如果消息无法解析为 General 格式，会回退到这里
	fmt.Printf("[Handler] HandleRawMessage called. Type: %d, Size: %d bytes\n", message.Type, len(message.Payload))
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

	wsClient := response.WebSocketClient.GetWebSocketClient()

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
	ctx := context.Background()
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
			if err := wsClient.SendText(ctx, string(messageJSON)); err != nil {
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

	fmt.Println("Test completed!")
}
