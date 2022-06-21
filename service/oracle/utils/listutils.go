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
