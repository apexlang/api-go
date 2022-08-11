package convert

import (
	"fmt"
)

const Package = "Convert"

//go:inline
func Ptr[T any](value T) *T {
	return &value
}

func Nillable[S, D any](value *S, convert func(S) D) *D {
	if value == nil {
		return nil
	}
	ret := convert(*value)
	return &ret
}

func Parse[S, D any](value *S, parse func(S) (D, error)) (*D, error) {
	if value == nil {
		return nil, nil
	}
	ret, err := parse(*value)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func StringPtr[S fmt.Stringer](value *S) *string {
	if value == nil {
		return nil
	}
	ret := (*value).String()
	return &ret
}

func Slice[S, D any](src []S, convert func(S) D) []D {
	if src == nil {
		return nil
	}
	dst := make([]D, len(src))
	for k, s := range src {
		dst[k] = convert(s)
	}
	return dst
}

func SlicePtr[S, D any](src []S, convert func(*S) D) []D {
	if src == nil {
		return nil
	}
	dst := make([]D, len(src))
	for k, s := range src {
		dst[k] = convert(&s)
	}
	return dst
}

func SliceRef[S, D any](src []S, convert func(S) *D) []D {
	if src == nil {
		return nil
	}
	dst := make([]D, len(src))
	for k, s := range src {
		dst[k] = *convert(s)
	}
	return dst
}

func Map[K comparable, S, D any](src map[K]S, convert func(S) D) map[K]D {
	if src == nil {
		return nil
	}
	dst := make(map[K]D, len(src))
	for k, s := range src {
		dst[k] = convert(s)
	}
	return dst
}

func MapRef[K comparable, S, D any](src map[K]S, convert func(S) *D) map[K]D {
	if src == nil {
		return nil
	}
	dst := make(map[K]D, len(src))
	for k, s := range src {
		dst[k] = *convert(s)
	}
	return dst
}

func MapPtr[K comparable, S, D any](src map[K]S, convert func(*S) D) map[K]D {
	if src == nil {
		return nil
	}
	dst := make(map[K]D, len(src))
	for k, s := range src {
		dst[k] = convert(&s)
	}
	return dst
}
