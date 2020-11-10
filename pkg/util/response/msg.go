package response

var responseMsg = map[int]string{
	ResponseOK:            "OK",
	ResponseError:         "Fail",
	ResponseInvalidParams: "Invalid request parameters",
	DBError:               "Database operation failed",
}

func ResponseMsg(code int) string {
	return responseMsg[code]
}
