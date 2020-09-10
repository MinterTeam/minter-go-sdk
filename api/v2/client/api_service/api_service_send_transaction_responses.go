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

// APIServiceSendTransactionReader is a Reader for the APIServiceSendTransaction structure.
type APIServiceSendTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceSendTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceSendTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceSendTransactionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceSendTransactionOK creates a APIServiceSendTransactionOK with default headers values
func NewAPIServiceSendTransactionOK() *APIServiceSendTransactionOK {
	return &APIServiceSendTransactionOK{}
}

/*APIServiceSendTransactionOK handles this case with default header values.

A successful response.
*/
type APIServiceSendTransactionOK struct {
	Payload *models.APIPbSendTransactionResponse
}

func (o *APIServiceSendTransactionOK) Error() string {
	return fmt.Sprintf("[GET /send_transaction/{tx}][%d] apiServiceSendTransactionOK  %+v", 200, o.Payload)
}

func (o *APIServiceSendTransactionOK) GetPayload() *models.APIPbSendTransactionResponse {
	return o.Payload
}

func (o *APIServiceSendTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbSendTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceSendTransactionDefault creates a APIServiceSendTransactionDefault with default headers values
func NewAPIServiceSendTransactionDefault(code int) *APIServiceSendTransactionDefault {
	return &APIServiceSendTransactionDefault{
		_statusCode: code,
	}
}

/*APIServiceSendTransactionDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceSendTransactionDefault struct {
	_statusCode int

	Payload *models.APIPbErrorBody
}

// Code gets the status code for the Api service send transaction default response
func (o *APIServiceSendTransactionDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceSendTransactionDefault) Error() string {
	return fmt.Sprintf("[GET /send_transaction/{tx}][%d] ApiService_SendTransaction default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceSendTransactionDefault) GetPayload() *models.APIPbErrorBody {
	return o.Payload
}

func (o *APIServiceSendTransactionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
