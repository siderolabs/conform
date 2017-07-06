package utilities

import (
	"fmt"
	"os"
	"strings"
)

// ExportConformVar exports variable prefixed with CONFORM_
func ExportConformVar(name, value string) (err error) {
	variable := fmt.Sprintf("CONFORM_%s", strings.ToUpper(name))
	err = os.Setenv(variable, value)

	return
}
