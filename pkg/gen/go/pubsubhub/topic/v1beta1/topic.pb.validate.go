// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pubsubhub/topic/v1beta1/topic.proto

package topicv1beta1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Message) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Payload

	// no validation rules for Attributes

	// no validation rules for MessageId

	if v, ok := interface{}(m.GetPublishTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "PublishTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// MessageValidationError is the validation error returned by Message.Validate
// if the designated constraints aren't met.
type MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageValidationError) ErrorName() string { return "MessageValidationError" }

// Error satisfies the builtin error interface
func (e MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageValidationError{}

// Validate checks the field values on Topic with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Topic) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Labels

	return nil
}

// TopicValidationError is the validation error returned by Topic.Validate if
// the designated constraints aren't met.
type TopicValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TopicValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TopicValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TopicValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TopicValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TopicValidationError) ErrorName() string { return "TopicValidationError" }

// Error satisfies the builtin error interface
func (e TopicValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTopic.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TopicValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TopicValidationError{}