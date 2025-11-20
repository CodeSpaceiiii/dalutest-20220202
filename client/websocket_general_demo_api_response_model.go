// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebsocketGeneralDemoApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WebsocketGeneralDemoApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WebsocketGeneralDemoApiResponse
	GetStatusCode() *int32
	SetBody(v *WebsocketGeneralDemoApiResponseBody) *WebsocketGeneralDemoApiResponse
	GetBody() *WebsocketGeneralDemoApiResponseBody
}

type WebsocketGeneralDemoApiResponse struct {
	WebSocketClient *dara.DefaultWebSocketClient `json:"websocketClient,omitempty" xml:"websocketClient,omitempty"`
}

// WebsocketGeneralDemoApiWebSocketGeneralUpstreamTextEvent represents the GeneralUpstreamTextEvent input event
type WebsocketGeneralDemoApiWebSocketGeneralUpstreamTextEvent struct {
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

func (s WebsocketGeneralDemoApiWebSocketGeneralUpstreamTextEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiWebSocketGeneralUpstreamTextEvent) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiWebSocketGeneralUpstreamTextEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketGeneralDemoApiWebSocketGeneralUpstreamBinaryEvent represents the GeneralUpstreamBinaryEvent input event
type WebsocketGeneralDemoApiWebSocketGeneralUpstreamBinaryEvent struct {
}

func (s WebsocketGeneralDemoApiWebSocketGeneralUpstreamBinaryEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiWebSocketGeneralUpstreamBinaryEvent) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiWebSocketGeneralUpstreamBinaryEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketGeneralDemoApiWebSocketUpstreamDefaultTextEvent represents the UpstreamDefaultTextEvent input event
type WebsocketGeneralDemoApiWebSocketUpstreamDefaultTextEvent struct {
}

func (s WebsocketGeneralDemoApiWebSocketUpstreamDefaultTextEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiWebSocketUpstreamDefaultTextEvent) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiWebSocketUpstreamDefaultTextEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketGeneralDemoApiWebSocketDownstreamDefaultTextEvent represents the DownstreamDefaultTextEvent output event
type WebsocketGeneralDemoApiWebSocketDownstreamDefaultTextEvent struct {
}

func (s WebsocketGeneralDemoApiWebSocketDownstreamDefaultTextEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiWebSocketDownstreamDefaultTextEvent) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiWebSocketDownstreamDefaultTextEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketGeneralDemoApiWebSocketGeneralDownstreamTextEvent represents the GeneralDownstreamTextEvent output event
type WebsocketGeneralDemoApiWebSocketGeneralDownstreamTextEvent struct {
	AudioId        *int32  `json:"audioId,omitempty" xml:"audioId,omitempty"`
	AudioType      *string `json:"audioType,omitempty" xml:"audioType,omitempty"`
	ProcessStatus  *string `json:"processStatus,omitempty" xml:"processStatus,omitempty"`
	AdditionalConf *struct {
		Timeout        *int64  `json:"timeout,omitempty" xml:"timeout,omitempty"`
		MaxAudioLength *string `json:"maxAudioLength,omitempty" xml:"maxAudioLength,omitempty"`
	} `json:"additionalConf,omitempty" xml:"additionalConf,omitempty"`
}

func (s WebsocketGeneralDemoApiWebSocketGeneralDownstreamTextEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiWebSocketGeneralDownstreamTextEvent) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiWebSocketGeneralDownstreamTextEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketGeneralDemoApiWebSocketGeneralDownstreamBinaryEvent represents the GeneralDownstreamBinaryEvent output event
type WebsocketGeneralDemoApiWebSocketGeneralDownstreamBinaryEvent struct {
}

func (s WebsocketGeneralDemoApiWebSocketGeneralDownstreamBinaryEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiWebSocketGeneralDownstreamBinaryEvent) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiWebSocketGeneralDownstreamBinaryEvent) Validate() error {
	return dara.Validate(s)
}
