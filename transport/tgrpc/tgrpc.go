package tgrpc

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/apexlang/api-go/errorz"
)

type RegisterFn func(server grpc.ServiceRegistrar)

func Register(registrar grpc.ServiceRegistrar, services ...RegisterFn) {
	for _, s := range services {
		s(registrar)
	}
}

func Error(err error) error {
	if err == nil {
		return nil
	}
	e := errorz.Translate(err)

	// TODO: advanced usage per https://jbrandhorst.com/post/grpc-errors/
	return status.Error(codes.Code(e.Code), e.Message)
}

func ConvertInputI8Ptr(val *int32) *int8 {
	if val == nil {
		return nil
	}
	ret := int8(*val)
	return &ret
}

func ConvertOutputI8Ptr(val *int8) *int32 {
	if val == nil {
		return nil
	}
	ret := int32(*val)
	return &ret
}

func ConvertInputI16Ptr(val *int32) *int16 {
	if val == nil {
		return nil
	}
	ret := int16(*val)
	return &ret
}

func ConvertOutputI16Ptr(val *int16) *int32 {
	if val == nil {
		return nil
	}
	ret := int32(*val)
	return &ret
}

func ConvertInputU8Ptr(val *uint32) *uint8 {
	if val == nil {
		return nil
	}
	ret := uint8(*val)
	return &ret
}

func ConvertInputTimestamp(val *timestamppb.Timestamp) *time.Time {
	if val == nil {
		return nil
	}
	ret := (*val).AsTime()
	return &ret
}

func ConvertOutputU8Ptr(val *uint8) *uint32 {
	if val == nil {
		return nil
	}
	ret := uint32(*val)
	return &ret
}

func ConvertInputU16Ptr(val *uint32) *uint16 {
	if val == nil {
		return nil
	}
	ret := uint16(*val)
	return &ret
}

func ConvertOutputU16Ptr(val *uint16) *uint32 {
	if val == nil {
		return nil
	}
	ret := uint32(*val)
	return &ret
}

func ConvertOutputTimestamp(val *time.Time) *timestamppb.Timestamp {
	if val == nil {
		return nil
	}
	return timestamppb.New(*val)
}
