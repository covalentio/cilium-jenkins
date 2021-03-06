// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NodeAddressingElement Addressing information
// swagger:model NodeAddressingElement

type NodeAddressingElement struct {

	// Node address type, one of HostName, ExternalIP or InternalIP
	AddressType string `json:"address-type,omitempty"`

	// Address pool to be used for local endpoints
	AllocRange string `json:"alloc-range,omitempty"`

	// True if address family is enabled
	Enabled bool `json:"enabled,omitempty"`

	// IP address of node
	IP string `json:"ip,omitempty"`
}

/* polymorph NodeAddressingElement address-type false */

/* polymorph NodeAddressingElement alloc-range false */

/* polymorph NodeAddressingElement enabled false */

/* polymorph NodeAddressingElement ip false */

// Validate validates this node addressing element
func (m *NodeAddressingElement) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *NodeAddressingElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeAddressingElement) UnmarshalBinary(b []byte) error {
	var res NodeAddressingElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
