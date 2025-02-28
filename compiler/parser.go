package compiler

import (
	"UASM/models"
)

func ParseAll(parser *models.Parser) error {

	tk := parser.Get(0)
	for ; tk != nil && tk.Kind != models.TOKEN_EOF; tk = parser.Get(0) {
		switch tk.Kind {
		case models.TOKEN_SLASH:
			if parser.Get(1).Kind == models.TOKEN_SLASH {
				err := ParseComment(parser)
				if err != nil {
					return err
				}
			}
			break
		case models.TOKEN_BREAK_LINE:
			parser.NextLine()
			parser.Consume(1)
			break
		default:
			break
		}
		parser.Consume(1)
	}
	return nil
}

func ParseComment(parser *models.Parser) error {
	// var comment string
	// h0 := parser.Get(0)
	// h1 := parser.Get(1)
	// if h0 != nil && h1 != nil && h0.Kind == TOKEN_SLASH && h1.Kind == TOKEN_SLASH {
	// 	parser.Consume(2)
	// } else {
	// 	return errors.New(models.UnexpectexTokenErrLabel)
	// }
	// for here.Kind != models.TOKEN_BREAK_LINE {
	// 	comment += string(here.Value)	
	// 	parser.Consume(1)
	// }
	// parser.NextLine()
	// parser.Consume(1)
	// parser.Inject(models.CommentStmt{
	// 	Pos:   parser.Where(),
	// 	Value: comment,
	// })
	return nil
}

func ParsePureInstruction(this *models.Parser) error {
	// h0 := parser.Get(0)
	// if h0 == nil {
	// 	return nil
	// }
	// switch h0.kind {
	// case TOKEN_INC, TOKEN_DEC, TOKEN_NEG, TOKEN_NOT, TOKEN_HLT:
	// 	parser.Inject(models.NewPureInstructionStmt(h0.kind, parser.Where()))
	// 	break
	// default:
	// 	return errors.New(models.UnexpectedTokenErrLabel)
	// }
	return nil
}

func ParseSingleInstruction(this *models.Parser) error {
	return nil
}

func ParseDoubleInstruction(this *models.Parser) error {
	return nil
}