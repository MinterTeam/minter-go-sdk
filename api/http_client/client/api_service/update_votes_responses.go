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

// UpdateVotesReader is a Reader for the UpdateVotes structure.
type UpdateVotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVotesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateVotesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateVotesOK creates a UpdateVotesOK with default headers values
func NewUpdateVotesOK() *UpdateVotesOK {
	return &UpdateVotesOK{}
}

/* UpdateVotesOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateVotesOK struct {
	Payload *models.UpdateVotesResponse
}

func (o *UpdateVotesOK) Error() string {
	return fmt.Sprintf("[GET /update_votes/{target_version}][%d] updateVotesOK  %+v", 200, o.Payload)
}
func (o *UpdateVotesOK) GetPayload() *models.UpdateVotesResponse {
	return o.Payload
}

func (o *UpdateVotesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateVotesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVotesDefault creates a UpdateVotesDefault with default headers values
func NewUpdateVotesDefault(code int) *UpdateVotesDefault {
	return &UpdateVotesDefault{
		_statusCode: code,
	}
}

/* UpdateVotesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateVotesDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the update votes default response
func (o *UpdateVotesDefault) Code() int {
	return o._statusCode
}

func (o *UpdateVotesDefault) Error() string {
	return fmt.Sprintf("[GET /update_votes/{target_version}][%d] UpdateVotes default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateVotesDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVotesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
