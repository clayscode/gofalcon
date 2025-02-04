// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewUpdateRuleGroupParams creates a new UpdateRuleGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRuleGroupParams() *UpdateRuleGroupParams {
	return &UpdateRuleGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRuleGroupParamsWithTimeout creates a new UpdateRuleGroupParams object
// with the ability to set a timeout on a request.
func NewUpdateRuleGroupParamsWithTimeout(timeout time.Duration) *UpdateRuleGroupParams {
	return &UpdateRuleGroupParams{
		timeout: timeout,
	}
}

// NewUpdateRuleGroupParamsWithContext creates a new UpdateRuleGroupParams object
// with the ability to set a context for a request.
func NewUpdateRuleGroupParamsWithContext(ctx context.Context) *UpdateRuleGroupParams {
	return &UpdateRuleGroupParams{
		Context: ctx,
	}
}

// NewUpdateRuleGroupParamsWithHTTPClient creates a new UpdateRuleGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRuleGroupParamsWithHTTPClient(client *http.Client) *UpdateRuleGroupParams {
	return &UpdateRuleGroupParams{
		HTTPClient: client,
	}
}

/* UpdateRuleGroupParams contains all the parameters to send to the API endpoint
   for the update rule group operation.

   Typically these are written to a http.Request.
*/
type UpdateRuleGroupParams struct {

	// Body.
	Body *models.FwmgrAPIRuleGroupModifyRequestV1

	/* Comment.

	   Audit log comment for this action
	*/
	Comment *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRuleGroupParams) WithDefaults() *UpdateRuleGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRuleGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update rule group params
func (o *UpdateRuleGroupParams) WithTimeout(timeout time.Duration) *UpdateRuleGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update rule group params
func (o *UpdateRuleGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update rule group params
func (o *UpdateRuleGroupParams) WithContext(ctx context.Context) *UpdateRuleGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update rule group params
func (o *UpdateRuleGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update rule group params
func (o *UpdateRuleGroupParams) WithHTTPClient(client *http.Client) *UpdateRuleGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update rule group params
func (o *UpdateRuleGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update rule group params
func (o *UpdateRuleGroupParams) WithBody(body *models.FwmgrAPIRuleGroupModifyRequestV1) *UpdateRuleGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update rule group params
func (o *UpdateRuleGroupParams) SetBody(body *models.FwmgrAPIRuleGroupModifyRequestV1) {
	o.Body = body
}

// WithComment adds the comment to the update rule group params
func (o *UpdateRuleGroupParams) WithComment(comment *string) *UpdateRuleGroupParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the update rule group params
func (o *UpdateRuleGroupParams) SetComment(comment *string) {
	o.Comment = comment
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRuleGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
