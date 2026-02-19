// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated from semantic convention specification. DO NOT EDIT.

package semconv // import "go.opentelemetry.io/otel/semconv/v1.30.0"

import "go.opentelemetry.io/otel/attribute"

// Namespace: test
const (
	// TestAttrKey is the attribute Key conforming to the "test.attr" semantic
	// conventions. It represents a test attribute.
	//
	// Type: string
	// RequirementLevel:
	// Stability: Stable
	//
	// Examples:
	TestAttrKey = attribute.Key("test.attr")
)

// TestAttr returns an attribute KeyValue conforming to the "test.attr" semantic
// conventions. It represents a test attribute.
func TestAttr(val string) attribute.KeyValue {
	return TestAttrKey.String(val)
}