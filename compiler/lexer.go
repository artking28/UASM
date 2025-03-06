package compiler

import (
	"UASM/models"
	"os"
	"strings"
)

func Tokenize(filename string) ([]models.Token, error) {

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(bytes) < 0 {
		return nil, models.GetEmptyFileErr(filename)
	}

	var ret []models.Token
	column, line := 0, 1
	isComment, runes := false, []rune(string(bytes))
	buffer := strings.Builder{}
	for i, run := range runes {
		column++
		if 'a' <= run && run <= 'z' ||
			'A' <= run && run <= 'Z' ||
			'0' <= run && run <= '9' {
			buffer.WriteRune(run)
			continue
		}
		if buffer.Len() > 0 {
			var tk models.Token
			pos := models.Pos{Line: int64(line), Column: int64(column - buffer.Len())}
			tk = models.NewToken(pos, models.TOKEN_ID, 1, []rune(buffer.String())...)
			models.AppendToken(&ret, models.ResolveTokenId(tk))
			buffer.Reset()
		}
		switch run {
		case '\n':
			line += 1
			column = 0
			isComment = false
			pos := models.Pos{Line: int64(line), Column: int64(column)}
			tk := models.NewToken(pos, models.TOKEN_BREAK_LINE, 1, run)
			models.AppendToken(&ret, tk)
			break
		case '\t':
			pos := models.Pos{Line: int64(line), Column: int64(column)}
			tk := models.NewToken(pos, models.TOKEN_TAB, 1, run)
			models.AppendToken(&ret, tk)
			break
		case ' ':
			pos := models.Pos{Line: int64(line), Column: int64(column)}
			tk := models.NewToken(pos, models.TOKEN_SPACE, 1, run)
			models.AppendToken(&ret, tk)
			break
		case ',':
			pos := models.Pos{Line: int64(line), Column: int64(column)}
			tk := models.NewToken(pos, models.TOKEN_COMMA, 1, run)
			models.AppendToken(&ret, tk)
			break
		case '/':
			pos := models.Pos{Line: int64(line), Column: int64(column)}
			tk := models.NewToken(pos, models.TOKEN_SLASH, 1, run)
			if runes[i+1] == '/' {
				isComment = true
				i += 1
				column++
				models.AppendToken(&ret, tk)
				break
			}
			models.AppendToken(&ret, tk)
			break
		default:
			if isComment {
				buffer.WriteRune(run)
				continue
			}
			return nil, models.GetUnexpectedTokenErr(filename, run, line)
		}
	}

	pos := models.Pos{Line: int64(line), Column: int64(column)}
	tk := models.NewToken(pos, models.TOKEN_EOF, 1, '0')
	models.AppendToken(&ret, tk)
	return ret, nil
}
