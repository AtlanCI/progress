package internal

import (
	"encoding/json"
	"strings"
)

type Mode int

const (
	ModeSpinner Mode = iota
	ModeProgress
	ModeDone
	ModeError
)

func (m *Mode) MarshalJSON() ([]byte, error) {
	var s string
	switch *m {
	case ModeSpinner:
		s = "running"
	case ModeProgress:
		s = "progress"
	case ModeDone:
		s = "done"
	case ModeError:
		s = "error"
	default:
		s = "unknown"
	}
	return json.Marshal(s)
}

func (m *Mode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	case "spinner", "running": // keep "spinner" for compatiability
		*m = ModeSpinner
	case "progress":
		*m = ModeProgress
	case "done":
		*m = ModeDone
	case "error":
		*m = ModeError
	default:
		panic("unknown mode")
	}

	return nil
}

func (m *Mode) String() string {
	var s string
	switch *m {
	case ModeSpinner:
		s = "running"
	case ModeProgress:
		s = "progress"
	case ModeDone:
		s = "done"
	case ModeError:
		s = "error"
	default:
		s = "unknown"
	}
	return s
}

type DisplayProps struct {
	Prefix string `json:"prefix,omitempty"`
	Suffix string `json:"suffix,omitempty"`
	Mode   Mode   `json:"mode,omitempty"`
}

func (dp *DisplayProps) String() string {
	var builder strings.Builder
	builder.WriteString("(" + dp.Mode.String() + ")")
	builder.WriteByte(' ')
	builder.WriteString(dp.Prefix + ":")
	builder.WriteByte(' ')
	builder.WriteString(dp.Suffix)
	return builder.String()
}
