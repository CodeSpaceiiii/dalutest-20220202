// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackRouteHttpProxyTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BackRouteHttpProxyTestResponseBody
	GetRequestId() *string
}

type BackRouteHttpProxyTestResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BackRouteHttpProxyTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BackRouteHttpProxyTestResponseBody) GoString() string {
	return s.String()
}

func (s *BackRouteHttpProxyTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BackRouteHttpProxyTestResponseBody) SetRequestId(v string) *BackRouteHttpProxyTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *BackRouteHttpProxyTestResponseBody) Validate() error {
	return dara.Validate(s)
}
