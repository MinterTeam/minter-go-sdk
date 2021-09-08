// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
)

// LimitOrderReader is a Reader for the LimitOrder structure.
type LimitOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LimitOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLimitOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLimitOrderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLimitOrderOK creates a LimitOrderOK with default headers values
func NewLimitOrderOK() *LimitOrderOK {
	return &LimitOrderOK{}
}

/* LimitOrderOK describes a response with status code 200, with default header values.

A successful response.
*/
type LimitOrderOK struct {
	Payload *models.LimitOrderResponse
}

func (o *LimitOrderOK) Error() string {
	return fmt.Sprintf("[GET /limit_order/{order_id}][%d] limitOrderOK  %+v", 200, o.Payload)
}
func (o *LimitOrderOK) GetPayload() *models.LimitOrderResponse {
	return o.Payload
}

func (o *LimitOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LimitOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLimitOrderDefault creates a LimitOrderDefault with default headers values
func NewLimitOrderDefault(code int) *LimitOrderDefault {
	return &LimitOrderDefault{
		_statusCode: code,
	}
}

/* LimitOrderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LimitOrderDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the limit order default response
func (o *LimitOrderDefault) Code() int {
	return o._statusCode
}

func (o *LimitOrderDefault) Error() string {
	return fmt.Sprintf("[GET /limit_order/{order_id}][%d] LimitOrder default  %+v", o._statusCode, o.Payload)
}
func (o *LimitOrderDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *LimitOrderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}