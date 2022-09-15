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

// NewRevokeClusterAdminTokenParams creates a new RevokeClusterAdminTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevokeClusterAdminTokenParams() *RevokeClusterAdminTokenParams {
	return &RevokeClusterAdminTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeClusterAdminTokenParamsWithTimeout creates a new RevokeClusterAdminTokenParams object
// with the ability to set a timeout on a request.
func NewRevokeClusterAdminTokenParamsWithTimeout(timeout time.Duration) *RevokeClusterAdminTokenParams {
	return &RevokeClusterAdminTokenParams{
		timeout: timeout,
	}
}

// NewRevokeClusterAdminTokenParamsWithContext creates a new RevokeClusterAdminTokenParams object
// with the ability to set a context for a request.
func NewRevokeClusterAdminTokenParamsWithContext(ctx context.Context) *RevokeClusterAdminTokenParams {
	return &RevokeClusterAdminTokenParams{
		Context: ctx,
	}
}

// NewRevokeClusterAdminTokenParamsWithHTTPClient creates a new RevokeClusterAdminTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevokeClusterAdminTokenParamsWithHTTPClient(client *http.Client) *RevokeClusterAdminTokenParams {
	return &RevokeClusterAdminTokenParams{
		HTTPClient: client,
	}
}

/*
RevokeClusterAdminTokenParams contains all the parameters to send to the API endpoint

	for the revoke cluster admin token operation.

	Typically these are written to a http.Request.
*/
type RevokeClusterAdminTokenParams struct {

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

// WithDefaults hydrates default values in the revoke cluster admin token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeClusterAdminTokenParams) WithDefaults() *RevokeClusterAdminTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revoke cluster admin token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeClusterAdminTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revoke cluster admin token params
func (o *RevokeClusterAdminTokenParams) WithTimeout(timeout time.Duration) *RevokeClusterAdminTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke cluster admin token params
func (o *RevokeClusterAdminTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke cluster admin token params
func (o *RevokeClusterAdminTokenParams) WithContext(ctx context.Context) *RevokeClusterAdminTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke cluster admin token params
func (o *RevokeClusterAdminTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke cluster admin token params
func (o *RevokeClusterAdminTokenParams) WithHTTPClient(client *http.Client) *RevokeClusterAdminTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke cluster admin token params
func (o *RevokeClusterAdminTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the revoke cluster admin token params
func (o *RevokeClusterAdminTokenParams) WithClusterID(clusterID string) *RevokeClusterAdminTokenParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the revoke cluster admin token params
func (o *RevokeClusterAdminTokenParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the revoke cluster admin token params
func (o *RevokeClusterAdminTokenParams) WithDC(dc string) *RevokeClusterAdminTokenParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the revoke cluster admin token params
func (o *RevokeClusterAdminTokenParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the revoke cluster admin token params
func (o *RevokeClusterAdminTokenParams) WithProjectID(projectID string) *RevokeClusterAdminTokenParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the revoke cluster admin token params
func (o *RevokeClusterAdminTokenParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeClusterAdminTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
