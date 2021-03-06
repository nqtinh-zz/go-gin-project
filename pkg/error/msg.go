package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "invalid params",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "auth check token fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "auth check token timeout",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
