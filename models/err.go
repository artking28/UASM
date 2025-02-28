package models

import "fmt"

type (
	ErrCode uint

	ErrLabel string

	Err struct {
		Code  ErrCode
		Label ErrLabel
		Msg   string
	}
)

func (this *Err) Error() string {
	return fmt.Sprintf("error %.4d | %s\n", this.Code, this.Msg)
}

const (
	NilPointerExceptionErrCode ErrCode = iota
	FileNotFoundErrCode
	EmptyFileErrCode
	DivideByZeroErrCode
	InvalidArgumentErrCode
	UnexpectedTokenErrCode
	InvalidMinmonicErrCode
	UnkownFailureErrCode

	NilPointerExceptionErrLabel ErrLabel = "error.nil.pointer"
	FileNotFoundErrLabel        ErrLabel = "error.file.not.found"
	EmptyFileErrLabel           ErrLabel = "error.empty.file"
	DivideByZeroErrLabel        ErrLabel = "error.divide.by.zero"
	InvalidArgumentErrLabel     ErrLabel = "error.invalid.argument"
	UnexpectedTokenErrLabel     ErrLabel = "error.unexpected.token"
	InvalidMinmonicErrLabel     ErrLabel = "error.invalid.minmonic"
	UnkownFailureErrLabel       ErrLabel = "error.unkown.failure"
)

func NewNilPointerExceptionErr() Err {
	return Err{
		Code:  NilPointerExceptionErrCode,
		Label: NilPointerExceptionErrLabel,
		Msg:   "",
	}
}

func NewDivideByZeroErr() Err {
	return Err{
		Code:  DivideByZeroErrCode,
		Label: DivideByZeroErrLabel,
		Msg:   "",
	}
}

func NewFileNotFoundErr() Err {
	return Err{
		Code:  FileNotFoundErrCode,
		Label: FileNotFoundErrLabel,
		Msg:   "",
	}
}

func NewInvalidArgumentErr() Err {
	return Err{
		Code:  InvalidArgumentErrCode,
		Label: InvalidArgumentErrLabel,
		Msg:   "",
	}
}

func NewEmptyFileErr() Err {
	return Err{
		Code:  EmptyFileErrCode,
		Label: EmptyFileErrLabel,
		Msg:   "",
	}
}

func NewUnexpectedTokenErr() Err {
	return Err{
		Code:  UnexpectedTokenErrCode,
		Label: UnexpectedTokenErrLabel,
		Msg:   "",
	}
}

func NewUnkownErr() Err {
	return Err{
		Code:  UnkownFailureErrCode,
		Label: UnkownFailureErrLabel,
		Msg:   "",
	}
}
