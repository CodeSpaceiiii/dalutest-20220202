// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebsocketServerExecuteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WebsocketServerExecuteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WebsocketServerExecuteResponse
	GetStatusCode() *int32
	SetBody(v *WebsocketServerExecuteResponseBody) *WebsocketServerExecuteResponse
	GetBody() *WebsocketServerExecuteResponseBody
}

type WebsocketServerExecuteResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebsocketServerExecuteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WebsocketServerExecuteResponse) String() string {
	return dara.Prettify(s)
}

func (s WebsocketServerExecuteResponse) GoString() string {
	return s.String()
}

func (s *WebsocketServerExecuteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WebsocketServerExecuteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WebsocketServerExecuteResponse) GetBody() *WebsocketServerExecuteResponseBody {
	return s.Body
}

func (s *WebsocketServerExecuteResponse) SetHeaders(v map[string]*string) *WebsocketServerExecuteResponse {
	s.Headers = v
	return s
}

func (s *WebsocketServerExecuteResponse) SetStatusCode(v int32) *WebsocketServerExecuteResponse {
	s.StatusCode = &v
	return s
}

func (s *WebsocketServerExecuteResponse) SetBody(v *WebsocketServerExecuteResponseBody) *WebsocketServerExecuteResponse {
	s.Body = v
	return s
}

func (s *WebsocketServerExecuteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
