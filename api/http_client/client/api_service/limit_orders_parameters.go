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

// NewLimitOrdersParams creates a new LimitOrdersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLimitOrdersParams() *LimitOrdersParams {
	return &LimitOrdersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLimitOrdersParamsWithTimeout creates a new LimitOrdersParams object
// with the ability to set a timeout on a request.
func NewLimitOrdersParamsWithTimeout(timeout time.Duration) *LimitOrdersParams {
	return &LimitOrdersParams{
		timeout: timeout,
	}
}

// NewLimitOrdersParamsWithContext creates a new LimitOrdersParams object
// with the ability to set a context for a request.
func NewLimitOrdersParamsWithContext(ctx context.Context) *LimitOrdersParams {
	return &LimitOrdersParams{
		Context: ctx,
	}
}

// NewLimitOrdersParamsWithHTTPClient creates a new LimitOrdersParams object
// with the ability to set a custom HTTPClient for a request.
func NewLimitOrdersParamsWithHTTPClient(client *http.Client) *LimitOrdersParams {
	return &LimitOrdersParams{
		HTTPClient: client,
	}
}

/* LimitOrdersParams contains all the parameters to send to the API endpoint
   for the limit orders operation.

   Typically these are written to a http.Request.
*/
type LimitOrdersParams struct {

	// Height.
	//
	// Format: uint64
	Height *uint64

	// Ids.
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the limit orders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LimitOrdersParams) WithDefaults() *LimitOrdersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the limit orders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LimitOrdersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the limit orders params
func (o *LimitOrdersParams) WithTimeout(timeout time.Duration) *LimitOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the limit orders params
func (o *LimitOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the limit orders params
func (o *LimitOrdersParams) WithContext(ctx context.Context) *LimitOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the limit orders params
func (o *LimitOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the limit orders params
func (o *LimitOrdersParams) WithHTTPClient(client *http.Client) *LimitOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the limit orders params
func (o *LimitOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the limit orders params
func (o *LimitOrdersParams) WithHeight(height *uint64) *LimitOrdersParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the limit orders params
func (o *LimitOrdersParams) SetHeight(height *uint64) {
	o.Height = height
}

// WithIds adds the ids to the limit orders params
func (o *LimitOrdersParams) WithIds(ids []string) *LimitOrdersParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the limit orders params
func (o *LimitOrdersParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *LimitOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamLimitOrders binds the parameter ids
func (o *LimitOrdersParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "multi"
	idsIS := swag.JoinByFormat(idsIC, "multi")

	return idsIS
}
