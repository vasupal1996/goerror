package goerror

import "encoding/json"

// Map object of the error
func Map(err error) map[string]interface{} {
	emptyCtx := context{}
	var errMap = make(map[string]interface{})
	if customErr, ok := err.(*goError); ok && customErr.context != emptyCtx {
		errMap["field"] = customErr.context.Key
		errMap["message"] = customErr.context.Value
		errMap["type"] = GetType(customErr)
		return errMap
	}
	errMap["message"] = err.Error()
	errMap["type"] = GetType(err)
	return errMap
}

// JSON convert error object into bytes json
func JSON(err error) []byte {
	jsonStr, _ := json.Marshal(Map(err))
	return jsonStr
}
