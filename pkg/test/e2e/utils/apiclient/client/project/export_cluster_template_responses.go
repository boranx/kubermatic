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

// ExportClusterTemplateReader is a Reader for the ExportClusterTemplate structure.
type ExportClusterTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportClusterTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportClusterTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExportClusterTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExportClusterTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExportClusterTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportClusterTemplateOK creates a ExportClusterTemplateOK with default headers values
func NewExportClusterTemplateOK() *ExportClusterTemplateOK {
	return &ExportClusterTemplateOK{}
}

/*
ExportClusterTemplateOK describes a response with status code 200, with default header values.

ClusterTemplate
*/
type ExportClusterTemplateOK struct {
	Payload *models.ClusterTemplate
}

// IsSuccess returns true when this export cluster template o k response has a 2xx status code
func (o *ExportClusterTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export cluster template o k response has a 3xx status code
func (o *ExportClusterTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export cluster template o k response has a 4xx status code
func (o *ExportClusterTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export cluster template o k response has a 5xx status code
func (o *ExportClusterTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export cluster template o k response a status code equal to that given
func (o *ExportClusterTemplateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExportClusterTemplateOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}/export][%d] exportClusterTemplateOK  %+v", 200, o.Payload)
}

func (o *ExportClusterTemplateOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}/export][%d] exportClusterTemplateOK  %+v", 200, o.Payload)
}

func (o *ExportClusterTemplateOK) GetPayload() *models.ClusterTemplate {
	return o.Payload
}

func (o *ExportClusterTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportClusterTemplateUnauthorized creates a ExportClusterTemplateUnauthorized with default headers values
func NewExportClusterTemplateUnauthorized() *ExportClusterTemplateUnauthorized {
	return &ExportClusterTemplateUnauthorized{}
}

/*
ExportClusterTemplateUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ExportClusterTemplateUnauthorized struct {
}

// IsSuccess returns true when this export cluster template unauthorized response has a 2xx status code
func (o *ExportClusterTemplateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export cluster template unauthorized response has a 3xx status code
func (o *ExportClusterTemplateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export cluster template unauthorized response has a 4xx status code
func (o *ExportClusterTemplateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this export cluster template unauthorized response has a 5xx status code
func (o *ExportClusterTemplateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this export cluster template unauthorized response a status code equal to that given
func (o *ExportClusterTemplateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ExportClusterTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}/export][%d] exportClusterTemplateUnauthorized ", 401)
}

func (o *ExportClusterTemplateUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}/export][%d] exportClusterTemplateUnauthorized ", 401)
}

func (o *ExportClusterTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportClusterTemplateForbidden creates a ExportClusterTemplateForbidden with default headers values
func NewExportClusterTemplateForbidden() *ExportClusterTemplateForbidden {
	return &ExportClusterTemplateForbidden{}
}

/*
ExportClusterTemplateForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ExportClusterTemplateForbidden struct {
}

// IsSuccess returns true when this export cluster template forbidden response has a 2xx status code
func (o *ExportClusterTemplateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export cluster template forbidden response has a 3xx status code
func (o *ExportClusterTemplateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export cluster template forbidden response has a 4xx status code
func (o *ExportClusterTemplateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this export cluster template forbidden response has a 5xx status code
func (o *ExportClusterTemplateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this export cluster template forbidden response a status code equal to that given
func (o *ExportClusterTemplateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ExportClusterTemplateForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}/export][%d] exportClusterTemplateForbidden ", 403)
}

func (o *ExportClusterTemplateForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}/export][%d] exportClusterTemplateForbidden ", 403)
}

func (o *ExportClusterTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportClusterTemplateDefault creates a ExportClusterTemplateDefault with default headers values
func NewExportClusterTemplateDefault(code int) *ExportClusterTemplateDefault {
	return &ExportClusterTemplateDefault{
		_statusCode: code,
	}
}

/*
ExportClusterTemplateDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ExportClusterTemplateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the export cluster template default response
func (o *ExportClusterTemplateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this export cluster template default response has a 2xx status code
func (o *ExportClusterTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this export cluster template default response has a 3xx status code
func (o *ExportClusterTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this export cluster template default response has a 4xx status code
func (o *ExportClusterTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this export cluster template default response has a 5xx status code
func (o *ExportClusterTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this export cluster template default response a status code equal to that given
func (o *ExportClusterTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExportClusterTemplateDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}/export][%d] exportClusterTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *ExportClusterTemplateDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clustertemplates/{template_id}/export][%d] exportClusterTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *ExportClusterTemplateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportClusterTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
