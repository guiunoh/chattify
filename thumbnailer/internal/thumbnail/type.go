package thumbnail

import "fmt"

type Rate int

const (
	_ Rate = iota
	RATE30
	RATE50
	RATE70
)

func StringToRate(s string) (Rate, error) {
	switch s {
	case "RATE30":
		return RATE30, nil
	case "RATE50":
		return RATE50, nil
	case "RATE70":
		return RATE70, nil
	}
	return 0, fmt.Errorf("unknown rate: %s", s)
}

func (r Rate) Value() float32 {
	switch r {
	case RATE30:
		return 0.3
	case RATE50:
		return 0.5
	case RATE70:
		return 0.7
	}

	return 1.0
}
