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

// UpdateProjectReader is a Reader for the UpdateProject structure.
type UpdateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewUpdateProjectNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateProjectOK creates a UpdateProjectOK with default headers values
func NewUpdateProjectOK() *UpdateProjectOK {
	return &UpdateProjectOK{}
}

/*
UpdateProjectOK describes a response with status code 200, with default header values.

Project
*/
type UpdateProjectOK struct {
	Payload *models.Project
}

// IsSuccess returns true when this update project o k response has a 2xx status code
func (o *UpdateProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update project o k response has a 3xx status code
func (o *UpdateProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project o k response has a 4xx status code
func (o *UpdateProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project o k response has a 5xx status code
func (o *UpdateProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update project o k response a status code equal to that given
func (o *UpdateProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateProjectOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}][%d] updateProjectOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}][%d] updateProjectOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *UpdateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectBadRequest creates a UpdateProjectBadRequest with default headers values
func NewUpdateProjectBadRequest() *UpdateProjectBadRequest {
	return &UpdateProjectBadRequest{}
}

/*
UpdateProjectBadRequest describes a response with status code 400, with default header values.

EmptyResponse is a empty response
*/
type UpdateProjectBadRequest struct {
}

// IsSuccess returns true when this update project bad request response has a 2xx status code
func (o *UpdateProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project bad request response has a 3xx status code
func (o *UpdateProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project bad request response has a 4xx status code
func (o *UpdateProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project bad request response has a 5xx status code
func (o *UpdateProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update project bad request response a status code equal to that given
func (o *UpdateProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateProjectBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}][%d] updateProjectBadRequest ", 400)
}

func (o *UpdateProjectBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}][%d] updateProjectBadRequest ", 400)
}

func (o *UpdateProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProjectNotFound creates a UpdateProjectNotFound with default headers values
func NewUpdateProjectNotFound() *UpdateProjectNotFound {
	return &UpdateProjectNotFound{}
}

/*
UpdateProjectNotFound describes a response with status code 404, with default header values.

EmptyResponse is a empty response
*/
type UpdateProjectNotFound struct {
}

// IsSuccess returns true when this update project not found response has a 2xx status code
func (o *UpdateProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project not found response has a 3xx status code
func (o *UpdateProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project not found response has a 4xx status code
func (o *UpdateProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project not found response has a 5xx status code
func (o *UpdateProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update project not found response a status code equal to that given
func (o *UpdateProjectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateProjectNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}][%d] updateProjectNotFound ", 404)
}

func (o *UpdateProjectNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}][%d] updateProjectNotFound ", 404)
}

func (o *UpdateProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProjectInternalServerError creates a UpdateProjectInternalServerError with default headers values
func NewUpdateProjectInternalServerError() *UpdateProjectInternalServerError {
	return &UpdateProjectInternalServerError{}
}

/*
UpdateProjectInternalServerError describes a response with status code 500, with default header values.

EmptyResponse is a empty response
*/
type UpdateProjectInternalServerError struct {
}

// IsSuccess returns true when this update project internal server error response has a 2xx status code
func (o *UpdateProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project internal server error response has a 3xx status code
func (o *UpdateProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project internal server error response has a 4xx status code
func (o *UpdateProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project internal server error response has a 5xx status code
func (o *UpdateProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update project internal server error response a status code equal to that given
func (o *UpdateProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateProjectInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}][%d] updateProjectInternalServerError ", 500)
}

func (o *UpdateProjectInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}][%d] updateProjectInternalServerError ", 500)
}

func (o *UpdateProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProjectNotImplemented creates a UpdateProjectNotImplemented with default headers values
func NewUpdateProjectNotImplemented() *UpdateProjectNotImplemented {
	return &UpdateProjectNotImplemented{}
}

/*
UpdateProjectNotImplemented describes a response with status code 501, with default header values.

EmptyResponse is a empty response
*/
type UpdateProjectNotImplemented struct {
}

// IsSuccess returns true when this update project not implemented response has a 2xx status code
func (o *UpdateProjectNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project not implemented response has a 3xx status code
func (o *UpdateProjectNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project not implemented response has a 4xx status code
func (o *UpdateProjectNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project not implemented response has a 5xx status code
func (o *UpdateProjectNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this update project not implemented response a status code equal to that given
func (o *UpdateProjectNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *UpdateProjectNotImplemented) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}][%d] updateProjectNotImplemented ", 501)
}

func (o *UpdateProjectNotImplemented) String() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}][%d] updateProjectNotImplemented ", 501)
}

func (o *UpdateProjectNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProjectDefault creates a UpdateProjectDefault with default headers values
func NewUpdateProjectDefault(code int) *UpdateProjectDefault {
	return &UpdateProjectDefault{
		_statusCode: code,
	}
}

/*
UpdateProjectDefault describes a response with status code -1, with default header values.

errorResponse
*/
type UpdateProjectDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update project default response
func (o *UpdateProjectDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update project default response has a 2xx status code
func (o *UpdateProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update project default response has a 3xx status code
func (o *UpdateProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update project default response has a 4xx status code
func (o *UpdateProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update project default response has a 5xx status code
func (o *UpdateProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update project default response a status code equal to that given
func (o *UpdateProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateProjectDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}][%d] updateProject default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateProjectDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}][%d] updateProject default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateProjectDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
