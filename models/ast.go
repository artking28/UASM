package models

import (
	"encoding/json"
	"fmt"
)

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
		Value string `json:"value"`
		StmtBase
	}

	LabelDeclStmt struct {
		LabelName string `json:"labelName"`
		StmtBase
	}

	JumpStmt struct {
		TargetLabelName string `json:"TargetLabelName"`
		JumpKind        string `json:"jumpKind"`
		StmtBase
	}

	PureInstructionStmt struct {
		Code TokenKindEnum `json:"code"`
		StmtBase
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

	fmt.Printf("%s\n", string(str))
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

func NewLabelDeclStmt(labelName string, pos Pos) LabelDeclStmt {
	return LabelDeclStmt{
		LabelName: labelName,
		StmtBase: StmtBase{
			Title: "LabelDeclStmt",
			Pos:   pos,
		},
	}
}

func NewJumpStmt(targetLabelName, jumpKind string, pos Pos) JumpStmt {
	return JumpStmt{
		TargetLabelName: targetLabelName,
		JumpKind:        jumpKind,
		StmtBase: StmtBase{
			Title: "JumpStmt",
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

func (this LabelDeclStmt) WriteMemASM() []byte {
	//TODO implement me
	panic("implement me")
	return []byte{}
}

func (j JumpStmt) WriteMemASM() []byte {
	//TODO implement me
	panic("implement me")
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
