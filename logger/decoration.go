package logger

// Decoration is decorator for terminal
type Decoration struct {
	Normal    string
	Bold      string
	Dim       string
	Italic    string
	Underline string
	Blinking  string
	Reverse   string
	Invisible string
	Yellow    string
	Important string
}

// NewDecoration Create New Decoration
func NewDecoration() (decoration Decoration) {
	return Decoration{
		Normal:    "\x1B[0m",
		Bold:      "\x1B[1m",
		Dim:       "\x1B[2m",
		Italic:    "\x1B[3m",
		Underline: "\x1B[4m",
		Blinking:  "\x1B[5m",
		Reverse:   "\x1B[6m",
		Invisible: "\x1B[7m",
		Yellow:    "\x1B[33m",
		Important: "\x1B[33m",
	}
}
