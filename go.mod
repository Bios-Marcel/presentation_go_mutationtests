module presentation_go_mutationtests

go 1.21.1

require (
	github.com/gtramontina/ooze v0.2.0
	github.com/stretchr/testify v1.8.4
)

// Windows support
replace github.com/gtramontina/ooze => ../ooze

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/hexops/gotextdiff v1.0.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
