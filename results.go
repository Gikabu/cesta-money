package money

type CurrencyResult struct {
	Code        string `json:"code"`
	NumericCode string `json:"numeric_code"`
	Grapheme    string `json:"grapheme"`
}

type AmountResult struct {
	CurrencyCode string         `json:"currency_code"`
	Currency     CurrencyResult `json:"currency"`
	Value        float64        `json:"value"`
	DisplayValue string         `json:"display_value"`
}

func (m *Money) AsResult() AmountResult {
	return AmountResult{
		CurrencyCode: m.currency.Code,
		Currency: CurrencyResult{
			Code:        m.currency.Code,
			NumericCode: m.currency.NumericCode,
			Grapheme:    m.currency.Grapheme,
		},
		Value:        m.AsMajorUnits(),
		DisplayValue: m.Display(),
	}
}
