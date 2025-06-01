package service

import (
	bo "GolangCUI/BO"
	model "GolangCUI/Model"
	utility "GolangCUI/Utility"
)

type InvestService struct {
}

func AddInvestment(invest model.Investment)error{

	err := bo.AddInvestment(invest)

	if err != nil{
		utility.ErrorLogger.Println(err.Error())
		if customErr, ok := err.(*utility.CustomError);ok{
			return customErr
		}
		return utility.NewCustomError(500,"Internal Server Error")
	}

	utility.InfoLogger.Println("Investment Added successfully")
	return nil
}

