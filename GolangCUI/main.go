package main

import (
	model "GolangCUI/Model"
	service "GolangCUI/Service"
	utility "GolangCUI/Utility"
	// "crypto/internal/fips140/nistec"
	"fmt"
	"time"
)

func main(){
	fmt.Println("Main from golangCUI")

	utility.InitLogger()

	// var invest *model.Investment
	invest := &model.Investment{}


	
	fmt.Println("Enter Client Id:")
	fmt.Scanln(&invest.ClientID)
	fmt.Println("Enter Invest Amount:")
	fmt.Scanln(&invest.InvestAmount)


	// fmt.Println("Enter Date Of Invest:")
	// fmt.Scanln(&invest.DateOfInvest)
	// Read DateOfInvest as string first
	var dateOfInvestStr string
	fmt.Println("Enter Date Of Invest (YYYY-MM-DD):")
	fmt.Scanln(&dateOfInvestStr)
	dateOfInvest, err := time.Parse("2006-01-02", dateOfInvestStr)
	if err != nil {
		fmt.Println("Invalid Date Of Invest format:", err)
		return
	}
	invest.DateOfInvest = dateOfInvest


	fmt.Println("Enter Investment Type:")
	fmt.Scanln(&invest.InvestmentType)
	fmt.Println("Enter Company name:")
	fmt.Scanln(&invest.CompanyName)
	fmt.Println("Enter Interest rate:")
	fmt.Scanln(&invest.InterestRate)
	fmt.Println("Enter Duration in months:")
	fmt.Scanln(&invest.DurationMonths)



	// fmt.Println("Enter maturity date:")
	// fmt.Scanln(&invest.MaturityDate)
	// Read MaturityDate as string and parse
	var maturityDateStr string
	fmt.Println("Enter Maturity Date (YYYY-MM-DD):")
	fmt.Scanln(&maturityDateStr)
	maturityDate, err := time.Parse("2006-01-02", maturityDateStr)
	if err != nil {
		fmt.Println("Invalid Maturity Date format:", err)
		return
	}
	invest.MaturityDate = maturityDate



	fmt.Println("Enter Current value:")
	fmt.Scanln(&invest.CurrentValue)
	fmt.Println("Enter status:")
	fmt.Scanln(&invest.Status)

	err = service.AddInvestment(*invest)

	if err != nil{
		fmt.Println(err.Error())
	}else {
		fmt.Println("Investment added successfully\n",invest.InvestID)
	}
		


	



}