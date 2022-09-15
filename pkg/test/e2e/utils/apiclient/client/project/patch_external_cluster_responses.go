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

// PatchExternalClusterReader is a Reader for the PatchExternalCluster structure.
type PatchExternalClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchExternalClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchExternalClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchExternalClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchExternalClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchExternalClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchExternalClusterOK creates a PatchExternalClusterOK with default headers values
func NewPatchExternalClusterOK() *PatchExternalClusterOK {
	return &PatchExternalClusterOK{}
}

/*
PatchExternalClusterOK describes a response with status code 200, with default header values.

ExternalCluster
*/
type PatchExternalClusterOK struct {
	Payload *models.ExternalCluster
}

// IsSuccess returns true when this patch external cluster o k response has a 2xx status code
func (o *PatchExternalClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch external cluster o k response has a 3xx status code
func (o *PatchExternalClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch external cluster o k response has a 4xx status code
func (o *PatchExternalClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch external cluster o k response has a 5xx status code
func (o *PatchExternalClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch external cluster o k response a status code equal to that given
func (o *PatchExternalClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchExternalClusterOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] patchExternalClusterOK  %+v", 200, o.Payload)
}

func (o *PatchExternalClusterOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] patchExternalClusterOK  %+v", 200, o.Payload)
}

func (o *PatchExternalClusterOK) GetPayload() *models.ExternalCluster {
	return o.Payload
}

func (o *PatchExternalClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchExternalClusterUnauthorized creates a PatchExternalClusterUnauthorized with default headers values
func NewPatchExternalClusterUnauthorized() *PatchExternalClusterUnauthorized {
	return &PatchExternalClusterUnauthorized{}
}

/*
PatchExternalClusterUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchExternalClusterUnauthorized struct {
}

// IsSuccess returns true when this patch external cluster unauthorized response has a 2xx status code
func (o *PatchExternalClusterUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch external cluster unauthorized response has a 3xx status code
func (o *PatchExternalClusterUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch external cluster unauthorized response has a 4xx status code
func (o *PatchExternalClusterUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch external cluster unauthorized response has a 5xx status code
func (o *PatchExternalClusterUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch external cluster unauthorized response a status code equal to that given
func (o *PatchExternalClusterUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchExternalClusterUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] patchExternalClusterUnauthorized ", 401)
}

func (o *PatchExternalClusterUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] patchExternalClusterUnauthorized ", 401)
}

func (o *PatchExternalClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchExternalClusterForbidden creates a PatchExternalClusterForbidden with default headers values
func NewPatchExternalClusterForbidden() *PatchExternalClusterForbidden {
	return &PatchExternalClusterForbidden{}
}

/*
PatchExternalClusterForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchExternalClusterForbidden struct {
}

// IsSuccess returns true when this patch external cluster forbidden response has a 2xx status code
func (o *PatchExternalClusterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch external cluster forbidden response has a 3xx status code
func (o *PatchExternalClusterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch external cluster forbidden response has a 4xx status code
func (o *PatchExternalClusterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch external cluster forbidden response has a 5xx status code
func (o *PatchExternalClusterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch external cluster forbidden response a status code equal to that given
func (o *PatchExternalClusterForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchExternalClusterForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] patchExternalClusterForbidden ", 403)
}

func (o *PatchExternalClusterForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] patchExternalClusterForbidden ", 403)
}

func (o *PatchExternalClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchExternalClusterDefault creates a PatchExternalClusterDefault with default headers values
func NewPatchExternalClusterDefault(code int) *PatchExternalClusterDefault {
	return &PatchExternalClusterDefault{
		_statusCode: code,
	}
}

/*
PatchExternalClusterDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchExternalClusterDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch external cluster default response
func (o *PatchExternalClusterDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this patch external cluster default response has a 2xx status code
func (o *PatchExternalClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch external cluster default response has a 3xx status code
func (o *PatchExternalClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch external cluster default response has a 4xx status code
func (o *PatchExternalClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch external cluster default response has a 5xx status code
func (o *PatchExternalClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch external cluster default response a status code equal to that given
func (o *PatchExternalClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PatchExternalClusterDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] patchExternalCluster default  %+v", o._statusCode, o.Payload)
}

func (o *PatchExternalClusterDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] patchExternalCluster default  %+v", o._statusCode, o.Payload)
}

func (o *PatchExternalClusterDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchExternalClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
