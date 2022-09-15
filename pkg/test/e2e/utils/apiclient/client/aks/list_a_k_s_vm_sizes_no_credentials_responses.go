// Code generated by go-swagger; DO NOT EDIT.

package aks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListAKSVMSizesNoCredentialsReader is a Reader for the ListAKSVMSizesNoCredentials structure.
type ListAKSVMSizesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAKSVMSizesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAKSVMSizesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAKSVMSizesNoCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAKSVMSizesNoCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListAKSVMSizesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAKSVMSizesNoCredentialsOK creates a ListAKSVMSizesNoCredentialsOK with default headers values
func NewListAKSVMSizesNoCredentialsOK() *ListAKSVMSizesNoCredentialsOK {
	return &ListAKSVMSizesNoCredentialsOK{}
}

/*
ListAKSVMSizesNoCredentialsOK describes a response with status code 200, with default header values.

AKSVMSizeList
*/
type ListAKSVMSizesNoCredentialsOK struct {
	Payload models.AKSVMSizeList
}

// IsSuccess returns true when this list a k s Vm sizes no credentials o k response has a 2xx status code
func (o *ListAKSVMSizesNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list a k s Vm sizes no credentials o k response has a 3xx status code
func (o *ListAKSVMSizesNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list a k s Vm sizes no credentials o k response has a 4xx status code
func (o *ListAKSVMSizesNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list a k s Vm sizes no credentials o k response has a 5xx status code
func (o *ListAKSVMSizesNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list a k s Vm sizes no credentials o k response a status code equal to that given
func (o *ListAKSVMSizesNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAKSVMSizesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/vmsizes][%d] listAKSVmSizesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListAKSVMSizesNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/vmsizes][%d] listAKSVmSizesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListAKSVMSizesNoCredentialsOK) GetPayload() models.AKSVMSizeList {
	return o.Payload
}

func (o *ListAKSVMSizesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAKSVMSizesNoCredentialsUnauthorized creates a ListAKSVMSizesNoCredentialsUnauthorized with default headers values
func NewListAKSVMSizesNoCredentialsUnauthorized() *ListAKSVMSizesNoCredentialsUnauthorized {
	return &ListAKSVMSizesNoCredentialsUnauthorized{}
}

/*
ListAKSVMSizesNoCredentialsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListAKSVMSizesNoCredentialsUnauthorized struct {
}

// IsSuccess returns true when this list a k s Vm sizes no credentials unauthorized response has a 2xx status code
func (o *ListAKSVMSizesNoCredentialsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list a k s Vm sizes no credentials unauthorized response has a 3xx status code
func (o *ListAKSVMSizesNoCredentialsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list a k s Vm sizes no credentials unauthorized response has a 4xx status code
func (o *ListAKSVMSizesNoCredentialsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list a k s Vm sizes no credentials unauthorized response has a 5xx status code
func (o *ListAKSVMSizesNoCredentialsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list a k s Vm sizes no credentials unauthorized response a status code equal to that given
func (o *ListAKSVMSizesNoCredentialsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListAKSVMSizesNoCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/vmsizes][%d] listAKSVmSizesNoCredentialsUnauthorized ", 401)
}

func (o *ListAKSVMSizesNoCredentialsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/vmsizes][%d] listAKSVmSizesNoCredentialsUnauthorized ", 401)
}

func (o *ListAKSVMSizesNoCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAKSVMSizesNoCredentialsForbidden creates a ListAKSVMSizesNoCredentialsForbidden with default headers values
func NewListAKSVMSizesNoCredentialsForbidden() *ListAKSVMSizesNoCredentialsForbidden {
	return &ListAKSVMSizesNoCredentialsForbidden{}
}

/*
ListAKSVMSizesNoCredentialsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListAKSVMSizesNoCredentialsForbidden struct {
}

// IsSuccess returns true when this list a k s Vm sizes no credentials forbidden response has a 2xx status code
func (o *ListAKSVMSizesNoCredentialsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list a k s Vm sizes no credentials forbidden response has a 3xx status code
func (o *ListAKSVMSizesNoCredentialsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list a k s Vm sizes no credentials forbidden response has a 4xx status code
func (o *ListAKSVMSizesNoCredentialsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list a k s Vm sizes no credentials forbidden response has a 5xx status code
func (o *ListAKSVMSizesNoCredentialsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list a k s Vm sizes no credentials forbidden response a status code equal to that given
func (o *ListAKSVMSizesNoCredentialsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListAKSVMSizesNoCredentialsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/vmsizes][%d] listAKSVmSizesNoCredentialsForbidden ", 403)
}

func (o *ListAKSVMSizesNoCredentialsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/vmsizes][%d] listAKSVmSizesNoCredentialsForbidden ", 403)
}

func (o *ListAKSVMSizesNoCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAKSVMSizesNoCredentialsDefault creates a ListAKSVMSizesNoCredentialsDefault with default headers values
func NewListAKSVMSizesNoCredentialsDefault(code int) *ListAKSVMSizesNoCredentialsDefault {
	return &ListAKSVMSizesNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListAKSVMSizesNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAKSVMSizesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list a k s VM sizes no credentials default response
func (o *ListAKSVMSizesNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list a k s VM sizes no credentials default response has a 2xx status code
func (o *ListAKSVMSizesNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list a k s VM sizes no credentials default response has a 3xx status code
func (o *ListAKSVMSizesNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list a k s VM sizes no credentials default response has a 4xx status code
func (o *ListAKSVMSizesNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list a k s VM sizes no credentials default response has a 5xx status code
func (o *ListAKSVMSizesNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list a k s VM sizes no credentials default response a status code equal to that given
func (o *ListAKSVMSizesNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAKSVMSizesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/vmsizes][%d] listAKSVMSizesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListAKSVMSizesNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/aks/vmsizes][%d] listAKSVMSizesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListAKSVMSizesNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAKSVMSizesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
