package pageaction

import (
	"fmt"
)

type Setup struct {
	PageAction
}

func (s *Setup) PostAction(args ...string) {
	s.NextPage = "suggestion"
	fmt.Println(s.NextPage)
}
