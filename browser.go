package uaparser

type Browser struct {
	Family string `json:"family"`
	Major  string `json:"major"`
	Minor  string `json:"minor"`
	Patch  string `json:"patch"`
}

func (this *Browser) Parse(matches []string, spec map[string]string) {
	countMatches := len(matches)

	{
		replacement, ok := spec["family_replacement"]
		index := 1
		if ok {
			this.Family = replacementReg.ReplaceAllStringFunc(replacement, func(s string) string {
				index := toInt(s[1:])
				if index >= 0 && index < countMatches {
					return matches[index]
				}
				return s
			})
		} else {
			if countMatches > index {
				this.Family = matches[index]
			}
		}
	}

	// v1
	{
		replacement, ok := spec["v1_replacement"]
		index := 2
		if ok {
			this.Major = replacementReg.ReplaceAllStringFunc(replacement, func(s string) string {
				index := toInt(s[1:])
				if index >= 0 && index < countMatches {
					return matches[index]
				}
				return s
			})
		} else {
			if countMatches > index {
				this.Major = matches[index]
			}
		}
	}

	// v2
	{
		replacement, ok := spec["v2_replacement"]
		index := 3
		if ok {
			this.Minor = replacementReg.ReplaceAllStringFunc(replacement, func(s string) string {
				index := toInt(s[1:])
				if index >= 0 && index < countMatches {
					return matches[index]
				}
				return s
			})
		} else {
			if countMatches > index {
				this.Minor = matches[index]
			}
		}
	}

	// v3
	{
		replacement, ok := spec["v3_replacement"]
		index := 4
		if ok {
			this.Patch = replacementReg.ReplaceAllStringFunc(replacement, func(s string) string {
				index := toInt(s[1:])
				if index >= 0 && index < countMatches {
					return matches[index]
				}
				return s
			})
		} else {
			if countMatches > index {
				this.Patch = matches[index]
			}
		}
	}
}
