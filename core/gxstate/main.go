package gxstate

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

type GXState struct {
	content map[string]interface{}
	Prefix  string
}

func NewGXState(raw string) (*GXState, error) {
	raw = strings.ReplaceAll(raw, "\\>", "&gt;")

	re := regexp.MustCompile(`MPW\d{4}`)
	prefix := re.FindString(raw)
	if prefix == "" {
		return nil, errors.New("prefix not found")
	}

	var content map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &content); err != nil {
		return nil, err
	}

	return &GXState{content: content, Prefix: prefix}, nil
}

func (gs *GXState) Get(key string) interface{} {
	return gs.content[key]
}

func (gs *GXState) GetWithPrefix(key string) interface{} {
	return gs.content[gs.Prefix+key]
}
