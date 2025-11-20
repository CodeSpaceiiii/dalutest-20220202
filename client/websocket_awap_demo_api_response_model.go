// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebsocketAwapDemoApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WebsocketAwapDemoApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WebsocketAwapDemoApiResponse
	GetStatusCode() *int32
	SetBody(v *WebsocketAwapDemoApiResponseBody) *WebsocketAwapDemoApiResponse
	GetBody() *WebsocketAwapDemoApiResponseBody
}

type WebsocketAwapDemoApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebsocketAwapDemoApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WebsocketAwapDemoApiResponse) String() string {
	return dara.Prettify(s)
}

func (s WebsocketAwapDemoApiResponse) GoString() string {
	return s.String()
}

func (s *WebsocketAwapDemoApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WebsocketAwapDemoApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WebsocketAwapDemoApiResponse) GetBody() *WebsocketAwapDemoApiResponseBody {
	return s.Body
}

func (s *WebsocketAwapDemoApiResponse) SetHeaders(v map[string]*string) *WebsocketAwapDemoApiResponse {
	s.Headers = v
	return s
}

func (s *WebsocketAwapDemoApiResponse) SetStatusCode(v int32) *WebsocketAwapDemoApiResponse {
	s.StatusCode = &v
	return s
}

func (s *WebsocketAwapDemoApiResponse) SetBody(v *WebsocketAwapDemoApiResponseBody) *WebsocketAwapDemoApiResponse {
	s.Body = v
	return s
}

func (s *WebsocketAwapDemoApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
