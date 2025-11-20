// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/darabonba-openapi/v2/websocketutils"
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
	WebSocketClient *websocketutils.WebSocketClient `json:"websocketClient,omitempty" xml:"websocketClient,omitempty"`
}

// WebsocketAwapDemoApiDataUpstreamTextEvent represents the UpstreamTextEvent input event
type WebsocketAwapDemoApiDataUpstreamTextEvent struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Object *struct {
		StrField  *string   `json:"strField,omitempty" xml:"strField,omitempty"`
		IntField  *int32    `json:"intField,omitempty" xml:"intField,omitempty"`
		InnerList []*string `json:"innerList,omitempty" xml:"innerList,omitempty"`
	} `json:"object,omitempty" xml:"object,omitempty"`
	List []*struct {
		Key2     *bool   `json:"boolField,omitempty" xml:"boolField,omitempty"`
		StrField *string `json:"strField,omitempty" xml:"strField,omitempty"`
	} `json:"list,omitempty" xml:"list,omitempty"`
	Map map[string]interface{} `json:"map,omitempty" xml:"map,omitempty"`
}

func (s WebsocketAwapDemoApiDataUpstreamTextEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketAwapDemoApiDataUpstreamTextEvent) GoString() string {
	return s.String()
}

func (s *WebsocketAwapDemoApiDataUpstreamTextEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketAwapDemoApiDataAckRequiredTextEvent represents the AckRequiredTextEvent input event
type WebsocketAwapDemoApiDataAckRequiredTextEvent struct {
}

func (s WebsocketAwapDemoApiDataAckRequiredTextEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketAwapDemoApiDataAckRequiredTextEvent) GoString() string {
	return s.String()
}

func (s *WebsocketAwapDemoApiDataAckRequiredTextEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketAwapDemoApiDataUpstreamBinaryEvent represents the UpstreamBinaryEvent input event
type WebsocketAwapDemoApiDataUpstreamBinaryEvent struct {
}

func (s WebsocketAwapDemoApiDataUpstreamBinaryEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketAwapDemoApiDataUpstreamBinaryEvent) GoString() string {
	return s.String()
}

func (s *WebsocketAwapDemoApiDataUpstreamBinaryEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketAwapDemoApiDataMessageReceiveEvent represents the MessageReceiveEvent output event
type WebsocketAwapDemoApiDataMessageReceiveEvent struct {
}

func (s WebsocketAwapDemoApiDataMessageReceiveEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketAwapDemoApiDataMessageReceiveEvent) GoString() string {
	return s.String()
}

func (s *WebsocketAwapDemoApiDataMessageReceiveEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketAwapDemoApiDataDownstreamTextEvent represents the DownstreamTextEvent output event
type WebsocketAwapDemoApiDataDownstreamTextEvent struct {
	AudioId        *int32  `json:"audioId,omitempty" xml:"audioId,omitempty"`
	AudioType      *string `json:"audioType,omitempty" xml:"audioType,omitempty"`
	ProcessStatus  *string `json:"processStatus,omitempty" xml:"processStatus,omitempty"`
	AdditionalConf *struct {
		Timeout        *int64  `json:"timeout,omitempty" xml:"timeout,omitempty"`
		MaxAudioLength *string `json:"maxAudioLength,omitempty" xml:"maxAudioLength,omitempty"`
	} `json:"additionalConf,omitempty" xml:"additionalConf,omitempty"`
}

func (s WebsocketAwapDemoApiDataDownstreamTextEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketAwapDemoApiDataDownstreamTextEvent) GoString() string {
	return s.String()
}

func (s *WebsocketAwapDemoApiDataDownstreamTextEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketAwapDemoApiDataDownstreamBinaryEvent represents the DownstreamBinaryEvent output event
type WebsocketAwapDemoApiDataDownstreamBinaryEvent struct {
}

func (s WebsocketAwapDemoApiDataDownstreamBinaryEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketAwapDemoApiDataDownstreamBinaryEvent) GoString() string {
	return s.String()
}

func (s *WebsocketAwapDemoApiDataDownstreamBinaryEvent) Validate() error {
	return dara.Validate(s)
}
