package uaparser

import (
	"fmt"
	"math/rand"
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
	//agent, found := p.Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.59 Safari/537.36")

	count := 10000
	for i := 0; i < count; i++ {
		r := rand.Int()
		p.Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.59 Safari/537.36 " + fmt.Sprintf("%d", r%12000))
	}

	cost := time.Since(before).Seconds()

	t.Log("qps:", float64(count)/cost)
}

func TestParser_Parse_Performance(t *testing.T) {
	p, err := NewParser(os.Getenv("GOPATH") + "/src/github.com/TeaWeb/uaparser/regexes.yaml")
	if err != nil {
		t.Fatal(err)
	}

	before := time.Now()

	count := 50000
	for i := 0; i < count; i ++ {
		p.ParseBrowser("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.59 Safari/537.36")
	}

	cost := time.Since(before).Seconds()

	t.Log("qps:", float64(count)/cost)
}

func TestParser_Keywords(t *testing.T) {
	p, err := NewParser(os.Getenv("GOPATH") + "/src/github.com/TeaWeb/uaparser/regexes.yaml")
	if err != nil {
		t.Fatal(err)
	}

	//pattern := "(ESPN)[%20| ]+Radio/(\\d+)\\.(\\d+)\\.(\\d+) CFNetwork"
	pattern := `(SE 2\.X) MetaSr (\d+)\.(\d+) map[family_replacement:Sogou Explorer`
	before := time.Now()
	result := p.parseKeywordsFromPattern(pattern)
	cost := time.Since(before).Seconds()
	t.Log(result)
	t.Log(1/cost, "qps")
}

func TestParser_Keywords2(t *testing.T) {
	p, err := NewParser(os.Getenv("GOPATH") + "/src/github.com/TeaWeb/uaparser/regexes.yaml")
	if err != nil {
		t.Fatal(err)
	}

	//pattern := "(ESPN)[%20| ]+Radio/(\\d+)\\.(\\d+)\\.(\\d+) CFNetwork"
	pattern := `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36`
	before := time.Now()
	result := p.parseKeywordsFromPattern(pattern)
	cost := time.Since(before).Seconds()
	t.Log(result)
	t.Log(1/cost, "qps")
}

func TestParser_Wget(t *testing.T) {
	userAgent := "Wget/1.0"

	p, err := NewParser(os.Getenv("GOPATH") + "/src/github.com/TeaWeb/uaparser/regexes.yaml")
	if err != nil {
		t.Fatal(err)
	}
	b, found := p.ParseBrowser(userAgent)
	if !found {
		t.Log("not found")
	} else {
		t.Logf("%#v", b)
	}

	o, found := p.ParseOS(userAgent)
	if !found {
		t.Log("not found")
	} else {
		t.Logf("%#v", o)
	}

	d, found := p.ParseDevice(userAgent)
	if !found {
		t.Log("not found")
	} else {
		t.Logf("%#v", d)
	}
}

func TestParser_Other(t *testing.T) {
	userAgent := "Hello/1.0"

	p, err := NewParser(os.Getenv("GOPATH") + "/src/github.com/TeaWeb/uaparser/regexes.yaml")
	if err != nil {
		t.Fatal(err)
	}
	b, found := p.Parse(userAgent)
	t.Log(b, found)
}
