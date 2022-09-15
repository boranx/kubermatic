// Code generated by go-swagger; DO NOT EDIT.

package eks

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

// NewListEKSClusterRolesParams creates a new ListEKSClusterRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListEKSClusterRolesParams() *ListEKSClusterRolesParams {
	return &ListEKSClusterRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListEKSClusterRolesParamsWithTimeout creates a new ListEKSClusterRolesParams object
// with the ability to set a timeout on a request.
func NewListEKSClusterRolesParamsWithTimeout(timeout time.Duration) *ListEKSClusterRolesParams {
	return &ListEKSClusterRolesParams{
		timeout: timeout,
	}
}

// NewListEKSClusterRolesParamsWithContext creates a new ListEKSClusterRolesParams object
// with the ability to set a context for a request.
func NewListEKSClusterRolesParamsWithContext(ctx context.Context) *ListEKSClusterRolesParams {
	return &ListEKSClusterRolesParams{
		Context: ctx,
	}
}

// NewListEKSClusterRolesParamsWithHTTPClient creates a new ListEKSClusterRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListEKSClusterRolesParamsWithHTTPClient(client *http.Client) *ListEKSClusterRolesParams {
	return &ListEKSClusterRolesParams{
		HTTPClient: client,
	}
}

/*
ListEKSClusterRolesParams contains all the parameters to send to the API endpoint

	for the list e k s cluster roles operation.

	Typically these are written to a http.Request.
*/
type ListEKSClusterRolesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list e k s cluster roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEKSClusterRolesParams) WithDefaults() *ListEKSClusterRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list e k s cluster roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEKSClusterRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list e k s cluster roles params
func (o *ListEKSClusterRolesParams) WithTimeout(timeout time.Duration) *ListEKSClusterRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list e k s cluster roles params
func (o *ListEKSClusterRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list e k s cluster roles params
func (o *ListEKSClusterRolesParams) WithContext(ctx context.Context) *ListEKSClusterRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list e k s cluster roles params
func (o *ListEKSClusterRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list e k s cluster roles params
func (o *ListEKSClusterRolesParams) WithHTTPClient(client *http.Client) *ListEKSClusterRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list e k s cluster roles params
func (o *ListEKSClusterRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListEKSClusterRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
