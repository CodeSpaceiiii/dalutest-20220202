// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebsocketServerExecuteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *WebsocketServerExecuteResponseBody
	GetRequestId() *string
}

type WebsocketServerExecuteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WebsocketServerExecuteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WebsocketServerExecuteResponseBody) GoString() string {
	return s.String()
}

func (s *WebsocketServerExecuteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebsocketServerExecuteResponseBody) SetRequestId(v string) *WebsocketServerExecuteResponseBody {
	s.RequestId = &v
	return s
}

func (s *WebsocketServerExecuteResponseBody) Validate() error {
	return dara.Validate(s)
}
