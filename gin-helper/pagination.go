package gin_helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 分页计算

type Pagination struct {
	Total  int64
	Offset int64
	Limit  int64
	Result interface{}
}

// 分页计算
func Paginator(page int64, limit int64) Pagination {
	p := Pagination{
		Limit:  limit,
		Offset: 0,
	}
	if page > 1 {
		p.Offset = (page - 1) * limit
	}
	return p
}

func WritePaginationResp(c *gin.Context, payload Pagination) {
	// always return http.StatusOK
	c.JSON(http.StatusOK, payload)
}
