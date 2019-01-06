package text

import "github.com/gxlog/gxlog/iface"

// All predefined headers here.
const (
	FullHeader = "{{time}} {{level}} {{file}}:{{line}} {{pkg}}.{{func}} " +
		"{{prefix}}[{{context}}] {{msg}}\n"
	CompactHeader = "{{time:time.us}} {{level}} {{file:1}}:{{line}} {{pkg}}.{{func}} " +
		"{{prefix}}[{{context}}] {{msg}}\n"
	SyslogHeader = "{{file:1}}:{{line}} {{pkg}}.{{func}} {{prefix}}[{{context}}] {{msg}}\n"
)

// A Config is used to configure a text formatter.
// A Config should be created with NewConfig.
type Config struct {
	// Header is the format specifier of a text formatter.
	// It is used to specify which and how the fields of a Record to be formatted.
	// The pattern of a field specifier is {{<name>[:property][%fmtstr]}}.
	// e.g. {{level:char}}, {{line%05d}}, {{pkg:1}}, {{context:list%40s}}
	// All fields have support for the fmtstr. If the fmtstr is not the default
	// one of a field, it will be passed to fmt.Sprintf to format the field and
	// this affects the performance a little.
	// The supported properties vary with fields.
	// All supported fields are as the follows:
	//            supported property          defaults      property examples
	// --------------------------------------------------------------------------
	//   time     <date|time>[.ms|.us|.ns]    date.us %s    date.ns, time
	//            layout that is supported                  time.RFC3339Nano
	//              by the time package                     02 Jan 06 15:04 -0700
	//   level    <full|char>                 full    %s    full, char
	//   file     <lastSegs>                  0       %s    0, 1, 2, ...
	//   line                                         %d
	//   pkg      <lastSegs>                  0       %s    0, 1, 2, ...
	//   func     <lastSegs>                  0       %s    0, 1, 2, ...
	//   prefix                                       %s
	//   context  <pair|list>                 pair    %s    pair, list
	//   msg                                          %s
	Header string
	// MinBufSize is the initial size of the internal buf of a formatter.
	// MinBufSize must not be negative.
	MinBufSize int
	// ColorMap is used to remap the color of each level.
	// The color of a level is left to be unchanged if it is not in the map.
	// By default, the color of level iface.Trace, iface.Debug and iface.Info is
	// Green, the color of level iface.Warn is Yellow, the color of level
	// iface.Error and iface.Fatal is Red and the color of a marked log is
	// Magenta no matter at which level it is.
	ColorMap map[iface.Level]Color
	// EnableColor enables colorization if it is true.
	EnableColor bool
}

// NewConfig creates a new Config. By default, the Header is the FullHeader,
// the MinBufSize is 256, the ColorMap is nil, the EnableColor is false.
func NewConfig() *Config {
	return &Config{
		Header:     FullHeader,
		MinBufSize: 256,
	}
}

// WithHeader sets the Header of the Config and returns the Config.
func (cfg *Config) WithHeader(header string) *Config {
	cfg.Header = header
	return cfg
}

// WithMinBufSize sets the MinBufSize of the Config and returns the Config.
func (cfg *Config) WithMinBufSize(size int) *Config {
	cfg.MinBufSize = size
	return cfg
}

// WithColorMap sets the ColorMap of the Config and returns the Config.
func (cfg *Config) WithColorMap(colorMap map[iface.Level]Color) *Config {
	cfg.ColorMap = colorMap
	return cfg
}

// WithEnableColor sets the EnableColor of the Config and returns the Config.
func (cfg *Config) WithEnableColor(enable bool) *Config {
	cfg.EnableColor = enable
	return cfg
}
