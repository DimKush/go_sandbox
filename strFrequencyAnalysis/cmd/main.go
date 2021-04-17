package main

import (
	"github.com/DimKush/go_sandbox/tree/main/strFrequencyAnalysis/internal/analyzer"
)

func main() {
	str := `In this chapter, we will develop a simple content generator. The customary
	HelloWorld example demonstrates the basic concepts of module programming,
	including the complete module structure, and use of the handler callback and
	request_rec.
	
	By the end of the chapter, we will have extended our HelloWorld module to report
	the full details of the request and response headers, the environment variables, and
	any data posted to the server, and we will be equipped to write content generator
	modules in situations where we might otherwise have used a CGI script or compa-
	rable extension
	`
	analyzer.Analyzer(str)

}
