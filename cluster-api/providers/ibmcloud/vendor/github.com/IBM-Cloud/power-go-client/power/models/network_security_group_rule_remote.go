// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkSecurityGroupRuleRemote network security group rule remote
//
// swagger:model NetworkSecurityGroupRuleRemote
type NetworkSecurityGroupRuleRemote struct {

	// The ID of the remote Network Address Group or Network Security Group the rules apply to. Not required for default-network-address-group
	ID string `json:"id,omitempty"`

	// The type of remote group the rules apply to
	// Enum: ["network-security-group","network-address-group","default-network-address-group"]
	Type string `json:"type,omitempty"`
}

// Validate validates this network security group rule remote
func (m *NetworkSecurityGroupRuleRemote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var networkSecurityGroupRuleRemoteTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["network-security-group","network-address-group","default-network-address-group"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkSecurityGroupRuleRemoteTypeTypePropEnum = append(networkSecurityGroupRuleRemoteTypeTypePropEnum, v)
	}
}

const (

	// NetworkSecurityGroupRuleRemoteTypeNetworkDashSecurityDashGroup captures enum value "network-security-group"
	NetworkSecurityGroupRuleRemoteTypeNetworkDashSecurityDashGroup string = "network-security-group"

	// NetworkSecurityGroupRuleRemoteTypeNetworkDashAddressDashGroup captures enum value "network-address-group"
	NetworkSecurityGroupRuleRemoteTypeNetworkDashAddressDashGroup string = "network-address-group"

	// NetworkSecurityGroupRuleRemoteTypeDefaultDashNetworkDashAddressDashGroup captures enum value "default-network-address-group"
	NetworkSecurityGroupRuleRemoteTypeDefaultDashNetworkDashAddressDashGroup string = "default-network-address-group"
)

// prop value enum
func (m *NetworkSecurityGroupRuleRemote) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkSecurityGroupRuleRemoteTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkSecurityGroupRuleRemote) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network security group rule remote based on context it is used
func (m *NetworkSecurityGroupRuleRemote) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkSecurityGroupRuleRemote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkSecurityGroupRuleRemote) UnmarshalBinary(b []byte) error {
	var res NetworkSecurityGroupRuleRemote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
