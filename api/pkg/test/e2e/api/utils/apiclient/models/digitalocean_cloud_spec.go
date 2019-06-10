// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DigitaloceanCloudSpec DigitaloceanCloudSpec specifies access data to DigitalOcean.
// swagger:model DigitaloceanCloudSpec
type DigitaloceanCloudSpec struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this digitalocean cloud spec
func (m *DigitaloceanCloudSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DigitaloceanCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DigitaloceanCloudSpec) UnmarshalBinary(b []byte) error {
	var res DigitaloceanCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}