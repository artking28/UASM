package compiler

import (
	"UASM/models"
)

func ParseAll(parser *models.Parser) {

	tk := parser.Get(0)
	for ; tk != nil && tk.Kind != models.TOKEN_EOF; tk = parser.Get(0) {
		switch tk.Kind {
		case models.TOKEN_SLASH:
			if parser.Get(1).Kind == models.TOKEN_SLASH {
				parser.Consume(2)
				ParseComment(parser)
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
}

func ParseComment(parser *models.Parser) {
	var comment string
	here := parser.Get(0)
	for here.Kind != models.TOKEN_BREAK_LINE {
		comment += string(here.Value)
		parser.Consume(1)
	}
	parser.Inject(models.CommentStmt{
		Pos:   parser.Where(),
		Value: comment,
	})
}

func ParseInstruction(this *models.Parser) {
}
