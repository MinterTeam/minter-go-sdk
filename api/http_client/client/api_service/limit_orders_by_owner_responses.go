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

// LimitOrdersByOwnerReader is a Reader for the LimitOrdersByOwner structure.
type LimitOrdersByOwnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LimitOrdersByOwnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLimitOrdersByOwnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLimitOrdersByOwnerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLimitOrdersByOwnerOK creates a LimitOrdersByOwnerOK with default headers values
func NewLimitOrdersByOwnerOK() *LimitOrdersByOwnerOK {
	return &LimitOrdersByOwnerOK{}
}

/* LimitOrdersByOwnerOK describes a response with status code 200, with default header values.

A successful response.
*/
type LimitOrdersByOwnerOK struct {
	Payload *models.LimitOrdersResponse
}

func (o *LimitOrdersByOwnerOK) Error() string {
	return fmt.Sprintf("[GET /limit_orders_by_owner/{address}][%d] limitOrdersByOwnerOK  %+v", 200, o.Payload)
}
func (o *LimitOrdersByOwnerOK) GetPayload() *models.LimitOrdersResponse {
	return o.Payload
}

func (o *LimitOrdersByOwnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LimitOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLimitOrdersByOwnerDefault creates a LimitOrdersByOwnerDefault with default headers values
func NewLimitOrdersByOwnerDefault(code int) *LimitOrdersByOwnerDefault {
	return &LimitOrdersByOwnerDefault{
		_statusCode: code,
	}
}

/* LimitOrdersByOwnerDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LimitOrdersByOwnerDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the limit orders by owner default response
func (o *LimitOrdersByOwnerDefault) Code() int {
	return o._statusCode
}

func (o *LimitOrdersByOwnerDefault) Error() string {
	return fmt.Sprintf("[GET /limit_orders_by_owner/{address}][%d] LimitOrdersByOwner default  %+v", o._statusCode, o.Payload)
}
func (o *LimitOrdersByOwnerDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *LimitOrdersByOwnerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
