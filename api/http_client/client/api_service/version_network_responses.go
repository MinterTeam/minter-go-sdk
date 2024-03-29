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

// VersionNetworkReader is a Reader for the VersionNetwork structure.
type VersionNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersionNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVersionNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVersionNetworkOK creates a VersionNetworkOK with default headers values
func NewVersionNetworkOK() *VersionNetworkOK {
	return &VersionNetworkOK{}
}

/* VersionNetworkOK describes a response with status code 200, with default header values.

A successful response.
*/
type VersionNetworkOK struct {
	Payload *models.VersionNetworkResponse
}

func (o *VersionNetworkOK) Error() string {
	return fmt.Sprintf("[GET /version_network][%d] versionNetworkOK  %+v", 200, o.Payload)
}
func (o *VersionNetworkOK) GetPayload() *models.VersionNetworkResponse {
	return o.Payload
}

func (o *VersionNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionNetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionNetworkDefault creates a VersionNetworkDefault with default headers values
func NewVersionNetworkDefault(code int) *VersionNetworkDefault {
	return &VersionNetworkDefault{
		_statusCode: code,
	}
}

/* VersionNetworkDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VersionNetworkDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the version network default response
func (o *VersionNetworkDefault) Code() int {
	return o._statusCode
}

func (o *VersionNetworkDefault) Error() string {
	return fmt.Sprintf("[GET /version_network][%d] VersionNetwork default  %+v", o._statusCode, o.Payload)
}
func (o *VersionNetworkDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *VersionNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
