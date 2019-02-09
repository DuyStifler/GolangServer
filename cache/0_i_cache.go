package cache

type ICaching interface {
	convertFromObjectToString(value interface{}) string
	convertFromStringToObject(str string) interface{}
}
