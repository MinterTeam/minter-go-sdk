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

// NewTestBlockParams creates a new TestBlockParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestBlockParams() *TestBlockParams {
	return &TestBlockParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestBlockParamsWithTimeout creates a new TestBlockParams object
// with the ability to set a timeout on a request.
func NewTestBlockParamsWithTimeout(timeout time.Duration) *TestBlockParams {
	return &TestBlockParams{
		timeout: timeout,
	}
}

// NewTestBlockParamsWithContext creates a new TestBlockParams object
// with the ability to set a context for a request.
func NewTestBlockParamsWithContext(ctx context.Context) *TestBlockParams {
	return &TestBlockParams{
		Context: ctx,
	}
}

// NewTestBlockParamsWithHTTPClient creates a new TestBlockParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestBlockParamsWithHTTPClient(client *http.Client) *TestBlockParams {
	return &TestBlockParams{
		HTTPClient: client,
	}
}

/* TestBlockParams contains all the parameters to send to the API endpoint
   for the test block operation.

   Typically these are written to a http.Request.
*/
type TestBlockParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test block params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestBlockParams) WithDefaults() *TestBlockParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test block params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestBlockParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test block params
func (o *TestBlockParams) WithTimeout(timeout time.Duration) *TestBlockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test block params
func (o *TestBlockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test block params
func (o *TestBlockParams) WithContext(ctx context.Context) *TestBlockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test block params
func (o *TestBlockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test block params
func (o *TestBlockParams) WithHTTPClient(client *http.Client) *TestBlockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test block params
func (o *TestBlockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TestBlockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
