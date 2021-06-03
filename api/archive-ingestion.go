package api

type ArchiveIngestion struct {
	SourceId  string `json:"sourceId"`
	Pending   int    `json:"pending"`
	Scanning  int    `json:"scanning"`
	Ingesting int    `json:"ingesting"`
	Failed    int    `json:"failed"`
	Succeeded int    `json:"succeeded"`
}

type ListArchiveIngestion struct {
	Data []ArchiveIngestion `json:"data"`
}
