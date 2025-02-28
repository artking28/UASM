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
	line, isComment, runes := 1, false, []rune(string(bytes))
	buffer := strings.Builder{}
	for i, run := range runes {

		if 'a' <= run && run <= 'z' ||
			'A' <= run && run <= 'Z' ||
			'0' <= run && run <= '9' {
			buffer.WriteRune(run)
			continue
		}
		tk := models.NewToken(models.TOKEN_ID, 1, []rune(buffer.String())...)
		models.AppendToken(&ret, tk)

		switch run {
		case '\n':
			line += 1
			isComment = false
			tk := models.NewToken(models.TOKEN_BREAK_LINE, 1, run)
			models.AppendToken(&ret, tk)
			break
		case '\t':
			tk := models.NewToken(models.TOKEN_TAB, 1, run)
			models.AppendToken(&ret, tk)
			break
		case ' ':
			tk := models.NewToken(models.TOKEN_SPACE, 1, run)
			models.AppendToken(&ret, tk)
			break
		case ',':
			tk := models.NewToken(models.TOKEN_COMMA, 1, run)
			models.AppendToken(&ret, tk)
			break
		case '/':
			tk := models.NewToken(models.TOKEN_SLASH, 1, run)
			if runes[i+1] == '/' {
				isComment = true
				models.AppendToken(&ret, tk)
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

	return ret, nil
}
