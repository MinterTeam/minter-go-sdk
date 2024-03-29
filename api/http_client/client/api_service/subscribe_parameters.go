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

// NewSubscribeParams creates a new SubscribeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscribeParams() *SubscribeParams {
	return &SubscribeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscribeParamsWithTimeout creates a new SubscribeParams object
// with the ability to set a timeout on a request.
func NewSubscribeParamsWithTimeout(timeout time.Duration) *SubscribeParams {
	return &SubscribeParams{
		timeout: timeout,
	}
}

// NewSubscribeParamsWithContext creates a new SubscribeParams object
// with the ability to set a context for a request.
func NewSubscribeParamsWithContext(ctx context.Context) *SubscribeParams {
	return &SubscribeParams{
		Context: ctx,
	}
}

// NewSubscribeParamsWithHTTPClient creates a new SubscribeParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscribeParamsWithHTTPClient(client *http.Client) *SubscribeParams {
	return &SubscribeParams{
		HTTPClient: client,
	}
}

/* SubscribeParams contains all the parameters to send to the API endpoint
   for the subscribe operation.

   Typically these are written to a http.Request.
*/
type SubscribeParams struct {

	/* Query.

	   tm.event = 'NewBlock' or tm.event = 'Tx'
	*/
	Query string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscribe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscribeParams) WithDefaults() *SubscribeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscribe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscribeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscribe params
func (o *SubscribeParams) WithTimeout(timeout time.Duration) *SubscribeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscribe params
func (o *SubscribeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscribe params
func (o *SubscribeParams) WithContext(ctx context.Context) *SubscribeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscribe params
func (o *SubscribeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscribe params
func (o *SubscribeParams) WithHTTPClient(client *http.Client) *SubscribeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscribe params
func (o *SubscribeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the subscribe params
func (o *SubscribeParams) WithQuery(query string) *SubscribeParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the subscribe params
func (o *SubscribeParams) SetQuery(query string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *SubscribeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param query
	qrQuery := o.Query
	qQuery := qrQuery
	if qQuery != "" {

		if err := r.SetQueryParam("query", qQuery); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
