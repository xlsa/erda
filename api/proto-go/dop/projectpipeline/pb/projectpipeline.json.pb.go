// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: projectpipeline.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*ListAppRequest)(nil)
var _ json.Unmarshaler = (*ListAppRequest)(nil)
var _ json.Marshaler = (*ListAppResponse)(nil)
var _ json.Unmarshaler = (*ListAppResponse)(nil)
var _ json.Marshaler = (*Application)(nil)
var _ json.Unmarshaler = (*Application)(nil)
var _ json.Marshaler = (*CreateProjectPipelineRequest)(nil)
var _ json.Unmarshaler = (*CreateProjectPipelineRequest)(nil)
var _ json.Marshaler = (*CreateProjectPipelineResponse)(nil)
var _ json.Unmarshaler = (*CreateProjectPipelineResponse)(nil)
var _ json.Marshaler = (*ProjectPipeline)(nil)
var _ json.Unmarshaler = (*ProjectPipeline)(nil)

// ListAppRequest implement json.Marshaler.
func (m *ListAppRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListAppRequest implement json.Marshaler.
func (m *ListAppRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListAppResponse implement json.Marshaler.
func (m *ListAppResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListAppResponse implement json.Marshaler.
func (m *ListAppResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Application implement json.Marshaler.
func (m *Application) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Application implement json.Marshaler.
func (m *Application) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateProjectPipelineRequest implement json.Marshaler.
func (m *CreateProjectPipelineRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateProjectPipelineRequest implement json.Marshaler.
func (m *CreateProjectPipelineRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateProjectPipelineResponse implement json.Marshaler.
func (m *CreateProjectPipelineResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateProjectPipelineResponse implement json.Marshaler.
func (m *CreateProjectPipelineResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ProjectPipeline implement json.Marshaler.
func (m *ProjectPipeline) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ProjectPipeline implement json.Marshaler.
func (m *ProjectPipeline) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}