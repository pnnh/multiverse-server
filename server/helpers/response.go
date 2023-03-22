package helpers

import (
	"github.com/pnnh/multiverse-cloud-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ResponseCode(gctx *gin.Context, code models.MCode) {
	message := models.CodeMessage(code)
	ResponseCodeMessageData(gctx, code, message, nil)
}

func ResponseCodeMessage(gctx *gin.Context, code models.MCode, message string) {
	ResponseCodeMessageData(gctx, code, message, nil)
}

func ResponseCodeMessageData(gctx *gin.Context, code models.MCode, message string, data interface{}) {
	jsonBody := gin.H{"code": code}
	if len(message) > 0 {
		jsonBody["message"] = message
	}
	if data != nil {
		jsonBody["data"] = data
	}
	gctx.JSON(http.StatusOK, jsonBody)
}

func ResponseCodeError(gctx *gin.Context, code models.MCode, err error) {
	if err != nil {
		logrus.Errorln("ResponseCodeError", gctx.FullPath(), err)
	}
	message := models.CodeMessage(code)
	ResponseCodeMessageData(gctx, code, message, nil)
}

func ResponseCodeMessageError(gctx *gin.Context, code models.MCode, message string, err error) {
	if err != nil {
		logrus.Errorln("ResponseCodeMessageError", gctx.FullPath(), message, err)
	}
	ResponseCodeMessageData(gctx, code, message, nil)
}

func ResponseMessageError(gctx *gin.Context, message string, err error) {
	if err != nil {
		logrus.Errorln("ResponseMessageError", message, err)
	}
	ResponseCodeMessageData(gctx, models.CodeError, message, nil)
}
