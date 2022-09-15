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

// DeleteNodeForClusterLegacyReader is a Reader for the DeleteNodeForClusterLegacy structure.
type DeleteNodeForClusterLegacyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNodeForClusterLegacyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNodeForClusterLegacyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNodeForClusterLegacyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteNodeForClusterLegacyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteNodeForClusterLegacyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNodeForClusterLegacyOK creates a DeleteNodeForClusterLegacyOK with default headers values
func NewDeleteNodeForClusterLegacyOK() *DeleteNodeForClusterLegacyOK {
	return &DeleteNodeForClusterLegacyOK{}
}

/*
DeleteNodeForClusterLegacyOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteNodeForClusterLegacyOK struct {
}

// IsSuccess returns true when this delete node for cluster legacy o k response has a 2xx status code
func (o *DeleteNodeForClusterLegacyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete node for cluster legacy o k response has a 3xx status code
func (o *DeleteNodeForClusterLegacyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete node for cluster legacy o k response has a 4xx status code
func (o *DeleteNodeForClusterLegacyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete node for cluster legacy o k response has a 5xx status code
func (o *DeleteNodeForClusterLegacyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete node for cluster legacy o k response a status code equal to that given
func (o *DeleteNodeForClusterLegacyOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteNodeForClusterLegacyOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}][%d] deleteNodeForClusterLegacyOK ", 200)
}

func (o *DeleteNodeForClusterLegacyOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}][%d] deleteNodeForClusterLegacyOK ", 200)
}

func (o *DeleteNodeForClusterLegacyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodeForClusterLegacyUnauthorized creates a DeleteNodeForClusterLegacyUnauthorized with default headers values
func NewDeleteNodeForClusterLegacyUnauthorized() *DeleteNodeForClusterLegacyUnauthorized {
	return &DeleteNodeForClusterLegacyUnauthorized{}
}

/*
DeleteNodeForClusterLegacyUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteNodeForClusterLegacyUnauthorized struct {
}

// IsSuccess returns true when this delete node for cluster legacy unauthorized response has a 2xx status code
func (o *DeleteNodeForClusterLegacyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete node for cluster legacy unauthorized response has a 3xx status code
func (o *DeleteNodeForClusterLegacyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete node for cluster legacy unauthorized response has a 4xx status code
func (o *DeleteNodeForClusterLegacyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete node for cluster legacy unauthorized response has a 5xx status code
func (o *DeleteNodeForClusterLegacyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete node for cluster legacy unauthorized response a status code equal to that given
func (o *DeleteNodeForClusterLegacyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteNodeForClusterLegacyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}][%d] deleteNodeForClusterLegacyUnauthorized ", 401)
}

func (o *DeleteNodeForClusterLegacyUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}][%d] deleteNodeForClusterLegacyUnauthorized ", 401)
}

func (o *DeleteNodeForClusterLegacyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodeForClusterLegacyForbidden creates a DeleteNodeForClusterLegacyForbidden with default headers values
func NewDeleteNodeForClusterLegacyForbidden() *DeleteNodeForClusterLegacyForbidden {
	return &DeleteNodeForClusterLegacyForbidden{}
}

/*
DeleteNodeForClusterLegacyForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteNodeForClusterLegacyForbidden struct {
}

// IsSuccess returns true when this delete node for cluster legacy forbidden response has a 2xx status code
func (o *DeleteNodeForClusterLegacyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete node for cluster legacy forbidden response has a 3xx status code
func (o *DeleteNodeForClusterLegacyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete node for cluster legacy forbidden response has a 4xx status code
func (o *DeleteNodeForClusterLegacyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete node for cluster legacy forbidden response has a 5xx status code
func (o *DeleteNodeForClusterLegacyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete node for cluster legacy forbidden response a status code equal to that given
func (o *DeleteNodeForClusterLegacyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteNodeForClusterLegacyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}][%d] deleteNodeForClusterLegacyForbidden ", 403)
}

func (o *DeleteNodeForClusterLegacyForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}][%d] deleteNodeForClusterLegacyForbidden ", 403)
}

func (o *DeleteNodeForClusterLegacyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodeForClusterLegacyDefault creates a DeleteNodeForClusterLegacyDefault with default headers values
func NewDeleteNodeForClusterLegacyDefault(code int) *DeleteNodeForClusterLegacyDefault {
	return &DeleteNodeForClusterLegacyDefault{
		_statusCode: code,
	}
}

/*
DeleteNodeForClusterLegacyDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteNodeForClusterLegacyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete node for cluster legacy default response
func (o *DeleteNodeForClusterLegacyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete node for cluster legacy default response has a 2xx status code
func (o *DeleteNodeForClusterLegacyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete node for cluster legacy default response has a 3xx status code
func (o *DeleteNodeForClusterLegacyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete node for cluster legacy default response has a 4xx status code
func (o *DeleteNodeForClusterLegacyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete node for cluster legacy default response has a 5xx status code
func (o *DeleteNodeForClusterLegacyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete node for cluster legacy default response a status code equal to that given
func (o *DeleteNodeForClusterLegacyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteNodeForClusterLegacyDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}][%d] deleteNodeForClusterLegacy default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNodeForClusterLegacyDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}][%d] deleteNodeForClusterLegacy default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNodeForClusterLegacyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteNodeForClusterLegacyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
