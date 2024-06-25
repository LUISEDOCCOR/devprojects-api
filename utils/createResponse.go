package utils

func CreateResponse(mode string, msg string) map[string]string {
	return map[string]string{
		"mode": mode,
		"msg":  msg,
	}
}
