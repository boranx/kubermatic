// Code generated by go-swagger; DO NOT EDIT.

package eks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListEKSVPCsNoCredentialsReader is a Reader for the ListEKSVPCsNoCredentials structure.
type ListEKSVPCsNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSVPCsNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSVPCsNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEKSVPCsNoCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListEKSVPCsNoCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListEKSVPCsNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSVPCsNoCredentialsOK creates a ListEKSVPCsNoCredentialsOK with default headers values
func NewListEKSVPCsNoCredentialsOK() *ListEKSVPCsNoCredentialsOK {
	return &ListEKSVPCsNoCredentialsOK{}
}

/*
ListEKSVPCsNoCredentialsOK describes a response with status code 200, with default header values.

EKSVPCList
*/
type ListEKSVPCsNoCredentialsOK struct {
	Payload models.EKSVPCList
}

// IsSuccess returns true when this list e k s v p cs no credentials o k response has a 2xx status code
func (o *ListEKSVPCsNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list e k s v p cs no credentials o k response has a 3xx status code
func (o *ListEKSVPCsNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s v p cs no credentials o k response has a 4xx status code
func (o *ListEKSVPCsNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list e k s v p cs no credentials o k response has a 5xx status code
func (o *ListEKSVPCsNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s v p cs no credentials o k response a status code equal to that given
func (o *ListEKSVPCsNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListEKSVPCsNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/vpcs][%d] listEKSVPCsNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListEKSVPCsNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/vpcs][%d] listEKSVPCsNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListEKSVPCsNoCredentialsOK) GetPayload() models.EKSVPCList {
	return o.Payload
}

func (o *ListEKSVPCsNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSVPCsNoCredentialsUnauthorized creates a ListEKSVPCsNoCredentialsUnauthorized with default headers values
func NewListEKSVPCsNoCredentialsUnauthorized() *ListEKSVPCsNoCredentialsUnauthorized {
	return &ListEKSVPCsNoCredentialsUnauthorized{}
}

/*
ListEKSVPCsNoCredentialsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListEKSVPCsNoCredentialsUnauthorized struct {
}

// IsSuccess returns true when this list e k s v p cs no credentials unauthorized response has a 2xx status code
func (o *ListEKSVPCsNoCredentialsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list e k s v p cs no credentials unauthorized response has a 3xx status code
func (o *ListEKSVPCsNoCredentialsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s v p cs no credentials unauthorized response has a 4xx status code
func (o *ListEKSVPCsNoCredentialsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list e k s v p cs no credentials unauthorized response has a 5xx status code
func (o *ListEKSVPCsNoCredentialsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s v p cs no credentials unauthorized response a status code equal to that given
func (o *ListEKSVPCsNoCredentialsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListEKSVPCsNoCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/vpcs][%d] listEKSVPCsNoCredentialsUnauthorized ", 401)
}

func (o *ListEKSVPCsNoCredentialsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/vpcs][%d] listEKSVPCsNoCredentialsUnauthorized ", 401)
}

func (o *ListEKSVPCsNoCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSVPCsNoCredentialsForbidden creates a ListEKSVPCsNoCredentialsForbidden with default headers values
func NewListEKSVPCsNoCredentialsForbidden() *ListEKSVPCsNoCredentialsForbidden {
	return &ListEKSVPCsNoCredentialsForbidden{}
}

/*
ListEKSVPCsNoCredentialsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListEKSVPCsNoCredentialsForbidden struct {
}

// IsSuccess returns true when this list e k s v p cs no credentials forbidden response has a 2xx status code
func (o *ListEKSVPCsNoCredentialsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list e k s v p cs no credentials forbidden response has a 3xx status code
func (o *ListEKSVPCsNoCredentialsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s v p cs no credentials forbidden response has a 4xx status code
func (o *ListEKSVPCsNoCredentialsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list e k s v p cs no credentials forbidden response has a 5xx status code
func (o *ListEKSVPCsNoCredentialsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s v p cs no credentials forbidden response a status code equal to that given
func (o *ListEKSVPCsNoCredentialsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListEKSVPCsNoCredentialsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/vpcs][%d] listEKSVPCsNoCredentialsForbidden ", 403)
}

func (o *ListEKSVPCsNoCredentialsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/vpcs][%d] listEKSVPCsNoCredentialsForbidden ", 403)
}

func (o *ListEKSVPCsNoCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSVPCsNoCredentialsDefault creates a ListEKSVPCsNoCredentialsDefault with default headers values
func NewListEKSVPCsNoCredentialsDefault(code int) *ListEKSVPCsNoCredentialsDefault {
	return &ListEKSVPCsNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListEKSVPCsNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSVPCsNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s v p cs no credentials default response
func (o *ListEKSVPCsNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list e k s v p cs no credentials default response has a 2xx status code
func (o *ListEKSVPCsNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list e k s v p cs no credentials default response has a 3xx status code
func (o *ListEKSVPCsNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list e k s v p cs no credentials default response has a 4xx status code
func (o *ListEKSVPCsNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list e k s v p cs no credentials default response has a 5xx status code
func (o *ListEKSVPCsNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list e k s v p cs no credentials default response a status code equal to that given
func (o *ListEKSVPCsNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListEKSVPCsNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/vpcs][%d] listEKSVPCsNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSVPCsNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/vpcs][%d] listEKSVPCsNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSVPCsNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSVPCsNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
