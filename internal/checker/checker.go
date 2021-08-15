package checker

const (
	ForContainer = 0
	ForService   = 1
)

type Checker interface {
	AssertRunning() error
	GetTitle() string
}
