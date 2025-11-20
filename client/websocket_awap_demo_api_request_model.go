// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebsocketAwapDemoApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuth(v string) *WebsocketAwapDemoApiRequest
	GetAuth() *string
	SetRespBody(v string) *WebsocketAwapDemoApiRequest
	GetRespBody() *string
	SetRespStatus(v string) *WebsocketAwapDemoApiRequest
	GetRespStatus() *string
	SetSleepMs(v string) *WebsocketAwapDemoApiRequest
	GetSleepMs() *string
}

type WebsocketAwapDemoApiRequest struct {
	Auth       *string `json:"auth,omitempty" xml:"auth,omitempty"`
	RespBody   *string `json:"respBody,omitempty" xml:"respBody,omitempty"`
	RespStatus *string `json:"respStatus,omitempty" xml:"respStatus,omitempty"`
	SleepMs    *string `json:"sleepMs,omitempty" xml:"sleepMs,omitempty"`
}

func (s WebsocketAwapDemoApiRequest) String() string {
	return dara.Prettify(s)
}

func (s WebsocketAwapDemoApiRequest) GoString() string {
	return s.String()
}

func (s *WebsocketAwapDemoApiRequest) GetAuth() *string {
	return s.Auth
}

func (s *WebsocketAwapDemoApiRequest) GetRespBody() *string {
	return s.RespBody
}

func (s *WebsocketAwapDemoApiRequest) GetRespStatus() *string {
	return s.RespStatus
}

func (s *WebsocketAwapDemoApiRequest) GetSleepMs() *string {
	return s.SleepMs
}

func (s *WebsocketAwapDemoApiRequest) SetAuth(v string) *WebsocketAwapDemoApiRequest {
	s.Auth = &v
	return s
}

func (s *WebsocketAwapDemoApiRequest) SetRespBody(v string) *WebsocketAwapDemoApiRequest {
	s.RespBody = &v
	return s
}

func (s *WebsocketAwapDemoApiRequest) SetRespStatus(v string) *WebsocketAwapDemoApiRequest {
	s.RespStatus = &v
	return s
}

func (s *WebsocketAwapDemoApiRequest) SetSleepMs(v string) *WebsocketAwapDemoApiRequest {
	s.SleepMs = &v
	return s
}

func (s *WebsocketAwapDemoApiRequest) Validate() error {
	return dara.Validate(s)
}
