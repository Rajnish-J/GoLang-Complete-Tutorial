package model

import "time"


type Investment struct {
    InvestID       int
    ClientID       int
    InvestAmount   float64
    DateOfInvest   time.Time
    InvestmentType string
    CompanyName    string
    InterestRate   float64
    DurationMonths int
    MaturityDate   time.Time
    CurrentValue   float64
    Status         string
}