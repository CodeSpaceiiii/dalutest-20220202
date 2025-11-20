// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/darabonba-openapi/v2/websocketutils"
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
	WebSocketClient *websocketutils.WebSocketClient `json:"websocketClient,omitempty" xml:"websocketClient,omitempty"`
}

// WebsocketGeneralDemoApiWebSocketGeneralUpstreamTextEvent represents the GeneralUpstreamTextEvent input event
type WebsocketGeneralDemoApiDataGeneralUpstreamTextEvent struct {
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

func (s WebsocketGeneralDemoApiDataGeneralUpstreamTextEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiDataGeneralUpstreamTextEvent) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiDataGeneralUpstreamTextEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketGeneralDemoApiWebSocketGeneralUpstreamBinaryEvent represents the GeneralUpstreamBinaryEvent input event
type WebsocketGeneralDemoApiDataGeneralUpstreamBinaryEvent struct {
}

func (s WebsocketGeneralDemoApiDataGeneralUpstreamBinaryEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiDataGeneralUpstreamBinaryEvent) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiDataGeneralUpstreamBinaryEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketGeneralDemoApiWebSocketUpstreamDefaultTextEvent represents the UpstreamDefaultTextEvent input event
type WebsocketGeneralDemoApiDataUpstreamDefaultTextEvent struct {
}

func (s WebsocketGeneralDemoApiDataUpstreamDefaultTextEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiDataUpstreamDefaultTextEvent) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiDataUpstreamDefaultTextEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketGeneralDemoApiWebSocketDownstreamDefaultTextEvent represents the DownstreamDefaultTextEvent output event
type WebsocketGeneralDemoApiDataDownstreamDefaultTextEvent struct {
}

func (s WebsocketGeneralDemoApiDataDownstreamDefaultTextEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiDataDownstreamDefaultTextEvent) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiDataDownstreamDefaultTextEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketGeneralDemoApiWebSocketGeneralDownstreamTextEvent represents the GeneralDownstreamTextEvent output event
type WebsocketGeneralDemoApiDataGeneralDownstreamTextEvent struct {
	AudioId        *int32  `json:"audioId,omitempty" xml:"audioId,omitempty"`
	AudioType      *string `json:"audioType,omitempty" xml:"audioType,omitempty"`
	ProcessStatus  *string `json:"processStatus,omitempty" xml:"processStatus,omitempty"`
	AdditionalConf *struct {
		Timeout        *int64  `json:"timeout,omitempty" xml:"timeout,omitempty"`
		MaxAudioLength *string `json:"maxAudioLength,omitempty" xml:"maxAudioLength,omitempty"`
	} `json:"additionalConf,omitempty" xml:"additionalConf,omitempty"`
}

func (s WebsocketGeneralDemoApiDataGeneralDownstreamTextEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiDataGeneralDownstreamTextEvent) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiDataGeneralDownstreamTextEvent) Validate() error {
	return dara.Validate(s)
}

// WebsocketGeneralDemoApiWebSocketGeneralDownstreamBinaryEvent represents the GeneralDownstreamBinaryEvent output event
type WebsocketGeneralDemoApiDataGeneralDownstreamBinaryEvent struct {
}

func (s WebsocketGeneralDemoApiDataGeneralDownstreamBinaryEvent) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiDataGeneralDownstreamBinaryEvent) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiDataGeneralDownstreamBinaryEvent) Validate() error {
	return dara.Validate(s)
}
