// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: aperture/policy/sync/v1/policy.proto

package syncv1

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

// Validate checks the field values on PolicyWrapper with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PolicyWrapper) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PolicyWrapper with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PolicyWrapperMultiError, or
// nil if none found.
func (m *PolicyWrapper) ValidateAll() error {
	return m.validate(true)
}

func (m *PolicyWrapper) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCommonAttributes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PolicyWrapperValidationError{
					field:  "CommonAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PolicyWrapperValidationError{
					field:  "CommonAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolicyWrapperValidationError{
				field:  "CommonAttributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPolicy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PolicyWrapperValidationError{
					field:  "Policy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PolicyWrapperValidationError{
					field:  "Policy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPolicy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolicyWrapperValidationError{
				field:  "Policy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPolicyMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PolicyWrapperValidationError{
					field:  "PolicyMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PolicyWrapperValidationError{
					field:  "PolicyMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPolicyMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolicyWrapperValidationError{
				field:  "PolicyMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PolicyWrapperMultiError(errors)
	}

	return nil
}

// PolicyWrapperMultiError is an error wrapping multiple validation errors
// returned by PolicyWrapper.ValidateAll() if the designated constraints
// aren't met.
type PolicyWrapperMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PolicyWrapperMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PolicyWrapperMultiError) AllErrors() []error { return m }

// PolicyWrapperValidationError is the validation error returned by
// PolicyWrapper.Validate if the designated constraints aren't met.
type PolicyWrapperValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolicyWrapperValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolicyWrapperValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolicyWrapperValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolicyWrapperValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolicyWrapperValidationError) ErrorName() string { return "PolicyWrapperValidationError" }

// Error satisfies the builtin error interface
func (e PolicyWrapperValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolicyWrapper.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolicyWrapperValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolicyWrapperValidationError{}

// Validate checks the field values on PolicyWrappers with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PolicyWrappers) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PolicyWrappers with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PolicyWrappersMultiError,
// or nil if none found.
func (m *PolicyWrappers) ValidateAll() error {
	return m.validate(true)
}

func (m *PolicyWrappers) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetPolicyWrappers()))
		i := 0
		for key := range m.GetPolicyWrappers() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetPolicyWrappers()[key]
			_ = val

			// no validation rules for PolicyWrappers[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, PolicyWrappersValidationError{
							field:  fmt.Sprintf("PolicyWrappers[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, PolicyWrappersValidationError{
							field:  fmt.Sprintf("PolicyWrappers[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return PolicyWrappersValidationError{
						field:  fmt.Sprintf("PolicyWrappers[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return PolicyWrappersMultiError(errors)
	}

	return nil
}

// PolicyWrappersMultiError is an error wrapping multiple validation errors
// returned by PolicyWrappers.ValidateAll() if the designated constraints
// aren't met.
type PolicyWrappersMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PolicyWrappersMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PolicyWrappersMultiError) AllErrors() []error { return m }

// PolicyWrappersValidationError is the validation error returned by
// PolicyWrappers.Validate if the designated constraints aren't met.
type PolicyWrappersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolicyWrappersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolicyWrappersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolicyWrappersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolicyWrappersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolicyWrappersValidationError) ErrorName() string { return "PolicyWrappersValidationError" }

// Error satisfies the builtin error interface
func (e PolicyWrappersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolicyWrappers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolicyWrappersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolicyWrappersValidationError{}
