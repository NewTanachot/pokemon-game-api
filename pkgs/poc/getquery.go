package poc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryTest struct {
	Test string `binding:"required" form:"test"`
}

func TestQuery(c *gin.Context) {
	var req QueryTest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"test": req.Test})
}
