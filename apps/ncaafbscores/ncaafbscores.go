// Package ncaafbscores provides details for the NCAA Football Scores applet.
package ncaafbscores

import (
	_ "embed"

	"tidbyt.dev/community/apps/manifest"
)

//go:embed ncaafb_scores.star
var source []byte

// New creates a new instance of the NCAA Football Scores applet.
func New() manifest.Manifest {
	return manifest.Manifest{
		ID:          "ncaafb-scores",
		Name:        "NCAA Football Scores",
		Author:      "LunchBox8484 & CaseyErbes",
		Summary:     "NCAA football scores",
		Desc:        "For slower scrolling of scores, add the app to your Tidbyt multiple times. Then, within each app instance, set 'Total Instances of App' to the amount of times you have it installed, and set 'App Instance Number' unique to each app instance.",
		FileName:    "ncaafb_scores.star",
		PackageName: "ncaafbscores",
		Source:  source,
	}
}
