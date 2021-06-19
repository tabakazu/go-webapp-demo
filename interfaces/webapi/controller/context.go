package controller

type Context interface {
	String(code int, s string) error
	JSON(code int, i interface{}) error
	Get(key string) interface{}
	Bind(i interface{}) error
	ApiSessionUserID() int
}
