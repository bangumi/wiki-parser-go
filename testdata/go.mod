module wiki

go 1.23.0

require (
	github.com/bangumi/wiki-parser-go v0.0.0
	github.com/goccy/go-yaml v1.19.2
	github.com/stretchr/testify v1.12.1
)

require go.yaml.in/yaml/v3 v3.0.5 // indirect

replace github.com/bangumi/wiki-parser-go => ../
