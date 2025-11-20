// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestWsHandshakeRoaAnonymousResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestWsHandshakeRoaAnonymousResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestWsHandshakeRoaAnonymousResponse
	GetStatusCode() *int32
	SetBody(v *TestWsHandshakeRoaAnonymousResponseBody) *TestWsHandshakeRoaAnonymousResponse
	GetBody() *TestWsHandshakeRoaAnonymousResponseBody
}

type TestWsHandshakeRoaAnonymousResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestWsHandshakeRoaAnonymousResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestWsHandshakeRoaAnonymousResponse) String() string {
	return dara.Prettify(s)
}

func (s TestWsHandshakeRoaAnonymousResponse) GoString() string {
	return s.String()
}

func (s *TestWsHandshakeRoaAnonymousResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestWsHandshakeRoaAnonymousResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestWsHandshakeRoaAnonymousResponse) GetBody() *TestWsHandshakeRoaAnonymousResponseBody {
	return s.Body
}

func (s *TestWsHandshakeRoaAnonymousResponse) SetHeaders(v map[string]*string) *TestWsHandshakeRoaAnonymousResponse {
	s.Headers = v
	return s
}

func (s *TestWsHandshakeRoaAnonymousResponse) SetStatusCode(v int32) *TestWsHandshakeRoaAnonymousResponse {
	s.StatusCode = &v
	return s
}

func (s *TestWsHandshakeRoaAnonymousResponse) SetBody(v *TestWsHandshakeRoaAnonymousResponseBody) *TestWsHandshakeRoaAnonymousResponse {
	s.Body = v
	return s
}

func (s *TestWsHandshakeRoaAnonymousResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
