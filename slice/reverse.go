package slice

import (
	"github.com/ljun20160606/gox/reflectx"
	"reflect"
)

func Reverse(s interface{}) error {
	valueOf := reflect.ValueOf(s)
	indirect := reflect.Indirect(valueOf)
	if indirect.Kind() != reflect.Slice {
		return reflectx.MustSlice
	}
	for i, j := 0, valueOf.Len()-1; i < j; i, j = i+1, j-1 {
		left := valueOf.Index(i)
		right := valueOf.Index(j)
		temp := reflect.ValueOf(left.Interface())
		left.Set(right)
		right.Set(temp)
	}
	return nil
}
