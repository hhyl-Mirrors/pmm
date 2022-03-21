// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChangePSMDBComponentsReader is a Reader for the ChangePSMDBComponents structure.
type ChangePSMDBComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangePSMDBComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangePSMDBComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangePSMDBComponentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangePSMDBComponentsOK creates a ChangePSMDBComponentsOK with default headers values
func NewChangePSMDBComponentsOK() *ChangePSMDBComponentsOK {
	return &ChangePSMDBComponentsOK{}
}

/*ChangePSMDBComponentsOK handles this case with default header values.

A successful response.
*/
type ChangePSMDBComponentsOK struct {
	Payload interface{}
}

func (o *ChangePSMDBComponentsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/ChangePSMDB][%d] changePsmdbComponentsOk  %+v", 200, o.Payload)
}

func (o *ChangePSMDBComponentsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangePSMDBComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangePSMDBComponentsDefault creates a ChangePSMDBComponentsDefault with default headers values
func NewChangePSMDBComponentsDefault(code int) *ChangePSMDBComponentsDefault {
	return &ChangePSMDBComponentsDefault{
		_statusCode: code,
	}
}

/*ChangePSMDBComponentsDefault handles this case with default header values.

An unexpected error response.
*/
type ChangePSMDBComponentsDefault struct {
	_statusCode int

	Payload *ChangePSMDBComponentsDefaultBody
}

// Code gets the status code for the change PSMDB components default response
func (o *ChangePSMDBComponentsDefault) Code() int {
	return o._statusCode
}

func (o *ChangePSMDBComponentsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/ChangePSMDB][%d] ChangePSMDBComponents default  %+v", o._statusCode, o.Payload)
}

func (o *ChangePSMDBComponentsDefault) GetPayload() *ChangePSMDBComponentsDefaultBody {
	return o.Payload
}

func (o *ChangePSMDBComponentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangePSMDBComponentsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangePSMDBComponentsBody change PSMDB components body
swagger:model ChangePSMDBComponentsBody
*/
type ChangePSMDBComponentsBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// mongod
	Mongod *ChangePSMDBComponentsParamsBodyMongod `json:"mongod,omitempty"`
}

// Validate validates this change PSMDB components body
func (o *ChangePSMDBComponentsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePSMDBComponentsBody) validateMongod(formats strfmt.Registry) error {

	if swag.IsZero(o.Mongod) { // not required
		return nil
	}

	if o.Mongod != nil {
		if err := o.Mongod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "mongod")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangePSMDBComponentsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePSMDBComponentsBody) UnmarshalBinary(b []byte) error {
	var res ChangePSMDBComponentsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangePSMDBComponentsDefaultBody change PSMDB components default body
swagger:model ChangePSMDBComponentsDefaultBody
*/
type ChangePSMDBComponentsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change PSMDB components default body
func (o *ChangePSMDBComponentsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePSMDBComponentsDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangePSMDBComponents default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangePSMDBComponentsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePSMDBComponentsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangePSMDBComponentsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangePSMDBComponentsParamsBodyMongod ChangeComponent contains fields to manage components.
swagger:model ChangePSMDBComponentsParamsBodyMongod
*/
type ChangePSMDBComponentsParamsBodyMongod struct {

	// default version
	DefaultVersion string `json:"default_version,omitempty"`

	// versions
	Versions []*ChangePSMDBComponentsParamsBodyMongodVersionsItems0 `json:"versions"`
}

// Validate validates this change PSMDB components params body mongod
func (o *ChangePSMDBComponentsParamsBodyMongod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePSMDBComponentsParamsBodyMongod) validateVersions(formats strfmt.Registry) error {

	if swag.IsZero(o.Versions) { // not required
		return nil
	}

	for i := 0; i < len(o.Versions); i++ {
		if swag.IsZero(o.Versions[i]) { // not required
			continue
		}

		if o.Versions[i] != nil {
			if err := o.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "mongod" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangePSMDBComponentsParamsBodyMongod) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePSMDBComponentsParamsBodyMongod) UnmarshalBinary(b []byte) error {
	var res ChangePSMDBComponentsParamsBodyMongod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangePSMDBComponentsParamsBodyMongodVersionsItems0 ComponentVersion contains operations which should be done with component version.
swagger:model ChangePSMDBComponentsParamsBodyMongodVersionsItems0
*/
type ChangePSMDBComponentsParamsBodyMongodVersionsItems0 struct {

	// version
	Version string `json:"version,omitempty"`

	// disable
	Disable bool `json:"disable,omitempty"`

	// enable
	Enable bool `json:"enable,omitempty"`
}

// Validate validates this change PSMDB components params body mongod versions items0
func (o *ChangePSMDBComponentsParamsBodyMongodVersionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangePSMDBComponentsParamsBodyMongodVersionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePSMDBComponentsParamsBodyMongodVersionsItems0) UnmarshalBinary(b []byte) error {
	var res ChangePSMDBComponentsParamsBodyMongodVersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
