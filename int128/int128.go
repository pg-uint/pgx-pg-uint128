package int128

import (
	"fmt"
	"go.shabbyrobe.org/num"
)

func FromString(str string) (num.I128, error) {
	res, isAccurate, err := num.I128FromString(str)
	if err != nil {
		return num.I128{}, err
	}

	if !isAccurate {
		return num.I128{}, fmt.Errorf("int128.FromString: parsing %q: value out of range", str)
	}

	return res, nil
}
