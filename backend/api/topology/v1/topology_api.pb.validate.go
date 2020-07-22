// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: topology/v1/topology_api.proto

package topologyv1

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _topology_api_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetTopologyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetTopologyRequest) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetQueries() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTopologyRequestValidationError{
					field:  fmt.Sprintf("Queries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetTopologyRequestValidationError is the validation error returned by
// GetTopologyRequest.Validate if the designated constraints aren't met.
type GetTopologyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTopologyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTopologyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTopologyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTopologyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTopologyRequestValidationError) ErrorName() string {
	return "GetTopologyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTopologyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTopologyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTopologyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTopologyRequestValidationError{}

// Validate checks the field values on GetTopologyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetTopologyResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTopologyResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetTopologyResponseValidationError is the validation error returned by
// GetTopologyResponse.Validate if the designated constraints aren't met.
type GetTopologyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTopologyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTopologyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTopologyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTopologyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTopologyResponseValidationError) ErrorName() string {
	return "GetTopologyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTopologyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTopologyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTopologyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTopologyResponseValidationError{}

// Validate checks the field values on FeatureQuery with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FeatureQuery) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// FeatureQueryValidationError is the validation error returned by
// FeatureQuery.Validate if the designated constraints aren't met.
type FeatureQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeatureQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeatureQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeatureQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeatureQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeatureQueryValidationError) ErrorName() string { return "FeatureQueryValidationError" }

// Error satisfies the builtin error interface
func (e FeatureQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeatureQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeatureQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeatureQueryValidationError{}

// Validate checks the field values on Constraint with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Constraint) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Operator

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConstraintValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ConstraintValidationError is the validation error returned by
// Constraint.Validate if the designated constraints aren't met.
type ConstraintValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConstraintValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConstraintValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConstraintValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConstraintValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConstraintValidationError) ErrorName() string { return "ConstraintValidationError" }

// Error satisfies the builtin error interface
func (e ConstraintValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConstraint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConstraintValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConstraintValidationError{}

// Validate checks the field values on MetadataQuery with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MetadataQuery) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetParams()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetadataQueryValidationError{
				field:  "Params",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Aggregation

	for idx, item := range m.GetConstraints() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MetadataQueryValidationError{
					field:  fmt.Sprintf("Constraints[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MetadataQueryValidationError is the validation error returned by
// MetadataQuery.Validate if the designated constraints aren't met.
type MetadataQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataQueryValidationError) ErrorName() string { return "MetadataQueryValidationError" }

// Error satisfies the builtin error interface
func (e MetadataQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadataQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataQueryValidationError{}

// Validate checks the field values on Query with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Query) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetFeatures() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QueryValidationError{
					field:  fmt.Sprintf("Features[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetNodeMetadata() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QueryValidationError{
					field:  fmt.Sprintf("NodeMetadata[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetEdgeMetadata() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QueryValidationError{
					field:  fmt.Sprintf("EdgeMetadata[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for SourceDepth

	// no validation rules for TargetDepth

	return nil
}

// QueryValidationError is the validation error returned by Query.Validate if
// the designated constraints aren't met.
type QueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryValidationError) ErrorName() string { return "QueryValidationError" }

// Error satisfies the builtin error interface
func (e QueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryValidationError{}

// Validate checks the field values on QueryResult with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *QueryResult) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QueryResultValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetQuery()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QueryResultValidationError{
				field:  "Query",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for key, val := range m.GetNodes() {
		_ = val

		// no validation rules for Nodes[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QueryResultValidationError{
					field:  fmt.Sprintf("Nodes[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for key, val := range m.GetEdges() {
		_ = val

		// no validation rules for Edges[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QueryResultValidationError{
					field:  fmt.Sprintf("Edges[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// QueryResultValidationError is the validation error returned by
// QueryResult.Validate if the designated constraints aren't met.
type QueryResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryResultValidationError) ErrorName() string { return "QueryResultValidationError" }

// Error satisfies the builtin error interface
func (e QueryResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryResultValidationError{}

// Validate checks the field values on Node with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Node) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Features

	for key, val := range m.GetMetadata() {
		_ = val

		// no validation rules for Metadata[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeValidationError{
					field:  fmt.Sprintf("Metadata[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// NodeValidationError is the validation error returned by Node.Validate if the
// designated constraints aren't met.
type NodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeValidationError) ErrorName() string { return "NodeValidationError" }

// Error satisfies the builtin error interface
func (e NodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeValidationError{}

// Validate checks the field values on Edge with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Edge) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for SourceNodeId

	// no validation rules for TargetNodeId

	for key, val := range m.GetMetadata() {
		_ = val

		// no validation rules for Metadata[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EdgeValidationError{
					field:  fmt.Sprintf("Metadata[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// EdgeValidationError is the validation error returned by Edge.Validate if the
// designated constraints aren't met.
type EdgeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EdgeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EdgeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EdgeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EdgeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EdgeValidationError) ErrorName() string { return "EdgeValidationError" }

// Error satisfies the builtin error interface
func (e EdgeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEdge.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EdgeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EdgeValidationError{}