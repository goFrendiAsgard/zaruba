package output

import (
	"strings"

	"github.com/state-alchemists/zaruba/strutil"
)

// Decoration is decorator for terminal
type Decoration struct {
	Normal           string
	Bold             string
	Faint            string
	Italic           string
	Underline        string
	BlinkSlow        string
	BlinkRapid       string
	Inverse          string
	Conceal          string
	CrossedOut       string
	Black            string
	Red              string
	Green            string
	Yellow           string
	Blue             string
	Magenta          string
	Cyan             string
	White            string
	BgBlack          string
	BgRed            string
	BgGreen          string
	BgYellow         string
	BgBlue           string
	BgMagenta        string
	BgCyan           string
	BgWhite          string
	NoStyle          string
	NoUnderline      string
	NoInverse        string
	NoColor          string
	colorIndex       int
	colorList        []string
	iconIndex        int
	iconList         []string
	ZarubaIcon       string
	SuccessIcon      string
	ErrorIcon        string
	StartIcon        string
	KillIcon         string
	InspectIcon      string
	RunIcon          string
	WorkerIcon       string
	ScriptIcon       string
	ConstructionIcon string
	ContainerIcon    string
	EmptyIcon        string
	Icon             func(string) string
}

func NewPlainDecoration() (d *Decoration) {
	return &Decoration{
		Normal:           "",
		Bold:             "",
		Faint:            "",
		Italic:           "",
		Underline:        "",
		BlinkSlow:        "",
		BlinkRapid:       "",
		Inverse:          "",
		Conceal:          "",
		CrossedOut:       "",
		Black:            "",
		Red:              "",
		Green:            "",
		Yellow:           "",
		Blue:             "",
		Magenta:          "",
		Cyan:             "",
		White:            "",
		BgBlack:          "",
		BgRed:            "",
		BgGreen:          "",
		BgYellow:         "",
		BgBlue:           "",
		BgMagenta:        "",
		BgCyan:           "",
		BgWhite:          "",
		NoStyle:          "",
		NoUnderline:      "",
		NoInverse:        "",
		NoColor:          "",
		ZarubaIcon:       "",
		SuccessIcon:      "",
		ErrorIcon:        "",
		StartIcon:        "",
		KillIcon:         "",
		InspectIcon:      "",
		RunIcon:          "",
		WorkerIcon:       "",
		ScriptIcon:       "",
		ConstructionIcon: "",
		ContainerIcon:    "",
		EmptyIcon:        "",
		colorIndex:       0,
		iconIndex:        0,
		iconList:         []string{""},
		colorList:        []string{""},
		Icon:             func(icon string) string { return "" },
	}
}

// NewDefaultDecoration Create New Decoration
func NewDefaultDecoration() (d *Decoration) {
	// source: https://gist.github.com/mxmerz/92e97cd27857a9ba787b
	d = &Decoration{
		Normal:           "\x1b[0m",
		Bold:             "\x1b[1m",
		Faint:            "\x1b[2m",
		Italic:           "\x1b[3m",
		Underline:        "\x1b[4m",
		BlinkSlow:        "\x1b[5m",
		BlinkRapid:       "\x1b[6m",
		Inverse:          "\x1b[7m",
		Conceal:          "\x1b[8m",
		CrossedOut:       "\x1b[9m",
		Black:            "\x1b[30m",
		Red:              "\x1b[31m",
		Green:            "\x1b[32m",
		Yellow:           "\x1b[33m",
		Blue:             "\x1b[34m",
		Magenta:          "\x1b[35m",
		Cyan:             "\x1b[36m",
		White:            "\x1b[37m",
		BgBlack:          "\x1b[40m",
		BgRed:            "\x1b[41m",
		BgGreen:          "\x1b[42m",
		BgYellow:         "\x1b[43m",
		BgBlue:           "\x1b[44m",
		BgMagenta:        "\x1b[45m",
		BgCyan:           "\x1b[46m",
		BgWhite:          "\x1b[47m",
		NoStyle:          "\x1b[0m",
		NoUnderline:      "\x1b[24m",
		NoInverse:        "\x1b[27m",
		NoColor:          "\x1b[39m",
		ZarubaIcon:       "🤖",
		SuccessIcon:      "🎉",
		ErrorIcon:        "🔥",
		StartIcon:        "🏁",
		KillIcon:         "🔪",
		InspectIcon:      "🔎",
		RunIcon:          "🚀",
		WorkerIcon:       "👷",
		ScriptIcon:       "📜",
		ConstructionIcon: "🚧",
		ContainerIcon:    "🐳",
		EmptyIcon:        "  ",
		colorIndex:       0,
		iconIndex:        0,
		iconList: []string{
			"🍏", "🍎", "🍌", "🍉", "🍇", "🍐", "🍊", "🍋", "🍓", "🍈", "🍒", "🍑", "🍍", "🥝", "🍅", "🍆", "🥑",
		},
		Icon: func(icon string) string { return icon },
	}
	d.colorList = []string{
		d.Green,
		d.Yellow,
		d.Blue,
		d.Magenta,
		d.Cyan,
		d.Bold + d.Green,
		d.Bold + d.Yellow,
		d.Bold + d.Blue,
		d.Bold + d.Magenta,
		d.Bold + d.Cyan,
	}
	return d
}

// NewColorlessDecoration Create New Decoration
func NewColorlessDecoration() (d *Decoration) {
	d = &Decoration{
		Normal:           "",
		Bold:             "",
		Faint:            "",
		Italic:           "",
		Underline:        "",
		BlinkSlow:        "",
		BlinkRapid:       "",
		Inverse:          "",
		Conceal:          "",
		CrossedOut:       "",
		Black:            "",
		Red:              "",
		Green:            "",
		Yellow:           "",
		Blue:             "",
		Magenta:          "",
		Cyan:             "",
		White:            "",
		BgBlack:          "",
		BgRed:            "",
		BgGreen:          "",
		BgYellow:         "",
		BgBlue:           "",
		BgMagenta:        "",
		BgCyan:           "",
		BgWhite:          "",
		NoStyle:          "",
		NoUnderline:      "",
		NoInverse:        "",
		NoColor:          "",
		ZarubaIcon:       "🤖",
		SuccessIcon:      "🎉",
		ErrorIcon:        "🔥",
		StartIcon:        "🏁",
		KillIcon:         "🔪",
		InspectIcon:      "🔎",
		RunIcon:          "🚀",
		WorkerIcon:       "👷",
		ScriptIcon:       "📜",
		ConstructionIcon: "🚧",
		ContainerIcon:    "🐳",
		EmptyIcon:        "  ",
		colorIndex:       0,
		iconIndex:        0,
		iconList: []string{
			"🍏", "🍎", "🍌", "🍉", "🍇", "🍐", "🍊", "🍋", "🍓", "🍈", "🍒", "🍑", "🍍", "🥝", "🍅", "🍆", "🥑",
		},
		Icon: func(icon string) string { return icon },
	}
	d.colorList = []string{
		d.Green,
		d.Yellow,
		d.Blue,
		d.Magenta,
		d.Cyan,
		d.Bold + d.Green,
		d.Bold + d.Yellow,
		d.Bold + d.Blue,
		d.Bold + d.Magenta,
		d.Bold + d.Cyan,
	}
	return d
}

// GenerateColor new color
func (d *Decoration) GenerateColor() string {
	if d.colorIndex >= len(d.colorList) {
		d.colorIndex = 0
	}
	color := d.colorList[d.colorIndex]
	d.colorIndex++
	return color
}

// GenerateIcon new icon
func (d *Decoration) GenerateIcon() string {
	if d.iconIndex >= len(d.iconList) {
		d.iconIndex = 0
	}
	icon := d.iconList[d.iconIndex]
	d.iconIndex++
	return icon
}

func (d *Decoration) ToShellVariables() string {
	shellVariableList := []string{
		strutil.StrShellVariable("_NORMAL", d.Normal),
		strutil.StrShellVariable("_BOLD", d.Bold),
		strutil.StrShellVariable("_FAINT", d.Faint),
		strutil.StrShellVariable("_ITALIC", d.Italic),
		strutil.StrShellVariable("_UNDERLINE", d.Underline),
		strutil.StrShellVariable("_BLINK_SLOW", d.BlinkSlow),
		strutil.StrShellVariable("_BLINK_RAPID", d.BlinkRapid),
		strutil.StrShellVariable("_INVERSE", d.Inverse),
		strutil.StrShellVariable("_CONCEAL", d.Conceal),
		strutil.StrShellVariable("_CROSSED_OUT", d.CrossedOut),
		strutil.StrShellVariable("_BLACK", d.Black),
		strutil.StrShellVariable("_RED", d.Red),
		strutil.StrShellVariable("_GREEN", d.Green),
		strutil.StrShellVariable("_YELLOW", d.Yellow),
		strutil.StrShellVariable("_BLUE", d.Blue),
		strutil.StrShellVariable("_MAGENTA", d.Magenta),
		strutil.StrShellVariable("_CYAN", d.Cyan),
		strutil.StrShellVariable("_WHITE", d.White),
		strutil.StrShellVariable("_BG_BLACK", d.BgBlack),
		strutil.StrShellVariable("_BG_RED", d.BgRed),
		strutil.StrShellVariable("_BG_GREEN", d.BgGreen),
		strutil.StrShellVariable("_BG_YELLOW", d.BgYellow),
		strutil.StrShellVariable("_BG_BLUE", d.BgBlue),
		strutil.StrShellVariable("_BG_MAGENTA", d.BgMagenta),
		strutil.StrShellVariable("_BG_CYAN", d.BgCyan),
		strutil.StrShellVariable("_BG_WHITE", d.BgWhite),
		strutil.StrShellVariable("_NO_UNDERLINE", d.NoUnderline),
		strutil.StrShellVariable("_NO_INVERSE", d.NoInverse),
		strutil.StrShellVariable("_NO_COLOR", d.NoColor),
		strutil.StrShellVariable("_ZARUBA_ICON", d.ZarubaIcon),
		strutil.StrShellVariable("_SUCCESS_ICON", d.SuccessIcon),
		strutil.StrShellVariable("_ERROR_ICON", d.ErrorIcon),
		strutil.StrShellVariable("_START_ICON", d.StartIcon),
		strutil.StrShellVariable("_KILL_ICON", d.KillIcon),
		strutil.StrShellVariable("_INSPECT_ICON", d.InspectIcon),
		strutil.StrShellVariable("_RUN_ICON", d.RunIcon),
		strutil.StrShellVariable("_WORKER_ICON", d.WorkerIcon),
		strutil.StrShellVariable("_SCRIPT_ICON", d.ScriptIcon),
		strutil.StrShellVariable("_CONSTRUCTION_ICON", d.ConstructionIcon),
		strutil.StrShellVariable("_CONTAINER_ICON", d.ContainerIcon),
		strutil.StrShellVariable("_EMPTY", d.EmptyIcon),
	}
	return strings.Join(shellVariableList, ";")
}
