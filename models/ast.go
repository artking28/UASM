package models

type (
	Pos struct {
		Line   int64
		Column int64
	}

	Ast struct {
		Statements []Stmt
	}

	Stmt interface {
		WriteMemASM() []byte
	}

	CommentStmt struct {
		Pos   Pos    `json:"pos"`
		Value string `json:"value"`
	}

	PureInstructionStmt struct {
		Pos  Pos           `json:"pos"`
		Code TokenKindEnum `json:"code"`
	}

	SingleInstructionStmt struct {
		PureInstructionStmt
		Left Token
	}

	DoubleInstructionStmt struct {
		SingleInstructionStmt
		Right Token
	}
)

func NewPureInstructionStmt(code TokenKindEnum, pos Pos) PureInstructionStmt {
	if code < TOKEN_GET {

	}
	return PureInstructionStmt{
		Code: code,
		Pos:  pos,
	}
}

func NewSingleInstructionStmt(code TokenKindEnum, pos Pos, left Token) SingleInstructionStmt {
	return SingleInstructionStmt{
		PureInstructionStmt: NewPureInstructionStmt(code, pos),
		Left:                left,
	}
}

func NewDoubleInstructionStmt(code TokenKindEnum, pos Pos, left, right Token) DoubleInstructionStmt {
	return DoubleInstructionStmt{
		SingleInstructionStmt: NewSingleInstructionStmt(code, pos, left),
		Right:                 right,
	}
}

func (this CommentStmt) WriteMemASM() []byte {
	return []byte{}
}
