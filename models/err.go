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

func (this Err) Error() string {
	return fmt.Sprintf("error %.4d | %s\n", this.Code, this.Msg)
}

const (
	NilPointerExceptionErrCode ErrCode = iota
	FileNotFoundErrCode
	EmptyFileErrCode
	DivideByZeroErrCode
	InvalidArgumentErrCode
	UnexpectedTokenErrCode
	ExpectedTokenErrCode
	InvalidMinmonicErrCode
	UnkownFailureErrCode

	NilPointerExceptionErrLabel ErrLabel = "error.nil.pointer"
	FileNotFoundErrLabel        ErrLabel = "error.file.not.found"
	EmptyFileErrLabel           ErrLabel = "error.empty.file"
	DivideByZeroErrLabel        ErrLabel = "error.divide.by.zero"
	InvalidArgumentErrLabel     ErrLabel = "error.invalid.argument"
	UnexpectedTokenErrLabel     ErrLabel = "error.unexpected.token"
	ExpectedTokenErrLabel       ErrLabel = "error.expected.token"
	InvalidMinmonicErrLabel     ErrLabel = "error.invalid.minmonic"
	UnkownFailureErrLabel       ErrLabel = "error.unkown.failure"
)

func GetNilPointerExceptionErr() Err {
	return Err{
		Code:  NilPointerExceptionErrCode,
		Label: NilPointerExceptionErrLabel,
		Msg:   "",
	}
}

func GetDivideByZeroErr() Err {
	return Err{
		Code:  DivideByZeroErrCode,
		Label: DivideByZeroErrLabel,
		Msg:   "",
	}
}

func GetFileNotFoundErr() Err {
	return Err{
		Code:  FileNotFoundErrCode,
		Label: FileNotFoundErrLabel,
		Msg:   "",
	}
}

func GetInvalidArgumentErr() Err {
	return Err{
		Code:  InvalidArgumentErrCode,
		Label: InvalidArgumentErrLabel,
		Msg:   "",
	}
}

func GetEmptyFileErr(filename string) Err {
	return Err{
		Code:  EmptyFileErrCode,
		Label: EmptyFileErrLabel,
		Msg:   fmt.Sprintf("The file '%s' is empty.", filename),
	}
}

func EexpectedTokenErr(filename string, line int) Err {
	return Err{
		Code:  UnexpectedTokenErrCode,
		Label: UnexpectedTokenErrLabel,
		Msg:   fmt.Sprintf("Missing token in the %dº line of the file '%s'.", line, filename),
	}
}

func GetUnexpectedTokenErr(filename string, char rune, line int) Err {
	return Err{
		Code:  UnexpectedTokenErrCode,
		Label: UnexpectedTokenErrLabel,
		Msg:   fmt.Sprintf("Unexpected token '%s' in the %dº line of the file '%s'.", string(char), line, filename),
	}
}

func GetUnkownErr() Err {
	return Err{
		Code:  UnkownFailureErrCode,
		Label: UnkownFailureErrLabel,
		Msg:   "",
	}
}
