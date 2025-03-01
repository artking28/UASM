package models

import "encoding/json"

type (
	Pos struct {
		Line   int64 `json:"line"`
		Column int64 `json:"column"`
	}

	Ast struct {
		Statements []Stmt `json:"statements"`
	}

	Stmt interface {
		WriteMemASM() []byte
	}

	StmtBase struct {
		Title string `json:"title"`
		Pos   Pos    `json:"pos"`
	}

	CommentStmt struct {
		StmtBase
		Value string `json:"value"`
	}

	PureInstructionStmt struct {
		StmtBase
		Code TokenKindEnum `json:"code"`
	}

	SingleInstructionStmt struct {
		PureInstructionStmt
		Left Token `json:"left"`
	}

	DoubleInstructionStmt struct {
		SingleInstructionStmt
		Right Token `json:"right"`
	}
)

func (this Ast) Inspect() {
	str, err := json.MarshalIndent(this, "", "   ")
	if err != nil {
		panic(err.Error())
	}

	println(string(str))
}

func NewCommentStmt(content string, pos Pos) CommentStmt {
	return CommentStmt{
		Value: content,
		StmtBase: StmtBase{
			Title: "CommentStmt",
			Pos:   pos,
		},
	}
}

func NewPureInstructionStmt(code TokenKindEnum, pos Pos) PureInstructionStmt {
	return PureInstructionStmt{
		Code: code,
		StmtBase: StmtBase{
			Title: "PureInstructionStmt",
			Pos:   pos,
		},
	}
}

func NewSingleInstructionStmt(code TokenKindEnum, pos Pos, left Token) SingleInstructionStmt {
	s := SingleInstructionStmt{
		PureInstructionStmt: NewPureInstructionStmt(code, pos),
		Left:                left,
	}
	s.StmtBase.Title = "SingleInstructionStmt"
	return s
}

func NewDoubleInstructionStmt(code TokenKindEnum, pos Pos, left, right Token) DoubleInstructionStmt {
	d := DoubleInstructionStmt{
		SingleInstructionStmt: NewSingleInstructionStmt(code, pos, left),
		Right:                 right,
	}
	d.StmtBase.Title = "DoubleInstructionStmt"
	return d
}

func (this CommentStmt) WriteMemASM() []byte {
	return []byte{}
}

func (this PureInstructionStmt) WriteMemASM() []byte {
	//TODO implement me
	panic("implement me")
	return []byte{}
}

func (this SingleInstructionStmt) WriteMemASM() []byte {
	//TODO implement me
	panic("implement me")
	return []byte{}
}

func (this DoubleInstructionStmt) WriteMemASM() []byte {
	//TODO implement me
	panic("implement me")
	return []byte{}
}
