package bo

import (
	dao "GolangCUI/DAO"
	model "GolangCUI/Model"
	utility "GolangCUI/Utility"
)

type InvetBO struct {
}

func AddInvestment(invest model.Investment)error{
	if invest.ClientID == 0{
		return utility.NewCustomError(400,"Client id not equal to zero")
	}

	return dao.AddInvestment(invest)
}

