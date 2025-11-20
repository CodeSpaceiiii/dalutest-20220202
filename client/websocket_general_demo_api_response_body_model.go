// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebsocketGeneralDemoApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *WebsocketGeneralDemoApiResponseBody
	GetRequestId() *string
}

type WebsocketGeneralDemoApiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WebsocketGeneralDemoApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiResponseBody) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebsocketGeneralDemoApiResponseBody) SetRequestId(v string) *WebsocketGeneralDemoApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *WebsocketGeneralDemoApiResponseBody) Validate() error {
	return dara.Validate(s)
}
