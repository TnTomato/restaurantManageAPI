package response

var responseMsg = map[int]string{
	OK:            "OK",
	Error:         "Fail",
	InvalidParams: "Invalid request parameters",
	NotFound:      "Resource not found",
	DBError:       "Database operation failed",
	DuplicatedName: "Duplicated name",
}

func ResponseMsg(code int) string {
	return responseMsg[code]
}
