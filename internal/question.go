package ps

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ronaudinho/ps/pkg/util"
)

type QuestionParser struct {
	ans   []string
	def   string
	exp   string
	regex *regexp.Regexp
	mp    map[string]Parser
	// warn []error // maybe can add warnings to show why question fallbacks to default answer
}

func NewQuestionParser(def, exp string, mp map[string]Parser) *QuestionParser {
	return &QuestionParser{
		def:   def,
		exp:   exp,
		regex: regexp.MustCompile(exp),
		mp:    mp,
	}
}

func (qp *QuestionParser) Parse(line string, store Store) error {
	var ans string
	for _, p := range qp.mp {
		reg := regexp.MustCompile(p.regexp())
		if reg.MatchString(line) {
			err := p.Parse(line, store)
			if err != nil {
				ans = qp.def
				continue
			}
			ans = p.Strout()
		}
	}
	if ans == "" {
		ans = qp.def
	}
	qp.ans = append(qp.ans, ans)
	return nil
}

func (qp *QuestionParser) regexp() string {
	return qp.exp
}

func (qp *QuestionParser) Strout() string {
	return strings.Join(qp.ans, "\n")
}

type HowManyQuestion struct {
	Intergalactic string `key:"intergalactic"`
	Item          string `key:"item"`
}

type HowManyQuestionParser struct {
	ans   string
	exp   string
	regex *regexp.Regexp
}

func NewHowManyQuestionParser(exp string) *HowManyQuestionParser {
	return &HowManyQuestionParser{
		exp:   exp,
		regex: regexp.MustCompile(exp),
	}
}

func (qp *HowManyQuestionParser) Parse(line string, store Store) error {
	m := util.MapRegex(line, qp.regex)
	var q HowManyQuestion
	util.Unmarshal(m, &q)
	u, err := store.Get([]string{"unit"}...)
	if err != nil {
		return err
	}
	uu, _ := u.(map[string]interface{})
	um := make(map[string]string, len(uu))
	for k, v := range uu {
		vv, _ := v.(string)
		um[k] = vv
	}
	p, err := store.Get([]string{"price", q.Item}...)
	if err != nil {
		return err
	}
	pp, _ := p.(float32)
	a, err := util.IntergalacticToInt(q.Intergalactic, " ", um)
	if err != nil {
		return err
	}
	qp.ans = fmt.Sprintf("%s %s is %.f Credits", q.Intergalactic, q.Item, float32(a)*pp)
	return nil
}

func (qp *HowManyQuestionParser) regexp() string {
	return qp.exp
}

func (qp *HowManyQuestionParser) Strout() string {
	return qp.ans
}

type HowMuchQuestion struct {
	Intergalactic string `key:"intergalactic"`
	Item          string `key:"item"`
}

type HowMuchQuestionParser struct {
	ans   string
	exp   string
	regex *regexp.Regexp
}

func NewHowMuchQuestionParser(exp string) *HowMuchQuestionParser {
	return &HowMuchQuestionParser{
		exp:   exp,
		regex: regexp.MustCompile(exp),
	}
}

func (qp *HowMuchQuestionParser) Parse(line string, store Store) error {
	m := util.MapRegex(line, qp.regex)
	var q HowMuchQuestion
	util.Unmarshal(m, &q)
	u, err := store.Get([]string{"unit"}...)
	if err != nil {
		return err
	}
	uu, _ := u.(map[string]interface{})
	um := make(map[string]string, len(uu))
	for k, v := range uu {
		vv, _ := v.(string)
		um[k] = vv
	}
	a, err := util.IntergalacticToInt(q.Intergalactic, " ", um)
	if err != nil {
		return err
	}
	qp.ans = fmt.Sprintf("%s is %d", q.Intergalactic, a)
	return nil
}

func (qp *HowMuchQuestionParser) regexp() string {
	return qp.exp
}

func (qp *HowMuchQuestionParser) Strout() string {
	return qp.ans
}
