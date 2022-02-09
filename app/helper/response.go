package helper

func ErrorResponseBuilder(msg string, errs interface{}) map[string]interface{} {
	res := map[string]interface{}{
		"status": "Failed",
		"msg":    msg,
		"err":    errs,
	}
	return res

}

func ResponseBuilder(msg string, data interface{}) map[string]interface{} {
	res := map[string]interface{}{
		"status": "OK",
		"msg":    msg,
		"data":   data,
	}
	return res
}
