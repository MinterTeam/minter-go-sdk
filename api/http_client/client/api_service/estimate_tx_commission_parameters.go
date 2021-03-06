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

// NewEstimateTxCommissionParams creates a new EstimateTxCommissionParams object
// with the default values initialized.
func NewEstimateTxCommissionParams() *EstimateTxCommissionParams {
	var ()
	return &EstimateTxCommissionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEstimateTxCommissionParamsWithTimeout creates a new EstimateTxCommissionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEstimateTxCommissionParamsWithTimeout(timeout time.Duration) *EstimateTxCommissionParams {
	var ()
	return &EstimateTxCommissionParams{

		timeout: timeout,
	}
}

// NewEstimateTxCommissionParamsWithContext creates a new EstimateTxCommissionParams object
// with the default values initialized, and the ability to set a context for a request
func NewEstimateTxCommissionParamsWithContext(ctx context.Context) *EstimateTxCommissionParams {
	var ()
	return &EstimateTxCommissionParams{

		Context: ctx,
	}
}

// NewEstimateTxCommissionParamsWithHTTPClient creates a new EstimateTxCommissionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEstimateTxCommissionParamsWithHTTPClient(client *http.Client) *EstimateTxCommissionParams {
	var ()
	return &EstimateTxCommissionParams{
		HTTPClient: client,
	}
}

/*EstimateTxCommissionParams contains all the parameters to send to the API endpoint
for the estimate tx commission operation typically these are written to a http.Request
*/
type EstimateTxCommissionParams struct {

	/*Height*/
	Height *uint64
	/*Tx*/
	Tx string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the estimate tx commission params
func (o *EstimateTxCommissionParams) WithTimeout(timeout time.Duration) *EstimateTxCommissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the estimate tx commission params
func (o *EstimateTxCommissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the estimate tx commission params
func (o *EstimateTxCommissionParams) WithContext(ctx context.Context) *EstimateTxCommissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the estimate tx commission params
func (o *EstimateTxCommissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the estimate tx commission params
func (o *EstimateTxCommissionParams) WithHTTPClient(client *http.Client) *EstimateTxCommissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the estimate tx commission params
func (o *EstimateTxCommissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the estimate tx commission params
func (o *EstimateTxCommissionParams) WithHeight(height *uint64) *EstimateTxCommissionParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the estimate tx commission params
func (o *EstimateTxCommissionParams) SetHeight(height *uint64) {
	o.Height = height
}

// WithTx adds the tx to the estimate tx commission params
func (o *EstimateTxCommissionParams) WithTx(tx string) *EstimateTxCommissionParams {
	o.SetTx(tx)
	return o
}

// SetTx adds the tx to the estimate tx commission params
func (o *EstimateTxCommissionParams) SetTx(tx string) {
	o.Tx = tx
}

// WriteToRequest writes these params to a swagger request
func (o *EstimateTxCommissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tx
	if err := r.SetPathParam("tx", o.Tx); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
