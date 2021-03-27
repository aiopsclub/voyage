package util

import (
	"bytes"
	"text/template"
)

type MasterHost struct {
	HostName string
	Ip       string
}

type MasterHostList struct {
	Hosts []*MasterHost
}

func (m *MasterHostList) MasterHostsInfo() (string, error) {
	var out bytes.Buffer
	const hostCentext = `
if ! grep "# generate by voyage" /etc/hosts &> /dev/null
then
cat >> /etc/hosts <<EOF
# generate by voyage
{{-  range .Hosts  }}	
{{ .Ip  }} {{ .HostName }}
{{- end }}
EOF
fi
`
	t := template.Must(template.New("hosts").Parse(hostCentext))
	err := t.Execute(&out, m)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
