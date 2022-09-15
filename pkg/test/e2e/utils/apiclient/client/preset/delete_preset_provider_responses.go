// Code generated by go-swagger; DO NOT EDIT.

package preset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// DeletePresetProviderReader is a Reader for the DeletePresetProvider structure.
type DeletePresetProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePresetProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePresetProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeletePresetProviderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePresetProviderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePresetProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePresetProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePresetProviderOK creates a DeletePresetProviderOK with default headers values
func NewDeletePresetProviderOK() *DeletePresetProviderOK {
	return &DeletePresetProviderOK{}
}

/*
DeletePresetProviderOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeletePresetProviderOK struct {
}

// IsSuccess returns true when this delete preset provider o k response has a 2xx status code
func (o *DeletePresetProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete preset provider o k response has a 3xx status code
func (o *DeletePresetProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete preset provider o k response has a 4xx status code
func (o *DeletePresetProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete preset provider o k response has a 5xx status code
func (o *DeletePresetProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete preset provider o k response a status code equal to that given
func (o *DeletePresetProviderOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeletePresetProviderOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}/provider/{provider_name}][%d] deletePresetProviderOK ", 200)
}

func (o *DeletePresetProviderOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}/provider/{provider_name}][%d] deletePresetProviderOK ", 200)
}

func (o *DeletePresetProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePresetProviderUnauthorized creates a DeletePresetProviderUnauthorized with default headers values
func NewDeletePresetProviderUnauthorized() *DeletePresetProviderUnauthorized {
	return &DeletePresetProviderUnauthorized{}
}

/*
DeletePresetProviderUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeletePresetProviderUnauthorized struct {
}

// IsSuccess returns true when this delete preset provider unauthorized response has a 2xx status code
func (o *DeletePresetProviderUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete preset provider unauthorized response has a 3xx status code
func (o *DeletePresetProviderUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete preset provider unauthorized response has a 4xx status code
func (o *DeletePresetProviderUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete preset provider unauthorized response has a 5xx status code
func (o *DeletePresetProviderUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete preset provider unauthorized response a status code equal to that given
func (o *DeletePresetProviderUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeletePresetProviderUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}/provider/{provider_name}][%d] deletePresetProviderUnauthorized ", 401)
}

func (o *DeletePresetProviderUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}/provider/{provider_name}][%d] deletePresetProviderUnauthorized ", 401)
}

func (o *DeletePresetProviderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePresetProviderForbidden creates a DeletePresetProviderForbidden with default headers values
func NewDeletePresetProviderForbidden() *DeletePresetProviderForbidden {
	return &DeletePresetProviderForbidden{}
}

/*
DeletePresetProviderForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeletePresetProviderForbidden struct {
}

// IsSuccess returns true when this delete preset provider forbidden response has a 2xx status code
func (o *DeletePresetProviderForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete preset provider forbidden response has a 3xx status code
func (o *DeletePresetProviderForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete preset provider forbidden response has a 4xx status code
func (o *DeletePresetProviderForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete preset provider forbidden response has a 5xx status code
func (o *DeletePresetProviderForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete preset provider forbidden response a status code equal to that given
func (o *DeletePresetProviderForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeletePresetProviderForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}/provider/{provider_name}][%d] deletePresetProviderForbidden ", 403)
}

func (o *DeletePresetProviderForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}/provider/{provider_name}][%d] deletePresetProviderForbidden ", 403)
}

func (o *DeletePresetProviderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePresetProviderNotFound creates a DeletePresetProviderNotFound with default headers values
func NewDeletePresetProviderNotFound() *DeletePresetProviderNotFound {
	return &DeletePresetProviderNotFound{}
}

/*
DeletePresetProviderNotFound describes a response with status code 404, with default header values.

EmptyResponse is a empty response
*/
type DeletePresetProviderNotFound struct {
}

// IsSuccess returns true when this delete preset provider not found response has a 2xx status code
func (o *DeletePresetProviderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete preset provider not found response has a 3xx status code
func (o *DeletePresetProviderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete preset provider not found response has a 4xx status code
func (o *DeletePresetProviderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete preset provider not found response has a 5xx status code
func (o *DeletePresetProviderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete preset provider not found response a status code equal to that given
func (o *DeletePresetProviderNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeletePresetProviderNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}/provider/{provider_name}][%d] deletePresetProviderNotFound ", 404)
}

func (o *DeletePresetProviderNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}/provider/{provider_name}][%d] deletePresetProviderNotFound ", 404)
}

func (o *DeletePresetProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePresetProviderDefault creates a DeletePresetProviderDefault with default headers values
func NewDeletePresetProviderDefault(code int) *DeletePresetProviderDefault {
	return &DeletePresetProviderDefault{
		_statusCode: code,
	}
}

/*
DeletePresetProviderDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeletePresetProviderDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete preset provider default response
func (o *DeletePresetProviderDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete preset provider default response has a 2xx status code
func (o *DeletePresetProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete preset provider default response has a 3xx status code
func (o *DeletePresetProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete preset provider default response has a 4xx status code
func (o *DeletePresetProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete preset provider default response has a 5xx status code
func (o *DeletePresetProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete preset provider default response a status code equal to that given
func (o *DeletePresetProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeletePresetProviderDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}/provider/{provider_name}][%d] deletePresetProvider default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePresetProviderDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presets/{preset_name}/provider/{provider_name}][%d] deletePresetProvider default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePresetProviderDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeletePresetProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
