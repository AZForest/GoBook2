package generalpurposeconverter

import (
	"fmt"
)

type pounds float64
type kilograms float64
type grams float64

func (p pounds) String() string { return fmt.Sprintf("%flbs", p)}
func (k kilograms) String() string { return fmt.Sprintf("%fkg", k)}
func (g grams) String() string { return fmt.Sprintf("%fg", g)}