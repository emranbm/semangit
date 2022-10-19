package utils

func GetResultOrPanicError[TResult interface{}](result TResult, error error) TResult {
	if error != nil {
		panic(error)
	}
	return result
}

func PanicError(error error) {
	if error != nil {
		panic(error)
	}
}
