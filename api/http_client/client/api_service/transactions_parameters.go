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

// NewTransactionsParams creates a new TransactionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTransactionsParams() *TransactionsParams {
	return &TransactionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTransactionsParamsWithTimeout creates a new TransactionsParams object
// with the ability to set a timeout on a request.
func NewTransactionsParamsWithTimeout(timeout time.Duration) *TransactionsParams {
	return &TransactionsParams{
		timeout: timeout,
	}
}

// NewTransactionsParamsWithContext creates a new TransactionsParams object
// with the ability to set a context for a request.
func NewTransactionsParamsWithContext(ctx context.Context) *TransactionsParams {
	return &TransactionsParams{
		Context: ctx,
	}
}

// NewTransactionsParamsWithHTTPClient creates a new TransactionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewTransactionsParamsWithHTTPClient(client *http.Client) *TransactionsParams {
	return &TransactionsParams{
		HTTPClient: client,
	}
}

/* TransactionsParams contains all the parameters to send to the API endpoint
   for the transactions operation.

   Typically these are written to a http.Request.
*/
type TransactionsParams struct {

	// Page.
	//
	// Format: int32
	// Default: 1
	Page *int32

	// PerPage.
	//
	// Format: int32
	// Default: 30
	PerPage *int32

	// Query.
	Query string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransactionsParams) WithDefaults() *TransactionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransactionsParams) SetDefaults() {
	var (
		pageDefault = int32(1)

		perPageDefault = int32(30)
	)

	val := TransactionsParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the transactions params
func (o *TransactionsParams) WithTimeout(timeout time.Duration) *TransactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transactions params
func (o *TransactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transactions params
func (o *TransactionsParams) WithContext(ctx context.Context) *TransactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transactions params
func (o *TransactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transactions params
func (o *TransactionsParams) WithHTTPClient(client *http.Client) *TransactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transactions params
func (o *TransactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the transactions params
func (o *TransactionsParams) WithPage(page *int32) *TransactionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the transactions params
func (o *TransactionsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the transactions params
func (o *TransactionsParams) WithPerPage(perPage *int32) *TransactionsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the transactions params
func (o *TransactionsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithQuery adds the query to the transactions params
func (o *TransactionsParams) WithQuery(query string) *TransactionsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the transactions params
func (o *TransactionsParams) SetQuery(query string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *TransactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

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
