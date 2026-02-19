// Code generated from semantic convention specification. DO NOT EDIT.

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package testconv provides types and functionality for OpenTelemetry semantic
// conventions in the "test" namespace.
package testconv

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

var (
	addOptPool = &sync.Pool{New: func() any { return &[]metric.AddOption{} }}
	recOptPool = &sync.Pool{New: func() any { return &[]metric.RecordOption{} }}
)

// Metric is an instrument used to record metric values conforming to the
// "test.metric" semantic conventions. It represents a test metric.
type Metric struct {
	metric.Int64Counter
}

var newMetricOpts = []metric.Int64CounterOption{
	metric.WithDescription("A test metric"),
	metric.WithUnit("1"),
}

// NewMetric returns a new Metric instrument.
func NewMetric(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (Metric, error) {
	// Check if the meter is nil.
	if m == nil {
		return Metric{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newMetricOpts
	} else {
		opt = append(opt, newMetricOpts...)
	}

	i, err := m.Int64Counter(
		"test.metric",
		opt...,
	)
	if err != nil {
		return Metric{noop.Int64Counter{}}, err
	}
	return Metric{i}, nil
}

// Inst returns the underlying metric instrument.
func (m Metric) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (Metric) Name() string {
	return "test.metric"
}

// Unit returns the semantic convention unit of the instrument
func (Metric) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (Metric) Description() string {
	return "A test metric"
}

// Add adds incr to the existing count for attrs.
//
// The attr is the a test attribute
func (m Metric) Add(
	ctx context.Context,
	incr int64,
	attr string,
	attrs ...attribute.KeyValue,
) {
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := addOptPool.Get().(*[]metric.AddOption)
	defer func() {
		*o = (*o)[:0]
		addOptPool.Put(o)
	}()

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs,
				attribute.String("test.attr", attr),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
func (m Metric) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	if set.Len() == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := addOptPool.Get().(*[]metric.AddOption)
	defer func() {
		*o = (*o)[:0]
		addOptPool.Put(o)
	}()

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Counter.Add(ctx, incr, *o...)
}
