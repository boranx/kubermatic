// Code generated by go-swagger; DO NOT EDIT.

package addon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListAddonsReader is a Reader for the ListAddons structure.
type ListAddonsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAddonsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAddonsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAddonsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAddonsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListAddonsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAddonsOK creates a ListAddonsOK with default headers values
func NewListAddonsOK() *ListAddonsOK {
	return &ListAddonsOK{}
}

/*
ListAddonsOK describes a response with status code 200, with default header values.

Addon
*/
type ListAddonsOK struct {
	Payload []*models.Addon
}

// IsSuccess returns true when this list addons o k response has a 2xx status code
func (o *ListAddonsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list addons o k response has a 3xx status code
func (o *ListAddonsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list addons o k response has a 4xx status code
func (o *ListAddonsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list addons o k response has a 5xx status code
func (o *ListAddonsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list addons o k response a status code equal to that given
func (o *ListAddonsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAddonsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] listAddonsOK  %+v", 200, o.Payload)
}

func (o *ListAddonsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] listAddonsOK  %+v", 200, o.Payload)
}

func (o *ListAddonsOK) GetPayload() []*models.Addon {
	return o.Payload
}

func (o *ListAddonsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAddonsUnauthorized creates a ListAddonsUnauthorized with default headers values
func NewListAddonsUnauthorized() *ListAddonsUnauthorized {
	return &ListAddonsUnauthorized{}
}

/*
ListAddonsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListAddonsUnauthorized struct {
}

// IsSuccess returns true when this list addons unauthorized response has a 2xx status code
func (o *ListAddonsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list addons unauthorized response has a 3xx status code
func (o *ListAddonsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list addons unauthorized response has a 4xx status code
func (o *ListAddonsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list addons unauthorized response has a 5xx status code
func (o *ListAddonsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list addons unauthorized response a status code equal to that given
func (o *ListAddonsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListAddonsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] listAddonsUnauthorized ", 401)
}

func (o *ListAddonsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] listAddonsUnauthorized ", 401)
}

func (o *ListAddonsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAddonsForbidden creates a ListAddonsForbidden with default headers values
func NewListAddonsForbidden() *ListAddonsForbidden {
	return &ListAddonsForbidden{}
}

/*
ListAddonsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListAddonsForbidden struct {
}

// IsSuccess returns true when this list addons forbidden response has a 2xx status code
func (o *ListAddonsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list addons forbidden response has a 3xx status code
func (o *ListAddonsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list addons forbidden response has a 4xx status code
func (o *ListAddonsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list addons forbidden response has a 5xx status code
func (o *ListAddonsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list addons forbidden response a status code equal to that given
func (o *ListAddonsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListAddonsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] listAddonsForbidden ", 403)
}

func (o *ListAddonsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] listAddonsForbidden ", 403)
}

func (o *ListAddonsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAddonsDefault creates a ListAddonsDefault with default headers values
func NewListAddonsDefault(code int) *ListAddonsDefault {
	return &ListAddonsDefault{
		_statusCode: code,
	}
}

/*
ListAddonsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAddonsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list addons default response
func (o *ListAddonsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list addons default response has a 2xx status code
func (o *ListAddonsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list addons default response has a 3xx status code
func (o *ListAddonsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list addons default response has a 4xx status code
func (o *ListAddonsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list addons default response has a 5xx status code
func (o *ListAddonsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list addons default response a status code equal to that given
func (o *ListAddonsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAddonsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] listAddons default  %+v", o._statusCode, o.Payload)
}

func (o *ListAddonsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] listAddons default  %+v", o._statusCode, o.Payload)
}

func (o *ListAddonsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAddonsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
