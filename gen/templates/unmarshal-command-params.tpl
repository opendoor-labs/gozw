{{with .}}i := 2{{end}}{{/* {{with}} here ensures we don't print this if there are no params */}}
{{ $keys := .OrderedKeys}}
{{ $command := .}}
{{ range $key := $keys}}
  {{ $type := $command.GetFieldType $key}}
  {{if eq $type "param"}}
    {{ $param := $command.GetParam $key }}
    {{template "unmarshal-command-param" $param}}
  {{else if eq $type "vg"}}
    {{ $vg := $command.GetVg $key }}
    {{ $isLast := IsLastKey $key $keys}}
    {{if eq $isLast true}}
      for i < len(payload) {
    {{else}}
      for i < int(cmd.{{$command.FindVgLengthVar $key $keys}}) {
    {{end}}
    {{ToGoNameLower $vg.Name}} := {{$command.GetStructName $command.CC}}{{ToGoName $vg.Name}} {}
    {{template "unmarshal-command-vg-params" $vg}}
    {{range $_, $param := $vg.Params}}
      {{ if eq $param.Type "STRUCT_BYTE" -}}
       // struct byte fields are assigned to the variant group when computed
      {{- else -}}
      {{ToGoNameLower $vg.Name}}.{{ToGoName $param.Name}} = {{ToGoNameLower $param.Name}}
      {{- end}}
    {{end}}
    cmd.{{ToGoName $vg.Name}} = append(cmd.{{ToGoName $vg.Name}}, {{ToGoNameLower $vg.Name}})
    }
  {{end}}
{{end}}
