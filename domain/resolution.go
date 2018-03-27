package domain

type ResolutionValueData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResolutionValue struct {
	Value *ResolutionValueData `json:"value"`
}

type ResolutionStatus struct {
	Code string `json:"code"`
}

type Resolution struct {
	Authority string             `json:"authority"`
	Status    *ResolutionStatus  `json:"status"`
	Values    []*ResolutionValue `json:"values"`
}

type Resolutions struct {
	ResolutionsPerAuthority []*Resolution `json:"resolutionsPerAuthority"`
}
