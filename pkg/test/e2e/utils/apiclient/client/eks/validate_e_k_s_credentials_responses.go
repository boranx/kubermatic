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

// ValidateEKSCredentialsReader is a Reader for the ValidateEKSCredentials structure.
type ValidateEKSCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateEKSCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateEKSCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewValidateEKSCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidateEKSCredentialsOK creates a ValidateEKSCredentialsOK with default headers values
func NewValidateEKSCredentialsOK() *ValidateEKSCredentialsOK {
	return &ValidateEKSCredentialsOK{}
}

/*
ValidateEKSCredentialsOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type ValidateEKSCredentialsOK struct {
}

// IsSuccess returns true when this validate e k s credentials o k response has a 2xx status code
func (o *ValidateEKSCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate e k s credentials o k response has a 3xx status code
func (o *ValidateEKSCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate e k s credentials o k response has a 4xx status code
func (o *ValidateEKSCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate e k s credentials o k response has a 5xx status code
func (o *ValidateEKSCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate e k s credentials o k response a status code equal to that given
func (o *ValidateEKSCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ValidateEKSCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/validatecredentials][%d] validateEKSCredentialsOK ", 200)
}

func (o *ValidateEKSCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/validatecredentials][%d] validateEKSCredentialsOK ", 200)
}

func (o *ValidateEKSCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateEKSCredentialsDefault creates a ValidateEKSCredentialsDefault with default headers values
func NewValidateEKSCredentialsDefault(code int) *ValidateEKSCredentialsDefault {
	return &ValidateEKSCredentialsDefault{
		_statusCode: code,
	}
}

/*
ValidateEKSCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ValidateEKSCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the validate e k s credentials default response
func (o *ValidateEKSCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this validate e k s credentials default response has a 2xx status code
func (o *ValidateEKSCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this validate e k s credentials default response has a 3xx status code
func (o *ValidateEKSCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this validate e k s credentials default response has a 4xx status code
func (o *ValidateEKSCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this validate e k s credentials default response has a 5xx status code
func (o *ValidateEKSCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this validate e k s credentials default response a status code equal to that given
func (o *ValidateEKSCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ValidateEKSCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/validatecredentials][%d] validateEKSCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ValidateEKSCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/validatecredentials][%d] validateEKSCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ValidateEKSCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ValidateEKSCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
