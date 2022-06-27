package response

import (
	"fmt"
	"time"
)

type DateTime time.Time

func (dt DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(dt)
	if t.IsZero() {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}