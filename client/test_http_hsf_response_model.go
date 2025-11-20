// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestHttpHsfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestHttpHsfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestHttpHsfResponse
	GetStatusCode() *int32
	SetBody(v *TestHttpHsfResponseBody) *TestHttpHsfResponse
	GetBody() *TestHttpHsfResponseBody
}

type TestHttpHsfResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestHttpHsfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestHttpHsfResponse) String() string {
	return dara.Prettify(s)
}

func (s TestHttpHsfResponse) GoString() string {
	return s.String()
}

func (s *TestHttpHsfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestHttpHsfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestHttpHsfResponse) GetBody() *TestHttpHsfResponseBody {
	return s.Body
}

func (s *TestHttpHsfResponse) SetHeaders(v map[string]*string) *TestHttpHsfResponse {
	s.Headers = v
	return s
}

func (s *TestHttpHsfResponse) SetStatusCode(v int32) *TestHttpHsfResponse {
	s.StatusCode = &v
	return s
}

func (s *TestHttpHsfResponse) SetBody(v *TestHttpHsfResponseBody) *TestHttpHsfResponse {
	s.Body = v
	return s
}

func (s *TestHttpHsfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
