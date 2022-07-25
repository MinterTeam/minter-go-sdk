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
	"github.com/go-openapi/swag"
)

// NewLimitOrdersByOwnerParams creates a new LimitOrdersByOwnerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLimitOrdersByOwnerParams() *LimitOrdersByOwnerParams {
	return &LimitOrdersByOwnerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLimitOrdersByOwnerParamsWithTimeout creates a new LimitOrdersByOwnerParams object
// with the ability to set a timeout on a request.
func NewLimitOrdersByOwnerParamsWithTimeout(timeout time.Duration) *LimitOrdersByOwnerParams {
	return &LimitOrdersByOwnerParams{
		timeout: timeout,
	}
}

// NewLimitOrdersByOwnerParamsWithContext creates a new LimitOrdersByOwnerParams object
// with the ability to set a context for a request.
func NewLimitOrdersByOwnerParamsWithContext(ctx context.Context) *LimitOrdersByOwnerParams {
	return &LimitOrdersByOwnerParams{
		Context: ctx,
	}
}

// NewLimitOrdersByOwnerParamsWithHTTPClient creates a new LimitOrdersByOwnerParams object
// with the ability to set a custom HTTPClient for a request.
func NewLimitOrdersByOwnerParamsWithHTTPClient(client *http.Client) *LimitOrdersByOwnerParams {
	return &LimitOrdersByOwnerParams{
		HTTPClient: client,
	}
}

/* LimitOrdersByOwnerParams contains all the parameters to send to the API endpoint
   for the limit orders by owner operation.

   Typically these are written to a http.Request.
*/
type LimitOrdersByOwnerParams struct {

	// Address.
	Address string

	// Height.
	//
	// Format: uint64
	Height *uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the limit orders by owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LimitOrdersByOwnerParams) WithDefaults() *LimitOrdersByOwnerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the limit orders by owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LimitOrdersByOwnerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the limit orders by owner params
func (o *LimitOrdersByOwnerParams) WithTimeout(timeout time.Duration) *LimitOrdersByOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the limit orders by owner params
func (o *LimitOrdersByOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the limit orders by owner params
func (o *LimitOrdersByOwnerParams) WithContext(ctx context.Context) *LimitOrdersByOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the limit orders by owner params
func (o *LimitOrdersByOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the limit orders by owner params
func (o *LimitOrdersByOwnerParams) WithHTTPClient(client *http.Client) *LimitOrdersByOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the limit orders by owner params
func (o *LimitOrdersByOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the limit orders by owner params
func (o *LimitOrdersByOwnerParams) WithAddress(address string) *LimitOrdersByOwnerParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the limit orders by owner params
func (o *LimitOrdersByOwnerParams) SetAddress(address string) {
	o.Address = address
}

// WithHeight adds the height to the limit orders by owner params
func (o *LimitOrdersByOwnerParams) WithHeight(height *uint64) *LimitOrdersByOwnerParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the limit orders by owner params
func (o *LimitOrdersByOwnerParams) SetHeight(height *uint64) {
	o.Height = height
}

// WriteToRequest writes these params to a swagger request
func (o *LimitOrdersByOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if o.Height != nil {

		// query param height
		var qrHeight uint64

		if o.Height != nil {
			qrHeight = *o.Height
		}
		qHeight := swag.FormatUint64(qrHeight)
		if qHeight != "" {

			if err := r.SetQueryParam("height", qHeight); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
