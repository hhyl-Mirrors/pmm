// Code generated by go-swagger; DO NOT EDIT.

package metrics_names

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

// GetMetricsNamesReader is a Reader for the GetMetricsNames structure.
type GetMetricsNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricsNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetricsNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMetricsNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMetricsNamesOK creates a GetMetricsNamesOK with default headers values
func NewGetMetricsNamesOK() *GetMetricsNamesOK {
	return &GetMetricsNamesOK{}
}

/*GetMetricsNamesOK handles this case with default header values.

A successful response.
*/
type GetMetricsNamesOK struct {
	Payload *GetMetricsNamesOKBody
}

func (o *GetMetricsNamesOK) Error() string {
	return fmt.Sprintf("[POST /v0/qan/GetMetricsNames][%d] getMetricsNamesOk  %+v", 200, o.Payload)
}

func (o *GetMetricsNamesOK) GetPayload() *GetMetricsNamesOKBody {
	return o.Payload
}

func (o *GetMetricsNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetMetricsNamesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsNamesDefault creates a GetMetricsNamesDefault with default headers values
func NewGetMetricsNamesDefault(code int) *GetMetricsNamesDefault {
	return &GetMetricsNamesDefault{
		_statusCode: code,
	}
}

/*GetMetricsNamesDefault handles this case with default header values.

An unexpected error response.
*/
type GetMetricsNamesDefault struct {
	_statusCode int

	Payload *GetMetricsNamesDefaultBody
}

// Code gets the status code for the get metrics names default response
func (o *GetMetricsNamesDefault) Code() int {
	return o._statusCode
}

func (o *GetMetricsNamesDefault) Error() string {
	return fmt.Sprintf("[POST /v0/qan/GetMetricsNames][%d] GetMetricsNames default  %+v", o._statusCode, o.Payload)
}

func (o *GetMetricsNamesDefault) GetPayload() *GetMetricsNamesDefaultBody {
	return o.Payload
}

func (o *GetMetricsNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetMetricsNamesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

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

/*GetMetricsNamesDefaultBody get metrics names default body
swagger:model GetMetricsNamesDefaultBody
*/
type GetMetricsNamesDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this get metrics names default body
func (o *GetMetricsNamesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMetricsNamesDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("GetMetricsNames default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMetricsNamesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMetricsNamesDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetMetricsNamesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetMetricsNamesOKBody MetricsNamesReply is map of stored metrics:
// key is root of metric name in db (Ex:. [m_]query_time[_sum]);
// value - Human readable name of metrics.
swagger:model GetMetricsNamesOKBody
*/
type GetMetricsNamesOKBody struct {

	// data
	Data map[string]string `json:"data,omitempty"`
}

// Validate validates this get metrics names OK body
func (o *GetMetricsNamesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetMetricsNamesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMetricsNamesOKBody) UnmarshalBinary(b []byte) error {
	var res GetMetricsNamesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
