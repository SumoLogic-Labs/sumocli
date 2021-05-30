package api

type CreateField struct {
	FieldName string `json:"fieldName"`
}

type DroppedField struct {
	FieldName string `json:"fieldName"`
}

type GetCapacityInformation struct {
	Quota     int `json:"quota"`
	Remaining int `json:"remaining"`
}

type GetDroppedFields struct {
	Data []DroppedField `json:"data"`
}

type GetFields struct {
	Data []Fields `json:"data"`
}

type Fields struct {
	FieldName string `json:"fieldName"`
	FieldId   string `json:"fieldId"`
	DataType  string `json:"dataType"`
	State     string `json:"state"`
}
