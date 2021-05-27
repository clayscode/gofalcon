// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Resources resources
//
// swagger:model .resources
type Resources struct {

	// default subscription id
	// Required: true
	DefaultSubscriptionID *string `json:"default_subscription_id"`

	// subscription ids
	// Required: true
	SubscriptionIds []string `json:"subscription_ids"`

	// tenant id
	// Required: true
	TenantID *string `json:"tenant_id"`
}

// Validate validates this resources
func (m *Resources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resources) validateDefaultSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("default_subscription_id", "body", m.DefaultSubscriptionID); err != nil {
		return err
	}

	return nil
}

func (m *Resources) validateSubscriptionIds(formats strfmt.Registry) error {

	if err := validate.Required("subscription_ids", "body", m.SubscriptionIds); err != nil {
		return err
	}

	return nil
}

func (m *Resources) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenant_id", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this resources based on context it is used
func (m *Resources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Resources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resources) UnmarshalBinary(b []byte) error {
	var res Resources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
