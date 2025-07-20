package gpcconv

import (
	"fmt"
)

type Pounds float64
type Kilograms float64
type Grams float64

func (p Pounds) String() string { return fmt.Sprintf("%flbs", p)}
func (k Kilograms) String() string { return fmt.Sprintf("%fkg", k)}
func (g Grams) String() string { return fmt.Sprintf("%fg", g)}