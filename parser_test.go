package uaparser

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestParser_Parse(t *testing.T) {
	p, err := NewParser(os.Getenv("GOPATH") + "/src/github.com/TeaWeb/uaparser/regexes.yaml")
	if err != nil {
		t.Fatal(err)
	}
	b, found := p.ParseBrowser("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.59 Safari/537.36")
	if !found {
		t.Log("not found")
	} else {
		t.Logf("%#v", b)
	}

	o, found := p.ParseOS("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.59 Safari/537.36")
	if !found {
		t.Log("not found")
	} else {
		t.Logf("%#v", o)
	}

	d, found := p.ParseDevice("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.59 Safari/537.36 iPhone OS/10")
	if !found {
		t.Log("not found")
	} else {
		t.Logf("%#v", d)
	}
}

func TestParser_Parse_Cost(t *testing.T) {
	p, err := NewParser(os.Getenv("GOPATH") + "/src/github.com/TeaWeb/uaparser/regexes.yaml")
	if err != nil {
		t.Fatal(err)
	}

	before := time.Now()
	agent, found := p.Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.59 Safari/537.36")
	cost := time.Since(before).Seconds()

	if !found {
		t.Log("not found")
	} else {
		log.Printf("%#v,\n %#v, \n %#v", agent.Browser, agent.OS, agent.Device)
	}

	t.Log("cost:", cost)
	t.Log("qps:", 1/cost)
}

func TestParser_Keywords(t *testing.T) {
	p, err := NewParser(os.Getenv("GOPATH") + "/src/github.com/TeaWeb/uaparser/regexes.yaml")
	if err != nil {
		t.Fatal(err)
	}

	//pattern := "(ESPN)[%20| ]+Radio/(\\d+)\\.(\\d+)\\.(\\d+) CFNetwork"
	pattern := `(SE 2\.X) MetaSr (\d+)\.(\d+) map[family_replacement:Sogou Explorer`
	t.Log(p.parseKeywordsFromPattern(pattern))
}
