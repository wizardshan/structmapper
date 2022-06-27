package response

import "fmt"

type Money int

func (m Money) MarshalJSON() ([]byte, error) {
	d := float64(m) * 0.01
	str := fmt.Sprintf("%.2f", d)
	return []byte(str), nil
}

