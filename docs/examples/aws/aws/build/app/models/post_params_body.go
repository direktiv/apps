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

// PostParamsBody post params body
//
// swagger:model postParamsBody
type PostParamsBody struct {

	// access key
	// Required: true
	AccessKey *string `json:"access-key"`

	// region
	Region string `json:"region,omitempty"`

	// secret key
	// Required: true
	SecretKey *string `json:"secret-key"`
}

// Validate validates this post params body
func (m *PostParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBody) validateAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("access-key", "body", m.AccessKey); err != nil {
		return err
	}

	return nil
}

func (m *PostParamsBody) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("secret-key", "body", m.SecretKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post params body based on context it is used
func (m *PostParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBody) UnmarshalBinary(b []byte) error {
	var res PostParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}