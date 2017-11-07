package diff

// Solver is an interface implemented by the diff algorithms.
type Solver interface {
	Solve() *DiffSolution
}

var (
	_ Solver = &HistogramDiffer{}
	_ Solver = &SequenceDiffer{}
)

func DiffToHTML(a, b string) (string, error) {
	buf, err := HistogramDiff(a, b).HTML()
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func DiffToText(a, b string) string {
	return HistogramDiff(a, b).Text()
}
