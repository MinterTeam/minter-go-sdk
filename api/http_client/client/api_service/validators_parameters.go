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

// NewValidatorsParams creates a new ValidatorsParams object
// with the default values initialized.
func NewValidatorsParams() *ValidatorsParams {
	var ()
	return &ValidatorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidatorsParamsWithTimeout creates a new ValidatorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidatorsParamsWithTimeout(timeout time.Duration) *ValidatorsParams {
	var ()
	return &ValidatorsParams{

		timeout: timeout,
	}
}

// NewValidatorsParamsWithContext creates a new ValidatorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidatorsParamsWithContext(ctx context.Context) *ValidatorsParams {
	var ()
	return &ValidatorsParams{

		Context: ctx,
	}
}

// NewValidatorsParamsWithHTTPClient creates a new ValidatorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidatorsParamsWithHTTPClient(client *http.Client) *ValidatorsParams {
	var ()
	return &ValidatorsParams{
		HTTPClient: client,
	}
}

/*ValidatorsParams contains all the parameters to send to the API endpoint
for the validators operation typically these are written to a http.Request
*/
type ValidatorsParams struct {

	/*Height*/
	Height *string
	/*Page*/
	Page *int32
	/*PerPage*/
	PerPage *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validators params
func (o *ValidatorsParams) WithTimeout(timeout time.Duration) *ValidatorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validators params
func (o *ValidatorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validators params
func (o *ValidatorsParams) WithContext(ctx context.Context) *ValidatorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validators params
func (o *ValidatorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validators params
func (o *ValidatorsParams) WithHTTPClient(client *http.Client) *ValidatorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validators params
func (o *ValidatorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the validators params
func (o *ValidatorsParams) WithHeight(height *string) *ValidatorsParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the validators params
func (o *ValidatorsParams) SetHeight(height *string) {
	o.Height = height
}

// WithPage adds the page to the validators params
func (o *ValidatorsParams) WithPage(page *int32) *ValidatorsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the validators params
func (o *ValidatorsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the validators params
func (o *ValidatorsParams) WithPerPage(perPage *int32) *ValidatorsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the validators params
func (o *ValidatorsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *ValidatorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Height != nil {

		// query param height
		var qrHeight string
		if o.Height != nil {
			qrHeight = *o.Height
		}
		qHeight := qrHeight
		if qHeight != "" {
			if err := r.SetQueryParam("height", qHeight); err != nil {
				return err
			}
		}

	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}