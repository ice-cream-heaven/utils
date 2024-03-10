package candy

import (
	"golang.org/x/exp/constraints"
	"math/rand"
	"sort"
)

func Max[T constraints.Ordered](ss []T) (max T) {
	if len(ss) == 0 {
		return
	}

	max = ss[0]
	for _, s := range ss {
		if s > max {
			max = s
		}
	}

	return
}

func Min[T constraints.Ordered](ss []T) (min T) {
	if len(ss) == 0 {
		return
	}

	min = ss[0]
	for _, s := range ss {
		if s < min {
			min = s
		}
	}

	return
}

func Unique[T constraints.Ordered](ss []T) (ret []T) {
	m := make(map[T]struct{})
	for _, s := range ss {
		m[s] = struct{}{}
	}

	for s := range m {
		ret = append(ret, s)
	}

	return
}

func Random[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[rand.Intn(len(ss))]
}

func Each[T any](ss []T, f func(T)) {
	for _, s := range ss {
		f(s)
	}
}

func Sort[T constraints.Ordered](ss []T) []T {
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]T, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

func SortUsing[T any](ss []T, less func(a, b T) bool) []T {
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]T, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

func Filter[T any](ss []T, f func(T) bool) (ret []T) {
	for _, s := range ss {
		if f(s) {
			ret = append(ret, s)
		}
	}

	return
}

func Map[T, U any](ss []T, f func(T) U) (ret []U) {
	for _, s := range ss {
		ret = append(ret, f(s))
	}

	return
}

func Contains[T constraints.Ordered](ss []T, s T) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}

	return false
}

func Sum[T constraints.Integer | constraints.Float](ss []T) (ret T) {
	for _, s := range ss {
		ret += s
	}

	return
}

func First[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[0]
}

func Last[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[len(ss)-1]
}

func Top[T any](ss []T, n int) (ret []T) {
	if n > len(ss) {
		n = len(ss)
	}

	for i := 0; i < n; i++ {
		ret = append(ret, ss[i])
	}

	return
}

func Average[T constraints.Integer | constraints.Float](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return Sum(ss) / T(len(ss))
}

func Reverse[T any](ss []T) (ret []T) {
	for i := len(ss) - 1; i >= 0; i-- {
		ret = append(ret, ss[i])
	}

	return
}
