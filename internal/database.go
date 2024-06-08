package internal

type Form struct {
	Name        string `json:"name" sql:"name"`
	Type        string `json:"type" sql:"type"`
	Bank        string `json:"bank" sql:"bank"`
	OpsYears    int    `json:"opYear" sql:"opsyears"`
	SSM         bool   `json:"ssm" sql:"ssm"`
	PrevGateway bool   `json:"paymentgateway" sql:"prevgateway"`
	ProdType    string `json:"prodType" sql:"prodtype"`
	StoreType   string `json:"storeType" sql:"storetype"`
	Inventory   bool   `json:"inventory" sql:"inventory"`
	Reference   string `json:"reference" sql:"reference"`
	SocMedia    string `json:"socMedia" sql:"socmedia"`
	Litigation  bool   `json:"litigation" sql:"litigation"`
	Score       int    `json:"score" sql:"score"`
}