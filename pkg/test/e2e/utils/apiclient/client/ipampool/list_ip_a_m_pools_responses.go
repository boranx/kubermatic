// Code generated by go-swagger; DO NOT EDIT.

package ipampool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListIPAMPoolsReader is a Reader for the ListIPAMPools structure.
type ListIPAMPoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIPAMPoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListIPAMPoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListIPAMPoolsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListIPAMPoolsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListIPAMPoolsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListIPAMPoolsOK creates a ListIPAMPoolsOK with default headers values
func NewListIPAMPoolsOK() *ListIPAMPoolsOK {
	return &ListIPAMPoolsOK{}
}

/*
ListIPAMPoolsOK describes a response with status code 200, with default header values.

IPAMPool
*/
type ListIPAMPoolsOK struct {
	Payload []*models.IPAMPool
}

// IsSuccess returns true when this list Ip a m pools o k response has a 2xx status code
func (o *ListIPAMPoolsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list Ip a m pools o k response has a 3xx status code
func (o *ListIPAMPoolsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list Ip a m pools o k response has a 4xx status code
func (o *ListIPAMPoolsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list Ip a m pools o k response has a 5xx status code
func (o *ListIPAMPoolsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list Ip a m pools o k response a status code equal to that given
func (o *ListIPAMPoolsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListIPAMPoolsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools][%d] listIpAMPoolsOK  %+v", 200, o.Payload)
}

func (o *ListIPAMPoolsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools][%d] listIpAMPoolsOK  %+v", 200, o.Payload)
}

func (o *ListIPAMPoolsOK) GetPayload() []*models.IPAMPool {
	return o.Payload
}

func (o *ListIPAMPoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIPAMPoolsUnauthorized creates a ListIPAMPoolsUnauthorized with default headers values
func NewListIPAMPoolsUnauthorized() *ListIPAMPoolsUnauthorized {
	return &ListIPAMPoolsUnauthorized{}
}

/*
ListIPAMPoolsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListIPAMPoolsUnauthorized struct {
}

// IsSuccess returns true when this list Ip a m pools unauthorized response has a 2xx status code
func (o *ListIPAMPoolsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list Ip a m pools unauthorized response has a 3xx status code
func (o *ListIPAMPoolsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list Ip a m pools unauthorized response has a 4xx status code
func (o *ListIPAMPoolsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list Ip a m pools unauthorized response has a 5xx status code
func (o *ListIPAMPoolsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list Ip a m pools unauthorized response a status code equal to that given
func (o *ListIPAMPoolsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListIPAMPoolsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools][%d] listIpAMPoolsUnauthorized ", 401)
}

func (o *ListIPAMPoolsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools][%d] listIpAMPoolsUnauthorized ", 401)
}

func (o *ListIPAMPoolsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListIPAMPoolsForbidden creates a ListIPAMPoolsForbidden with default headers values
func NewListIPAMPoolsForbidden() *ListIPAMPoolsForbidden {
	return &ListIPAMPoolsForbidden{}
}

/*
ListIPAMPoolsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListIPAMPoolsForbidden struct {
}

// IsSuccess returns true when this list Ip a m pools forbidden response has a 2xx status code
func (o *ListIPAMPoolsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list Ip a m pools forbidden response has a 3xx status code
func (o *ListIPAMPoolsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list Ip a m pools forbidden response has a 4xx status code
func (o *ListIPAMPoolsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list Ip a m pools forbidden response has a 5xx status code
func (o *ListIPAMPoolsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list Ip a m pools forbidden response a status code equal to that given
func (o *ListIPAMPoolsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListIPAMPoolsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools][%d] listIpAMPoolsForbidden ", 403)
}

func (o *ListIPAMPoolsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools][%d] listIpAMPoolsForbidden ", 403)
}

func (o *ListIPAMPoolsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListIPAMPoolsDefault creates a ListIPAMPoolsDefault with default headers values
func NewListIPAMPoolsDefault(code int) *ListIPAMPoolsDefault {
	return &ListIPAMPoolsDefault{
		_statusCode: code,
	}
}

/*
ListIPAMPoolsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListIPAMPoolsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list IP a m pools default response
func (o *ListIPAMPoolsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list IP a m pools default response has a 2xx status code
func (o *ListIPAMPoolsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list IP a m pools default response has a 3xx status code
func (o *ListIPAMPoolsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list IP a m pools default response has a 4xx status code
func (o *ListIPAMPoolsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list IP a m pools default response has a 5xx status code
func (o *ListIPAMPoolsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list IP a m pools default response a status code equal to that given
func (o *ListIPAMPoolsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListIPAMPoolsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools][%d] listIPAMPools default  %+v", o._statusCode, o.Payload)
}

func (o *ListIPAMPoolsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/ipampools][%d] listIPAMPools default  %+v", o._statusCode, o.Payload)
}

func (o *ListIPAMPoolsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListIPAMPoolsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
