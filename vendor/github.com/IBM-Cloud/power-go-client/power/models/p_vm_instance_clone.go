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

// PVMInstanceClone p VM instance clone
//
// swagger:model PVMInstanceClone
type PVMInstanceClone struct {

	// The name of the SSH key pair provided to the server for authenticating users (looked up in the tenant's list of keys)
	KeyPairName string `json:"keyPairName,omitempty"`

	// Amount of memory allocated (in GB)
	Memory *float64 `json:"memory,omitempty"`

	// Name of the server to create
	// Required: true
	Name *string `json:"name"`

	// The pvm instance networks information
	// Required: true
	Networks []*PVMInstanceAddNetwork `json:"networks"`

	// Processor type (dedicated, shared, capped)
	// Enum: [dedicated shared capped]
	ProcType *string `json:"procType,omitempty"`

	// Number of processors allocated
	Processors *float64 `json:"processors,omitempty"`

	// The pvm instance Software Licenses
	SoftwareLicenses *SoftwareLicenses `json:"softwareLicenses,omitempty"`

	// List of volume IDs
	VolumeIDs []string `json:"volumeIDs"`
}

// Validate validates this p VM instance clone
func (m *PVMInstanceClone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareLicenses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceClone) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceClone) validateNetworks(formats strfmt.Registry) error {

	if err := validate.Required("networks", "body", m.Networks); err != nil {
		return err
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var pVmInstanceCloneTypeProcTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dedicated","shared","capped"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceCloneTypeProcTypePropEnum = append(pVmInstanceCloneTypeProcTypePropEnum, v)
	}
}

const (

	// PVMInstanceCloneProcTypeDedicated captures enum value "dedicated"
	PVMInstanceCloneProcTypeDedicated string = "dedicated"

	// PVMInstanceCloneProcTypeShared captures enum value "shared"
	PVMInstanceCloneProcTypeShared string = "shared"

	// PVMInstanceCloneProcTypeCapped captures enum value "capped"
	PVMInstanceCloneProcTypeCapped string = "capped"
)

// prop value enum
func (m *PVMInstanceClone) validateProcTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pVmInstanceCloneTypeProcTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstanceClone) validateProcType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProcTypeEnum("procType", "body", *m.ProcType); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceClone) validateSoftwareLicenses(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareLicenses) { // not required
		return nil
	}

	if m.SoftwareLicenses != nil {
		if err := m.SoftwareLicenses.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareLicenses")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softwareLicenses")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this p VM instance clone based on the context it is used
func (m *PVMInstanceClone) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareLicenses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceClone) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Networks); i++ {

		if m.Networks[i] != nil {

			if swag.IsZero(m.Networks[i]) { // not required
				return nil
			}

			if err := m.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PVMInstanceClone) contextValidateSoftwareLicenses(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareLicenses != nil {

		if swag.IsZero(m.SoftwareLicenses) { // not required
			return nil
		}

		if err := m.SoftwareLicenses.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareLicenses")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softwareLicenses")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceClone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceClone) UnmarshalBinary(b []byte) error {
	var res PVMInstanceClone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
