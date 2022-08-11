package convert

import (
	"fmt"

	"github.com/apexlang/api-go/errorz"
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

func NillableErr[S, D any](value *S, parse func(S) (D, error)) (*D, error) {
	if value == nil {
		return nil, nil
	}
	ret, err := parse(*value)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func NillableEt[S, D any](e *errorz.Tracker, value *S, parse func(S) (D, error)) *D {
	if value == nil {
		return nil
	}
	ret, err := parse(*value)
	if err != nil {
		e.Append(err)
		return nil
	}
	return &ret
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

func SliceErr[S, D any](src []S, convert func(S) (D, error)) ([]D, error) {
	if src == nil {
		return nil, nil
	}
	dst := make([]D, len(src))
	for k, s := range src {
		var err error
		dst[k], err = convert(s)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func SliceEt[S, D any](e *errorz.Tracker, src []S, convert func(S) (D, error)) []D {
	if src == nil {
		return nil
	}
	dst := make([]D, len(src))
	for k, s := range src {
		var err error
		dst[k], err = convert(s)
		if err != nil {
			e.Append(err)
			return nil
		}
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

func SlicePtrErr[S, D any](src []S, convert func(*S) (D, error)) ([]D, error) {
	if src == nil {
		return nil, nil
	}
	dst := make([]D, len(src))
	for k, s := range src {
		var err error
		dst[k], err = convert(&s)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func SlicePtrRt[S, D any](e *errorz.Tracker, src []S, convert func(*S) (D, error)) []D {
	if src == nil {
		return nil
	}
	dst := make([]D, len(src))
	for k, s := range src {
		var err error
		dst[k], err = convert(&s)
		if err != nil {
			e.Append(err)
			return nil
		}
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

func SliceRefErr[S, D any](src []S, convert func(S) (*D, error)) ([]D, error) {
	if src == nil {
		return nil, nil
	}
	dst := make([]D, len(src))
	for k, s := range src {
		val, err := convert(s)
		if err != nil {
			return nil, err
		}
		dst[k] = *val
	}
	return dst, nil
}

func SliceRefEt[S, D any](e *errorz.Tracker, src []S, convert func(S) (*D, error)) []D {
	if src == nil {
		return nil
	}
	dst := make([]D, len(src))
	for k, s := range src {
		val, err := convert(s)
		if err != nil {
			e.Append(err)
			return nil
		}
		dst[k] = *val
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

func MapErr[K comparable, S, D any](src map[K]S, convert func(S) (D, error)) (map[K]D, error) {
	if src == nil {
		return nil, nil
	}
	dst := make(map[K]D, len(src))
	for k, s := range src {
		var err error
		dst[k], err = convert(s)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func MapEt[K comparable, S, D any](e *errorz.Tracker, src map[K]S, convert func(S) (D, error)) map[K]D {
	if src == nil {
		return nil
	}
	dst := make(map[K]D, len(src))
	for k, s := range src {
		var err error
		dst[k], err = convert(s)
		if err != nil {
			e.Append(err)
			return nil
		}
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

func MapRefErr[K comparable, S, D any](src map[K]S, convert func(S) (*D, error)) (map[K]D, error) {
	if src == nil {
		return nil, nil
	}
	dst := make(map[K]D, len(src))
	for k, s := range src {
		val, err := convert(s)
		if err != nil {
			return nil, err
		}
		dst[k] = *val
	}
	return dst, nil
}

func MapRefEt[K comparable, S, D any](e *errorz.Tracker, src map[K]S, convert func(S) (*D, error)) map[K]D {
	if src == nil {
		return nil
	}
	dst := make(map[K]D, len(src))
	for k, s := range src {
		val, err := convert(s)
		if err != nil {
			e.Append(err)
			return nil
		}
		dst[k] = *val
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

func MapPtrErr[K comparable, S, D any](src map[K]S, convert func(*S) (D, error)) (map[K]D, error) {
	if src == nil {
		return nil, nil
	}
	dst := make(map[K]D, len(src))
	for k, s := range src {
		var err error
		dst[k], err = convert(&s)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func MapPtrEt[K comparable, S, D any](e *errorz.Tracker, src map[K]S, convert func(*S) (D, error)) map[K]D {
	if src == nil {
		return nil
	}
	dst := make(map[K]D, len(src))
	for k, s := range src {
		var err error
		dst[k], err = convert(&s)
		if err != nil {
			e.Append(err)
			return nil
		}
	}
	return dst
}
