package utils

func FromInterfaceToSliceMap[U string, R interface{}](s interface{}) []map[U]R {
	slice := make([]map[U]R, 0)

	if interfaceSlice, ok := s.([]interface{}); ok {
		for _, item := range interfaceSlice {
			if item, ok := item.(map[U]R); ok {
				slice = append(slice, item)
			}
		}
	}

	return slice
}
