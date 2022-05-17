// Package http_api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.10.1 DO NOT EDIT.
package http_api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GameMakeAction request with any body
	GameMakeActionWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	GameMakeAction(ctx context.Context, body GameMakeActionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GameGetState request
	GameGetState(ctx context.Context, params *GameGetStateParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PlayerGet request
	PlayerGet(ctx context.Context, params *PlayerGetParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PlayerCreate request with any body
	PlayerCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PlayerCreate(ctx context.Context, body PlayerCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PlayerStats request
	PlayerStats(ctx context.Context, params *PlayerStatsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GameMakeActionWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGameMakeActionRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GameMakeAction(ctx context.Context, body GameMakeActionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGameMakeActionRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GameGetState(ctx context.Context, params *GameGetStateParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGameGetStateRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PlayerGet(ctx context.Context, params *PlayerGetParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPlayerGetRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PlayerCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPlayerCreateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PlayerCreate(ctx context.Context, body PlayerCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPlayerCreateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PlayerStats(ctx context.Context, params *PlayerStatsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPlayerStatsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGameMakeActionRequest calls the generic GameMakeAction builder with application/json body
func NewGameMakeActionRequest(server string, body GameMakeActionJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGameMakeActionRequestWithBody(server, "application/json", bodyReader)
}

// NewGameMakeActionRequestWithBody generates requests for GameMakeAction with any type of body
func NewGameMakeActionRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/game/action")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGameGetStateRequest generates requests for GameGetState
func NewGameGetStateRequest(server string, params *GameGetStateParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/game/state")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "login", runtime.ParamLocationQuery, params.Login); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPlayerGetRequest generates requests for PlayerGet
func NewPlayerGetRequest(server string, params *PlayerGetParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/player")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "login", runtime.ParamLocationQuery, params.Login); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPlayerCreateRequest calls the generic PlayerCreate builder with application/json body
func NewPlayerCreateRequest(server string, body PlayerCreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPlayerCreateRequestWithBody(server, "application/json", bodyReader)
}

// NewPlayerCreateRequestWithBody generates requests for PlayerCreate with any type of body
func NewPlayerCreateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/player/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewPlayerStatsRequest generates requests for PlayerStats
func NewPlayerStatsRequest(server string, params *PlayerStatsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/player/stats")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "login", runtime.ParamLocationQuery, params.Login); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GameMakeAction request with any body
	GameMakeActionWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GameMakeActionResponse, error)

	GameMakeActionWithResponse(ctx context.Context, body GameMakeActionJSONRequestBody, reqEditors ...RequestEditorFn) (*GameMakeActionResponse, error)

	// GameGetState request
	GameGetStateWithResponse(ctx context.Context, params *GameGetStateParams, reqEditors ...RequestEditorFn) (*GameGetStateResponse, error)

	// PlayerGet request
	PlayerGetWithResponse(ctx context.Context, params *PlayerGetParams, reqEditors ...RequestEditorFn) (*PlayerGetResponse, error)

	// PlayerCreate request with any body
	PlayerCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PlayerCreateResponse, error)

	PlayerCreateWithResponse(ctx context.Context, body PlayerCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*PlayerCreateResponse, error)

	// PlayerStats request
	PlayerStatsWithResponse(ctx context.Context, params *PlayerStatsParams, reqEditors ...RequestEditorFn) (*PlayerStatsResponse, error)
}

type GameMakeActionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r GameMakeActionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GameMakeActionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GameGetStateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GameState
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r GameGetStateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GameGetStateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PlayerGetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Player
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r PlayerGetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PlayerGetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PlayerCreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Player
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r PlayerCreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PlayerCreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PlayerStatsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Stats
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r PlayerStatsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PlayerStatsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GameMakeActionWithBodyWithResponse request with arbitrary body returning *GameMakeActionResponse
func (c *ClientWithResponses) GameMakeActionWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GameMakeActionResponse, error) {
	rsp, err := c.GameMakeActionWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGameMakeActionResponse(rsp)
}

func (c *ClientWithResponses) GameMakeActionWithResponse(ctx context.Context, body GameMakeActionJSONRequestBody, reqEditors ...RequestEditorFn) (*GameMakeActionResponse, error) {
	rsp, err := c.GameMakeAction(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGameMakeActionResponse(rsp)
}

// GameGetStateWithResponse request returning *GameGetStateResponse
func (c *ClientWithResponses) GameGetStateWithResponse(ctx context.Context, params *GameGetStateParams, reqEditors ...RequestEditorFn) (*GameGetStateResponse, error) {
	rsp, err := c.GameGetState(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGameGetStateResponse(rsp)
}

// PlayerGetWithResponse request returning *PlayerGetResponse
func (c *ClientWithResponses) PlayerGetWithResponse(ctx context.Context, params *PlayerGetParams, reqEditors ...RequestEditorFn) (*PlayerGetResponse, error) {
	rsp, err := c.PlayerGet(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePlayerGetResponse(rsp)
}

// PlayerCreateWithBodyWithResponse request with arbitrary body returning *PlayerCreateResponse
func (c *ClientWithResponses) PlayerCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PlayerCreateResponse, error) {
	rsp, err := c.PlayerCreateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePlayerCreateResponse(rsp)
}

func (c *ClientWithResponses) PlayerCreateWithResponse(ctx context.Context, body PlayerCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*PlayerCreateResponse, error) {
	rsp, err := c.PlayerCreate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePlayerCreateResponse(rsp)
}

// PlayerStatsWithResponse request returning *PlayerStatsResponse
func (c *ClientWithResponses) PlayerStatsWithResponse(ctx context.Context, params *PlayerStatsParams, reqEditors ...RequestEditorFn) (*PlayerStatsResponse, error) {
	rsp, err := c.PlayerStats(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePlayerStatsResponse(rsp)
}

// ParseGameMakeActionResponse parses an HTTP response from a GameMakeActionWithResponse call
func ParseGameMakeActionResponse(rsp *http.Response) (*GameMakeActionResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GameMakeActionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGameGetStateResponse parses an HTTP response from a GameGetStateWithResponse call
func ParseGameGetStateResponse(rsp *http.Response) (*GameGetStateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GameGetStateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GameState
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParsePlayerGetResponse parses an HTTP response from a PlayerGetWithResponse call
func ParsePlayerGetResponse(rsp *http.Response) (*PlayerGetResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PlayerGetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Player
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParsePlayerCreateResponse parses an HTTP response from a PlayerCreateWithResponse call
func ParsePlayerCreateResponse(rsp *http.Response) (*PlayerCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PlayerCreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Player
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParsePlayerStatsResponse parses an HTTP response from a PlayerStatsWithResponse call
func ParsePlayerStatsResponse(rsp *http.Response) (*PlayerStatsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PlayerStatsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Stats
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
