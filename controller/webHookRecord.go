package controller

import (
	"FEDeployService/database"
	"FEDeployService/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebHookRecordFind(c *gin.Context)  {

	id := c.Query("id")

	var result []models.WebHookRecord

	if database.DB.Where("repository_id = ?", id).Order("created_at desc").Find(&result).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    database.DB.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
		"msg":    "查询成功",
	})
}