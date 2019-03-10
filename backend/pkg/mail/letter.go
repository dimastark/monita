package mail

import (
	"text/template"
)

const letter = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
    <head>
        <meta name="viewport" content="width=device-width" />
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
        <title>Monita report</title>
        <style type="text/css">
            body {
                margin: 0 auto;
                width: 450px;
                font-size: 16px;
                font-family: sans-serif;
            }

            table {
                margin: 50px 0 50px 0;
            }

            .header {
                height: 40px;
                text-align: center;
                text-transform: uppercase;
                font-size: 24px;
                font-weight: bold;
            }

            .label {
                color: #247ad6;
                padding: 0 30px;
                width: 1px;
                white-space: nowrap;
                vertical-align: top;
            }

            .data {
                word-break: break-all;
            }

            .updates tr {
                line-height: 30px;
            }

            .updates td {
                padding-top: 15px;
            }
        </style>
    </head>
    <body>
        <table width="100%" border="0" cellspacing="0" cellpadding="0">
            <tr class="header">
                <td>Monita {{ .ReportType }} report</td>
            </tr>
        </table>

        {{ range .Updates }}
        <table class="updates" width="100%" border="0" cellspacing="0" cellpadding="0">
            <tr>
                <td class="label">ID:</td>
                <td>{{ .ID }}</td>
            </tr>
            <tr>
                <td class="label">Name:</td>
                <td>{{ .Name }}</td>
            </tr>
            <tr>
                <td class="label">Old Data:</td>
                <td class="data">{{ .OldData }}</td>
            </tr>
            <tr>
                <td class="label">New Data:</td>
                <td class="data">{{ .NewData }}</td>
            </tr>
		</table>
        {{ end }}
    </body>
</html>
`

var t = template.Must(template.New("letter").Parse(letter))
