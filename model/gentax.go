package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type TaxData struct {
	Id              int     `json:"id" db:"id"`
	Month           int     `json:"month" db:"month"`
	Year            int     `json:"year" db:"year"`
	DocDate         string  `json:"doc_date" db:"doc_date"`
	DocNo           string  `json:"doc_no" db:"doc_no"`
	ArCode          string  `json:"ar_code" db:"ar_code"`
	ArName          string  `json:"ar_name" db:"ar_name"`
	TaxNo           string  `json:"tax_no" db:"tax_no"`
	Address         string  `json:"address" db:"address"`
	BeforeTaxAmount float64 `json:"before_tax_amount" db:"before_tax_amount"`
	TaxAmount       float64 `json:"tax_amount" db:"tax_amount"`
	TotalAmount     float64 `json:"total_amount" db:"total_amount"`
}

func (t *TaxData) GenTaxData(begindate string, enddate string, SendTaxAmount float64, db *sqlx.DB) error {

	var vDay int;
	sql := `select count(docdate) as day1 from (select distinct docdate from dbo.bcarinvoice where  docdate between '01/03/2018' and '31/03/2018') as aa)`
	err := db.Get(vDay, sql)
	if err != nil {
		fmt.Println("Count Day =", err.Error())
		return err
	}
	return nil

}
