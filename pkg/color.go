package pkg

import (
	"regexp"

	"github.com/fatih/color"
)

func customWriter(w string) *color.Color {

	var (
		ErrorColor   = color.New(color.FgRed)
		SuccessColor = color.New(color.FgGreen)
		NormalColor  = color.New(color.FgBlue)
	)

	// Regexp Conditions
	match_error, error := regexp.MatchString(".+FAILED|.+ROLLBACK|^ROLLBACK", w)
	errorHandle(error)
	match_succes, error := regexp.MatchString("CREATE_COMPLETE|IMPORT_COMPLETE|UPDATE_COMPLETE|DELETE_COMPLETE", w)
	errorHandle(error)

	if match_error {
		return ErrorColor
	} else if match_succes {
		return SuccessColor
	} else {
		return NormalColor
	}
}
