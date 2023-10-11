package utils

import "reflect"

func MapStruct[S any, D any](source S) D {
	v := reflect.ValueOf(source)
	typeOfS := v.Type()

	var instanceD = new(D)

	typeOfD := reflect.ValueOf(instanceD).Elem()

	for i := 0; i < v.NumField(); i++ {
		fieldName := typeOfS.Field(i).Name
		fieldValue := v.Field(i)

		dField := typeOfD.FieldByName(fieldName)
		if dField.IsValid() {
			if dField.CanSet() {
				if dField.Kind() == fieldValue.Kind() {
					dField.Set(fieldValue)
				}
			}
		}
	}

	return *instanceD
}
