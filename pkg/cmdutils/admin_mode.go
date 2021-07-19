package cmdutils

func AdminMode(isAdminMode bool) string {
	if isAdminMode == true {
		return "true"
	} else {
		return "false"
	}
}
