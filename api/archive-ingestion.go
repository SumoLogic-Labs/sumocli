package api

type ArchiveIngestion struct {
	SourceId  string `json:"sourceId"`
	Pending   int    `json:"pending"`
	Scanning  int    `json:"scanning"`
	Ingesting int    `json:"ingesting"`
	Failed    int    `json:"failed"`
	Succeeded int    `json:"succeeded"`
}

type CreateArchiveIngestion struct {
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type CreateArchiveIngestionResponse struct {
	Name                 string `json:"name"`
	StartTime            string `json:"startTime"`
	EndTime              string `json:"endTime"`
	Id                   string `json:"id"`
	TotalObjectsScanned  int    `json:"totalObjectsScanned"`
	TotalObjectsIngested int    `json:"totalObjectsIngested"`
	TotalBytesIngested   int    `json:"totalBytesIngested"`
	Status               string `json:"status"`
	CreatedAt            string `json:"createdAt"`
	CreatedBy            string `json:"createdBy"`
}

type GetArchiveIngestion struct {
	Data []CreateArchiveIngestionResponse `json:"data"`
}

type ListArchiveIngestion struct {
	Data []ArchiveIngestion `json:"data"`
}
