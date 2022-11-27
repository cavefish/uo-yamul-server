package assertions

import "fmt"

func (f assertionsFor) Equals(expected any, actual any) {
	if expected != actual {
		f.t.Error(fmt.Sprintf("expected:\t%v\tactual:\t%v", expected, actual))
	}
}

func (f assertionsFor) EqualList(expected []byte, actual []byte) {
	le := len(expected)
	la := len(actual)
	if le != la {
		f.t.Error(fmt.Sprintf("actual is %d bigger than expected", la-le))
	}
	for i := 0; i < le; i++ {
		if expected[i] != actual[i] {
			f.t.Error(fmt.Sprintf("idx:\t%d\texpected:\t%v\tactual:\t%v", i, expected[i], actual[i]))
		}
	}
}
