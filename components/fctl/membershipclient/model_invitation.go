/*
Membership API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membershipclient

import (
	"encoding/json"
	"time"
)

// checks if the Invitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invitation{}

// Invitation struct for Invitation
type Invitation struct {
	Id string `json:"id"`
	OrganizationId string `json:"organizationId"`
	UserEmail string `json:"userEmail"`
	Status string `json:"status"`
	CreationDate time.Time `json:"creationDate"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// User roles
	// Deprecated
	Roles []string `json:"roles,omitempty"`
	StackClaims []StackClaimsInner `json:"stackClaims,omitempty"`
}

// NewInvitation instantiates a new Invitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitation(id string, organizationId string, userEmail string, status string, creationDate time.Time) *Invitation {
	this := Invitation{}
	this.Id = id
	this.OrganizationId = organizationId
	this.UserEmail = userEmail
	this.Status = status
	this.CreationDate = creationDate
	return &this
}

// NewInvitationWithDefaults instantiates a new Invitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationWithDefaults() *Invitation {
	this := Invitation{}
	return &this
}

// GetId returns the Id field value
func (o *Invitation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Invitation) SetId(v string) {
	o.Id = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *Invitation) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *Invitation) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetUserEmail returns the UserEmail field value
func (o *Invitation) GetUserEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEmail, true
}

// SetUserEmail sets field value
func (o *Invitation) SetUserEmail(v string) {
	o.UserEmail = v
}

// GetStatus returns the Status field value
func (o *Invitation) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Invitation) SetStatus(v string) {
	o.Status = v
}

// GetCreationDate returns the CreationDate field value
func (o *Invitation) GetCreationDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetCreationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationDate, true
}

// SetCreationDate sets field value
func (o *Invitation) SetCreationDate(v time.Time) {
	o.CreationDate = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Invitation) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Invitation) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Invitation) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
// Deprecated
func (o *Invitation) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Invitation) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Invitation) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
// Deprecated
func (o *Invitation) SetRoles(v []string) {
	o.Roles = v
}

// GetStackClaims returns the StackClaims field value if set, zero value otherwise.
func (o *Invitation) GetStackClaims() []StackClaimsInner {
	if o == nil || IsNil(o.StackClaims) {
		var ret []StackClaimsInner
		return ret
	}
	return o.StackClaims
}

// GetStackClaimsOk returns a tuple with the StackClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetStackClaimsOk() ([]StackClaimsInner, bool) {
	if o == nil || IsNil(o.StackClaims) {
		return nil, false
	}
	return o.StackClaims, true
}

// HasStackClaims returns a boolean if a field has been set.
func (o *Invitation) HasStackClaims() bool {
	if o != nil && !IsNil(o.StackClaims) {
		return true
	}

	return false
}

// SetStackClaims gets a reference to the given []StackClaimsInner and assigns it to the StackClaims field.
func (o *Invitation) SetStackClaims(v []StackClaimsInner) {
	o.StackClaims = v
}

func (o Invitation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["userEmail"] = o.UserEmail
	toSerialize["status"] = o.Status
	toSerialize["creationDate"] = o.CreationDate
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.StackClaims) {
		toSerialize["stackClaims"] = o.StackClaims
	}
	return toSerialize, nil
}

type NullableInvitation struct {
	value *Invitation
	isSet bool
}

func (v NullableInvitation) Get() *Invitation {
	return v.value
}

func (v *NullableInvitation) Set(val *Invitation) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitation(val *Invitation) *NullableInvitation {
	return &NullableInvitation{value: val, isSet: true}
}

func (v NullableInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


