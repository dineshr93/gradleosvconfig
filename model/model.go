package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type OSVData struct {
	Results []Results `json:"results,omitempty"`
}
type Source struct {
	Path string `json:"path,omitempty"`
	Type string `json:"type,omitempty"`
}
type Package struct {
	Name      string `json:"name,omitempty"`
	Version   string `json:"version,omitempty"`
	Ecosystem string `json:"ecosystem,omitempty"`
	Purl      string `json:"purl,omitempty"`
}

type Events struct {
	Introduced string `json:"introduced,omitempty"`
	Fixed      string `json:"fixed,omitempty"`
}
type Ranges struct {
	Type   string   `json:"type,omitempty"`
	Events []Events `json:"events,omitempty"`
}

type Affected struct {
	Package          Package          `json:"package,omitempty"`
	Ranges           []Ranges         `json:"ranges,omitempty"`
	Versions         []string         `json:"versions,omitempty"`
	DatabaseSpecific DatabaseSpecific `json:"database_specific,omitempty"`
}
type References struct {
	Type string `json:"type,omitempty"`
	URL  string `json:"url,omitempty"`
}
type DatabaseSpecific struct {
	CweIds                        []string  `json:"cwe_ids,omitempty"`
	GithubReviewed                bool      `json:"github_reviewed,omitempty"`
	GithubReviewedAt              time.Time `json:"github_reviewed_at,omitempty"`
	NvdPublishedAt                time.Time `json:"nvd_published_at,omitempty"`
	Severity                      string    `json:"severity,omitempty"`
	LastKnownAffectedVersionRange string    `json:"last_known_affected_version_range,omitempty"`
	Source                        string    `json:"source,omitempty"`
}
type Vulnerabilities struct {
	SchemaVersion    string           `json:"schema_version,omitempty"`
	ID               string           `json:"id,omitempty"`
	Modified         time.Time        `json:"modified,omitempty"`
	Published        time.Time        `json:"published,omitempty"`
	Aliases          []string         `json:"aliases,omitempty"`
	Summary          string           `json:"summary,omitempty"`
	Details          string           `json:"details,omitempty"`
	Affected         []Affected       `json:"affected,omitempty"`
	References       []References     `json:"references,omitempty"`
	DatabaseSpecific DatabaseSpecific `json:"database_specific,omitempty"`
}
type Groups struct {
	Ids []string `json:"ids,omitempty"`
}
type Packages struct {
	Package         Package           `json:"package,omitempty"`
	Vulnerabilities []Vulnerabilities `json:"vulnerabilities,omitempty"`
	Groups          []Groups          `json:"groups,omitempty"`
}
type Results struct {
	Source   Source     `json:"source,omitempty"`
	Packages []Packages `json:"packages,omitempty"`
}

func (t *OSVData) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}
func (s *OSVData) PrintVuls() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "pkg name"},
			{Align: simpletable.AlignCenter, Text: "version"},
			{Align: simpletable.AlignCenter, Text: "fixed"},
			{Align: simpletable.AlignCenter, Text: "severity"},
			{Align: simpletable.AlignCenter, Text: "vuln source"},
		},
	}
	// {Align: simpletable.AlignCenter, Text: "Checksums"},
	// {Align: simpletable.AlignCenter, Text: "Algorithm - Checksums"},
	var pn, pv, fixed, severity, source string
	var len_pkgs, len_vulns, total_vulns int
	var cells [][]*simpletable.Cell
	for i, v := range s.Results {
		_ = i
		len_pkgs = len(v.Packages)
		for n_pkg, pkg := range v.Packages {
			// _ = n_pkg

			pn = pkg.Package.Name
			pv = pkg.Package.Version

			len_vulns = len(pkg.Vulnerabilities)
			total_vulns += len_vulns
			for n_vuln, vuln := range pkg.Vulnerabilities {
				_ = n_vuln
				for n_aff, aff := range vuln.Affected {
					_ = n_aff
					for n_rngs, rngs := range aff.Ranges {
						_ = n_rngs

						for n_evnts, evnts := range rngs.Events {
							_ = n_evnts
							fixed = evnts.Fixed
						}
					}
					source = aff.DatabaseSpecific.Source

				}
				severity = vuln.DatabaseSpecific.Severity
				cells = append(cells, *&[]*simpletable.Cell{
					{Text: fmt.Sprintf("%d", n_pkg)},
					{Text: pn},
					{Text: pv},
					{Text: fixed},
					{Text: severity},
					{Text: source},
				})
			}

		}

	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 6, Text: blue(fmt.Sprintf("There are %d pkgs with %d vulnerabilities", len_pkgs, total_vulns))},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()

}
