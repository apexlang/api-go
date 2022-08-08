package convert

//go:inline
func Ptr[T any](value T) *T {
	return &value
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
