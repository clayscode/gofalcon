// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResponsesDeviceControlSettingsV1 responses device control settings v1
//
// swagger:model responses.DeviceControlSettingsV1
type ResponsesDeviceControlSettingsV1 struct {

	// classes
	// Required: true
	Classes []*ResponsesDeviceControlPolicyClassSettingsV1 `json:"classes"`

	// Does the end user receives a notification when the policy is violated
	// Required: true
	// Enum: [TRUE FALSE]
	EndUserNotification *string `json:"end_user_notification"`

	// [How] is this policy enforced
	// Required: true
	// Enum: [ENFORCED MONITORY_ONLY]
	EnforcementMode *string `json:"enforcement_mode"`
}

// Validate validates this responses device control settings v1
func (m *ResponsesDeviceControlSettingsV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClasses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndUserNotification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnforcementMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesDeviceControlSettingsV1) validateClasses(formats strfmt.Registry) error {

	if err := validate.Required("classes", "body", m.Classes); err != nil {
		return err
	}

	for i := 0; i < len(m.Classes); i++ {
		if swag.IsZero(m.Classes[i]) { // not required
			continue
		}

		if m.Classes[i] != nil {
			if err := m.Classes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("classes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("classes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var responsesDeviceControlSettingsV1TypeEndUserNotificationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TRUE","FALSE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		responsesDeviceControlSettingsV1TypeEndUserNotificationPropEnum = append(responsesDeviceControlSettingsV1TypeEndUserNotificationPropEnum, v)
	}
}

const (

	// ResponsesDeviceControlSettingsV1EndUserNotificationTRUE captures enum value "TRUE"
	ResponsesDeviceControlSettingsV1EndUserNotificationTRUE string = "TRUE"

	// ResponsesDeviceControlSettingsV1EndUserNotificationFALSE captures enum value "FALSE"
	ResponsesDeviceControlSettingsV1EndUserNotificationFALSE string = "FALSE"
)

// prop value enum
func (m *ResponsesDeviceControlSettingsV1) validateEndUserNotificationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, responsesDeviceControlSettingsV1TypeEndUserNotificationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResponsesDeviceControlSettingsV1) validateEndUserNotification(formats strfmt.Registry) error {

	if err := validate.Required("end_user_notification", "body", m.EndUserNotification); err != nil {
		return err
	}

	// value enum
	if err := m.validateEndUserNotificationEnum("end_user_notification", "body", *m.EndUserNotification); err != nil {
		return err
	}

	return nil
}

var responsesDeviceControlSettingsV1TypeEnforcementModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENFORCED","MONITORY_ONLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		responsesDeviceControlSettingsV1TypeEnforcementModePropEnum = append(responsesDeviceControlSettingsV1TypeEnforcementModePropEnum, v)
	}
}

const (

	// ResponsesDeviceControlSettingsV1EnforcementModeENFORCED captures enum value "ENFORCED"
	ResponsesDeviceControlSettingsV1EnforcementModeENFORCED string = "ENFORCED"

	// ResponsesDeviceControlSettingsV1EnforcementModeMONITORYONLY captures enum value "MONITORY_ONLY"
	ResponsesDeviceControlSettingsV1EnforcementModeMONITORYONLY string = "MONITORY_ONLY"
)

// prop value enum
func (m *ResponsesDeviceControlSettingsV1) validateEnforcementModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, responsesDeviceControlSettingsV1TypeEnforcementModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResponsesDeviceControlSettingsV1) validateEnforcementMode(formats strfmt.Registry) error {

	if err := validate.Required("enforcement_mode", "body", m.EnforcementMode); err != nil {
		return err
	}

	// value enum
	if err := m.validateEnforcementModeEnum("enforcement_mode", "body", *m.EnforcementMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this responses device control settings v1 based on the context it is used
func (m *ResponsesDeviceControlSettingsV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClasses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesDeviceControlSettingsV1) contextValidateClasses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Classes); i++ {

		if m.Classes[i] != nil {
			if err := m.Classes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("classes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("classes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponsesDeviceControlSettingsV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponsesDeviceControlSettingsV1) UnmarshalBinary(b []byte) error {
	var res ResponsesDeviceControlSettingsV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
