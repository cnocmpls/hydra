// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OAuth2AccessRequest Requester is a token endpoint's request context.
//
// swagger:model oAuth2AccessRequest
type OAuth2AccessRequest struct {

	// ClientID is the identifier of the OAuth 2.0 client.
	ClientID string `json:"client_id,omitempty"`

	// GrantTypes is the requests grant types.
	GrantTypes []string `json:"grant_types"`

	// GrantedAudience is the list of audiences granted to the OAuth 2.0 client.
	GrantedAudience []string `json:"granted_audience"`

	// GrantedScopes is the list of scopes granted to the OAuth 2.0 client.
	GrantedScopes []string `json:"granted_scopes"`
}

// Validate validates this o auth2 access request
func (m *OAuth2AccessRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this o auth2 access request based on context it is used
func (m *OAuth2AccessRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OAuth2AccessRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuth2AccessRequest) UnmarshalBinary(b []byte) error {
	var res OAuth2AccessRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
