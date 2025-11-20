// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebsocketAwapDemoApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *WebsocketAwapDemoApiResponseBody
	GetRequestId() *string
}

type WebsocketAwapDemoApiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WebsocketAwapDemoApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WebsocketAwapDemoApiResponseBody) GoString() string {
	return s.String()
}

func (s *WebsocketAwapDemoApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebsocketAwapDemoApiResponseBody) SetRequestId(v string) *WebsocketAwapDemoApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *WebsocketAwapDemoApiResponseBody) Validate() error {
	return dara.Validate(s)
}
