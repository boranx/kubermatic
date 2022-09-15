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

// DeleteRoleReader is a Reader for the DeleteRole structure.
type DeleteRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRoleOK creates a DeleteRoleOK with default headers values
func NewDeleteRoleOK() *DeleteRoleOK {
	return &DeleteRoleOK{}
}

/*
DeleteRoleOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteRoleOK struct {
}

// IsSuccess returns true when this delete role o k response has a 2xx status code
func (o *DeleteRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete role o k response has a 3xx status code
func (o *DeleteRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role o k response has a 4xx status code
func (o *DeleteRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete role o k response has a 5xx status code
func (o *DeleteRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete role o k response a status code equal to that given
func (o *DeleteRoleOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteRoleOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] deleteRoleOK ", 200)
}

func (o *DeleteRoleOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] deleteRoleOK ", 200)
}

func (o *DeleteRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoleUnauthorized creates a DeleteRoleUnauthorized with default headers values
func NewDeleteRoleUnauthorized() *DeleteRoleUnauthorized {
	return &DeleteRoleUnauthorized{}
}

/*
DeleteRoleUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteRoleUnauthorized struct {
}

// IsSuccess returns true when this delete role unauthorized response has a 2xx status code
func (o *DeleteRoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete role unauthorized response has a 3xx status code
func (o *DeleteRoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role unauthorized response has a 4xx status code
func (o *DeleteRoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete role unauthorized response has a 5xx status code
func (o *DeleteRoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete role unauthorized response a status code equal to that given
func (o *DeleteRoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteRoleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] deleteRoleUnauthorized ", 401)
}

func (o *DeleteRoleUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] deleteRoleUnauthorized ", 401)
}

func (o *DeleteRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoleForbidden creates a DeleteRoleForbidden with default headers values
func NewDeleteRoleForbidden() *DeleteRoleForbidden {
	return &DeleteRoleForbidden{}
}

/*
DeleteRoleForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteRoleForbidden struct {
}

// IsSuccess returns true when this delete role forbidden response has a 2xx status code
func (o *DeleteRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete role forbidden response has a 3xx status code
func (o *DeleteRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role forbidden response has a 4xx status code
func (o *DeleteRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete role forbidden response has a 5xx status code
func (o *DeleteRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete role forbidden response a status code equal to that given
func (o *DeleteRoleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRoleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] deleteRoleForbidden ", 403)
}

func (o *DeleteRoleForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] deleteRoleForbidden ", 403)
}

func (o *DeleteRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoleDefault creates a DeleteRoleDefault with default headers values
func NewDeleteRoleDefault(code int) *DeleteRoleDefault {
	return &DeleteRoleDefault{
		_statusCode: code,
	}
}

/*
DeleteRoleDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete role default response
func (o *DeleteRoleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete role default response has a 2xx status code
func (o *DeleteRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete role default response has a 3xx status code
func (o *DeleteRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete role default response has a 4xx status code
func (o *DeleteRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete role default response has a 5xx status code
func (o *DeleteRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete role default response a status code equal to that given
func (o *DeleteRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteRoleDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] deleteRole default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRoleDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] deleteRole default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRoleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
