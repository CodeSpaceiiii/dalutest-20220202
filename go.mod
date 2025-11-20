module github.com/CodeSpaceiiii/dalutest-20220202

go 1.14

require (
	github.com/alibabacloud-go/darabonba-openapi/v2 v2.1.13
	github.com/alibabacloud-go/tea v1.3.13
)

replace (
	github.com/alibabacloud-go/darabonba-openapi/v2 => ../darabonba-openapi/golang
	github.com/alibabacloud-go/tea => ../tea
)