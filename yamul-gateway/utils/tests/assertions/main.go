package assertions

import "testing"

type assertionsFor struct {
	t *testing.T
}

func For(t *testing.T) assertionsFor {
	return assertionsFor{t: t}
}
