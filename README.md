## User Agent Parser For TeaWeb 
More faster than other implementations. 

## Usage
~~~go
p, err := "/path/to/regexes.yaml")
if err != nil {
    t.Fatal(err)
}

agent, found := p.Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.59 Safari/537.36")

if found {
    log.Printf("%#v,\n %#v, \n %#v", agent.Browser, agent.OS, agent.Device)
}
~~~

## Links
* `UserAgent` data comes from：[https://github.com/ua-parser/uap-core](https://github.com/ua-parser/uap-core)
* [Specification](https://github.com/ua-parser/uap-core/blob/master/docs/specification.md)
