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

// WaitListsReader is a Reader for the WaitLists structure.
type WaitListsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaitListsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaitListsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaitListsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaitListsOK creates a WaitListsOK with default headers values
func NewWaitListsOK() *WaitListsOK {
	return &WaitListsOK{}
}

/* WaitListsOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaitListsOK struct {
	Payload *models.WaitListsResponse
}

func (o *WaitListsOK) Error() string {
	return fmt.Sprintf("[GET /waitlists][%d] waitListsOK  %+v", 200, o.Payload)
}
func (o *WaitListsOK) GetPayload() *models.WaitListsResponse {
	return o.Payload
}

func (o *WaitListsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WaitListsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaitListsDefault creates a WaitListsDefault with default headers values
func NewWaitListsDefault(code int) *WaitListsDefault {
	return &WaitListsDefault{
		_statusCode: code,
	}
}

/* WaitListsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaitListsDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the wait lists default response
func (o *WaitListsDefault) Code() int {
	return o._statusCode
}

func (o *WaitListsDefault) Error() string {
	return fmt.Sprintf("[GET /waitlists][%d] WaitLists default  %+v", o._statusCode, o.Payload)
}
func (o *WaitListsDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *WaitListsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}