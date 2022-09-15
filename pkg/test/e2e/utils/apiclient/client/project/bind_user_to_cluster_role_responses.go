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

// BindUserToClusterRoleReader is a Reader for the BindUserToClusterRole structure.
type BindUserToClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BindUserToClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBindUserToClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBindUserToClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBindUserToClusterRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBindUserToClusterRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBindUserToClusterRoleOK creates a BindUserToClusterRoleOK with default headers values
func NewBindUserToClusterRoleOK() *BindUserToClusterRoleOK {
	return &BindUserToClusterRoleOK{}
}

/*
BindUserToClusterRoleOK describes a response with status code 200, with default header values.

ClusterRoleBinding
*/
type BindUserToClusterRoleOK struct {
	Payload *models.ClusterRoleBinding
}

// IsSuccess returns true when this bind user to cluster role o k response has a 2xx status code
func (o *BindUserToClusterRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bind user to cluster role o k response has a 3xx status code
func (o *BindUserToClusterRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bind user to cluster role o k response has a 4xx status code
func (o *BindUserToClusterRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bind user to cluster role o k response has a 5xx status code
func (o *BindUserToClusterRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bind user to cluster role o k response a status code equal to that given
func (o *BindUserToClusterRoleOK) IsCode(code int) bool {
	return code == 200
}

func (o *BindUserToClusterRoleOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleOK  %+v", 200, o.Payload)
}

func (o *BindUserToClusterRoleOK) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleOK  %+v", 200, o.Payload)
}

func (o *BindUserToClusterRoleOK) GetPayload() *models.ClusterRoleBinding {
	return o.Payload
}

func (o *BindUserToClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterRoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindUserToClusterRoleUnauthorized creates a BindUserToClusterRoleUnauthorized with default headers values
func NewBindUserToClusterRoleUnauthorized() *BindUserToClusterRoleUnauthorized {
	return &BindUserToClusterRoleUnauthorized{}
}

/*
BindUserToClusterRoleUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type BindUserToClusterRoleUnauthorized struct {
}

// IsSuccess returns true when this bind user to cluster role unauthorized response has a 2xx status code
func (o *BindUserToClusterRoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bind user to cluster role unauthorized response has a 3xx status code
func (o *BindUserToClusterRoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bind user to cluster role unauthorized response has a 4xx status code
func (o *BindUserToClusterRoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this bind user to cluster role unauthorized response has a 5xx status code
func (o *BindUserToClusterRoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this bind user to cluster role unauthorized response a status code equal to that given
func (o *BindUserToClusterRoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BindUserToClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleUnauthorized ", 401)
}

func (o *BindUserToClusterRoleUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleUnauthorized ", 401)
}

func (o *BindUserToClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBindUserToClusterRoleForbidden creates a BindUserToClusterRoleForbidden with default headers values
func NewBindUserToClusterRoleForbidden() *BindUserToClusterRoleForbidden {
	return &BindUserToClusterRoleForbidden{}
}

/*
BindUserToClusterRoleForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type BindUserToClusterRoleForbidden struct {
}

// IsSuccess returns true when this bind user to cluster role forbidden response has a 2xx status code
func (o *BindUserToClusterRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bind user to cluster role forbidden response has a 3xx status code
func (o *BindUserToClusterRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bind user to cluster role forbidden response has a 4xx status code
func (o *BindUserToClusterRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this bind user to cluster role forbidden response has a 5xx status code
func (o *BindUserToClusterRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this bind user to cluster role forbidden response a status code equal to that given
func (o *BindUserToClusterRoleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BindUserToClusterRoleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleForbidden ", 403)
}

func (o *BindUserToClusterRoleForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleForbidden ", 403)
}

func (o *BindUserToClusterRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBindUserToClusterRoleDefault creates a BindUserToClusterRoleDefault with default headers values
func NewBindUserToClusterRoleDefault(code int) *BindUserToClusterRoleDefault {
	return &BindUserToClusterRoleDefault{
		_statusCode: code,
	}
}

/*
BindUserToClusterRoleDefault describes a response with status code -1, with default header values.

errorResponse
*/
type BindUserToClusterRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the bind user to cluster role default response
func (o *BindUserToClusterRoleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this bind user to cluster role default response has a 2xx status code
func (o *BindUserToClusterRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this bind user to cluster role default response has a 3xx status code
func (o *BindUserToClusterRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this bind user to cluster role default response has a 4xx status code
func (o *BindUserToClusterRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this bind user to cluster role default response has a 5xx status code
func (o *BindUserToClusterRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this bind user to cluster role default response a status code equal to that given
func (o *BindUserToClusterRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *BindUserToClusterRoleDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRole default  %+v", o._statusCode, o.Payload)
}

func (o *BindUserToClusterRoleDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRole default  %+v", o._statusCode, o.Payload)
}

func (o *BindUserToClusterRoleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BindUserToClusterRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
