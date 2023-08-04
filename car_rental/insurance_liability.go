package main

type LiabilityInsurance struct {
	Insurance
	LiabilityAmount float64
}

// NewLiabilityInsurance creates a new liability insurance instance
func NewLiabilityInsurance() *LiabilityInsurance {
	return &LiabilityInsurance{
		Insurance: Insurance{
			Provider:   "SureGuard",
			Coverage:   "ThirdParty",
			CostPerDay: 18.0,
		},
		LiabilityAmount: 100000.0,
	}
}
