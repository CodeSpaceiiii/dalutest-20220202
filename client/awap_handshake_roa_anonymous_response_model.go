// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAwapHandshakeRoaAnonymousResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AwapHandshakeRoaAnonymousResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AwapHandshakeRoaAnonymousResponse
	GetStatusCode() *int32
	SetBody(v *AwapHandshakeRoaAnonymousResponseBody) *AwapHandshakeRoaAnonymousResponse
	GetBody() *AwapHandshakeRoaAnonymousResponseBody
}

type AwapHandshakeRoaAnonymousResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AwapHandshakeRoaAnonymousResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AwapHandshakeRoaAnonymousResponse) String() string {
	return dara.Prettify(s)
}

func (s AwapHandshakeRoaAnonymousResponse) GoString() string {
	return s.String()
}

func (s *AwapHandshakeRoaAnonymousResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AwapHandshakeRoaAnonymousResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AwapHandshakeRoaAnonymousResponse) GetBody() *AwapHandshakeRoaAnonymousResponseBody {
	return s.Body
}

func (s *AwapHandshakeRoaAnonymousResponse) SetHeaders(v map[string]*string) *AwapHandshakeRoaAnonymousResponse {
	s.Headers = v
	return s
}

func (s *AwapHandshakeRoaAnonymousResponse) SetStatusCode(v int32) *AwapHandshakeRoaAnonymousResponse {
	s.StatusCode = &v
	return s
}

func (s *AwapHandshakeRoaAnonymousResponse) SetBody(v *AwapHandshakeRoaAnonymousResponseBody) *AwapHandshakeRoaAnonymousResponse {
	s.Body = v
	return s
}

func (s *AwapHandshakeRoaAnonymousResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
