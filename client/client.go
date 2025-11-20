// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dalutest"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AwapHandshakeRoaAnonymousResponse
func (client *Client) AwapHandshakeRoaAnonymousWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AwapHandshakeRoaAnonymousResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("AwapHandshakeRoaAnonymous"),
		Version:     dara.String("2022-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ws/proxy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AwapHandshakeRoaAnonymousResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return AwapHandshakeRoaAnonymousResponse
func (client *Client) AwapHandshakeRoaAnonymous() (_result *AwapHandshakeRoaAnonymousResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AwapHandshakeRoaAnonymousResponse{}
	_body, _err := client.AwapHandshakeRoaAnonymousWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BackRouteHttpProxyTestResponse
func (client *Client) BackRouteHttpProxyTestWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BackRouteHttpProxyTestResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("BackRouteHttpProxyTest"),
		Version:     dara.String("2022-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/back-route/proxy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BackRouteHttpProxyTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return BackRouteHttpProxyTestResponse
func (client *Client) BackRouteHttpProxyTest() (_result *BackRouteHttpProxyTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BackRouteHttpProxyTestResponse{}
	_body, _err := client.BackRouteHttpProxyTestWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestHttpHsfResponse
func (client *Client) TestHttpHsfWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TestHttpHsfResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestHttpHsf"),
		Version:     dara.String("2022-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/test123"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TestHttpHsfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return TestHttpHsfResponse
func (client *Client) TestHttpHsf() (_result *TestHttpHsfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TestHttpHsfResponse{}
	_body, _err := client.TestHttpHsfWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestWsHandshakeRoaResponse
func (client *Client) TestWsHandshakeRoaWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TestWsHandshakeRoaResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestWsHandshakeRoa"),
		Version:     dara.String("2022-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ws/proxy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TestWsHandshakeRoaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return TestWsHandshakeRoaResponse
func (client *Client) TestWsHandshakeRoa() (_result *TestWsHandshakeRoaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TestWsHandshakeRoaResponse{}
	_body, _err := client.TestWsHandshakeRoaWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestWsHandshakeRoaAnonymousResponse
func (client *Client) TestWsHandshakeRoaAnonymousWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TestWsHandshakeRoaAnonymousResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestWsHandshakeRoaAnonymous"),
		Version:     dara.String("2022-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ws/proxy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TestWsHandshakeRoaAnonymousResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return TestWsHandshakeRoaAnonymousResponse
func (client *Client) TestWsHandshakeRoaAnonymous() (_result *TestWsHandshakeRoaAnonymousResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TestWsHandshakeRoaAnonymousResponse{}
	_body, _err := client.TestWsHandshakeRoaAnonymousWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - WebsocketAwapDemoApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WebsocketAwapDemoApiResponse
func (client *Client) WebsocketAwapDemoApiWithOptions(request *WebsocketAwapDemoApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *WebsocketAwapDemoApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Auth) {
		query["auth"] = request.Auth
	}

	if !dara.IsNil(request.RespBody) {
		query["respBody"] = request.RespBody
	}

	if !dara.IsNil(request.RespStatus) {
		query["respStatus"] = request.RespStatus
	}

	if !dara.IsNil(request.SleepMs) {
		query["sleepMs"] = request.SleepMs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WebsocketAwapDemoApi"),
		Version:     dara.String("2022-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ws/awap-demo-api"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &WebsocketAwapDemoApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - WebsocketAwapDemoApiRequest
//
// @return WebsocketAwapDemoApiResponse
func (client *Client) WebsocketAwapDemoApi(request *WebsocketAwapDemoApiRequest) (_result *WebsocketAwapDemoApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WebsocketAwapDemoApiResponse{}
	_body, _err := client.WebsocketAwapDemoApiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - WebsocketGeneralDemoApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WebsocketGeneralDemoApiResponse
// / need update
func (client *Client) WebsocketGeneralDemoApiWithOptions(request *WebsocketGeneralDemoApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *WebsocketGeneralDemoApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Auth) {
		query["auth"] = request.Auth
	}

	if !dara.IsNil(request.RespBody) {
		query["respBody"] = request.RespBody
	}

	if !dara.IsNil(request.RespStatus) {
		query["respStatus"] = request.RespStatus
	}

	if !dara.IsNil(request.SleepMs) {
		query["sleepMs"] = request.SleepMs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:  dara.String("WebsocketGeneralDemoApi"),
		Version: dara.String("2022-02-02"),
		// 注意，上层生成器再这里必须转换为全小写的ws和wss，否则会报错
		Protocol:    dara.String("wss"),
		Pathname:    dara.String("/ws/general-demo-api"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &WebsocketGeneralDemoApiResponse{}
	_body, _err := client.DoRequest(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	// 这里不能继续直接使用convert，需要使用新的类，但握手的信息，header还是要能返回
	// 确保body存在wsClient
	wsClient, ok := _body["wsClient"].(*dara.DefaultWebSocketClient)
	if !ok {
		return _result, _err
	}
	_result.WebSocketClient = wsClient
	return _result, nil
}

// @param request - WebsocketGeneralDemoApiRequest
//
// @return WebsocketGeneralDemoApiResponse
func (client *Client) WebsocketGeneralDemoApi(request *WebsocketGeneralDemoApiRequest) (_result *WebsocketGeneralDemoApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WebsocketGeneralDemoApiResponse{}
	_body, _err := client.WebsocketGeneralDemoApiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - WebsocketServerExecuteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WebsocketServerExecuteResponse
func (client *Client) WebsocketServerExecuteWithOptions(request *WebsocketServerExecuteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *WebsocketServerExecuteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		query["action"] = request.Action
	}

	if !dara.IsNil(request.CloseReason) {
		query["closeReason"] = request.CloseReason
	}

	if !dara.IsNil(request.CloseStatus) {
		query["closeStatus"] = request.CloseStatus
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WebsocketServerExecute"),
		Version:     dara.String("2022-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ws_server/execute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &WebsocketServerExecuteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - WebsocketServerExecuteRequest
//
// @return WebsocketServerExecuteResponse
func (client *Client) WebsocketServerExecute(request *WebsocketServerExecuteRequest) (_result *WebsocketServerExecuteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WebsocketServerExecuteResponse{}
	_body, _err := client.WebsocketServerExecuteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
