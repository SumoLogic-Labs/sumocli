package cmdutils

func GenerateFieldsMap(fieldNames []string, fieldValues []string) map[string]string {
	if len(fieldNames) > 0 && len(fieldValues) > 0 {
		fieldsMap := make(map[string]string)
		for i, _ := range fieldNames {
			fieldsMap[fieldNames[i]] = fieldValues[i]
			i++
		}
		return fieldsMap
	} else {
		return nil
	}
}
