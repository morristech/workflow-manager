package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Job job
// swagger:model Job
type Job struct {

	// attempts
	Attempts []*JobAttempt `json:"attempts"`

	// container
	Container string `json:"container,omitempty"`

	// created at
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// input
	Input string `json:"input,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// output
	Output string `json:"output,omitempty"`

	// queue
	Queue string `json:"queue,omitempty"`

	// started at
	StartedAt strfmt.DateTime `json:"startedAt,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// state resource
	StateResource *StateResource `json:"stateResource,omitempty"`

	// status
	Status JobStatus `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`

	// stopped at
	StoppedAt strfmt.DateTime `json:"stoppedAt,omitempty"`
}

// Validate validates this job
func (m *Job) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttempts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStateResource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Job) validateAttempts(formats strfmt.Registry) error {

	if swag.IsZero(m.Attempts) { // not required
		return nil
	}

	for i := 0; i < len(m.Attempts); i++ {

		if swag.IsZero(m.Attempts[i]) { // not required
			continue
		}

		if m.Attempts[i] != nil {

			if err := m.Attempts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attempts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Job) validateStateResource(formats strfmt.Registry) error {

	if swag.IsZero(m.StateResource) { // not required
		return nil
	}

	if m.StateResource != nil {

		if err := m.StateResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stateResource")
			}
			return err
		}
	}

	return nil
}

func (m *Job) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}
