package response

var responseMsg = map[int]string{
	OK:            "OK",
	Error:         "Fail",
	InvalidParams: "Invalid request parameters",
	DBError:       "Database operation failed",
}

func ResponseMsg(code int) string {
	return responseMsg[code]
}
