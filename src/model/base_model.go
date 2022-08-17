package model

import (
	"time"

	"twodo.app/condo/utility"
)

type AbstractBaseDb struct {
	ID          string    `json:"id"`
	CreatedDate time.Time `json:"createdDate"`
}

func newAbstractBaseDb() AbstractBaseDb {
	return AbstractBaseDb{ID: utility.NewUuid(), CreatedDate: time.Now()}
}

type AbstractBaseAttachment struct {
	AbstractBaseDb
	AttachmentPath string `json:"attachmentPath"`
	ContentType    string `json:"contentType"`
	Filename       string `json:"filename"`
}

func newAbstractBaseAttachment() AbstractBaseAttachment {
	return AbstractBaseAttachment{AbstractBaseDb: newAbstractBaseDb()}
}
