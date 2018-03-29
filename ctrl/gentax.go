package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/itnopadol/gen_email/model"
	"strconv"
	"net/http"
)

func GenTaxData(c *gin.Context) {
	c.Keys = headerKeys
	send_tax_amount := c.Request.URL.Query().Get("send_tax_amount")
	begin_date := c.Request.URL.Query().Get("begin_date")
	end_date := c.Request.URL.Query().Get("end_date")

	var tax float64

	tax, _ = strconv.ParseFloat(send_tax_amount, 64)

	g := new(model.TaxData)
	err := g.GenTaxData(begin_date, end_date, tax, Dbc)

	rs := resp.
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error :" + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = g
		c.JSON(http.StatusOK, rs)
	}

}
