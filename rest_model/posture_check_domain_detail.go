// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostureCheckDomainDetail posture check domain detail
//
// swagger:model PostureCheckDomainDetail
type PostureCheckDomainDetail struct {
	linksField Links

	createdAtField *strfmt.DateTime

	descriptionField *string

	idField *string

	nameField *string

	tagsField Tags

	typeField *string

	updatedAtField *strfmt.DateTime

	versionField *int64

	// domains
	// Required: true
	// Min Items: 1
	Domains []string `json:"domains"`
}

// Links gets the links of this subtype
func (m *PostureCheckDomainDetail) Links() Links {
	return m.linksField
}

// SetLinks sets the links of this subtype
func (m *PostureCheckDomainDetail) SetLinks(val Links) {
	m.linksField = val
}

// CreatedAt gets the created at of this subtype
func (m *PostureCheckDomainDetail) CreatedAt() *strfmt.DateTime {
	return m.createdAtField
}

// SetCreatedAt sets the created at of this subtype
func (m *PostureCheckDomainDetail) SetCreatedAt(val *strfmt.DateTime) {
	m.createdAtField = val
}

// Description gets the description of this subtype
func (m *PostureCheckDomainDetail) Description() *string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *PostureCheckDomainDetail) SetDescription(val *string) {
	m.descriptionField = val
}

// ID gets the id of this subtype
func (m *PostureCheckDomainDetail) ID() *string {
	return m.idField
}

// SetID sets the id of this subtype
func (m *PostureCheckDomainDetail) SetID(val *string) {
	m.idField = val
}

// Name gets the name of this subtype
func (m *PostureCheckDomainDetail) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *PostureCheckDomainDetail) SetName(val *string) {
	m.nameField = val
}

// Tags gets the tags of this subtype
func (m *PostureCheckDomainDetail) Tags() Tags {
	return m.tagsField
}

// SetTags sets the tags of this subtype
func (m *PostureCheckDomainDetail) SetTags(val Tags) {
	m.tagsField = val
}

// Type gets the type of this subtype
func (m *PostureCheckDomainDetail) Type() *string {
	return m.typeField
}

// SetType sets the type of this subtype
func (m *PostureCheckDomainDetail) SetType(val *string) {
	m.typeField = val
}

// TypeID gets the type Id of this subtype
func (m *PostureCheckDomainDetail) TypeID() string {
	return "DOMAIN"
}

// SetTypeID sets the type Id of this subtype
func (m *PostureCheckDomainDetail) SetTypeID(val string) {
}

// UpdatedAt gets the updated at of this subtype
func (m *PostureCheckDomainDetail) UpdatedAt() *strfmt.DateTime {
	return m.updatedAtField
}

// SetUpdatedAt sets the updated at of this subtype
func (m *PostureCheckDomainDetail) SetUpdatedAt(val *strfmt.DateTime) {
	m.updatedAtField = val
}

// Version gets the version of this subtype
func (m *PostureCheckDomainDetail) Version() *int64 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *PostureCheckDomainDetail) SetVersion(val *int64) {
	m.versionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PostureCheckDomainDetail) UnmarshalJSON(raw []byte) error {
	var data struct {

		// domains
		// Required: true
		// Min Items: 1
		Domains []string `json:"domains"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Links Links `json:"_links"`

		CreatedAt *strfmt.DateTime `json:"createdAt"`

		Description *string `json:"description"`

		ID *string `json:"id"`

		Name *string `json:"name"`

		Tags Tags `json:"tags"`

		Type *string `json:"type"`

		TypeID string `json:"typeId"`

		UpdatedAt *strfmt.DateTime `json:"updatedAt"`

		Version *int64 `json:"version"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result PostureCheckDomainDetail

	result.linksField = base.Links

	result.createdAtField = base.CreatedAt

	result.descriptionField = base.Description

	result.idField = base.ID

	result.nameField = base.Name

	result.tagsField = base.Tags

	result.typeField = base.Type

	if base.TypeID != result.TypeID() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid typeId value: %q", base.TypeID)
	}
	result.updatedAtField = base.UpdatedAt

	result.versionField = base.Version

	result.Domains = data.Domains

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PostureCheckDomainDetail) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// domains
		// Required: true
		// Min Items: 1
		Domains []string `json:"domains"`
	}{

		Domains: m.Domains,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Links Links `json:"_links"`

		CreatedAt *strfmt.DateTime `json:"createdAt"`

		Description *string `json:"description"`

		ID *string `json:"id"`

		Name *string `json:"name"`

		Tags Tags `json:"tags"`

		Type *string `json:"type"`

		TypeID string `json:"typeId"`

		UpdatedAt *strfmt.DateTime `json:"updatedAt"`

		Version *int64 `json:"version"`
	}{

		Links: m.Links(),

		CreatedAt: m.CreatedAt(),

		Description: m.Description(),

		ID: m.ID(),

		Name: m.Name(),

		Tags: m.Tags(),

		Type: m.Type(),

		TypeID: m.TypeID(),

		UpdatedAt: m.UpdatedAt(),

		Version: m.Version(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this posture check domain detail
func (m *PostureCheckDomainDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureCheckDomainDetail) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links()); err != nil {
		return err
	}

	if err := m.Links().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("_links")
		}
		return err
	}

	return nil
}

func (m *PostureCheckDomainDetail) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt()); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckDomainDetail) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckDomainDetail) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckDomainDetail) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckDomainDetail) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags()); err != nil {
		return err
	}

	if err := m.Tags().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tags")
		}
		return err
	}

	return nil
}

func (m *PostureCheckDomainDetail) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckDomainDetail) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt()); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckDomainDetail) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckDomainDetail) validateDomains(formats strfmt.Registry) error {

	if err := validate.Required("domains", "body", m.Domains); err != nil {
		return err
	}

	iDomainsSize := int64(len(m.Domains))

	if err := validate.MinItems("domains", "body", iDomainsSize, 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostureCheckDomainDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostureCheckDomainDetail) UnmarshalBinary(b []byte) error {
	var res PostureCheckDomainDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
