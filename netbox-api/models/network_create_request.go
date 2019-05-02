// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkCreateRequest network create request
// swagger:model NetworkCreateRequest
type NetworkCreateRequest struct {

	// The description of this prefix
	Description string `json:"description,omitempty"`

	// The id of the partition
	Partition string `json:"partition,omitempty"`

	// The prefixes to create in the netbox
	// Required: true
	Prefixes []string `json:"prefixes"`

	// The name of the project to assign this prefix to
	Project string `json:"project,omitempty"`

	// The name of the tenant to assign this prefix to
	Tenant string `json:"tenant,omitempty"`
}

// Validate validates this network create request
func (m *NetworkCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrefixes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkCreateRequest) validatePrefixes(formats strfmt.Registry) error {

	if err := validate.Required("prefixes", "body", m.Prefixes); err != nil {
		return err
	}

	for i := 0; i < len(m.Prefixes); i++ {

		if err := validate.MinLength("prefixes"+"."+strconv.Itoa(i), "body", string(m.Prefixes[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkCreateRequest) UnmarshalBinary(b []byte) error {
	var res NetworkCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
