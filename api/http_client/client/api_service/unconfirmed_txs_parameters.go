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

// NewUnconfirmedTxsParams creates a new UnconfirmedTxsParams object
// with the default values initialized.
func NewUnconfirmedTxsParams() *UnconfirmedTxsParams {
	var (
		limitDefault = int32(30)
	)
	return &UnconfirmedTxsParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUnconfirmedTxsParamsWithTimeout creates a new UnconfirmedTxsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnconfirmedTxsParamsWithTimeout(timeout time.Duration) *UnconfirmedTxsParams {
	var (
		limitDefault = int32(30)
	)
	return &UnconfirmedTxsParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewUnconfirmedTxsParamsWithContext creates a new UnconfirmedTxsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnconfirmedTxsParamsWithContext(ctx context.Context) *UnconfirmedTxsParams {
	var (
		limitDefault = int32(30)
	)
	return &UnconfirmedTxsParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewUnconfirmedTxsParamsWithHTTPClient creates a new UnconfirmedTxsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnconfirmedTxsParamsWithHTTPClient(client *http.Client) *UnconfirmedTxsParams {
	var (
		limitDefault = int32(30)
	)
	return &UnconfirmedTxsParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*UnconfirmedTxsParams contains all the parameters to send to the API endpoint
for the unconfirmed txs operation typically these are written to a http.Request
*/
type UnconfirmedTxsParams struct {

	/*Limit*/
	Limit *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the unconfirmed txs params
func (o *UnconfirmedTxsParams) WithTimeout(timeout time.Duration) *UnconfirmedTxsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unconfirmed txs params
func (o *UnconfirmedTxsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unconfirmed txs params
func (o *UnconfirmedTxsParams) WithContext(ctx context.Context) *UnconfirmedTxsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unconfirmed txs params
func (o *UnconfirmedTxsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unconfirmed txs params
func (o *UnconfirmedTxsParams) WithHTTPClient(client *http.Client) *UnconfirmedTxsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unconfirmed txs params
func (o *UnconfirmedTxsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the unconfirmed txs params
func (o *UnconfirmedTxsParams) WithLimit(limit *int32) *UnconfirmedTxsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the unconfirmed txs params
func (o *UnconfirmedTxsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *UnconfirmedTxsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param page
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
