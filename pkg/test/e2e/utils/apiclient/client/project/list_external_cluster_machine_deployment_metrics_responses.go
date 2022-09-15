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

// ListExternalClusterMachineDeploymentMetricsReader is a Reader for the ListExternalClusterMachineDeploymentMetrics structure.
type ListExternalClusterMachineDeploymentMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListExternalClusterMachineDeploymentMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListExternalClusterMachineDeploymentMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListExternalClusterMachineDeploymentMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListExternalClusterMachineDeploymentMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListExternalClusterMachineDeploymentMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListExternalClusterMachineDeploymentMetricsOK creates a ListExternalClusterMachineDeploymentMetricsOK with default headers values
func NewListExternalClusterMachineDeploymentMetricsOK() *ListExternalClusterMachineDeploymentMetricsOK {
	return &ListExternalClusterMachineDeploymentMetricsOK{}
}

/*
ListExternalClusterMachineDeploymentMetricsOK describes a response with status code 200, with default header values.

NodeMetric
*/
type ListExternalClusterMachineDeploymentMetricsOK struct {
	Payload []*models.NodeMetric
}

// IsSuccess returns true when this list external cluster machine deployment metrics o k response has a 2xx status code
func (o *ListExternalClusterMachineDeploymentMetricsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list external cluster machine deployment metrics o k response has a 3xx status code
func (o *ListExternalClusterMachineDeploymentMetricsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list external cluster machine deployment metrics o k response has a 4xx status code
func (o *ListExternalClusterMachineDeploymentMetricsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list external cluster machine deployment metrics o k response has a 5xx status code
func (o *ListExternalClusterMachineDeploymentMetricsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list external cluster machine deployment metrics o k response a status code equal to that given
func (o *ListExternalClusterMachineDeploymentMetricsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListExternalClusterMachineDeploymentMetricsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics][%d] listExternalClusterMachineDeploymentMetricsOK  %+v", 200, o.Payload)
}

func (o *ListExternalClusterMachineDeploymentMetricsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics][%d] listExternalClusterMachineDeploymentMetricsOK  %+v", 200, o.Payload)
}

func (o *ListExternalClusterMachineDeploymentMetricsOK) GetPayload() []*models.NodeMetric {
	return o.Payload
}

func (o *ListExternalClusterMachineDeploymentMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExternalClusterMachineDeploymentMetricsUnauthorized creates a ListExternalClusterMachineDeploymentMetricsUnauthorized with default headers values
func NewListExternalClusterMachineDeploymentMetricsUnauthorized() *ListExternalClusterMachineDeploymentMetricsUnauthorized {
	return &ListExternalClusterMachineDeploymentMetricsUnauthorized{}
}

/*
ListExternalClusterMachineDeploymentMetricsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterMachineDeploymentMetricsUnauthorized struct {
}

// IsSuccess returns true when this list external cluster machine deployment metrics unauthorized response has a 2xx status code
func (o *ListExternalClusterMachineDeploymentMetricsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list external cluster machine deployment metrics unauthorized response has a 3xx status code
func (o *ListExternalClusterMachineDeploymentMetricsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list external cluster machine deployment metrics unauthorized response has a 4xx status code
func (o *ListExternalClusterMachineDeploymentMetricsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list external cluster machine deployment metrics unauthorized response has a 5xx status code
func (o *ListExternalClusterMachineDeploymentMetricsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list external cluster machine deployment metrics unauthorized response a status code equal to that given
func (o *ListExternalClusterMachineDeploymentMetricsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListExternalClusterMachineDeploymentMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics][%d] listExternalClusterMachineDeploymentMetricsUnauthorized ", 401)
}

func (o *ListExternalClusterMachineDeploymentMetricsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics][%d] listExternalClusterMachineDeploymentMetricsUnauthorized ", 401)
}

func (o *ListExternalClusterMachineDeploymentMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterMachineDeploymentMetricsForbidden creates a ListExternalClusterMachineDeploymentMetricsForbidden with default headers values
func NewListExternalClusterMachineDeploymentMetricsForbidden() *ListExternalClusterMachineDeploymentMetricsForbidden {
	return &ListExternalClusterMachineDeploymentMetricsForbidden{}
}

/*
ListExternalClusterMachineDeploymentMetricsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterMachineDeploymentMetricsForbidden struct {
}

// IsSuccess returns true when this list external cluster machine deployment metrics forbidden response has a 2xx status code
func (o *ListExternalClusterMachineDeploymentMetricsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list external cluster machine deployment metrics forbidden response has a 3xx status code
func (o *ListExternalClusterMachineDeploymentMetricsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list external cluster machine deployment metrics forbidden response has a 4xx status code
func (o *ListExternalClusterMachineDeploymentMetricsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list external cluster machine deployment metrics forbidden response has a 5xx status code
func (o *ListExternalClusterMachineDeploymentMetricsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list external cluster machine deployment metrics forbidden response a status code equal to that given
func (o *ListExternalClusterMachineDeploymentMetricsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListExternalClusterMachineDeploymentMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics][%d] listExternalClusterMachineDeploymentMetricsForbidden ", 403)
}

func (o *ListExternalClusterMachineDeploymentMetricsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics][%d] listExternalClusterMachineDeploymentMetricsForbidden ", 403)
}

func (o *ListExternalClusterMachineDeploymentMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterMachineDeploymentMetricsDefault creates a ListExternalClusterMachineDeploymentMetricsDefault with default headers values
func NewListExternalClusterMachineDeploymentMetricsDefault(code int) *ListExternalClusterMachineDeploymentMetricsDefault {
	return &ListExternalClusterMachineDeploymentMetricsDefault{
		_statusCode: code,
	}
}

/*
ListExternalClusterMachineDeploymentMetricsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListExternalClusterMachineDeploymentMetricsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list external cluster machine deployment metrics default response
func (o *ListExternalClusterMachineDeploymentMetricsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list external cluster machine deployment metrics default response has a 2xx status code
func (o *ListExternalClusterMachineDeploymentMetricsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list external cluster machine deployment metrics default response has a 3xx status code
func (o *ListExternalClusterMachineDeploymentMetricsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list external cluster machine deployment metrics default response has a 4xx status code
func (o *ListExternalClusterMachineDeploymentMetricsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list external cluster machine deployment metrics default response has a 5xx status code
func (o *ListExternalClusterMachineDeploymentMetricsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list external cluster machine deployment metrics default response a status code equal to that given
func (o *ListExternalClusterMachineDeploymentMetricsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListExternalClusterMachineDeploymentMetricsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics][%d] listExternalClusterMachineDeploymentMetrics default  %+v", o._statusCode, o.Payload)
}

func (o *ListExternalClusterMachineDeploymentMetricsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics][%d] listExternalClusterMachineDeploymentMetrics default  %+v", o._statusCode, o.Payload)
}

func (o *ListExternalClusterMachineDeploymentMetricsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListExternalClusterMachineDeploymentMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
