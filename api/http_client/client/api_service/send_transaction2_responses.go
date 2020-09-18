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

// SendTransaction2Reader is a Reader for the SendTransaction2 structure.
type SendTransaction2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendTransaction2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendTransaction2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSendTransaction2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSendTransaction2OK creates a SendTransaction2OK with default headers values
func NewSendTransaction2OK() *SendTransaction2OK {
	return &SendTransaction2OK{}
}

/*SendTransaction2OK handles this case with default header values.

A successful response.
*/
type SendTransaction2OK struct {
	Payload *models.SendTransactionResponse
}

func (o *SendTransaction2OK) Error() string {
	return fmt.Sprintf("[POST /send_transaction][%d] sendTransaction2OK  %+v", 200, o.Payload)
}

func (o *SendTransaction2OK) GetPayload() *models.SendTransactionResponse {
	return o.Payload
}

func (o *SendTransaction2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SendTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendTransaction2Default creates a SendTransaction2Default with default headers values
func NewSendTransaction2Default(code int) *SendTransaction2Default {
	return &SendTransaction2Default{
		_statusCode: code,
	}
}

/*SendTransaction2Default handles this case with default header values.

An unexpected error response
*/
type SendTransaction2Default struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the send transaction2 default response
func (o *SendTransaction2Default) Code() int {
	return o._statusCode
}

func (o *SendTransaction2Default) Error() string {
	return fmt.Sprintf("[POST /send_transaction][%d] SendTransaction2 default  %+v", o._statusCode, o.Payload)
}

func (o *SendTransaction2Default) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *SendTransaction2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}