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

// CandidatesReader is a Reader for the Candidates structure.
type CandidatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CandidatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCandidatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCandidatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCandidatesOK creates a CandidatesOK with default headers values
func NewCandidatesOK() *CandidatesOK {
	return &CandidatesOK{}
}

/* CandidatesOK describes a response with status code 200, with default header values.

A successful response.
*/
type CandidatesOK struct {
	Payload *models.CandidatesResponse
}

func (o *CandidatesOK) Error() string {
	return fmt.Sprintf("[GET /candidates][%d] candidatesOK  %+v", 200, o.Payload)
}
func (o *CandidatesOK) GetPayload() *models.CandidatesResponse {
	return o.Payload
}

func (o *CandidatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CandidatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCandidatesDefault creates a CandidatesDefault with default headers values
func NewCandidatesDefault(code int) *CandidatesDefault {
	return &CandidatesDefault{
		_statusCode: code,
	}
}

/* CandidatesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CandidatesDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the candidates default response
func (o *CandidatesDefault) Code() int {
	return o._statusCode
}

func (o *CandidatesDefault) Error() string {
	return fmt.Sprintf("[GET /candidates][%d] Candidates default  %+v", o._statusCode, o.Payload)
}
func (o *CandidatesDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CandidatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
