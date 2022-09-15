// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// CreateGatekeeperConfigReader is a Reader for the CreateGatekeeperConfig structure.
type CreateGatekeeperConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGatekeeperConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateGatekeeperConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateGatekeeperConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateGatekeeperConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateGatekeeperConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateGatekeeperConfigCreated creates a CreateGatekeeperConfigCreated with default headers values
func NewCreateGatekeeperConfigCreated() *CreateGatekeeperConfigCreated {
	return &CreateGatekeeperConfigCreated{}
}

/*
CreateGatekeeperConfigCreated describes a response with status code 201, with default header values.

GatekeeperConfig
*/
type CreateGatekeeperConfigCreated struct {
	Payload *models.GatekeeperConfig
}

// IsSuccess returns true when this create gatekeeper config created response has a 2xx status code
func (o *CreateGatekeeperConfigCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create gatekeeper config created response has a 3xx status code
func (o *CreateGatekeeperConfigCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gatekeeper config created response has a 4xx status code
func (o *CreateGatekeeperConfigCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create gatekeeper config created response has a 5xx status code
func (o *CreateGatekeeperConfigCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create gatekeeper config created response a status code equal to that given
func (o *CreateGatekeeperConfigCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateGatekeeperConfigCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] createGatekeeperConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateGatekeeperConfigCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] createGatekeeperConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateGatekeeperConfigCreated) GetPayload() *models.GatekeeperConfig {
	return o.Payload
}

func (o *CreateGatekeeperConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatekeeperConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGatekeeperConfigUnauthorized creates a CreateGatekeeperConfigUnauthorized with default headers values
func NewCreateGatekeeperConfigUnauthorized() *CreateGatekeeperConfigUnauthorized {
	return &CreateGatekeeperConfigUnauthorized{}
}

/*
CreateGatekeeperConfigUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateGatekeeperConfigUnauthorized struct {
}

// IsSuccess returns true when this create gatekeeper config unauthorized response has a 2xx status code
func (o *CreateGatekeeperConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create gatekeeper config unauthorized response has a 3xx status code
func (o *CreateGatekeeperConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gatekeeper config unauthorized response has a 4xx status code
func (o *CreateGatekeeperConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create gatekeeper config unauthorized response has a 5xx status code
func (o *CreateGatekeeperConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create gatekeeper config unauthorized response a status code equal to that given
func (o *CreateGatekeeperConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateGatekeeperConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] createGatekeeperConfigUnauthorized ", 401)
}

func (o *CreateGatekeeperConfigUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] createGatekeeperConfigUnauthorized ", 401)
}

func (o *CreateGatekeeperConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGatekeeperConfigForbidden creates a CreateGatekeeperConfigForbidden with default headers values
func NewCreateGatekeeperConfigForbidden() *CreateGatekeeperConfigForbidden {
	return &CreateGatekeeperConfigForbidden{}
}

/*
CreateGatekeeperConfigForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateGatekeeperConfigForbidden struct {
}

// IsSuccess returns true when this create gatekeeper config forbidden response has a 2xx status code
func (o *CreateGatekeeperConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create gatekeeper config forbidden response has a 3xx status code
func (o *CreateGatekeeperConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gatekeeper config forbidden response has a 4xx status code
func (o *CreateGatekeeperConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create gatekeeper config forbidden response has a 5xx status code
func (o *CreateGatekeeperConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create gatekeeper config forbidden response a status code equal to that given
func (o *CreateGatekeeperConfigForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateGatekeeperConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] createGatekeeperConfigForbidden ", 403)
}

func (o *CreateGatekeeperConfigForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] createGatekeeperConfigForbidden ", 403)
}

func (o *CreateGatekeeperConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGatekeeperConfigDefault creates a CreateGatekeeperConfigDefault with default headers values
func NewCreateGatekeeperConfigDefault(code int) *CreateGatekeeperConfigDefault {
	return &CreateGatekeeperConfigDefault{
		_statusCode: code,
	}
}

/*
CreateGatekeeperConfigDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateGatekeeperConfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create gatekeeper config default response
func (o *CreateGatekeeperConfigDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create gatekeeper config default response has a 2xx status code
func (o *CreateGatekeeperConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create gatekeeper config default response has a 3xx status code
func (o *CreateGatekeeperConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create gatekeeper config default response has a 4xx status code
func (o *CreateGatekeeperConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create gatekeeper config default response has a 5xx status code
func (o *CreateGatekeeperConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create gatekeeper config default response a status code equal to that given
func (o *CreateGatekeeperConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateGatekeeperConfigDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] createGatekeeperConfig default  %+v", o._statusCode, o.Payload)
}

func (o *CreateGatekeeperConfigDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] createGatekeeperConfig default  %+v", o._statusCode, o.Payload)
}

func (o *CreateGatekeeperConfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateGatekeeperConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
