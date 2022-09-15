// Code generated by go-swagger; DO NOT EDIT.

package rulegroup

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

// NewListAdminRuleGroupsParams creates a new ListAdminRuleGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAdminRuleGroupsParams() *ListAdminRuleGroupsParams {
	return &ListAdminRuleGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAdminRuleGroupsParamsWithTimeout creates a new ListAdminRuleGroupsParams object
// with the ability to set a timeout on a request.
func NewListAdminRuleGroupsParamsWithTimeout(timeout time.Duration) *ListAdminRuleGroupsParams {
	return &ListAdminRuleGroupsParams{
		timeout: timeout,
	}
}

// NewListAdminRuleGroupsParamsWithContext creates a new ListAdminRuleGroupsParams object
// with the ability to set a context for a request.
func NewListAdminRuleGroupsParamsWithContext(ctx context.Context) *ListAdminRuleGroupsParams {
	return &ListAdminRuleGroupsParams{
		Context: ctx,
	}
}

// NewListAdminRuleGroupsParamsWithHTTPClient creates a new ListAdminRuleGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAdminRuleGroupsParamsWithHTTPClient(client *http.Client) *ListAdminRuleGroupsParams {
	return &ListAdminRuleGroupsParams{
		HTTPClient: client,
	}
}

/*
ListAdminRuleGroupsParams contains all the parameters to send to the API endpoint

	for the list admin rule groups operation.

	Typically these are written to a http.Request.
*/
type ListAdminRuleGroupsParams struct {

	// SeedName.
	SeedName string

	// Type.
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list admin rule groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAdminRuleGroupsParams) WithDefaults() *ListAdminRuleGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list admin rule groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAdminRuleGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list admin rule groups params
func (o *ListAdminRuleGroupsParams) WithTimeout(timeout time.Duration) *ListAdminRuleGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list admin rule groups params
func (o *ListAdminRuleGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list admin rule groups params
func (o *ListAdminRuleGroupsParams) WithContext(ctx context.Context) *ListAdminRuleGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list admin rule groups params
func (o *ListAdminRuleGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list admin rule groups params
func (o *ListAdminRuleGroupsParams) WithHTTPClient(client *http.Client) *ListAdminRuleGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list admin rule groups params
func (o *ListAdminRuleGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSeedName adds the seedName to the list admin rule groups params
func (o *ListAdminRuleGroupsParams) WithSeedName(seedName string) *ListAdminRuleGroupsParams {
	o.SetSeedName(seedName)
	return o
}

// SetSeedName adds the seedName to the list admin rule groups params
func (o *ListAdminRuleGroupsParams) SetSeedName(seedName string) {
	o.SeedName = seedName
}

// WithType adds the typeVar to the list admin rule groups params
func (o *ListAdminRuleGroupsParams) WithType(typeVar *string) *ListAdminRuleGroupsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the list admin rule groups params
func (o *ListAdminRuleGroupsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ListAdminRuleGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param seed_name
	if err := r.SetPathParam("seed_name", o.SeedName); err != nil {
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
