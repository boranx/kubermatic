// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListUserReader is a Reader for the ListUser structure.
type ListUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListUserOK creates a ListUserOK with default headers values
func NewListUserOK() *ListUserOK {
	return &ListUserOK{}
}

/*
ListUserOK describes a response with status code 200, with default header values.

User
*/
type ListUserOK struct {
	Payload []*models.User
}

// IsSuccess returns true when this list user o k response has a 2xx status code
func (o *ListUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list user o k response has a 3xx status code
func (o *ListUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user o k response has a 4xx status code
func (o *ListUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list user o k response has a 5xx status code
func (o *ListUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list user o k response a status code equal to that given
func (o *ListUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users][%d] listUserOK  %+v", 200, o.Payload)
}

func (o *ListUserOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users][%d] listUserOK  %+v", 200, o.Payload)
}

func (o *ListUserOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *ListUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserUnauthorized creates a ListUserUnauthorized with default headers values
func NewListUserUnauthorized() *ListUserUnauthorized {
	return &ListUserUnauthorized{}
}

/*
ListUserUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListUserUnauthorized struct {
}

// IsSuccess returns true when this list user unauthorized response has a 2xx status code
func (o *ListUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user unauthorized response has a 3xx status code
func (o *ListUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user unauthorized response has a 4xx status code
func (o *ListUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list user unauthorized response has a 5xx status code
func (o *ListUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list user unauthorized response a status code equal to that given
func (o *ListUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users][%d] listUserUnauthorized ", 401)
}

func (o *ListUserUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/users][%d] listUserUnauthorized ", 401)
}

func (o *ListUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUserForbidden creates a ListUserForbidden with default headers values
func NewListUserForbidden() *ListUserForbidden {
	return &ListUserForbidden{}
}

/*
ListUserForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListUserForbidden struct {
}

// IsSuccess returns true when this list user forbidden response has a 2xx status code
func (o *ListUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user forbidden response has a 3xx status code
func (o *ListUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user forbidden response has a 4xx status code
func (o *ListUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list user forbidden response has a 5xx status code
func (o *ListUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list user forbidden response a status code equal to that given
func (o *ListUserForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListUserForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users][%d] listUserForbidden ", 403)
}

func (o *ListUserForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/users][%d] listUserForbidden ", 403)
}

func (o *ListUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUserDefault creates a ListUserDefault with default headers values
func NewListUserDefault(code int) *ListUserDefault {
	return &ListUserDefault{
		_statusCode: code,
	}
}

/*
ListUserDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListUserDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list user default response
func (o *ListUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list user default response has a 2xx status code
func (o *ListUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list user default response has a 3xx status code
func (o *ListUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list user default response has a 4xx status code
func (o *ListUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list user default response has a 5xx status code
func (o *ListUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list user default response a status code equal to that given
func (o *ListUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListUserDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/users][%d] listUser default  %+v", o._statusCode, o.Payload)
}

func (o *ListUserDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/users][%d] listUser default  %+v", o._statusCode, o.Payload)
}

func (o *ListUserDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
