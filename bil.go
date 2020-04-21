package bil

import (
	"fmt"
	"regexp"
	"strings"
)

var bil *Bil

type Bil struct {
	strict bool
}

func init() {
	bil = New()
}

func New() *Bil {
	return &Bil{strict: true}
}

func SetStrict(strict bool) { bil.SetStrict(strict) }
func (bil *Bil) SetStrict(strict bool) {
	bil.strict = strict
}

func Eval(text string, vars map[string]string) (string, error) { return bil.Eval(text, vars) }
func (bil *Bil) Eval(text string, vars map[string]string) (string, error) {
	for k, v := range vars {
		search := fmt.Sprintf("${%s}", k)

		if bil.strict && !strings.Contains(text, search) {
			return "", fmt.Errorf("variable %s doesn't exists", k)
		}

		text = strings.Replace(text, search, v, -1)
	}

	re := regexp.MustCompile(`\${.*?}`)
	if unfilled := re.FindAll([]byte(text), -1); len(unfilled) > 0 {
		return "", fmt.Errorf("unfilled variable is remaining")
	}

	return text, nil
}
