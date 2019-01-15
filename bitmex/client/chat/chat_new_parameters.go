// Code generated by go-swagger; DO NOT EDIT.

package chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewChatNewParams creates a new ChatNewParams object
// with the default values initialized.
func NewChatNewParams() *ChatNewParams {
	var (
		channelIDDefault = float64(1)
	)
	return &ChatNewParams{
		ChannelID: &channelIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewChatNewParamsWithTimeout creates a new ChatNewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChatNewParamsWithTimeout(timeout time.Duration) *ChatNewParams {
	var (
		channelIDDefault = float64(1)
	)
	return &ChatNewParams{
		ChannelID: &channelIDDefault,

		timeout: timeout,
	}
}

// NewChatNewParamsWithContext creates a new ChatNewParams object
// with the default values initialized, and the ability to set a context for a request
func NewChatNewParamsWithContext(ctx context.Context) *ChatNewParams {
	var (
		channelIdDefault = float64(1)
	)
	return &ChatNewParams{
		ChannelID: &channelIdDefault,

		Context: ctx,
	}
}

// NewChatNewParamsWithHTTPClient creates a new ChatNewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChatNewParamsWithHTTPClient(client *http.Client) *ChatNewParams {
	var (
		channelIdDefault = float64(1)
	)
	return &ChatNewParams{
		ChannelID:  &channelIdDefault,
		HTTPClient: client,
	}
}

/*ChatNewParams contains all the parameters to send to the API endpoint
for the chat new operation typically these are written to a http.Request
*/
type ChatNewParams struct {

	/*ChannelID
	  Channel to post to. Default 1 (English).

	*/
	ChannelID *float64
	/*Message*/
	Message string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the chat new params
func (o *ChatNewParams) WithTimeout(timeout time.Duration) *ChatNewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chat new params
func (o *ChatNewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chat new params
func (o *ChatNewParams) WithContext(ctx context.Context) *ChatNewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chat new params
func (o *ChatNewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chat new params
func (o *ChatNewParams) WithHTTPClient(client *http.Client) *ChatNewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chat new params
func (o *ChatNewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelID adds the channelID to the chat new params
func (o *ChatNewParams) WithChannelID(channelID *float64) *ChatNewParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the chat new params
func (o *ChatNewParams) SetChannelID(channelID *float64) {
	o.ChannelID = channelID
}

// WithMessage adds the message to the chat new params
func (o *ChatNewParams) WithMessage(message string) *ChatNewParams {
	o.SetMessage(message)
	return o
}

// SetMessage adds the message to the chat new params
func (o *ChatNewParams) SetMessage(message string) {
	o.Message = message
}

// WriteToRequest writes these params to a swagger request
func (o *ChatNewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChannelID != nil {

		// form param channelID
		var frChannelID float64
		if o.ChannelID != nil {
			frChannelID = *o.ChannelID
		}
		fChannelID := swag.FormatFloat64(frChannelID)
		if fChannelID != "" {
			if err := r.SetFormParam("channelID", fChannelID); err != nil {
				return err
			}
		}

	}

	// form param message
	frMessage := o.Message
	fMessage := frMessage
	if fMessage != "" {
		if err := r.SetFormParam("message", fMessage); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
