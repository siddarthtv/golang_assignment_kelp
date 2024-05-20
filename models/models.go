package models

type Financials struct {
	CompanyID     int     `db:"companyID"`
	OperatingCost float64 `db:"operating_cost"`
	Revenue       float64
	Profit        float64
	Loss          float64
}

type Sales struct {
	CompanyID  int     `db:"companyID"`
	GrossSales float64 `db:"gross_sales"`
}

type EmpStats struct {
	CompanyID   int     `db:"companyID"`
	TotalCount  int     `db:"total_count"`
	MaleCount   int     `db:"male_count"`
	FemaleCount int     `db:"female_count"`
	AverageAge  float64 `db:"average_age"`
}
