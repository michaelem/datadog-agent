// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019 Datadog, Inc.

//+build !zlib

package jsonstream

import (
	"github.com/DataDog/datadog-agent/pkg/forwarder"
	"github.com/DataDog/datadog-agent/pkg/serializer/marshaler"
)

// PayloadBuilder can not be used without zlib
type PayloadBuilder struct{}

// NewPayloadBuilder can not be used without zlib
func NewPayloadBuilder() *PayloadBuilder { return nil }

// Build can not be used without zlib
func (b *PayloadBuilder) Build(m marshaler.StreamJSONMarshaler) (forwarder.Payloads, error) {
	return nil, nil
}
