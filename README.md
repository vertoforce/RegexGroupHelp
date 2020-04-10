# Regex Group Helper

I use this function so much I just made a quick library for it.  It helps find regex groups in a regex expression

## Usage

```go
regex := regexp.MustCompile(`(?P<testGroup>test) data`)

foundGroups := FindRegexGroups(regex, "this is test data")
fmt.Printf("%v\n", foundGroups)

// Output: map[:[test data] testGroup:[test]]
```
