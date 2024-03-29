// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
)

// SubscribeReader is a Reader for the Subscribe structure.
type SubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSubscribeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubscribeOK creates a SubscribeOK with default headers values
func NewSubscribeOK() *SubscribeOK {
	return &SubscribeOK{}
}

/* SubscribeOK describes a response with status code 200, with default header values.

A successful response.(streaming responses)
*/
type SubscribeOK struct {
	Payload *SubscribeOKBody
}

func (o *SubscribeOK) Error() string {
	return fmt.Sprintf("[GET /subscribe][%d] subscribeOK  %+v", 200, o.Payload)
}
func (o *SubscribeOK) GetPayload() *SubscribeOKBody {
	return o.Payload
}

func (o *SubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SubscribeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscribeDefault creates a SubscribeDefault with default headers values
func NewSubscribeDefault(code int) *SubscribeDefault {
	return &SubscribeDefault{
		_statusCode: code,
	}
}

/* SubscribeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SubscribeDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the subscribe default response
func (o *SubscribeDefault) Code() int {
	return o._statusCode
}

func (o *SubscribeDefault) Error() string {
	return fmt.Sprintf("[GET /subscribe][%d] Subscribe default  %+v", o._statusCode, o.Payload)
}
func (o *SubscribeDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *SubscribeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SubscribeOKBody Stream result of SubscribeResponse
swagger:model SubscribeOKBody
*/
type SubscribeOKBody struct {

	// error
	Error interface{} `json:"error,omitempty"`

	// result
	Result *models.SubscribeResponse `json:"result,omitempty"`
}

// Validate validates this subscribe o k body
func (o *SubscribeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SubscribeOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscribeOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this subscribe o k body based on the context it is used
func (o *SubscribeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SubscribeOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscribeOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SubscribeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubscribeOKBody) UnmarshalBinary(b []byte) error {
	var res SubscribeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
