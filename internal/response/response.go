package response

import "github.com/gin-gonic/gin"

// IDRes -.
type IDRes struct {
	ID int64 `json:"id"`
}

// RowAffectRest -.
type RowAffectRes struct {
	RowAffected int64 `json:"row_affected"`
}

// ResponseByID -.
func ResponseByID(c *gin.Context, code int, id int64) {
	c.JSON(code, IDRes{ID: id})
}

// ResponseByRowAffect -.
func ResponseByRowAffect(c *gin.Context, code int, rowAffected int64) {
	c.JSON(code, RowAffectRes{RowAffected: rowAffected})
}
