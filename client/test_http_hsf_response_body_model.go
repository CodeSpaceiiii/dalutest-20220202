// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestHttpHsfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TestHttpHsfResponseBody
	GetRequestId() *string
}

type TestHttpHsfResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s TestHttpHsfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestHttpHsfResponseBody) GoString() string {
	return s.String()
}

func (s *TestHttpHsfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestHttpHsfResponseBody) SetRequestId(v string) *TestHttpHsfResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestHttpHsfResponseBody) Validate() error {
	return dara.Validate(s)
}
