// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderBookL2 order book l2
// swagger:model OrderBookL2
type OrderBookL2 struct {

	// id
	// Required: true
	ID *int64 `json:"id"`

	// price
	Price float64 `json:"price,omitempty"`

	// side
	// Required: true
	Side *string `json:"side"`

	// size
	Size int64 `json:"size,omitempty"`

	// symbol
	// Required: true
	Symbol *string `json:"symbol"`
}

// Validate validates this order book l2
func (m *OrderBookL2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSide(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSymbol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderBookL2) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *OrderBookL2) validateSide(formats strfmt.Registry) error {

	if err := validate.Required("side", "body", m.Side); err != nil {
		return err
	}

	return nil
}

func (m *OrderBookL2) validateSymbol(formats strfmt.Registry) error {

	if err := validate.Required("symbol", "body", m.Symbol); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderBookL2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderBookL2) UnmarshalBinary(b []byte) error {
	var res OrderBookL2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
