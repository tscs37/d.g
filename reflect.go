package discorddotgo

import "reflect"

func getTypeFrom(i interface{}) string {
	return reflect.TypeOf(i).String()
}
