package utils

import "reflect"

//func remove(slice []int, s int) []int {
//	return append(slice[:s], slice[s+1:]...)
//}

func Remove(slice []interface{}, s int) []interface{} {
	return append(slice[:s], slice[s+1:]...)
}

func RemoveItemByIndex(slice interface{}, i int) {
	v := reflect.ValueOf(slice).Elem()
	v.Set(reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len())))
}

func SliceContains(array interface{}, value interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}
