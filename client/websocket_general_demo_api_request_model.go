// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebsocketGeneralDemoApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuth(v string) *WebsocketGeneralDemoApiRequest
	GetAuth() *string
	SetRespBody(v string) *WebsocketGeneralDemoApiRequest
	GetRespBody() *string
	SetRespStatus(v string) *WebsocketGeneralDemoApiRequest
	GetRespStatus() *string
	SetSleepMs(v string) *WebsocketGeneralDemoApiRequest
	GetSleepMs() *string
}

type WebsocketGeneralDemoApiRequest struct {
	Auth       *string `json:"auth,omitempty" xml:"auth,omitempty"`
	RespBody   *string `json:"respBody,omitempty" xml:"respBody,omitempty"`
	RespStatus *string `json:"respStatus,omitempty" xml:"respStatus,omitempty"`
	SleepMs    *string `json:"sleepMs,omitempty" xml:"sleepMs,omitempty"`
}

func (s WebsocketGeneralDemoApiRequest) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiRequest) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiRequest) GetAuth() *string {
	return s.Auth
}

func (s *WebsocketGeneralDemoApiRequest) GetRespBody() *string {
	return s.RespBody
}

func (s *WebsocketGeneralDemoApiRequest) GetRespStatus() *string {
	return s.RespStatus
}

func (s *WebsocketGeneralDemoApiRequest) GetSleepMs() *string {
	return s.SleepMs
}

func (s *WebsocketGeneralDemoApiRequest) SetAuth(v string) *WebsocketGeneralDemoApiRequest {
	s.Auth = &v
	return s
}

func (s *WebsocketGeneralDemoApiRequest) SetRespBody(v string) *WebsocketGeneralDemoApiRequest {
	s.RespBody = &v
	return s
}

func (s *WebsocketGeneralDemoApiRequest) SetRespStatus(v string) *WebsocketGeneralDemoApiRequest {
	s.RespStatus = &v
	return s
}

func (s *WebsocketGeneralDemoApiRequest) SetSleepMs(v string) *WebsocketGeneralDemoApiRequest {
	s.SleepMs = &v
	return s
}

func (s *WebsocketGeneralDemoApiRequest) Validate() error {
	return dara.Validate(s)
}
