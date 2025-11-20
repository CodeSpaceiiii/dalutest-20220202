// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/darabonba-openapi/v2/websocketUtils"
	"github.com/alibabacloud-go/tea/dara"
)

type iWebsocketAwapDemoApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WebsocketAwapDemoApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WebsocketAwapDemoApiResponse
	GetStatusCode() *int32
	SetBody(v *WebsocketAwapDemoApiResponseBody) *WebsocketAwapDemoApiResponse
	GetBody() *WebsocketAwapDemoApiResponseBody
}

// 实现的枚举类
const (
	// Upstream event types (client -> server)
	WebsocketAwapDemoApiMessageTypeUpstreamTextEvent    dara.AwapMessageType = "UpstreamTextEvent"
	WebsocketAwapDemoApiMessageTypeUpstreamBinaryEvent  dara.AwapMessageType = "UpstreamBinaryEvent"
	WebsocketAwapDemoApiMessageTypeAckRequiredTextEvent dara.AwapMessageType = "AckRequiredTextEvent"

	// Downstream event types (server -> client)
	WebsocketAwapDemoApiMessageTypeMessageReceiveEvent   dara.AwapMessageType = "MessageReceiveEvent"
	WebsocketAwapDemoApiMessageTypeDownstreamTextEvent   dara.AwapMessageType = "DownstreamTextEvent"
	WebsocketAwapDemoApiMessageTypeDownstreamBinaryEvent dara.AwapMessageType = "DownstreamBinaryEvent"

	// Control message types (server -> client)
	WebsocketAwapDemoApiMessageTypeReconnect dara.GeneralMessageType = "RECONNECT" // Server-initiated graceful reconnection
)

type WebsocketAwapDemoApiResponse struct {
	WebsockeWebSocketClient *websocketUtils.WebSocketClient `json:"websocketClient,omitempty" xml:"websocketClient,omitempty"`
}
