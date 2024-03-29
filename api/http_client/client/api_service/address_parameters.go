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

// NewAddressParams creates a new AddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressParams() *AddressParams {
	return &AddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressParamsWithTimeout creates a new AddressParams object
// with the ability to set a timeout on a request.
func NewAddressParamsWithTimeout(timeout time.Duration) *AddressParams {
	return &AddressParams{
		timeout: timeout,
	}
}

// NewAddressParamsWithContext creates a new AddressParams object
// with the ability to set a context for a request.
func NewAddressParamsWithContext(ctx context.Context) *AddressParams {
	return &AddressParams{
		Context: ctx,
	}
}

// NewAddressParamsWithHTTPClient creates a new AddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressParamsWithHTTPClient(client *http.Client) *AddressParams {
	return &AddressParams{
		HTTPClient: client,
	}
}

/* AddressParams contains all the parameters to send to the API endpoint
   for the address operation.

   Typically these are written to a http.Request.
*/
type AddressParams struct {

	// Address.
	Address string

	// Delegated.
	Delegated *bool

	// Height.
	//
	// Format: uint64
	Height *uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressParams) WithDefaults() *AddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressParams) SetDefaults() {
	var (
		delegatedDefault = bool(false)
	)

	val := AddressParams{
		Delegated: &delegatedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the address params
func (o *AddressParams) WithTimeout(timeout time.Duration) *AddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address params
func (o *AddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address params
func (o *AddressParams) WithContext(ctx context.Context) *AddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address params
func (o *AddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address params
func (o *AddressParams) WithHTTPClient(client *http.Client) *AddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address params
func (o *AddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the address params
func (o *AddressParams) WithAddress(address string) *AddressParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the address params
func (o *AddressParams) SetAddress(address string) {
	o.Address = address
}

// WithDelegated adds the delegated to the address params
func (o *AddressParams) WithDelegated(delegated *bool) *AddressParams {
	o.SetDelegated(delegated)
	return o
}

// SetDelegated adds the delegated to the address params
func (o *AddressParams) SetDelegated(delegated *bool) {
	o.Delegated = delegated
}

// WithHeight adds the height to the address params
func (o *AddressParams) WithHeight(height *uint64) *AddressParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the address params
func (o *AddressParams) SetHeight(height *uint64) {
	o.Height = height
}

// WriteToRequest writes these params to a swagger request
func (o *AddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if o.Delegated != nil {

		// query param delegated
		var qrDelegated bool

		if o.Delegated != nil {
			qrDelegated = *o.Delegated
		}
		qDelegated := swag.FormatBool(qrDelegated)
		if qDelegated != "" {

			if err := r.SetQueryParam("delegated", qDelegated); err != nil {
				return err
			}
		}
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
