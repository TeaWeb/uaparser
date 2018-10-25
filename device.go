package uaparser

type Device struct {
	Family string `json:"family"`
	Brand  string `json:"brand"`
	Model  string `json:"model"`
}

func (this *Device) Parse(matches []string, spec map[string]string) {
	countMatches := len(matches)

	{
		replacement, ok := spec["device_replacement"]
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
		replacement, ok := spec["brand_replacement"]
		index := 2
		if ok {
			this.Brand = replacementReg.ReplaceAllStringFunc(replacement, func(s string) string {
				index := toInt(s[1:])
				if index >= 0 && index < countMatches {
					return matches[index]
				}
				return s
			})
		} else {
			if countMatches > index {
				this.Brand = matches[index]
			}
		}
	}

	// v2
	{
		replacement, ok := spec["model_replacement"]
		index := 3
		if ok {
			this.Model = replacementReg.ReplaceAllStringFunc(replacement, func(s string) string {
				index := toInt(s[1:])
				if index >= 0 && index < countMatches {
					return matches[index]
				}
				return s
			})
		} else {
			if countMatches > index {
				this.Model = matches[index]
			}
		}
	}
}
