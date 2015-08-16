// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package {{.CommandClass.GetPackageName}}

// {{.Help}}
{{$version := .CommandClass.Version}}
{{$typeName := (ToGoName .Command.Name) "V" $version}}
type {{$typeName}} struct {
  {{range $_, $param := .Command.Params}}
    {{template "command-struct-fields.tpl" $param}}
  {{end}}
}

func (cmd *{{$typeName}}) UnmarshalBinary(payload []byte) error {
  {{template "unmarshal-command-params.tpl" .Command.Params}}

  return nil
}

func (cmd *{{$typeName}}) MarshalBinary() (payload []byte, err error) {
  {{template "marshal-command-params.tpl" .Command.Params}}
  return
}