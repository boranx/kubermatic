// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListGCPZonesReader is a Reader for the ListGCPZones structure.
type ListGCPZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGCPZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGCPZonesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPZonesOK creates a ListGCPZonesOK with default headers values
func NewListGCPZonesOK() *ListGCPZonesOK {
	return &ListGCPZonesOK{}
}

/*
ListGCPZonesOK describes a response with status code 200, with default header values.

GCPZoneList
*/
type ListGCPZonesOK struct {
	Payload models.GCPZoneList
}

// IsSuccess returns true when this list g c p zones o k response has a 2xx status code
func (o *ListGCPZonesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list g c p zones o k response has a 3xx status code
func (o *ListGCPZonesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g c p zones o k response has a 4xx status code
func (o *ListGCPZonesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list g c p zones o k response has a 5xx status code
func (o *ListGCPZonesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list g c p zones o k response a status code equal to that given
func (o *ListGCPZonesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListGCPZonesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/{dc}/zones][%d] listGCPZonesOK  %+v", 200, o.Payload)
}

func (o *ListGCPZonesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/{dc}/zones][%d] listGCPZonesOK  %+v", 200, o.Payload)
}

func (o *ListGCPZonesOK) GetPayload() models.GCPZoneList {
	return o.Payload
}

func (o *ListGCPZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPZonesDefault creates a ListGCPZonesDefault with default headers values
func NewListGCPZonesDefault(code int) *ListGCPZonesDefault {
	return &ListGCPZonesDefault{
		_statusCode: code,
	}
}

/*
ListGCPZonesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGCPZonesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g c p zones default response
func (o *ListGCPZonesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list g c p zones default response has a 2xx status code
func (o *ListGCPZonesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list g c p zones default response has a 3xx status code
func (o *ListGCPZonesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list g c p zones default response has a 4xx status code
func (o *ListGCPZonesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list g c p zones default response has a 5xx status code
func (o *ListGCPZonesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list g c p zones default response a status code equal to that given
func (o *ListGCPZonesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListGCPZonesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/{dc}/zones][%d] listGCPZones default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPZonesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/gcp/{dc}/zones][%d] listGCPZones default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPZonesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGCPZonesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
