package WebsocketAwapDemoApi

import "github.com/alibabacloud-go/tea/dara"

// 实现的枚举类
const (
	// Upstream event types (client -> server)
	UpstreamTextEvent_MessageType    dara.AwapMessageType = "UpstreamTextEvent"
	UpstreamBinaryEvent_MessageType  dara.AwapMessageType = "UpstreamBinaryEvent"
	AckRequiredTextEvent_MessageType dara.AwapMessageType = "AckRequiredTextEvent"

	// Downstream event types (server -> client)
	ReceiveEvent_MessageType          dara.AwapMessageType = "MessageReceiveEvent"
	DownstreamTextEvent_MessageType   dara.AwapMessageType = "DownstreamTextEvent"
	DownstreamBinaryEvent_MessageType dara.AwapMessageType = "DownstreamBinaryEvent"

	// Control message types (server -> client)
	Reconnect_MessageType dara.GeneralMessageType = "RECONNECT" // Server-initiated graceful reconnection
)

type UpstreamTextEvent struct {
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

type AckRequiredTextEvent interface{}

type UpstreamBinaryEvent interface{}

type MessageReceiveEvent interface{}

type DownstreamTextEvent struct {
	AudioId        *int32  `json:"audioId,omitempty" xml:"audioId,omitempty"`
	AudioType      *string `json:"audioType,omitempty" xml:"audioType,omitempty"`
	ProcessStatus  *string `json:"processStatus,omitempty" xml:"processStatus,omitempty"`
	AdditionalConf *struct {
		Timeout        *int64  `json:"timeout,omitempty" xml:"timeout,omitempty"`
		MaxAudioLength *string `json:"maxAudioLength,omitempty" xml:"maxAudioLength,omitempty"`
	} `json:"additionalConf,omitempty" xml:"additionalConf,omitempty"`
}

type DownstreamBinaryEvent interface{}
