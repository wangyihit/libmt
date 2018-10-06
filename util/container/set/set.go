package set

import "github.com/deckarep/golang-set"

func NewMapSetFromStringSlice(ss []string) mapset.Set {
	is := make([]interface{}, len(ss))
	for i, v := range ss {
		is[i] = v
	}
	return mapset.NewSetFromSlice(is)
}

func NewMapSetFromInt64Slice(ss []int64) mapset.Set {
	is := make([]interface{}, len(ss))
	for i, v := range ss {
		is[i] = v
	}
	return mapset.NewSetFromSlice(is)
}
