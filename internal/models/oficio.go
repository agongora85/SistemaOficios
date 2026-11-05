package models

type Oficio struct {
	ID             int    `json:"id"`
	Folio          string `json:"folio"`
	Asunto         string `json:"asunto"`
	Remitente      string `json:"remitente"`
	FechaRecepcion string `json:"fecha_recepcion"`
	Estatus        string `json:"estatus"`
}
