package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"twodo.app/condo/model"
	"twodo.app/condo/utility"
)

type waterBillService struct {
	Get  func(*gin.Context)
	Post func(*gin.Context)
}

var waterBill = waterBillService{
	Get: func(c *gin.Context) {
		db, err := DB.Create()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, model.NewErrorInfo(err))
			return
		}
		var waterBills []model.WaterBillInfo
		result := db.Find(&waterBills)
		if result.Error != nil {
			c.IndentedJSON(http.StatusInternalServerError, model.NewErrorInfo(err))
			return
		}
		c.IndentedJSON(http.StatusOK, waterBills)
	},
	Post: func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if utility.String.ToSha512(authorization) != "670506c9b67375007e1b50d698085e54bcd6e8bc6a7fece2a907ff42b83a92698147460dee66c1856ac9777875c7f0e5b1edac57486e7a841311ea6beeb235a2" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var form model.AbstractAttachmentRequest[model.WaterBillInfo]
		if err := c.ShouldBind(&form); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, model.NewErrorInfo(err))
			return
		}

		file := form.File
		path := fmt.Sprintf("%v/%v-%v", utility.AttachmentPath, utility.NewUuid(), file.Filename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, model.NewErrorInfo(err))
			return
		}

		if err := mail.SendMail(path, "", fmt.Sprintf("ค่าน้ำประปา %v", form.Data.Note)); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, model.NewErrorInfo(err))
			return
		}

		waterBill := model.NewWaterBillInfo()
		waterBill.AttachmentPath = path
		waterBill.ContentType = file.Header.Get("Content-Type")
		waterBill.Filename = file.Filename
		waterBill.Note = form.Data.Note

		db, err := DB.Create()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, model.NewErrorInfo(err))
			return
		}
		db.Create(&waterBill)

		c.IndentedJSON(http.StatusCreated, waterBill)
	},
}
