package dao

import (
	// dao "GolangCUI/DAO"
	model "GolangCUI/Model"
	utility "GolangCUI/Utility"
	"database/sql"
	"fmt"
)

type InvestDAO struct {
	db *sql.DB
}

func AddInvestment(invest model.Investment)error{

	db := utility.GetDBConnection()
	defer db.Close()

	query := "Insert into investments (client_id, invest_amount, date_of_invest, investment_type, company_name, interest_rate, duration_months, maturity_date, current_value, status) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := db.Exec(query,invest.ClientID,invest.InvestAmount,invest.DateOfInvest,invest.InvestmentType,invest.CompanyName,invest.InterestRate,invest.DurationMonths,invest.MaturityDate,invest.CurrentValue,invest.Status)
	if err != nil{
		utility.ErrorLogger.Println("Insert error:", err)
		return utility.NewCustomError(500,"Failed to add investment")
	}

	utility.InfoLogger.Printf("%s", fmt.Sprintf("Investment (investId=%d) added succesfully.",invest.InvestID))
	return nil
}

