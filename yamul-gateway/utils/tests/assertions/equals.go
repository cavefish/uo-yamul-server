package assertions

import "fmt"

func (f assertionsFor) Equals(expected any, actual any) {
	if expected != actual {
		f.t.Error(fmt.Sprintf("expected:\t%v\tactual:\t%v", expected, actual))
	}
}
