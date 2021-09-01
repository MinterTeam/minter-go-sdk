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

// NewWaitListParams creates a new WaitListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaitListParams() *WaitListParams {
	return &WaitListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaitListParamsWithTimeout creates a new WaitListParams object
// with the ability to set a timeout on a request.
func NewWaitListParamsWithTimeout(timeout time.Duration) *WaitListParams {
	return &WaitListParams{
		timeout: timeout,
	}
}

// NewWaitListParamsWithContext creates a new WaitListParams object
// with the ability to set a context for a request.
func NewWaitListParamsWithContext(ctx context.Context) *WaitListParams {
	return &WaitListParams{
		Context: ctx,
	}
}

// NewWaitListParamsWithHTTPClient creates a new WaitListParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaitListParamsWithHTTPClient(client *http.Client) *WaitListParams {
	return &WaitListParams{
		HTTPClient: client,
	}
}

/* WaitListParams contains all the parameters to send to the API endpoint
   for the wait list operation.

   Typically these are written to a http.Request.
*/
type WaitListParams struct {

	// Address.
	Address string

	// Height.
	//
	// Format: uint64
	Height *uint64

	// PublicKey.
	PublicKey *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the wait list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaitListParams) WithDefaults() *WaitListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the wait list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaitListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the wait list params
func (o *WaitListParams) WithTimeout(timeout time.Duration) *WaitListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the wait list params
func (o *WaitListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the wait list params
func (o *WaitListParams) WithContext(ctx context.Context) *WaitListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the wait list params
func (o *WaitListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the wait list params
func (o *WaitListParams) WithHTTPClient(client *http.Client) *WaitListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the wait list params
func (o *WaitListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the wait list params
func (o *WaitListParams) WithAddress(address string) *WaitListParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the wait list params
func (o *WaitListParams) SetAddress(address string) {
	o.Address = address
}

// WithHeight adds the height to the wait list params
func (o *WaitListParams) WithHeight(height *uint64) *WaitListParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the wait list params
func (o *WaitListParams) SetHeight(height *uint64) {
	o.Height = height
}

// WithPublicKey adds the publicKey to the wait list params
func (o *WaitListParams) WithPublicKey(publicKey *string) *WaitListParams {
	o.SetPublicKey(publicKey)
	return o
}

// SetPublicKey adds the publicKey to the wait list params
func (o *WaitListParams) SetPublicKey(publicKey *string) {
	o.PublicKey = publicKey
}

// WriteToRequest writes these params to a swagger request
func (o *WaitListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PublicKey != nil {

		// query param public_key
		var qrPublicKey string

		if o.PublicKey != nil {
			qrPublicKey = *o.PublicKey
		}
		qPublicKey := qrPublicKey
		if qPublicKey != "" {

			if err := r.SetQueryParam("public_key", qPublicKey); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
