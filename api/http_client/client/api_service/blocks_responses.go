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

// BlocksReader is a Reader for the Blocks structure.
type BlocksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BlocksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBlocksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBlocksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBlocksOK creates a BlocksOK with default headers values
func NewBlocksOK() *BlocksOK {
	return &BlocksOK{}
}

/* BlocksOK describes a response with status code 200, with default header values.

A successful response.
*/
type BlocksOK struct {
	Payload *models.BlocksResponse
}

func (o *BlocksOK) Error() string {
	return fmt.Sprintf("[GET /blocks][%d] blocksOK  %+v", 200, o.Payload)
}
func (o *BlocksOK) GetPayload() *models.BlocksResponse {
	return o.Payload
}

func (o *BlocksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlocksResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBlocksDefault creates a BlocksDefault with default headers values
func NewBlocksDefault(code int) *BlocksDefault {
	return &BlocksDefault{
		_statusCode: code,
	}
}

/* BlocksDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BlocksDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the blocks default response
func (o *BlocksDefault) Code() int {
	return o._statusCode
}

func (o *BlocksDefault) Error() string {
	return fmt.Sprintf("[GET /blocks][%d] Blocks default  %+v", o._statusCode, o.Payload)
}
func (o *BlocksDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *BlocksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
