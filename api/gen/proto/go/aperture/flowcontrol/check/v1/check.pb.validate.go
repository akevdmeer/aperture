// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: aperture/flowcontrol/check/v1/check.proto

package checkv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CheckRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CheckRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CheckRequestMultiError, or
// nil if none found.
func (m *CheckRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ControlPoint

	// no validation rules for Labels

	// no validation rules for Tokens

	if len(errors) > 0 {
		return CheckRequestMultiError(errors)
	}

	return nil
}

// CheckRequestMultiError is an error wrapping multiple validation errors
// returned by CheckRequest.ValidateAll() if the designated constraints aren't met.
type CheckRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckRequestMultiError) AllErrors() []error { return m }

// CheckRequestValidationError is the validation error returned by
// CheckRequest.Validate if the designated constraints aren't met.
type CheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckRequestValidationError) ErrorName() string { return "CheckRequestValidationError" }

// Error satisfies the builtin error interface
func (e CheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckRequestValidationError{}

// Validate checks the field values on CheckResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CheckResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CheckResponseMultiError, or
// nil if none found.
func (m *CheckResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStart()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CheckResponseValidationError{
					field:  "Start",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CheckResponseValidationError{
					field:  "Start",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckResponseValidationError{
				field:  "Start",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEnd()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CheckResponseValidationError{
					field:  "End",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CheckResponseValidationError{
					field:  "End",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnd()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckResponseValidationError{
				field:  "End",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ControlPoint

	// no validation rules for TelemetryFlowLabels

	// no validation rules for DecisionType

	// no validation rules for RejectReason

	for idx, item := range m.GetClassifierInfos() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CheckResponseValidationError{
						field:  fmt.Sprintf("ClassifierInfos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CheckResponseValidationError{
						field:  fmt.Sprintf("ClassifierInfos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResponseValidationError{
					field:  fmt.Sprintf("ClassifierInfos[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetFluxMeterInfos() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CheckResponseValidationError{
						field:  fmt.Sprintf("FluxMeterInfos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CheckResponseValidationError{
						field:  fmt.Sprintf("FluxMeterInfos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResponseValidationError{
					field:  fmt.Sprintf("FluxMeterInfos[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetLimiterDecisions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CheckResponseValidationError{
						field:  fmt.Sprintf("LimiterDecisions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CheckResponseValidationError{
						field:  fmt.Sprintf("LimiterDecisions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResponseValidationError{
					field:  fmt.Sprintf("LimiterDecisions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CheckResponseMultiError(errors)
	}

	return nil
}

// CheckResponseMultiError is an error wrapping multiple validation errors
// returned by CheckResponse.ValidateAll() if the designated constraints
// aren't met.
type CheckResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckResponseMultiError) AllErrors() []error { return m }

// CheckResponseValidationError is the validation error returned by
// CheckResponse.Validate if the designated constraints aren't met.
type CheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResponseValidationError) ErrorName() string { return "CheckResponseValidationError" }

// Error satisfies the builtin error interface
func (e CheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResponseValidationError{}

// Validate checks the field values on ClassifierInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClassifierInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClassifierInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClassifierInfoMultiError,
// or nil if none found.
func (m *ClassifierInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *ClassifierInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PolicyName

	// no validation rules for PolicyHash

	// no validation rules for ClassifierIndex

	// no validation rules for Error

	if len(errors) > 0 {
		return ClassifierInfoMultiError(errors)
	}

	return nil
}

// ClassifierInfoMultiError is an error wrapping multiple validation errors
// returned by ClassifierInfo.ValidateAll() if the designated constraints
// aren't met.
type ClassifierInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClassifierInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClassifierInfoMultiError) AllErrors() []error { return m }

// ClassifierInfoValidationError is the validation error returned by
// ClassifierInfo.Validate if the designated constraints aren't met.
type ClassifierInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClassifierInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClassifierInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClassifierInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClassifierInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClassifierInfoValidationError) ErrorName() string { return "ClassifierInfoValidationError" }

// Error satisfies the builtin error interface
func (e ClassifierInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClassifierInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClassifierInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClassifierInfoValidationError{}

// Validate checks the field values on LimiterDecision with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LimiterDecision) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LimiterDecision with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LimiterDecisionMultiError, or nil if none found.
func (m *LimiterDecision) ValidateAll() error {
	return m.validate(true)
}

func (m *LimiterDecision) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PolicyName

	// no validation rules for PolicyHash

	// no validation rules for ComponentId

	// no validation rules for Dropped

	// no validation rules for Reason

	switch v := m.Details.(type) {
	case *LimiterDecision_RateLimiterInfo_:
		if v == nil {
			err := LimiterDecisionValidationError{
				field:  "Details",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetRateLimiterInfo()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LimiterDecisionValidationError{
						field:  "RateLimiterInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LimiterDecisionValidationError{
						field:  "RateLimiterInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRateLimiterInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LimiterDecisionValidationError{
					field:  "RateLimiterInfo",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LimiterDecision_ConcurrencyLimiterInfo_:
		if v == nil {
			err := LimiterDecisionValidationError{
				field:  "Details",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetConcurrencyLimiterInfo()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LimiterDecisionValidationError{
						field:  "ConcurrencyLimiterInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LimiterDecisionValidationError{
						field:  "ConcurrencyLimiterInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetConcurrencyLimiterInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LimiterDecisionValidationError{
					field:  "ConcurrencyLimiterInfo",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LimiterDecision_FlowRegulatorInfo_:
		if v == nil {
			err := LimiterDecisionValidationError{
				field:  "Details",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetFlowRegulatorInfo()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LimiterDecisionValidationError{
						field:  "FlowRegulatorInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LimiterDecisionValidationError{
						field:  "FlowRegulatorInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetFlowRegulatorInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LimiterDecisionValidationError{
					field:  "FlowRegulatorInfo",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return LimiterDecisionMultiError(errors)
	}

	return nil
}

// LimiterDecisionMultiError is an error wrapping multiple validation errors
// returned by LimiterDecision.ValidateAll() if the designated constraints
// aren't met.
type LimiterDecisionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LimiterDecisionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LimiterDecisionMultiError) AllErrors() []error { return m }

// LimiterDecisionValidationError is the validation error returned by
// LimiterDecision.Validate if the designated constraints aren't met.
type LimiterDecisionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LimiterDecisionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LimiterDecisionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LimiterDecisionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LimiterDecisionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LimiterDecisionValidationError) ErrorName() string { return "LimiterDecisionValidationError" }

// Error satisfies the builtin error interface
func (e LimiterDecisionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLimiterDecision.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LimiterDecisionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LimiterDecisionValidationError{}

// Validate checks the field values on FluxMeterInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FluxMeterInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FluxMeterInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FluxMeterInfoMultiError, or
// nil if none found.
func (m *FluxMeterInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *FluxMeterInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FluxMeterName

	if len(errors) > 0 {
		return FluxMeterInfoMultiError(errors)
	}

	return nil
}

// FluxMeterInfoMultiError is an error wrapping multiple validation errors
// returned by FluxMeterInfo.ValidateAll() if the designated constraints
// aren't met.
type FluxMeterInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FluxMeterInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FluxMeterInfoMultiError) AllErrors() []error { return m }

// FluxMeterInfoValidationError is the validation error returned by
// FluxMeterInfo.Validate if the designated constraints aren't met.
type FluxMeterInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FluxMeterInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FluxMeterInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FluxMeterInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FluxMeterInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FluxMeterInfoValidationError) ErrorName() string { return "FluxMeterInfoValidationError" }

// Error satisfies the builtin error interface
func (e FluxMeterInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFluxMeterInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FluxMeterInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FluxMeterInfoValidationError{}

// Validate checks the field values on LimiterDecision_RateLimiterInfo with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LimiterDecision_RateLimiterInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LimiterDecision_RateLimiterInfo with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// LimiterDecision_RateLimiterInfoMultiError, or nil if none found.
func (m *LimiterDecision_RateLimiterInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *LimiterDecision_RateLimiterInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Remaining

	// no validation rules for Current

	// no validation rules for Label

	if len(errors) > 0 {
		return LimiterDecision_RateLimiterInfoMultiError(errors)
	}

	return nil
}

// LimiterDecision_RateLimiterInfoMultiError is an error wrapping multiple
// validation errors returned by LimiterDecision_RateLimiterInfo.ValidateAll()
// if the designated constraints aren't met.
type LimiterDecision_RateLimiterInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LimiterDecision_RateLimiterInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LimiterDecision_RateLimiterInfoMultiError) AllErrors() []error { return m }

// LimiterDecision_RateLimiterInfoValidationError is the validation error
// returned by LimiterDecision_RateLimiterInfo.Validate if the designated
// constraints aren't met.
type LimiterDecision_RateLimiterInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LimiterDecision_RateLimiterInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LimiterDecision_RateLimiterInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LimiterDecision_RateLimiterInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LimiterDecision_RateLimiterInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LimiterDecision_RateLimiterInfoValidationError) ErrorName() string {
	return "LimiterDecision_RateLimiterInfoValidationError"
}

// Error satisfies the builtin error interface
func (e LimiterDecision_RateLimiterInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLimiterDecision_RateLimiterInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LimiterDecision_RateLimiterInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LimiterDecision_RateLimiterInfoValidationError{}

// Validate checks the field values on LimiterDecision_ConcurrencyLimiterInfo
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *LimiterDecision_ConcurrencyLimiterInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// LimiterDecision_ConcurrencyLimiterInfo with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// LimiterDecision_ConcurrencyLimiterInfoMultiError, or nil if none found.
func (m *LimiterDecision_ConcurrencyLimiterInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *LimiterDecision_ConcurrencyLimiterInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WorkloadIndex

	if len(errors) > 0 {
		return LimiterDecision_ConcurrencyLimiterInfoMultiError(errors)
	}

	return nil
}

// LimiterDecision_ConcurrencyLimiterInfoMultiError is an error wrapping
// multiple validation errors returned by
// LimiterDecision_ConcurrencyLimiterInfo.ValidateAll() if the designated
// constraints aren't met.
type LimiterDecision_ConcurrencyLimiterInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LimiterDecision_ConcurrencyLimiterInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LimiterDecision_ConcurrencyLimiterInfoMultiError) AllErrors() []error { return m }

// LimiterDecision_ConcurrencyLimiterInfoValidationError is the validation
// error returned by LimiterDecision_ConcurrencyLimiterInfo.Validate if the
// designated constraints aren't met.
type LimiterDecision_ConcurrencyLimiterInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LimiterDecision_ConcurrencyLimiterInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LimiterDecision_ConcurrencyLimiterInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LimiterDecision_ConcurrencyLimiterInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LimiterDecision_ConcurrencyLimiterInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LimiterDecision_ConcurrencyLimiterInfoValidationError) ErrorName() string {
	return "LimiterDecision_ConcurrencyLimiterInfoValidationError"
}

// Error satisfies the builtin error interface
func (e LimiterDecision_ConcurrencyLimiterInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLimiterDecision_ConcurrencyLimiterInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LimiterDecision_ConcurrencyLimiterInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LimiterDecision_ConcurrencyLimiterInfoValidationError{}

// Validate checks the field values on LimiterDecision_FlowRegulatorInfo with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *LimiterDecision_FlowRegulatorInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LimiterDecision_FlowRegulatorInfo
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// LimiterDecision_FlowRegulatorInfoMultiError, or nil if none found.
func (m *LimiterDecision_FlowRegulatorInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *LimiterDecision_FlowRegulatorInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Label

	if len(errors) > 0 {
		return LimiterDecision_FlowRegulatorInfoMultiError(errors)
	}

	return nil
}

// LimiterDecision_FlowRegulatorInfoMultiError is an error wrapping multiple
// validation errors returned by
// LimiterDecision_FlowRegulatorInfo.ValidateAll() if the designated
// constraints aren't met.
type LimiterDecision_FlowRegulatorInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LimiterDecision_FlowRegulatorInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LimiterDecision_FlowRegulatorInfoMultiError) AllErrors() []error { return m }

// LimiterDecision_FlowRegulatorInfoValidationError is the validation error
// returned by LimiterDecision_FlowRegulatorInfo.Validate if the designated
// constraints aren't met.
type LimiterDecision_FlowRegulatorInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LimiterDecision_FlowRegulatorInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LimiterDecision_FlowRegulatorInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LimiterDecision_FlowRegulatorInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LimiterDecision_FlowRegulatorInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LimiterDecision_FlowRegulatorInfoValidationError) ErrorName() string {
	return "LimiterDecision_FlowRegulatorInfoValidationError"
}

// Error satisfies the builtin error interface
func (e LimiterDecision_FlowRegulatorInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLimiterDecision_FlowRegulatorInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LimiterDecision_FlowRegulatorInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LimiterDecision_FlowRegulatorInfoValidationError{}
