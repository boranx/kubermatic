// Code generated by go-swagger; DO NOT EDIT.

package eks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListEKSAMITypesReader is a Reader for the ListEKSAMITypes structure.
type ListEKSAMITypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSAMITypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSAMITypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEKSAMITypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListEKSAMITypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListEKSAMITypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSAMITypesOK creates a ListEKSAMITypesOK with default headers values
func NewListEKSAMITypesOK() *ListEKSAMITypesOK {
	return &ListEKSAMITypesOK{}
}

/*
ListEKSAMITypesOK describes a response with status code 200, with default header values.

EKSAMITypeList
*/
type ListEKSAMITypesOK struct {
	Payload models.EKSAMITypeList
}

// IsSuccess returns true when this list e k s a m i types o k response has a 2xx status code
func (o *ListEKSAMITypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list e k s a m i types o k response has a 3xx status code
func (o *ListEKSAMITypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s a m i types o k response has a 4xx status code
func (o *ListEKSAMITypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list e k s a m i types o k response has a 5xx status code
func (o *ListEKSAMITypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s a m i types o k response a status code equal to that given
func (o *ListEKSAMITypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListEKSAMITypesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/eks/amitypes][%d] listEKSAMITypesOK  %+v", 200, o.Payload)
}

func (o *ListEKSAMITypesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/eks/amitypes][%d] listEKSAMITypesOK  %+v", 200, o.Payload)
}

func (o *ListEKSAMITypesOK) GetPayload() models.EKSAMITypeList {
	return o.Payload
}

func (o *ListEKSAMITypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSAMITypesUnauthorized creates a ListEKSAMITypesUnauthorized with default headers values
func NewListEKSAMITypesUnauthorized() *ListEKSAMITypesUnauthorized {
	return &ListEKSAMITypesUnauthorized{}
}

/*
ListEKSAMITypesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListEKSAMITypesUnauthorized struct {
}

// IsSuccess returns true when this list e k s a m i types unauthorized response has a 2xx status code
func (o *ListEKSAMITypesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list e k s a m i types unauthorized response has a 3xx status code
func (o *ListEKSAMITypesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s a m i types unauthorized response has a 4xx status code
func (o *ListEKSAMITypesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list e k s a m i types unauthorized response has a 5xx status code
func (o *ListEKSAMITypesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s a m i types unauthorized response a status code equal to that given
func (o *ListEKSAMITypesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListEKSAMITypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/eks/amitypes][%d] listEKSAMITypesUnauthorized ", 401)
}

func (o *ListEKSAMITypesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/eks/amitypes][%d] listEKSAMITypesUnauthorized ", 401)
}

func (o *ListEKSAMITypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSAMITypesForbidden creates a ListEKSAMITypesForbidden with default headers values
func NewListEKSAMITypesForbidden() *ListEKSAMITypesForbidden {
	return &ListEKSAMITypesForbidden{}
}

/*
ListEKSAMITypesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListEKSAMITypesForbidden struct {
}

// IsSuccess returns true when this list e k s a m i types forbidden response has a 2xx status code
func (o *ListEKSAMITypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list e k s a m i types forbidden response has a 3xx status code
func (o *ListEKSAMITypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s a m i types forbidden response has a 4xx status code
func (o *ListEKSAMITypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list e k s a m i types forbidden response has a 5xx status code
func (o *ListEKSAMITypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s a m i types forbidden response a status code equal to that given
func (o *ListEKSAMITypesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListEKSAMITypesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/eks/amitypes][%d] listEKSAMITypesForbidden ", 403)
}

func (o *ListEKSAMITypesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/eks/amitypes][%d] listEKSAMITypesForbidden ", 403)
}

func (o *ListEKSAMITypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSAMITypesDefault creates a ListEKSAMITypesDefault with default headers values
func NewListEKSAMITypesDefault(code int) *ListEKSAMITypesDefault {
	return &ListEKSAMITypesDefault{
		_statusCode: code,
	}
}

/*
ListEKSAMITypesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSAMITypesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s a m i types default response
func (o *ListEKSAMITypesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list e k s a m i types default response has a 2xx status code
func (o *ListEKSAMITypesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list e k s a m i types default response has a 3xx status code
func (o *ListEKSAMITypesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list e k s a m i types default response has a 4xx status code
func (o *ListEKSAMITypesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list e k s a m i types default response has a 5xx status code
func (o *ListEKSAMITypesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list e k s a m i types default response a status code equal to that given
func (o *ListEKSAMITypesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListEKSAMITypesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/eks/amitypes][%d] listEKSAMITypes default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSAMITypesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/eks/amitypes][%d] listEKSAMITypes default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSAMITypesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSAMITypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
