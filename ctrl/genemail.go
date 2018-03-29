package ctrl

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/itnopadol/gen_email/model"
)

func GenEmail(c *gin.Context) {
	c.Keys = headerKeys
	access_token := c.Request.URL.Query().Get("access_token")
	ar_code := c.Request.URL.Query().Get("ar_code")
	ar_name := c.Request.URL.Query().Get("ar_name")
	doc_no := c.Request.URL.Query().Get("doc_no")
	email := c.Request.URL.Query().Get("email")

	g := new(model.GenEmail)
	err := g.GenEmailAuto(access_token, ar_code, ar_name, doc_no, email)
	fmt.Println("Ctrl Send Email ")

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = nil
		c.JSON(http.StatusOK, rs)
	}

}
