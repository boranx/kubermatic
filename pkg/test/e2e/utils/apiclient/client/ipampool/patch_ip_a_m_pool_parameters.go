// Code generated by go-swagger; DO NOT EDIT.

package ipampool

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

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// NewPatchIPAMPoolParams creates a new PatchIPAMPoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchIPAMPoolParams() *PatchIPAMPoolParams {
	return &PatchIPAMPoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchIPAMPoolParamsWithTimeout creates a new PatchIPAMPoolParams object
// with the ability to set a timeout on a request.
func NewPatchIPAMPoolParamsWithTimeout(timeout time.Duration) *PatchIPAMPoolParams {
	return &PatchIPAMPoolParams{
		timeout: timeout,
	}
}

// NewPatchIPAMPoolParamsWithContext creates a new PatchIPAMPoolParams object
// with the ability to set a context for a request.
func NewPatchIPAMPoolParamsWithContext(ctx context.Context) *PatchIPAMPoolParams {
	return &PatchIPAMPoolParams{
		Context: ctx,
	}
}

// NewPatchIPAMPoolParamsWithHTTPClient creates a new PatchIPAMPoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchIPAMPoolParamsWithHTTPClient(client *http.Client) *PatchIPAMPoolParams {
	return &PatchIPAMPoolParams{
		HTTPClient: client,
	}
}

/*
PatchIPAMPoolParams contains all the parameters to send to the API endpoint

	for the patch IP a m pool operation.

	Typically these are written to a http.Request.
*/
type PatchIPAMPoolParams struct {

	// Body.
	Body *models.IPAMPool

	// IpampoolName.
	IPAMPoolName string

	// SeedName.
	SeedName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch IP a m pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchIPAMPoolParams) WithDefaults() *PatchIPAMPoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch IP a m pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchIPAMPoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch IP a m pool params
func (o *PatchIPAMPoolParams) WithTimeout(timeout time.Duration) *PatchIPAMPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch IP a m pool params
func (o *PatchIPAMPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch IP a m pool params
func (o *PatchIPAMPoolParams) WithContext(ctx context.Context) *PatchIPAMPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch IP a m pool params
func (o *PatchIPAMPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch IP a m pool params
func (o *PatchIPAMPoolParams) WithHTTPClient(client *http.Client) *PatchIPAMPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch IP a m pool params
func (o *PatchIPAMPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch IP a m pool params
func (o *PatchIPAMPoolParams) WithBody(body *models.IPAMPool) *PatchIPAMPoolParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch IP a m pool params
func (o *PatchIPAMPoolParams) SetBody(body *models.IPAMPool) {
	o.Body = body
}

// WithIPAMPoolName adds the ipampoolName to the patch IP a m pool params
func (o *PatchIPAMPoolParams) WithIPAMPoolName(ipampoolName string) *PatchIPAMPoolParams {
	o.SetIPAMPoolName(ipampoolName)
	return o
}

// SetIPAMPoolName adds the ipampoolName to the patch IP a m pool params
func (o *PatchIPAMPoolParams) SetIPAMPoolName(ipampoolName string) {
	o.IPAMPoolName = ipampoolName
}

// WithSeedName adds the seedName to the patch IP a m pool params
func (o *PatchIPAMPoolParams) WithSeedName(seedName string) *PatchIPAMPoolParams {
	o.SetSeedName(seedName)
	return o
}

// SetSeedName adds the seedName to the patch IP a m pool params
func (o *PatchIPAMPoolParams) SetSeedName(seedName string) {
	o.SeedName = seedName
}

// WriteToRequest writes these params to a swagger request
func (o *PatchIPAMPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param ipampool_name
	if err := r.SetPathParam("ipampool_name", o.IPAMPoolName); err != nil {
		return err
	}

	// path param seed_name
	if err := r.SetPathParam("seed_name", o.SeedName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
