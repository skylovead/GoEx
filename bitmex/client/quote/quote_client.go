// Code generated by go-swagger; DO NOT EDIT.

package quote

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new quote API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quote API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
QuoteGet gets quotes
*/
func (a *Client) QuoteGet(params *QuoteGetParams) (*QuoteGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Quote.get",
		Method:             "GET",
		PathPattern:        "/quote",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuoteGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QuoteGetOK), nil

}

/*
QuoteGetBucketed gets previous quotes in time buckets
*/
func (a *Client) QuoteGetBucketed(params *QuoteGetBucketedParams) (*QuoteGetBucketedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGetBucketedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Quote.getBucketed",
		Method:             "GET",
		PathPattern:        "/quote/bucketed",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuoteGetBucketedReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QuoteGetBucketedOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
