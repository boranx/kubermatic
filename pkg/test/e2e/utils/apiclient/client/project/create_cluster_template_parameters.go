// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewCreateClusterTemplateParams creates a new CreateClusterTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateClusterTemplateParams() *CreateClusterTemplateParams {
	return &CreateClusterTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterTemplateParamsWithTimeout creates a new CreateClusterTemplateParams object
// with the ability to set a timeout on a request.
func NewCreateClusterTemplateParamsWithTimeout(timeout time.Duration) *CreateClusterTemplateParams {
	return &CreateClusterTemplateParams{
		timeout: timeout,
	}
}

// NewCreateClusterTemplateParamsWithContext creates a new CreateClusterTemplateParams object
// with the ability to set a context for a request.
func NewCreateClusterTemplateParamsWithContext(ctx context.Context) *CreateClusterTemplateParams {
	return &CreateClusterTemplateParams{
		Context: ctx,
	}
}

// NewCreateClusterTemplateParamsWithHTTPClient creates a new CreateClusterTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateClusterTemplateParamsWithHTTPClient(client *http.Client) *CreateClusterTemplateParams {
	return &CreateClusterTemplateParams{
		HTTPClient: client,
	}
}

/*
CreateClusterTemplateParams contains all the parameters to send to the API endpoint

	for the create cluster template operation.

	Typically these are written to a http.Request.
*/
type CreateClusterTemplateParams struct {

	// Body.
	Body CreateClusterTemplateBody

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cluster template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterTemplateParams) WithDefaults() *CreateClusterTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cluster template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cluster template params
func (o *CreateClusterTemplateParams) WithTimeout(timeout time.Duration) *CreateClusterTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster template params
func (o *CreateClusterTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster template params
func (o *CreateClusterTemplateParams) WithContext(ctx context.Context) *CreateClusterTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster template params
func (o *CreateClusterTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster template params
func (o *CreateClusterTemplateParams) WithHTTPClient(client *http.Client) *CreateClusterTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster template params
func (o *CreateClusterTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create cluster template params
func (o *CreateClusterTemplateParams) WithBody(body CreateClusterTemplateBody) *CreateClusterTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cluster template params
func (o *CreateClusterTemplateParams) SetBody(body CreateClusterTemplateBody) {
	o.Body = body
}

// WithProjectID adds the projectID to the create cluster template params
func (o *CreateClusterTemplateParams) WithProjectID(projectID string) *CreateClusterTemplateParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create cluster template params
func (o *CreateClusterTemplateParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
