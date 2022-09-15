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

// NewValidateEKSCredentialsParams creates a new ValidateEKSCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateEKSCredentialsParams() *ValidateEKSCredentialsParams {
	return &ValidateEKSCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateEKSCredentialsParamsWithTimeout creates a new ValidateEKSCredentialsParams object
// with the ability to set a timeout on a request.
func NewValidateEKSCredentialsParamsWithTimeout(timeout time.Duration) *ValidateEKSCredentialsParams {
	return &ValidateEKSCredentialsParams{
		timeout: timeout,
	}
}

// NewValidateEKSCredentialsParamsWithContext creates a new ValidateEKSCredentialsParams object
// with the ability to set a context for a request.
func NewValidateEKSCredentialsParamsWithContext(ctx context.Context) *ValidateEKSCredentialsParams {
	return &ValidateEKSCredentialsParams{
		Context: ctx,
	}
}

// NewValidateEKSCredentialsParamsWithHTTPClient creates a new ValidateEKSCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateEKSCredentialsParamsWithHTTPClient(client *http.Client) *ValidateEKSCredentialsParams {
	return &ValidateEKSCredentialsParams{
		HTTPClient: client,
	}
}

/*
ValidateEKSCredentialsParams contains all the parameters to send to the API endpoint

	for the validate e k s credentials operation.

	Typically these are written to a http.Request.
*/
type ValidateEKSCredentialsParams struct {

	// AccessKeyID.
	AccessKeyID *string

	// Credential.
	Credential *string

	// Region.
	Region *string

	// SecretAccessKey.
	SecretAccessKey *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate e k s credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateEKSCredentialsParams) WithDefaults() *ValidateEKSCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate e k s credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateEKSCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) WithTimeout(timeout time.Duration) *ValidateEKSCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) WithContext(ctx context.Context) *ValidateEKSCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) WithHTTPClient(client *http.Client) *ValidateEKSCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKeyID adds the accessKeyID to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) WithAccessKeyID(accessKeyID *string) *ValidateEKSCredentialsParams {
	o.SetAccessKeyID(accessKeyID)
	return o
}

// SetAccessKeyID adds the accessKeyId to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) SetAccessKeyID(accessKeyID *string) {
	o.AccessKeyID = accessKeyID
}

// WithCredential adds the credential to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) WithCredential(credential *string) *ValidateEKSCredentialsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithRegion adds the region to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) WithRegion(region *string) *ValidateEKSCredentialsParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) SetRegion(region *string) {
	o.Region = region
}

// WithSecretAccessKey adds the secretAccessKey to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) WithSecretAccessKey(secretAccessKey *string) *ValidateEKSCredentialsParams {
	o.SetSecretAccessKey(secretAccessKey)
	return o
}

// SetSecretAccessKey adds the secretAccessKey to the validate e k s credentials params
func (o *ValidateEKSCredentialsParams) SetSecretAccessKey(secretAccessKey *string) {
	o.SecretAccessKey = secretAccessKey
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateEKSCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessKeyID != nil {

		// header param AccessKeyID
		if err := r.SetHeaderParam("AccessKeyID", *o.AccessKeyID); err != nil {
			return err
		}
	}

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.Region != nil {

		// header param Region
		if err := r.SetHeaderParam("Region", *o.Region); err != nil {
			return err
		}
	}

	if o.SecretAccessKey != nil {

		// header param SecretAccessKey
		if err := r.SetHeaderParam("SecretAccessKey", *o.SecretAccessKey); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
