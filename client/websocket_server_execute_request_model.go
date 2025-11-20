// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebsocketServerExecuteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *WebsocketServerExecuteRequest
	GetAction() *string
	SetCloseReason(v string) *WebsocketServerExecuteRequest
	GetCloseReason() *string
	SetCloseStatus(v string) *WebsocketServerExecuteRequest
	GetCloseStatus() *string
	SetSessionId(v string) *WebsocketServerExecuteRequest
	GetSessionId() *string
}

type WebsocketServerExecuteRequest struct {
	// This parameter is required.
	Action      *string `json:"action,omitempty" xml:"action,omitempty"`
	CloseReason *string `json:"closeReason,omitempty" xml:"closeReason,omitempty"`
	CloseStatus *string `json:"closeStatus,omitempty" xml:"closeStatus,omitempty"`
	// This parameter is required.
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s WebsocketServerExecuteRequest) String() string {
	return dara.Prettify(s)
}

func (s WebsocketServerExecuteRequest) GoString() string {
	return s.String()
}

func (s *WebsocketServerExecuteRequest) GetAction() *string {
	return s.Action
}

func (s *WebsocketServerExecuteRequest) GetCloseReason() *string {
	return s.CloseReason
}

func (s *WebsocketServerExecuteRequest) GetCloseStatus() *string {
	return s.CloseStatus
}

func (s *WebsocketServerExecuteRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *WebsocketServerExecuteRequest) SetAction(v string) *WebsocketServerExecuteRequest {
	s.Action = &v
	return s
}

func (s *WebsocketServerExecuteRequest) SetCloseReason(v string) *WebsocketServerExecuteRequest {
	s.CloseReason = &v
	return s
}

func (s *WebsocketServerExecuteRequest) SetCloseStatus(v string) *WebsocketServerExecuteRequest {
	s.CloseStatus = &v
	return s
}

func (s *WebsocketServerExecuteRequest) SetSessionId(v string) *WebsocketServerExecuteRequest {
	s.SessionId = &v
	return s
}

func (s *WebsocketServerExecuteRequest) Validate() error {
	return dara.Validate(s)
}
