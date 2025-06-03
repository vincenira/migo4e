package pkg2

import (
	"fmt"

	"github.com/vincenira/migo4e/expery/go-workspaces/md1/pckg2"
)

func AdderMrMrs(name string, char byte) string {
	smartName := ""
	switch char {
	case 'f':
		smartName = fmt.Sprintf("Mrs %s", name)
	case 'm':
		smartName = fmt.Sprintf("Mr %s", name)
	default:
		smartName = fmt.Sprintf("Mr/Mrs %s", name)
	}
	return pckg2.Greeter(smartName)
}
