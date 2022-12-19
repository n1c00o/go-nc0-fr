package tmpl

import "html/template"

// IndexTmpl is the HTML template to generate for the index page of the static
// site (route "/").
//
// Variables:
// - Hostname should be the hostname of the host, for example go.nc0.fr
// - Prefixes should be a list of module's import prefix, including the
// hostname, for instance go.nc0.fr/foo
var IndexTmpl = template.Must(
	template.New("index").Parse(`<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <title>go.nc0.fr</title>
        <meta name="title" content="go.nc0.fr">
		<meta name="description" content="Index of Go modules hosted on go.nc0.fr">
        <meta name="robots" content="index,follow">
    </head>
    <body>
		<h1>{{.Hostname}}</h1>
		<ul>
			{{range .Prefixes}}<li><a href="https://pkg.go.dev/{{.}}">{{.}}</a></li>{{end}}
		</ul>
    </body>
</html>`),
)

// ModuleTmpl is the HTML template to generate for the page of a module.
//
// Variables:
// - Hostname should be the hostname of the host, for example go.nc0.fr
// - Prefix should be the module's import prefix, including the hostname,
// for instance go.nc0.fr/foo
// - Repo should be the module's repository
// - VCS should be the VCS tool of the repository
// - Dir should be the URI that lists the files of a directory
// - File should be the URI with the content of a file
var ModuleTmpl = template.Must(
	template.New("module").Parse(`<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
			<meta name="go-import" content="{{.Prefix}} {{.VCS}} {{.Repo}}">
			<meta name="go-source" content="{{.Prefix}} {{.Repo}} {{.Dir}} {{.File}}">
			<meta http-equiv="refresh" content="0; url=https://pkg.go.dev/{{.Prefix}}">
			<title>{{.Prefix}}</title>
		</head>
		<body>
			<p>
				Nothing to see here, <a href="https://pkg.go.dev/{{.Prefix}}">see the package on pkg.go.dev</a>.
			</p>
		</body>
	</html>`),
)
