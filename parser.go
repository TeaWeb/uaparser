package uaparser

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

var replacementReg = regexp.MustCompile(`\$\d`)
var keywordReg = regexp.MustCompile(`[a-zA-Z]+`)

type UserAgent struct {
	Device  *Device  `json:"device"`
	OS      *OS      `json:"os"`
	Browser *Browser `json:"browser"`
}

type Parser struct {
	path string

	browserPatterns [][]interface{} // [ reg, spec ]
	osPatterns      [][]interface{}
	devicePatterns  [][]interface{}

	browserKeywords map[string][]interface{} // keyword => [ reg, spec ]
	osKeywords      map[string][]interface{} // keyword => [ reg, spec ]
	deviceKeywords  map[string][]interface{} // keyword => [ reg, spec ]
}

func NewParser(path string) (*Parser, error) {
	p := &Parser{
		path:            path,
		browserPatterns: [][]interface{}{},
		osPatterns:      [][]interface{}{},
		devicePatterns:  [][]interface{}{},

		browserKeywords: map[string][]interface{}{},
		osKeywords:      map[string][]interface{}{},
		deviceKeywords:  map[string][]interface{}{},
	}

	err := p.init()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (this *Parser) init() (err error) {
	data, err := ioutil.ReadFile(this.path)
	if err != nil {
		return
	}

	groups := map[string][]map[string]string{}

	err = yaml.Unmarshal(data, &groups)
	if err != nil {
		return
	}

	for key, patterns := range groups {
		for index, pattern := range patterns {
			regex, ok := pattern["regex"]
			if !ok {
				continue
			}
			delete(pattern, "regex")
			reg, err := regexp.Compile(regex)
			if err != nil {
				return err
			}

			if key == "user_agent_parsers" {
				this.browserPatterns = append(this.browserPatterns, []interface{}{reg, pattern, index})

				for _, keyword := range this.parseKeywordsFromPattern(regex) {
					list, found := this.browserKeywords[keyword]
					if found {
						list = append(list, []interface{}{reg, pattern, index})
					} else {
						list = []interface{}{[]interface{}{reg, pattern, index}}
					}
					this.browserKeywords[keyword] = list
				}
			} else if key == "os_parsers" {
				this.osPatterns = append(this.osPatterns, []interface{}{reg, pattern})

				for _, keyword := range this.parseKeywordsFromPattern(regex) {
					list, found := this.osKeywords[keyword]
					if found {
						list = append(list, []interface{}{reg, pattern, index})
					} else {
						list = []interface{}{[]interface{}{reg, pattern, index}}
					}
					this.osKeywords[keyword] = list
				}
			} else if key == "device_parsers" {
				this.devicePatterns = append(this.devicePatterns, []interface{}{reg, pattern})

				for _, keyword := range this.parseKeywordsFromPattern(regex) {
					list, found := this.deviceKeywords[keyword]
					if found {
						list = append(list, []interface{}{reg, pattern, index})
					} else {
						list = []interface{}{[]interface{}{reg, pattern, index}}
					}
					this.deviceKeywords[keyword] = list
				}
			}
		}
	}

	return nil
}

func (this *Parser) Parse(userAgentString string) (userAgent *UserAgent, found bool) {
	browser, found := this.ParseBrowser(userAgentString)
	if !found {
		return nil, false
	}

	userAgent = &UserAgent{}
	userAgent.Browser = browser

	os, found := this.ParseOS(userAgentString)
	if found {
		userAgent.OS = os
	}

	device, found := this.ParseDevice(userAgentString)
	if found {
		userAgent.Device = device
	}

	return userAgent, true
}

func (this *Parser) ParseBrowser(userAgentString string) (browser *Browser, found bool) {
	found = this.parseUserAgentKeywords(userAgentString, this.browserKeywords, func(matches []string, spec map[string]string) {
		browser = &Browser{}
		browser.Parse(matches, spec)
	})
	return
}

func (this *Parser) ParseOS(userAgentString string) (os *OS, found bool) {
	found = this.parseUserAgentKeywords(userAgentString, this.osKeywords, func(matches []string, spec map[string]string) {
		os = &OS{}
		os.Parse(matches, spec)
	})
	return
}

func (this *Parser) ParseDevice(userAgentString string) (device *Device, found bool) {
	found = this.parseUserAgentKeywords(userAgentString, this.deviceKeywords, func(matches []string, spec map[string]string) {
		device = &Device{}
		device.Parse(matches, spec)
	})
	return
}

func (this *Parser) parseKeywordsFromPattern(pattern string) []string {
	pattern = strings.Replace(pattern, "$", " ", -1)
	pattern = strings.Replace(pattern, "\\d", " ", -1)
	pattern = strings.ToLower(pattern)
	results := []string{}
	for _, s := range keywordReg.FindAllString(pattern, -1) {
		if len(s) <= 2 {
			continue
		}
		results = append(results, s)
	}
	return results
}

func (this *Parser) parseUserAgentKeywords(userAgentString string, keywordsMapping map[string][]interface{}, callback func(matches []string, spec map[string]string)) (found bool) {
	keywords := this.parseKeywordsFromPattern(userAgentString)
	patterns := []interface{}{}
	foundKeywords := false
	for _, keyword := range keywords {
		patternsArray, found := keywordsMapping[keyword]
		if !found {
			continue
		}
		patterns = append(patterns, patternsArray ...)
		foundKeywords = true
	}
	if foundKeywords {
		sort.Slice(patterns, func(i, j int) bool {
			return patterns[i].([]interface{})[2].(int) < patterns[j].([]interface{})[2].(int)
		})

		for _, setting := range patterns {
			reg := setting.([]interface{})[0].(*regexp.Regexp)
			spec := setting.([]interface{})[1].(map[string]string)

			if !reg.MatchString(userAgentString) {
				continue
			}

			matches := reg.FindStringSubmatch(userAgentString)
			callback(matches, spec)

			return true
		}
	}
	return false
}
