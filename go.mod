module github.com/CodeSpaceiiii/dalutest-20220202

go 1.14

require (
	github.com/alibabacloud-go/darabonba-openapi/v2 v2.1.13
	github.com/alibabacloud-go/tea v1.3.13
	github.com/aliyun/credentials-go v1.4.8
)

replace (
	github.com/alibabacloud-go/darabonba-openapi/v2 => ../darabonba-openapi/golang
	github.com/alibabacloud-go/tea => ../tea
	github.com/alibabacloud-go/tea-utils/v2 => ../tea-utils
)
