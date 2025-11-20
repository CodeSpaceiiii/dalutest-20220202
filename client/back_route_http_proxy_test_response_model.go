// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackRouteHttpProxyTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BackRouteHttpProxyTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BackRouteHttpProxyTestResponse
	GetStatusCode() *int32
	SetBody(v *BackRouteHttpProxyTestResponseBody) *BackRouteHttpProxyTestResponse
	GetBody() *BackRouteHttpProxyTestResponseBody
}

type BackRouteHttpProxyTestResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BackRouteHttpProxyTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BackRouteHttpProxyTestResponse) String() string {
	return dara.Prettify(s)
}

func (s BackRouteHttpProxyTestResponse) GoString() string {
	return s.String()
}

func (s *BackRouteHttpProxyTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BackRouteHttpProxyTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BackRouteHttpProxyTestResponse) GetBody() *BackRouteHttpProxyTestResponseBody {
	return s.Body
}

func (s *BackRouteHttpProxyTestResponse) SetHeaders(v map[string]*string) *BackRouteHttpProxyTestResponse {
	s.Headers = v
	return s
}

func (s *BackRouteHttpProxyTestResponse) SetStatusCode(v int32) *BackRouteHttpProxyTestResponse {
	s.StatusCode = &v
	return s
}

func (s *BackRouteHttpProxyTestResponse) SetBody(v *BackRouteHttpProxyTestResponseBody) *BackRouteHttpProxyTestResponse {
	s.Body = v
	return s
}

func (s *BackRouteHttpProxyTestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
