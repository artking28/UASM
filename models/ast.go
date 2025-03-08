package models

import (
	"UASM/neander"
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
		WriteMemASM() []uint16
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

func (this SingleInstructionStmt) GetLeftASUint16() uint16 {
	return uint16(this.Left.Value[0])
}

func NewDoubleInstructionStmt(code TokenKindEnum, pos Pos, left, right Token) DoubleInstructionStmt {
	d := DoubleInstructionStmt{
		SingleInstructionStmt: NewSingleInstructionStmt(code, pos, left),
		Right:                 right,
	}
	d.StmtBase.Title = "DoubleInstructionStmt"
	return d
}

func (this DoubleInstructionStmt) GetRightASUint16() uint16 {
	return uint16(this.Left.Value[0])
}

func (this CommentStmt) WriteMemASM() []uint16 {
	return []uint16{}
}

func (this LabelDeclStmt) WriteMemASM() []uint16 {
	//TODO implement me
	panic("implement me LabelDeclStmt WriteMemASM")
	return []uint16{}
}

func (this JumpStmt) WriteMemASM() []uint16 {
	//TODO implement me
	panic("implement me JumpStmt WriteMemASM")
	return []uint16{}
}

func (this PureInstructionStmt) WriteMemASM() (ret []uint16) {
	switch this.Code {
	case TOKEN_INC:
		ret = append(ret, neander.ADD, OneValue)
		break
	case TOKEN_DEC:
		ret = append(ret, neander.ADD, MinusOneValue)
		break
	case TOKEN_NEG:
		ret = append(ret, neander.NOT, neander.ADD, OneValue)
		break
	case TOKEN_NOT:
		ret = append(ret, neander.NOT)
		break
	case TOKEN_HLT:
		ret = append(ret, neander.HLT)
		break
	default:
		//TODO implement me
		panic("implement me switch default branch in PureInstructionStmt WriteMemASM implementation")
	}
	return ret
}

func (this SingleInstructionStmt) WriteMemASM() (ret []uint16) {
	switch this.Code {
	case TOKEN_GET:
		if this.Left.Kind == TOKEN_NUMBER {
			ret = append(ret, neander.LDA, this.GetLeftASUint16())
			break
		}
		ret = append(ret, neander.LDA, this.GetLeftASUint16())
		break
	case TOKEN_SET:
		//TODO implement me
		panic("implement me, TOKEN_SET materialize")
		break
	case TOKEN_MUL:
		leftArg := this.GetLeftASUint16()
		ret = append(ret, GetBuiltinMulFunc(uint16(len(ret)), leftArg)...)
		break
	case TOKEN_ADD:
		ret = append(ret, neander.ADD, this.GetLeftASUint16())
		break
	case TOKEN_AND:
		ret = append(ret, neander.AND, this.GetLeftASUint16())
		break
	case TOKEN_ORR:
		ret = append(ret, neander.OR, this.GetLeftASUint16())
		break
	case TOKEN_XOR:
		//ret = append(ret, neander.XOR, this.GetLeftASUint16())
		break
	case TOKEN_CMP:
		//TODO implement me
		panic("implement me, TOKEN_CMP materialize")
	default:
		//TODO implement me
		panic("implement me switch default branch in PureInstructionStmt WriteMemASM implementation")
	}
	return ret
}

func (this DoubleInstructionStmt) WriteMemASM() (ret []uint16) {
	//TODO implement me
	panic("implement me")
	return ret
}
