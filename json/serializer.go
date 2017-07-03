package json

import (
	"bytes"
	"strconv"
	"time"
)

type Serializer struct {
	buff  bytes.Buffer
	first bool
}

const OPEN_LITERAL = '\u007B'
const SEPARATOR = ','
const TRUE = "true"
const FALSE = "false"
const CLOSED_LITERAL = '\u007D'
const KEY_SEPARATOR = ':'
const ARRAY_SEPARATOR = ','

const QUOTE = '"'
const ARRAY_START_QUOTE = '\u005B'
const ARRAY_CLOSE_QUOTE = '\u005D'

func New() *Serializer {
	b := &Serializer{bytes.Buffer{}, true}
	b.buff.WriteRune(OPEN_LITERAL)
	return b
}

func (serializer *Serializer) isFirst() (bool) {
	isFirst := serializer.first
	if isFirst {
		serializer.first = false
	}
	return isFirst
}
func (serializer *Serializer) quote() (*Serializer) {
	serializer.buff.WriteRune(QUOTE)
	return serializer
}

func (serializer *Serializer) SetJsonObject(key string, objectSerializer *Serializer) (*Serializer) {
	serializer.setKey(key).buff.WriteString(objectSerializer.Serialize())
	return serializer
}
func (serializer *Serializer) fieldSeparator() (*Serializer) {
	if !serializer.isFirst() {
		serializer.buff.WriteRune(SEPARATOR)
	}
	return serializer
}
func (serializer *Serializer) setKey(key string) (*Serializer) {
	serializer.fieldSeparator().quote()
	serializer.buff.WriteString(key)
	serializer.quote()
	serializer.buff.WriteRune(KEY_SEPARATOR)
	return serializer
}

func (serializer *Serializer) SetString(key string, val string) (*Serializer) {
	serializer.setKey(key)
	serializer.quote()
	serializer.buff.WriteString(val)
	serializer.quote()
	return serializer
}

func (serializer *Serializer) SetInt(key string, val int) (*Serializer) {
	serializer.setKey(key).buff.WriteString(strconv.Itoa(val))
	return serializer
}

func (serializer *Serializer) SetTime(key string, val time.Time) (*Serializer) {
	serializer.setKey(key).quote().buff.WriteString(val.Format(time.RFC3339))
	serializer.quote()
	return serializer
}

func (serializer *Serializer) setString(key string, val string) (*Serializer) {
	serializer.buff.WriteString(val)
	return serializer
}
func (serializer *Serializer) setInt(val int) (*Serializer) {
	return serializer
}

func (serializer *Serializer) SetIntArray(key string, arr []int) (*Serializer) {
	serializer.setKey(key)
	serializer.buff.WriteRune(ARRAY_START_QUOTE)
	for i, val := range arr {
		if i != 0 {
			serializer.buff.WriteRune(ARRAY_SEPARATOR)
		}
		serializer.buff.WriteString(strconv.Itoa(val))
	}
	serializer.buff.WriteRune(ARRAY_CLOSE_QUOTE)
	return serializer
}

func (serializer *Serializer) SetStringArray(key string, arr []string) (*Serializer) {
	serializer.setKey(key)
	serializer.buff.WriteRune(ARRAY_START_QUOTE)
	for i, val := range arr {
		if i != 0 {
			serializer.buff.WriteRune(ARRAY_SEPARATOR)
		}
		serializer.quote()
		serializer.buff.WriteString(val)
		serializer.quote()
	}
	serializer.buff.WriteRune(ARRAY_CLOSE_QUOTE)
	return serializer
}

func (serializer *Serializer) SetTimeArray(key string, arr []time.Time) (*Serializer) {
	serializer.setKey(key)
	serializer.buff.WriteRune(ARRAY_START_QUOTE)
	for i, val := range arr {
		if i != 0 {
			serializer.buff.WriteRune(ARRAY_SEPARATOR)
		}
		serializer.quote()
		serializer.buff.WriteString(strconv.FormatInt(val.Unix(), 10))
		serializer.quote()
	}
	serializer.buff.WriteRune(ARRAY_CLOSE_QUOTE)
	return serializer
}

func (serializer *Serializer) SetJsonObjectArray(key string, arr []*Serializer) (*Serializer) {
	serializer.setKey(key)
	serializer.buff.WriteRune(ARRAY_START_QUOTE)
	for i, val := range arr {
		if i != 0 {
			serializer.buff.WriteRune(ARRAY_SEPARATOR)
		}
		serializer.buff.WriteString(val.Serialize())
	}
	serializer.buff.WriteRune(ARRAY_CLOSE_QUOTE)
	return serializer
}

func (serializer *Serializer) SetEpochTime(key string, val time.Time) (*Serializer) {
	serializer.setKey(key).buff.WriteString(strconv.FormatInt(val.Unix(), 10))
	return serializer
}
func (serializer *Serializer) SetBoolean(key string, boolean bool) (*Serializer) {
	if boolean {
		serializer.setKey(key).buff.WriteString(TRUE)
	} else {
		serializer.setKey(key).buff.WriteString(FALSE)
	}
	return serializer
}

func (serializer *Serializer) Serialize() string {
	serializer.buff.WriteRune(CLOSED_LITERAL)
	return serializer.buff.String()
}
