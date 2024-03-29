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

// FrozenReader is a Reader for the Frozen structure.
type FrozenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FrozenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFrozenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFrozenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFrozenOK creates a FrozenOK with default headers values
func NewFrozenOK() *FrozenOK {
	return &FrozenOK{}
}

/* FrozenOK describes a response with status code 200, with default header values.

A successful response.
*/
type FrozenOK struct {
	Payload *models.FrozenResponse
}

func (o *FrozenOK) Error() string {
	return fmt.Sprintf("[GET /frozen/{address}][%d] frozenOK  %+v", 200, o.Payload)
}
func (o *FrozenOK) GetPayload() *models.FrozenResponse {
	return o.Payload
}

func (o *FrozenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrozenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrozenDefault creates a FrozenDefault with default headers values
func NewFrozenDefault(code int) *FrozenDefault {
	return &FrozenDefault{
		_statusCode: code,
	}
}

/* FrozenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type FrozenDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the frozen default response
func (o *FrozenDefault) Code() int {
	return o._statusCode
}

func (o *FrozenDefault) Error() string {
	return fmt.Sprintf("[GET /frozen/{address}][%d] Frozen default  %+v", o._statusCode, o.Payload)
}
func (o *FrozenDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *FrozenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
