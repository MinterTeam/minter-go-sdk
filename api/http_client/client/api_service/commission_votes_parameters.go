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

// NewCommissionVotesParams creates a new CommissionVotesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCommissionVotesParams() *CommissionVotesParams {
	return &CommissionVotesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCommissionVotesParamsWithTimeout creates a new CommissionVotesParams object
// with the ability to set a timeout on a request.
func NewCommissionVotesParamsWithTimeout(timeout time.Duration) *CommissionVotesParams {
	return &CommissionVotesParams{
		timeout: timeout,
	}
}

// NewCommissionVotesParamsWithContext creates a new CommissionVotesParams object
// with the ability to set a context for a request.
func NewCommissionVotesParamsWithContext(ctx context.Context) *CommissionVotesParams {
	return &CommissionVotesParams{
		Context: ctx,
	}
}

// NewCommissionVotesParamsWithHTTPClient creates a new CommissionVotesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCommissionVotesParamsWithHTTPClient(client *http.Client) *CommissionVotesParams {
	return &CommissionVotesParams{
		HTTPClient: client,
	}
}

/* CommissionVotesParams contains all the parameters to send to the API endpoint
   for the commission votes operation.

   Typically these are written to a http.Request.
*/
type CommissionVotesParams struct {

	// Height.
	//
	// Format: uint64
	Height *uint64

	// TargetVersion.
	//
	// Format: uint64
	TargetVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the commission votes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommissionVotesParams) WithDefaults() *CommissionVotesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the commission votes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommissionVotesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the commission votes params
func (o *CommissionVotesParams) WithTimeout(timeout time.Duration) *CommissionVotesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the commission votes params
func (o *CommissionVotesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the commission votes params
func (o *CommissionVotesParams) WithContext(ctx context.Context) *CommissionVotesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the commission votes params
func (o *CommissionVotesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the commission votes params
func (o *CommissionVotesParams) WithHTTPClient(client *http.Client) *CommissionVotesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the commission votes params
func (o *CommissionVotesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the commission votes params
func (o *CommissionVotesParams) WithHeight(height *uint64) *CommissionVotesParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the commission votes params
func (o *CommissionVotesParams) SetHeight(height *uint64) {
	o.Height = height
}

// WithTargetVersion adds the targetVersion to the commission votes params
func (o *CommissionVotesParams) WithTargetVersion(targetVersion string) *CommissionVotesParams {
	o.SetTargetVersion(targetVersion)
	return o
}

// SetTargetVersion adds the targetVersion to the commission votes params
func (o *CommissionVotesParams) SetTargetVersion(targetVersion string) {
	o.TargetVersion = targetVersion
}

// WriteToRequest writes these params to a swagger request
func (o *CommissionVotesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param target_version
	if err := r.SetPathParam("target_version", o.TargetVersion); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
