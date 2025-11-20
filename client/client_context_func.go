// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AwapHandshakeRoaAnonymousResponse
func (client *Client) AwapHandshakeRoaAnonymousWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AwapHandshakeRoaAnonymousResponse, _err error) {
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
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BackRouteHttpProxyTestResponse
func (client *Client) BackRouteHttpProxyTestWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BackRouteHttpProxyTestResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestHttpHsfResponse
func (client *Client) TestHttpHsfWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TestHttpHsfResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestWsHandshakeRoaResponse
func (client *Client) TestWsHandshakeRoaWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TestWsHandshakeRoaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestWsHandshakeRoaAnonymousResponse
func (client *Client) TestWsHandshakeRoaAnonymousWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TestWsHandshakeRoaAnonymousResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - WebsocketAwapDemoApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WebsocketAwapDemoApiResponse
func (client *Client) WebsocketAwapDemoApiWithContext(ctx context.Context, request *WebsocketAwapDemoApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *WebsocketAwapDemoApiResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - WebsocketGeneralDemoApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WebsocketGeneralDemoApiResponse
func (client *Client) WebsocketGeneralDemoApiWithContext(ctx context.Context, request *WebsocketGeneralDemoApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *WebsocketGeneralDemoApiResponse, _err error) {
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
		Action:      dara.String("WebsocketGeneralDemoApi"),
		Version:     dara.String("2022-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ws/general-demo-api"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &WebsocketGeneralDemoApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - WebsocketServerExecuteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WebsocketServerExecuteResponse
func (client *Client) WebsocketServerExecuteWithContext(ctx context.Context, request *WebsocketServerExecuteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *WebsocketServerExecuteResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
