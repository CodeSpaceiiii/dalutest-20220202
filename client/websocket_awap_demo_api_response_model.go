// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/darabonba-openapi/v2/websocketutils"
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
	WebSocketClient *websocketutils.WebSocketClient `json:"websocketClient,omitempty" xml:"websocketClient,omitempty"`
}
