package base

type xHandlerFunc func()

// MethodHandler is function base to handle all method instead use Database connection
func MethodHandler(handlerFunc2 xHandlerFunc) {
	defer DbInstantiation().Close()
	handlerFunc2()
}