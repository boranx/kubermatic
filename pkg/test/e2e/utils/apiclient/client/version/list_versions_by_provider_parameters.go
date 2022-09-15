// Code generated by go-swagger; DO NOT EDIT.

package version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListVersionsByProviderParams creates a new ListVersionsByProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListVersionsByProviderParams() *ListVersionsByProviderParams {
	return &ListVersionsByProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListVersionsByProviderParamsWithTimeout creates a new ListVersionsByProviderParams object
// with the ability to set a timeout on a request.
func NewListVersionsByProviderParamsWithTimeout(timeout time.Duration) *ListVersionsByProviderParams {
	return &ListVersionsByProviderParams{
		timeout: timeout,
	}
}

// NewListVersionsByProviderParamsWithContext creates a new ListVersionsByProviderParams object
// with the ability to set a context for a request.
func NewListVersionsByProviderParamsWithContext(ctx context.Context) *ListVersionsByProviderParams {
	return &ListVersionsByProviderParams{
		Context: ctx,
	}
}

// NewListVersionsByProviderParamsWithHTTPClient creates a new ListVersionsByProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewListVersionsByProviderParamsWithHTTPClient(client *http.Client) *ListVersionsByProviderParams {
	return &ListVersionsByProviderParams{
		HTTPClient: client,
	}
}

/*
ListVersionsByProviderParams contains all the parameters to send to the API endpoint

	for the list versions by provider operation.

	Typically these are written to a http.Request.
*/
type ListVersionsByProviderParams struct {

	// ProviderName.
	ProviderName string

	/* Type.

	   Type is deprecated and not used anymore.
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list versions by provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVersionsByProviderParams) WithDefaults() *ListVersionsByProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list versions by provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVersionsByProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list versions by provider params
func (o *ListVersionsByProviderParams) WithTimeout(timeout time.Duration) *ListVersionsByProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list versions by provider params
func (o *ListVersionsByProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list versions by provider params
func (o *ListVersionsByProviderParams) WithContext(ctx context.Context) *ListVersionsByProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list versions by provider params
func (o *ListVersionsByProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list versions by provider params
func (o *ListVersionsByProviderParams) WithHTTPClient(client *http.Client) *ListVersionsByProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list versions by provider params
func (o *ListVersionsByProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProviderName adds the providerName to the list versions by provider params
func (o *ListVersionsByProviderParams) WithProviderName(providerName string) *ListVersionsByProviderParams {
	o.SetProviderName(providerName)
	return o
}

// SetProviderName adds the providerName to the list versions by provider params
func (o *ListVersionsByProviderParams) SetProviderName(providerName string) {
	o.ProviderName = providerName
}

// WithType adds the typeVar to the list versions by provider params
func (o *ListVersionsByProviderParams) WithType(typeVar *string) *ListVersionsByProviderParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the list versions by provider params
func (o *ListVersionsByProviderParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ListVersionsByProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param provider_name
	if err := r.SetPathParam("provider_name", o.ProviderName); err != nil {
		return err
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
