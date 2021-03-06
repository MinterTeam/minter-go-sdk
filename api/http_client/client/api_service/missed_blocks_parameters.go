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

// NewMissedBlocksParams creates a new MissedBlocksParams object
// with the default values initialized.
func NewMissedBlocksParams() *MissedBlocksParams {
	var ()
	return &MissedBlocksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMissedBlocksParamsWithTimeout creates a new MissedBlocksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMissedBlocksParamsWithTimeout(timeout time.Duration) *MissedBlocksParams {
	var ()
	return &MissedBlocksParams{

		timeout: timeout,
	}
}

// NewMissedBlocksParamsWithContext creates a new MissedBlocksParams object
// with the default values initialized, and the ability to set a context for a request
func NewMissedBlocksParamsWithContext(ctx context.Context) *MissedBlocksParams {
	var ()
	return &MissedBlocksParams{

		Context: ctx,
	}
}

// NewMissedBlocksParamsWithHTTPClient creates a new MissedBlocksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMissedBlocksParamsWithHTTPClient(client *http.Client) *MissedBlocksParams {
	var ()
	return &MissedBlocksParams{
		HTTPClient: client,
	}
}

/*MissedBlocksParams contains all the parameters to send to the API endpoint
for the missed blocks operation typically these are written to a http.Request
*/
type MissedBlocksParams struct {

	/*Height*/
	Height *uint64
	/*PublicKey*/
	PublicKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the missed blocks params
func (o *MissedBlocksParams) WithTimeout(timeout time.Duration) *MissedBlocksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the missed blocks params
func (o *MissedBlocksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the missed blocks params
func (o *MissedBlocksParams) WithContext(ctx context.Context) *MissedBlocksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the missed blocks params
func (o *MissedBlocksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the missed blocks params
func (o *MissedBlocksParams) WithHTTPClient(client *http.Client) *MissedBlocksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the missed blocks params
func (o *MissedBlocksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the missed blocks params
func (o *MissedBlocksParams) WithHeight(height *uint64) *MissedBlocksParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the missed blocks params
func (o *MissedBlocksParams) SetHeight(height *uint64) {
	o.Height = height
}

// WithPublicKey adds the publicKey to the missed blocks params
func (o *MissedBlocksParams) WithPublicKey(publicKey string) *MissedBlocksParams {
	o.SetPublicKey(publicKey)
	return o
}

// SetPublicKey adds the publicKey to the missed blocks params
func (o *MissedBlocksParams) SetPublicKey(publicKey string) {
	o.PublicKey = publicKey
}

// WriteToRequest writes these params to a swagger request
func (o *MissedBlocksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param public_key
	if err := r.SetPathParam("public_key", o.PublicKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
