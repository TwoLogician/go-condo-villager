package model

import "mime/multipart"

type AbstractAttachmentRequest[K any] struct {
	Data K                     `form:"data"`
	File *multipart.FileHeader `form:"file"`
}
