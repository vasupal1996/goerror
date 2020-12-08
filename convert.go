package goerror

import "encoding/json"

// goMap object of the error
func goMap(err error) map[string]interface{} {
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

// goJSON convert error object into bytes json
func goJSON(err error) []byte {
	jsonStr, _ := json.Marshal(goMap(err))
	return jsonStr
}
