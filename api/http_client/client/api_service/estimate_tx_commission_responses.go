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

// EstimateTxCommissionReader is a Reader for the EstimateTxCommission structure.
type EstimateTxCommissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EstimateTxCommissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEstimateTxCommissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEstimateTxCommissionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEstimateTxCommissionOK creates a EstimateTxCommissionOK with default headers values
func NewEstimateTxCommissionOK() *EstimateTxCommissionOK {
	return &EstimateTxCommissionOK{}
}

/* EstimateTxCommissionOK describes a response with status code 200, with default header values.

A successful response.
*/
type EstimateTxCommissionOK struct {
	Payload *models.EstimateTxCommissionResponse
}

func (o *EstimateTxCommissionOK) Error() string {
	return fmt.Sprintf("[GET /estimate_tx_commission/{tx}][%d] estimateTxCommissionOK  %+v", 200, o.Payload)
}
func (o *EstimateTxCommissionOK) GetPayload() *models.EstimateTxCommissionResponse {
	return o.Payload
}

func (o *EstimateTxCommissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EstimateTxCommissionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEstimateTxCommissionDefault creates a EstimateTxCommissionDefault with default headers values
func NewEstimateTxCommissionDefault(code int) *EstimateTxCommissionDefault {
	return &EstimateTxCommissionDefault{
		_statusCode: code,
	}
}

/* EstimateTxCommissionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EstimateTxCommissionDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the estimate tx commission default response
func (o *EstimateTxCommissionDefault) Code() int {
	return o._statusCode
}

func (o *EstimateTxCommissionDefault) Error() string {
	return fmt.Sprintf("[GET /estimate_tx_commission/{tx}][%d] EstimateTxCommission default  %+v", o._statusCode, o.Payload)
}
func (o *EstimateTxCommissionDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *EstimateTxCommissionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
