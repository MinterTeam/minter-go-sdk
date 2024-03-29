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

// MissedBlocksReader is a Reader for the MissedBlocks structure.
type MissedBlocksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MissedBlocksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMissedBlocksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMissedBlocksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMissedBlocksOK creates a MissedBlocksOK with default headers values
func NewMissedBlocksOK() *MissedBlocksOK {
	return &MissedBlocksOK{}
}

/* MissedBlocksOK describes a response with status code 200, with default header values.

A successful response.
*/
type MissedBlocksOK struct {
	Payload *models.MissedBlocksResponse
}

func (o *MissedBlocksOK) Error() string {
	return fmt.Sprintf("[GET /missed_blocks/{public_key}][%d] missedBlocksOK  %+v", 200, o.Payload)
}
func (o *MissedBlocksOK) GetPayload() *models.MissedBlocksResponse {
	return o.Payload
}

func (o *MissedBlocksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MissedBlocksResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMissedBlocksDefault creates a MissedBlocksDefault with default headers values
func NewMissedBlocksDefault(code int) *MissedBlocksDefault {
	return &MissedBlocksDefault{
		_statusCode: code,
	}
}

/* MissedBlocksDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type MissedBlocksDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the missed blocks default response
func (o *MissedBlocksDefault) Code() int {
	return o._statusCode
}

func (o *MissedBlocksDefault) Error() string {
	return fmt.Sprintf("[GET /missed_blocks/{public_key}][%d] MissedBlocks default  %+v", o._statusCode, o.Payload)
}
func (o *MissedBlocksDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MissedBlocksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
