package dao

import model "GolangCUI/Model"

type Invest interface {
	AddInvestment(invest model.Investment)error

}