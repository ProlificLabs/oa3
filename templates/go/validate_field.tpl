{{- /* Top level object is a Schema, Name is the name of the local var */ -}}
{{- $name := $.Name}}
ers = nil
{{- if $.Object.MaxLength}}
if err := primitives.ValidateMaxLength({{$name}}, {{$.Object.MaxLength}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if $.Object.MinLength}}
if err := primitives.ValidateMinLength({{$name}}, {{$.Object.MinLength}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if and $.Object.Maximum (or (eq $.Object.Type "number") (eq $.Object.Type "integer"))}}
if err := primitives.ValidateMaxNumber({{$name}}, {{$.Object.Maximum}}, {{$.Object.ExclusiveMaximum}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if and $.Object.Maximum (eq $.Object.Type "string") (eq (printf $.Object.Format) "decimal") (eq $.Params.decimaltype "shopspring")}}
if err := primitives.ValidateMaxShopspringDecimal({{$name}}, primitives.NewDecimalFromFloat({{$.Object.Maximum}}), {{$.Object.ExclusiveMaximum}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if and $.Object.Minimum (or (eq $.Object.Type "number") (eq $.Object.Type "integer"))}}
if err := primitives.ValidateMinNumber({{$name}}, {{$.Object.Minimum}}, {{$.Object.ExclusiveMinimum}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if and $.Object.Minimum (eq $.Object.Type "string") (eq (printf $.Object.Format) "decimal") (eq $.Params.decimaltype "shopspring")}}
if err := primitives.ValidateMinShopspringDecimal({{$name}}, primitives.NewDecimalFromFloat({{$.Object.Minimum}}), {{$.Object.ExclusiveMinimum}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if $.Object.MultipleOf}}
if err := primitives.ValidateMultipleOf{{if eq $.Object.Type "integer"}}Int{{else}}Float{{end}}({{$name}}, {{$.Object.MultipleOf}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if $.Object.MaxItems}}
if err := primitives.ValidateMaxItems({{$name}}, {{$.Object.MaxItems}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if $.Object.MinItems}}
if err := primitives.ValidateMinItems({{$name}}, {{$.Object.MinItems}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if $.Object.MaxProperties}}
if err := primitives.ValidateMaxProperties({{$name}}, {{$.Object.MaxProperties}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if $.Object.MinProperties}}
if err := primitives.ValidateMinProperties({{$name}}, {{$.Object.MinProperties}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if $.Object.Pattern }}
if err := primitives.ValidatePattern({{$name}}, {{printf $.Object.Pattern}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if and $.Object.Format (eq (printf $.Object.Format) "uuid") }}
if err := primitives.ValidateFormatUUIDv4(string({{$name}})); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if and $.Object.Format (eq (printf $.Object.Format) "decimal") (eq $.Params.decimaltype "string") }}
if err := primitives.ValidateFormatDecimal({{$name}}); err != nil {
    ers = append(ers, err)
}
{{- end -}}
{{- if and $.Object.Enum (gt (len $.Object.Enum) 0) }}
if err := primitives.ValidateEnum({{$name}}, []string{
    {{- range $i, $v := $.Object.Enum -}}
        {{- if ne "string" (typeOf .) -}}
        {{- else -}}
        {{if gt $i 0}}, {{end}}{{printf "%q" $v}}{{end -}}
    {{- end -}}
        }); err != nil {
    ers = append(ers, err)
}
{{end -}}
