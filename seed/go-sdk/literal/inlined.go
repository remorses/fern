// This file was auto-generated by Fern from our API Definition.

package literal

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/literal/fern/internal"
)

type SendLiteralsInlinedRequest struct {
	Context           *string             `json:"context,omitempty" url:"-"`
	Query             string              `json:"query" url:"-"`
	Temperature       *float64            `json:"temperature,omitempty" url:"-"`
	AliasedContext    SomeAliasedLiteral  `json:"aliasedContext,omitempty" url:"-"`
	MaybeContext      *SomeAliasedLiteral `json:"maybeContext,omitempty" url:"-"`
	ObjectWithLiteral *ATopLevelLiteral   `json:"objectWithLiteral,omitempty" url:"-"`
	prompt            string
	stream            bool
}

func (s *SendLiteralsInlinedRequest) Prompt() string {
	return s.prompt
}

func (s *SendLiteralsInlinedRequest) Stream() bool {
	return s.stream
}

func (s *SendLiteralsInlinedRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler SendLiteralsInlinedRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*s = SendLiteralsInlinedRequest(body)
	s.prompt = "You are a helpful assistant"
	s.stream = false
	return nil
}

func (s *SendLiteralsInlinedRequest) MarshalJSON() ([]byte, error) {
	type embed SendLiteralsInlinedRequest
	var marshaler = struct {
		embed
		Prompt string `json:"prompt"`
		Stream bool   `json:"stream"`
	}{
		embed:  embed(*s),
		Prompt: "You are a helpful assistant",
		Stream: false,
	}
	return json.Marshal(marshaler)
}

type ANestedLiteral struct {
	myLiteral string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ANestedLiteral) MyLiteral() string {
	return a.myLiteral
}

func (a *ANestedLiteral) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ANestedLiteral) UnmarshalJSON(data []byte) error {
	type embed ANestedLiteral
	var unmarshaler = struct {
		embed
		MyLiteral string `json:"myLiteral"`
	}{
		embed: embed(*a),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*a = ANestedLiteral(unmarshaler.embed)
	if unmarshaler.MyLiteral != "How super cool" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "How super cool", unmarshaler.MyLiteral)
	}
	a.myLiteral = unmarshaler.MyLiteral
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "myLiteral")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ANestedLiteral) MarshalJSON() ([]byte, error) {
	type embed ANestedLiteral
	var marshaler = struct {
		embed
		MyLiteral string `json:"myLiteral"`
	}{
		embed:     embed(*a),
		MyLiteral: "How super cool",
	}
	return json.Marshal(marshaler)
}

func (a *ANestedLiteral) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type ATopLevelLiteral struct {
	NestedLiteral *ANestedLiteral `json:"nestedLiteral,omitempty" url:"nestedLiteral,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ATopLevelLiteral) GetNestedLiteral() *ANestedLiteral {
	if a == nil {
		return nil
	}
	return a.NestedLiteral
}

func (a *ATopLevelLiteral) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ATopLevelLiteral) UnmarshalJSON(data []byte) error {
	type unmarshaler ATopLevelLiteral
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ATopLevelLiteral(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ATopLevelLiteral) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type SomeAliasedLiteral = string
