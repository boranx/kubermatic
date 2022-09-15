// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListGCPSizesNoCredentialsReader is a Reader for the ListGCPSizesNoCredentials structure.
type ListGCPSizesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPSizesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGCPSizesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGCPSizesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPSizesNoCredentialsOK creates a ListGCPSizesNoCredentialsOK with default headers values
func NewListGCPSizesNoCredentialsOK() *ListGCPSizesNoCredentialsOK {
	return &ListGCPSizesNoCredentialsOK{}
}

/*
ListGCPSizesNoCredentialsOK describes a response with status code 200, with default header values.

GCPMachineSizeList
*/
type ListGCPSizesNoCredentialsOK struct {
	Payload models.GCPMachineSizeList
}

// IsSuccess returns true when this list g c p sizes no credentials o k response has a 2xx status code
func (o *ListGCPSizesNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list g c p sizes no credentials o k response has a 3xx status code
func (o *ListGCPSizesNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g c p sizes no credentials o k response has a 4xx status code
func (o *ListGCPSizesNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list g c p sizes no credentials o k response has a 5xx status code
func (o *ListGCPSizesNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list g c p sizes no credentials o k response a status code equal to that given
func (o *ListGCPSizesNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListGCPSizesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/sizes][%d] listGCPSizesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListGCPSizesNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/sizes][%d] listGCPSizesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListGCPSizesNoCredentialsOK) GetPayload() models.GCPMachineSizeList {
	return o.Payload
}

func (o *ListGCPSizesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPSizesNoCredentialsDefault creates a ListGCPSizesNoCredentialsDefault with default headers values
func NewListGCPSizesNoCredentialsDefault(code int) *ListGCPSizesNoCredentialsDefault {
	return &ListGCPSizesNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListGCPSizesNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGCPSizesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g c p sizes no credentials default response
func (o *ListGCPSizesNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list g c p sizes no credentials default response has a 2xx status code
func (o *ListGCPSizesNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list g c p sizes no credentials default response has a 3xx status code
func (o *ListGCPSizesNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list g c p sizes no credentials default response has a 4xx status code
func (o *ListGCPSizesNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list g c p sizes no credentials default response has a 5xx status code
func (o *ListGCPSizesNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list g c p sizes no credentials default response a status code equal to that given
func (o *ListGCPSizesNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListGCPSizesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/sizes][%d] listGCPSizesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPSizesNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/sizes][%d] listGCPSizesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPSizesNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGCPSizesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
