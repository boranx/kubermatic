// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// CreateOIDCKubeconfigReader is a Reader for the CreateOIDCKubeconfig structure.
type CreateOIDCKubeconfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOIDCKubeconfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOIDCKubeconfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateOIDCKubeconfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOIDCKubeconfigOK creates a CreateOIDCKubeconfigOK with default headers values
func NewCreateOIDCKubeconfigOK() *CreateOIDCKubeconfigOK {
	return &CreateOIDCKubeconfigOK{}
}

/*
CreateOIDCKubeconfigOK describes a response with status code 200, with default header values.

Kubeconfig is a clusters kubeconfig
*/
type CreateOIDCKubeconfigOK struct {
	Payload []uint8
}

// IsSuccess returns true when this create o Id c kubeconfig o k response has a 2xx status code
func (o *CreateOIDCKubeconfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create o Id c kubeconfig o k response has a 3xx status code
func (o *CreateOIDCKubeconfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create o Id c kubeconfig o k response has a 4xx status code
func (o *CreateOIDCKubeconfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create o Id c kubeconfig o k response has a 5xx status code
func (o *CreateOIDCKubeconfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create o Id c kubeconfig o k response a status code equal to that given
func (o *CreateOIDCKubeconfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateOIDCKubeconfigOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/kubeconfig][%d] createOIdCKubeconfigOK  %+v", 200, o.Payload)
}

func (o *CreateOIDCKubeconfigOK) String() string {
	return fmt.Sprintf("[GET /api/v1/kubeconfig][%d] createOIdCKubeconfigOK  %+v", 200, o.Payload)
}

func (o *CreateOIDCKubeconfigOK) GetPayload() []uint8 {
	return o.Payload
}

func (o *CreateOIDCKubeconfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOIDCKubeconfigDefault creates a CreateOIDCKubeconfigDefault with default headers values
func NewCreateOIDCKubeconfigDefault(code int) *CreateOIDCKubeconfigDefault {
	return &CreateOIDCKubeconfigDefault{
		_statusCode: code,
	}
}

/*
CreateOIDCKubeconfigDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateOIDCKubeconfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create o ID c kubeconfig default response
func (o *CreateOIDCKubeconfigDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create o ID c kubeconfig default response has a 2xx status code
func (o *CreateOIDCKubeconfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create o ID c kubeconfig default response has a 3xx status code
func (o *CreateOIDCKubeconfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create o ID c kubeconfig default response has a 4xx status code
func (o *CreateOIDCKubeconfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create o ID c kubeconfig default response has a 5xx status code
func (o *CreateOIDCKubeconfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create o ID c kubeconfig default response a status code equal to that given
func (o *CreateOIDCKubeconfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateOIDCKubeconfigDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/kubeconfig][%d] createOIDCKubeconfig default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOIDCKubeconfigDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/kubeconfig][%d] createOIDCKubeconfig default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOIDCKubeconfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateOIDCKubeconfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
