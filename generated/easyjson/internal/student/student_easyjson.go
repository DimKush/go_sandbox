// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package student

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonB83d7b77DecodeGithubComDimKushGoSandboxTreeMainGeneratedEasyjsonInternalStudent(in *jlexer.Lexer, out *Student) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "FirstName":
			out.FirstName = string(in.String())
		case "LastName":
			out.LastName = string(in.String())
		case "Age":
			out.Age = int(in.Int())
		case "Marks":
			if in.IsNull() {
				in.Skip()
				out.Marks = nil
			} else {
				in.Delim('[')
				if out.Marks == nil {
					if !in.IsDelim(']') {
						out.Marks = make([]Descilpine, 0, 2)
					} else {
						out.Marks = []Descilpine{}
					}
				} else {
					out.Marks = (out.Marks)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Descilpine
					(v1).UnmarshalEasyJSON(in)
					out.Marks = append(out.Marks, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB83d7b77EncodeGithubComDimKushGoSandboxTreeMainGeneratedEasyjsonInternalStudent(out *jwriter.Writer, in Student) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"FirstName\":"
		out.RawString(prefix[1:])
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"LastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"Age\":"
		out.RawString(prefix)
		out.Int(int(in.Age))
	}
	{
		const prefix string = ",\"Marks\":"
		out.RawString(prefix)
		if in.Marks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Marks {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Student) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB83d7b77EncodeGithubComDimKushGoSandboxTreeMainGeneratedEasyjsonInternalStudent(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Student) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB83d7b77EncodeGithubComDimKushGoSandboxTreeMainGeneratedEasyjsonInternalStudent(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Student) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB83d7b77DecodeGithubComDimKushGoSandboxTreeMainGeneratedEasyjsonInternalStudent(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Student) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB83d7b77DecodeGithubComDimKushGoSandboxTreeMainGeneratedEasyjsonInternalStudent(l, v)
}
func easyjsonB83d7b77DecodeGithubComDimKushGoSandboxTreeMainGeneratedEasyjsonInternalStudent1(in *jlexer.Lexer, out *Descilpine) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Descilpine":
			out.Descilpine = string(in.String())
		case "Mark":
			out.Mark = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB83d7b77EncodeGithubComDimKushGoSandboxTreeMainGeneratedEasyjsonInternalStudent1(out *jwriter.Writer, in Descilpine) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Descilpine\":"
		out.RawString(prefix[1:])
		out.String(string(in.Descilpine))
	}
	{
		const prefix string = ",\"Mark\":"
		out.RawString(prefix)
		out.Int(int(in.Mark))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Descilpine) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB83d7b77EncodeGithubComDimKushGoSandboxTreeMainGeneratedEasyjsonInternalStudent1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Descilpine) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB83d7b77EncodeGithubComDimKushGoSandboxTreeMainGeneratedEasyjsonInternalStudent1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Descilpine) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB83d7b77DecodeGithubComDimKushGoSandboxTreeMainGeneratedEasyjsonInternalStudent1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Descilpine) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB83d7b77DecodeGithubComDimKushGoSandboxTreeMainGeneratedEasyjsonInternalStudent1(l, v)
}
