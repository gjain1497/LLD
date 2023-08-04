package main

type BelongingInsurance struct {
	Insurance
	MaxCoverageAmount float64
}

func NewBelongingInsurance() *BelongingInsurance {
	return &BelongingInsurance{
		Insurance: Insurance{
			Provider:   "SafeProtect",
			Coverage:   "Property",
			CostPerDay: 12.0,
		},
		MaxCoverageAmount: 3000.0,
	}
}
