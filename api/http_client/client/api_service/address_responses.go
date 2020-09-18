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

// AddressReader is a Reader for the Address structure.
type AddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddressOK creates a AddressOK with default headers values
func NewAddressOK() *AddressOK {
	return &AddressOK{}
}

/*AddressOK handles this case with default header values.

A successful response.
*/
type AddressOK struct {
	Payload *models.AddressResponse
}

func (o *AddressOK) Error() string {
	return fmt.Sprintf("[GET /address/{address}][%d] addressOK  %+v", 200, o.Payload)
}

func (o *AddressOK) GetPayload() *models.AddressResponse {
	return o.Payload
}

func (o *AddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressDefault creates a AddressDefault with default headers values
func NewAddressDefault(code int) *AddressDefault {
	return &AddressDefault{
		_statusCode: code,
	}
}

/*AddressDefault handles this case with default header values.

An unexpected error response
*/
type AddressDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the address default response
func (o *AddressDefault) Code() int {
	return o._statusCode
}

func (o *AddressDefault) Error() string {
	return fmt.Sprintf("[GET /address/{address}][%d] Address default  %+v", o._statusCode, o.Payload)
}

func (o *AddressDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *AddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}