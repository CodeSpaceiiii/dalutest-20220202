// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestWsHandshakeRoaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestWsHandshakeRoaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestWsHandshakeRoaResponse
	GetStatusCode() *int32
	SetBody(v *TestWsHandshakeRoaResponseBody) *TestWsHandshakeRoaResponse
	GetBody() *TestWsHandshakeRoaResponseBody
}

type TestWsHandshakeRoaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestWsHandshakeRoaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestWsHandshakeRoaResponse) String() string {
	return dara.Prettify(s)
}

func (s TestWsHandshakeRoaResponse) GoString() string {
	return s.String()
}

func (s *TestWsHandshakeRoaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestWsHandshakeRoaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestWsHandshakeRoaResponse) GetBody() *TestWsHandshakeRoaResponseBody {
	return s.Body
}

func (s *TestWsHandshakeRoaResponse) SetHeaders(v map[string]*string) *TestWsHandshakeRoaResponse {
	s.Headers = v
	return s
}

func (s *TestWsHandshakeRoaResponse) SetStatusCode(v int32) *TestWsHandshakeRoaResponse {
	s.StatusCode = &v
	return s
}

func (s *TestWsHandshakeRoaResponse) SetBody(v *TestWsHandshakeRoaResponseBody) *TestWsHandshakeRoaResponse {
	s.Body = v
	return s
}

func (s *TestWsHandshakeRoaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
