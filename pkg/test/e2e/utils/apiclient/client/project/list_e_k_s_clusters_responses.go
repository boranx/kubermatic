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

// ListEKSClustersReader is a Reader for the ListEKSClusters structure.
type ListEKSClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListEKSClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSClustersOK creates a ListEKSClustersOK with default headers values
func NewListEKSClustersOK() *ListEKSClustersOK {
	return &ListEKSClustersOK{}
}

/*
ListEKSClustersOK describes a response with status code 200, with default header values.

EKSClusterList
*/
type ListEKSClustersOK struct {
	Payload models.EKSClusterList
}

// IsSuccess returns true when this list e k s clusters o k response has a 2xx status code
func (o *ListEKSClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list e k s clusters o k response has a 3xx status code
func (o *ListEKSClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s clusters o k response has a 4xx status code
func (o *ListEKSClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list e k s clusters o k response has a 5xx status code
func (o *ListEKSClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s clusters o k response a status code equal to that given
func (o *ListEKSClustersOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListEKSClustersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/clusters][%d] listEKSClustersOK  %+v", 200, o.Payload)
}

func (o *ListEKSClustersOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/clusters][%d] listEKSClustersOK  %+v", 200, o.Payload)
}

func (o *ListEKSClustersOK) GetPayload() models.EKSClusterList {
	return o.Payload
}

func (o *ListEKSClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSClustersDefault creates a ListEKSClustersDefault with default headers values
func NewListEKSClustersDefault(code int) *ListEKSClustersDefault {
	return &ListEKSClustersDefault{
		_statusCode: code,
	}
}

/*
ListEKSClustersDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSClustersDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s clusters default response
func (o *ListEKSClustersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list e k s clusters default response has a 2xx status code
func (o *ListEKSClustersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list e k s clusters default response has a 3xx status code
func (o *ListEKSClustersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list e k s clusters default response has a 4xx status code
func (o *ListEKSClustersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list e k s clusters default response has a 5xx status code
func (o *ListEKSClustersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list e k s clusters default response a status code equal to that given
func (o *ListEKSClustersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListEKSClustersDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/clusters][%d] listEKSClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSClustersDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/eks/clusters][%d] listEKSClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSClustersDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
