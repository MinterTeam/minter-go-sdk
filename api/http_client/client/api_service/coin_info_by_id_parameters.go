// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCoinInfoByIDParams creates a new CoinInfoByIDParams object
// with the default values initialized.
func NewCoinInfoByIDParams() *CoinInfoByIDParams {
	var ()
	return &CoinInfoByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCoinInfoByIDParamsWithTimeout creates a new CoinInfoByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCoinInfoByIDParamsWithTimeout(timeout time.Duration) *CoinInfoByIDParams {
	var ()
	return &CoinInfoByIDParams{

		timeout: timeout,
	}
}

// NewCoinInfoByIDParamsWithContext creates a new CoinInfoByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewCoinInfoByIDParamsWithContext(ctx context.Context) *CoinInfoByIDParams {
	var ()
	return &CoinInfoByIDParams{

		Context: ctx,
	}
}

// NewCoinInfoByIDParamsWithHTTPClient creates a new CoinInfoByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCoinInfoByIDParamsWithHTTPClient(client *http.Client) *CoinInfoByIDParams {
	var ()
	return &CoinInfoByIDParams{
		HTTPClient: client,
	}
}

/*CoinInfoByIDParams contains all the parameters to send to the API endpoint
for the coin info by Id operation typically these are written to a http.Request
*/
type CoinInfoByIDParams struct {

	/*Height*/
	Height *string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the coin info by Id params
func (o *CoinInfoByIDParams) WithTimeout(timeout time.Duration) *CoinInfoByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the coin info by Id params
func (o *CoinInfoByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the coin info by Id params
func (o *CoinInfoByIDParams) WithContext(ctx context.Context) *CoinInfoByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the coin info by Id params
func (o *CoinInfoByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the coin info by Id params
func (o *CoinInfoByIDParams) WithHTTPClient(client *http.Client) *CoinInfoByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the coin info by Id params
func (o *CoinInfoByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the coin info by Id params
func (o *CoinInfoByIDParams) WithHeight(height *string) *CoinInfoByIDParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the coin info by Id params
func (o *CoinInfoByIDParams) SetHeight(height *string) {
	o.Height = height
}

// WithID adds the id to the coin info by Id params
func (o *CoinInfoByIDParams) WithID(id string) *CoinInfoByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the coin info by Id params
func (o *CoinInfoByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CoinInfoByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Height != nil {

		// query param height
		var qrHeight string
		if o.Height != nil {
			qrHeight = *o.Height
		}
		qHeight := qrHeight
		if qHeight != "" {
			if err := r.SetQueryParam("height", qHeight); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
