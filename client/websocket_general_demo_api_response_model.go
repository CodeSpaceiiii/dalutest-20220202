// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/darabonba-openapi/v2/websocketutils"
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
	WebSocketClient *websocketutils.WebSocketClient `json:"websocketClient,omitempty" xml:"websocketClient,omitempty"`
}
