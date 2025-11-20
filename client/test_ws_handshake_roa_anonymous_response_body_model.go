// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestWsHandshakeRoaAnonymousResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TestWsHandshakeRoaAnonymousResponseBody
	GetRequestId() *string
}

type TestWsHandshakeRoaAnonymousResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s TestWsHandshakeRoaAnonymousResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestWsHandshakeRoaAnonymousResponseBody) GoString() string {
	return s.String()
}

func (s *TestWsHandshakeRoaAnonymousResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestWsHandshakeRoaAnonymousResponseBody) SetRequestId(v string) *TestWsHandshakeRoaAnonymousResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestWsHandshakeRoaAnonymousResponseBody) Validate() error {
	return dara.Validate(s)
}
