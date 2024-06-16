package utils

import "reflect"

// Filter lọc các phần tử của slice theo điều kiện cho trước
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map ánh xạ các phần tử của slice sang một dạng khác
func Map[T any, U any](slice []T, mapper func(T) U) []U {
	var result []U
	for _, item := range slice {
		result = append(result, mapper(item))
	}
	return result
}

// Contains checks if a slice contains the given element.
func Contains(slice interface{}, element interface{}) bool {
	// Kiểm tra xem slice có phải là kiểu slice không
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		panic("Provided slice parameter is not a slice")
	}

	// Duyệt qua từng phần tử trong slice và so sánh với element
	for i := 0; i < sliceValue.Len(); i++ {
		if reflect.DeepEqual(sliceValue.Index(i).Interface(), element) {
			return true
		}
	}

	return false
}
