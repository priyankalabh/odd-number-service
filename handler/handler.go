package handler

import (
	"net/http"
	"odd-number-service/number"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OddHandler(c *gin.Context) {
	param := c.Query("max")
	if len(param) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Query param max is not given ",
		})
		return
	}
	num, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	result := number.FindOdds(num)
	c.JSON(http.StatusOK, gin.H{
		"oddNumbers": result,
	})

}