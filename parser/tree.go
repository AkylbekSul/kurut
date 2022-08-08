package parser

type (
	// TODO: understand wtf is this
	SyntaxTree struct {
		Scopes []*Scope
	}

	StatementType int

	Statement struct {
		Type  StatementType
		Scope *Scope
	}

	Expression[T any] struct {
		Left  T
		Right T
		Op    string
	}

	Identity[T any] struct {
		Name  string
		Value T
	}

	Scope struct {
		Identities []*Identity[any]
		Children   []*Scope
		Statements []*Statement
	}
)

const (
	IF StatementType = iota
	FOR
	PRINT // TODO: remove this statement for something like SYSCALL
)
