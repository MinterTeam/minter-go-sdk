// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MinterTeam/minter-go-sdk/v2/api/v2/models"
)

// APIServiceHaltsReader is a Reader for the APIServiceHalts structure.
type APIServiceHaltsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceHaltsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceHaltsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceHaltsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceHaltsOK creates a APIServiceHaltsOK with default headers values
func NewAPIServiceHaltsOK() *APIServiceHaltsOK {
	return &APIServiceHaltsOK{}
}

/*APIServiceHaltsOK handles this case with default header values.

A successful response.
*/
type APIServiceHaltsOK struct {
	Payload *models.APIPbHaltsResponse
}

func (o *APIServiceHaltsOK) Error() string {
	return fmt.Sprintf("[GET /halts][%d] apiServiceHaltsOK  %+v", 200, o.Payload)
}

func (o *APIServiceHaltsOK) GetPayload() *models.APIPbHaltsResponse {
	return o.Payload
}

func (o *APIServiceHaltsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbHaltsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceHaltsDefault creates a APIServiceHaltsDefault with default headers values
func NewAPIServiceHaltsDefault(code int) *APIServiceHaltsDefault {
	return &APIServiceHaltsDefault{
		_statusCode: code,
	}
}

/*APIServiceHaltsDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceHaltsDefault struct {
	_statusCode int

	Payload *models.APIPbErrorBody
}

// Code gets the status code for the Api service halts default response
func (o *APIServiceHaltsDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceHaltsDefault) Error() string {
	return fmt.Sprintf("[GET /halts][%d] ApiService_Halts default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceHaltsDefault) GetPayload() *models.APIPbErrorBody {
	return o.Payload
}

func (o *APIServiceHaltsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
