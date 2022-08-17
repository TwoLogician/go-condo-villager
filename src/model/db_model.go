package model

type WaterBillInfo struct {
	AbstractBaseAttachment
	Note string `json:"note"`
	Url  string `json:"url"`
}

func NewWaterBillInfo() WaterBillInfo {
	return WaterBillInfo{AbstractBaseAttachment: newAbstractBaseAttachment()}
}
