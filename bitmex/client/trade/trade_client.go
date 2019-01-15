// Code generated by go-swagger; DO NOT EDIT.

package trade

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new trade API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for trade API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TradeGet gets trades

Please note that indices (symbols starting with `.`) post trades at intervals to the trade feed. These have a `size` of 0 and are used only to indicate a changing price.

See [the FIX Spec](http://www.onixs.biz/fix-dictionary/5.0.SP2/msgType_AE_6569.html) for explanations of these fields.
*/
func (a *Client) TradeGet(params *TradeGetParams) (*TradeGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTradeGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Trade.get",
		Method:             "GET",
		PathPattern:        "/trade",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TradeGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TradeGetOK), nil

}

/*
TradeGetBucketed gets previous trades in time buckets
*/
func (a *Client) TradeGetBucketed(params *TradeGetBucketedParams) (*TradeGetBucketedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTradeGetBucketedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Trade.getBucketed",
		Method:             "GET",
		PathPattern:        "/trade/bucketed",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TradeGetBucketedReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TradeGetBucketedOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
