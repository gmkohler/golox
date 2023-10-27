package golox

type Scanner interface {
	ScanTokens() ([]Token, error)
}
