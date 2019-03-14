package mail

import (
	"text/template"
)

const (
	mainStyle      = "margin: 0 auto; width: 450px; font-size: 16px; font-family: sans-serif;"
	tableStyle     = "margin: 30px 0 30px 0;"
	headerStyle    = "height: 40px; text-align: center; text-transform: uppercase; font-size: 24px;"
	labelStyle     = "color: #247ad6; padding: 0 30px; width: 1px; white-space: nowrap; vertical-align: top;"
	dataStyle      = "padding-top: 15px; word-break: break-all;"
	updatesTrStyle = "line-height: 30px;"

	letter = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
    <head>
        <meta name="viewport" content="width=device-width" />
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
        <title>Monita report</title>
    </head>
    <body><main style="` + mainStyle + `">
        <table style="` + tableStyle + `" width="100%" border="0" cellspacing="0" cellpadding="0">
            <tr style="` + headerStyle + `">
                <td><b>Monita {{ .ReportType }} report</b></td>
            </tr>
        </table>

        {{ range .Updates }}
        <table style="` + tableStyle + `" width="100%" border="0" cellspacing="0" cellpadding="0">
            <tr style="` + updatesTrStyle + `">
                <td style="` + labelStyle + `">ID:</td>
                <td>{{ .ID }}</td>
            </tr>
            <tr style="` + updatesTrStyle + `">
                <td style="` + labelStyle + `">Name:</td>
                <td>{{ .Name }}</td>
            </tr>
            <tr style="` + updatesTrStyle + `">
                <td style="` + labelStyle + `">Old Data:</td>
                <td class="` + dataStyle + `">{{ .OldData }}</td>
            </tr>
            <tr style="` + updatesTrStyle + `">
                <td style="` + labelStyle + `">New Data:</td>
                <td class="` + dataStyle + `">{{ .NewData }}</td>
            </tr>
		</table>
        {{ end }}
    </main></body>
</html>
`
)

var t = template.Must(template.New("letter").Parse(letter))
