// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebsocketGeneralDemoApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WebsocketGeneralDemoApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WebsocketGeneralDemoApiResponse
	GetStatusCode() *int32
	SetBody(v *WebsocketGeneralDemoApiResponseBody) *WebsocketGeneralDemoApiResponse
	GetBody() *WebsocketGeneralDemoApiResponseBody
}

type WebsocketGeneralDemoApiResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebsocketGeneralDemoApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WebsocketGeneralDemoApiResponse) String() string {
	return dara.Prettify(s)
}

func (s WebsocketGeneralDemoApiResponse) GoString() string {
	return s.String()
}

func (s *WebsocketGeneralDemoApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WebsocketGeneralDemoApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WebsocketGeneralDemoApiResponse) GetBody() *WebsocketGeneralDemoApiResponseBody {
	return s.Body
}

func (s *WebsocketGeneralDemoApiResponse) SetHeaders(v map[string]*string) *WebsocketGeneralDemoApiResponse {
	s.Headers = v
	return s
}

func (s *WebsocketGeneralDemoApiResponse) SetStatusCode(v int32) *WebsocketGeneralDemoApiResponse {
	s.StatusCode = &v
	return s
}

func (s *WebsocketGeneralDemoApiResponse) SetBody(v *WebsocketGeneralDemoApiResponseBody) *WebsocketGeneralDemoApiResponse {
	s.Body = v
	return s
}

func (s *WebsocketGeneralDemoApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
