// Code generated by go-swagger; DO NOT EDIT.

package azure

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

// NewListAzureSizesNoCredentialsParams creates a new ListAzureSizesNoCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAzureSizesNoCredentialsParams() *ListAzureSizesNoCredentialsParams {
	return &ListAzureSizesNoCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAzureSizesNoCredentialsParamsWithTimeout creates a new ListAzureSizesNoCredentialsParams object
// with the ability to set a timeout on a request.
func NewListAzureSizesNoCredentialsParamsWithTimeout(timeout time.Duration) *ListAzureSizesNoCredentialsParams {
	return &ListAzureSizesNoCredentialsParams{
		timeout: timeout,
	}
}

// NewListAzureSizesNoCredentialsParamsWithContext creates a new ListAzureSizesNoCredentialsParams object
// with the ability to set a context for a request.
func NewListAzureSizesNoCredentialsParamsWithContext(ctx context.Context) *ListAzureSizesNoCredentialsParams {
	return &ListAzureSizesNoCredentialsParams{
		Context: ctx,
	}
}

// NewListAzureSizesNoCredentialsParamsWithHTTPClient creates a new ListAzureSizesNoCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAzureSizesNoCredentialsParamsWithHTTPClient(client *http.Client) *ListAzureSizesNoCredentialsParams {
	return &ListAzureSizesNoCredentialsParams{
		HTTPClient: client,
	}
}

/*
ListAzureSizesNoCredentialsParams contains all the parameters to send to the API endpoint

	for the list azure sizes no credentials operation.

	Typically these are written to a http.Request.
*/
type ListAzureSizesNoCredentialsParams struct {

	// ClusterID.
	ClusterID string

	// Dc.
	DC string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list azure sizes no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAzureSizesNoCredentialsParams) WithDefaults() *ListAzureSizesNoCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list azure sizes no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAzureSizesNoCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list azure sizes no credentials params
func (o *ListAzureSizesNoCredentialsParams) WithTimeout(timeout time.Duration) *ListAzureSizesNoCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list azure sizes no credentials params
func (o *ListAzureSizesNoCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list azure sizes no credentials params
func (o *ListAzureSizesNoCredentialsParams) WithContext(ctx context.Context) *ListAzureSizesNoCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list azure sizes no credentials params
func (o *ListAzureSizesNoCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list azure sizes no credentials params
func (o *ListAzureSizesNoCredentialsParams) WithHTTPClient(client *http.Client) *ListAzureSizesNoCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list azure sizes no credentials params
func (o *ListAzureSizesNoCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list azure sizes no credentials params
func (o *ListAzureSizesNoCredentialsParams) WithClusterID(clusterID string) *ListAzureSizesNoCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list azure sizes no credentials params
func (o *ListAzureSizesNoCredentialsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list azure sizes no credentials params
func (o *ListAzureSizesNoCredentialsParams) WithDC(dc string) *ListAzureSizesNoCredentialsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list azure sizes no credentials params
func (o *ListAzureSizesNoCredentialsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list azure sizes no credentials params
func (o *ListAzureSizesNoCredentialsParams) WithProjectID(projectID string) *ListAzureSizesNoCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list azure sizes no credentials params
func (o *ListAzureSizesNoCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAzureSizesNoCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
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
