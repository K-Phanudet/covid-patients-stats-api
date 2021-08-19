package models

type PatientInfo struct {
	ConfirmDate    string `json:"ConfirmDate"`
	No             int    `json:"No"`
	Age            int    `json:"Age"`
	Gender         string `json:"Gender"`
	Nation         string `json:"Nation"`
	NationEn       string `json:"NationEn"`
	Province       string `json:"Province"`
	ProvinceId     uint16 `json:"ProvinceId"`
	District       string `json:"District"`
	ProvinceEn     string `json:"ProvinceEn"`
	StatQuarantine uint8  `json:"StatQuarantine"`
}

type Patients struct {
	Data []PatientInfo `json:"Data"`
}
