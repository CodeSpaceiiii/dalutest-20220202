// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAwapHandshakeRoaAnonymousResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AwapHandshakeRoaAnonymousResponseBody
	GetRequestId() *string
}

type AwapHandshakeRoaAnonymousResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AwapHandshakeRoaAnonymousResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AwapHandshakeRoaAnonymousResponseBody) GoString() string {
	return s.String()
}

func (s *AwapHandshakeRoaAnonymousResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AwapHandshakeRoaAnonymousResponseBody) SetRequestId(v string) *AwapHandshakeRoaAnonymousResponseBody {
	s.RequestId = &v
	return s
}

func (s *AwapHandshakeRoaAnonymousResponseBody) Validate() error {
	return dara.Validate(s)
}
