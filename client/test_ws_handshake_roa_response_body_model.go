// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestWsHandshakeRoaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TestWsHandshakeRoaResponseBody
	GetRequestId() *string
}

type TestWsHandshakeRoaResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s TestWsHandshakeRoaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestWsHandshakeRoaResponseBody) GoString() string {
	return s.String()
}

func (s *TestWsHandshakeRoaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestWsHandshakeRoaResponseBody) SetRequestId(v string) *TestWsHandshakeRoaResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestWsHandshakeRoaResponseBody) Validate() error {
	return dara.Validate(s)
}
