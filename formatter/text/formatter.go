// Package text implements a text formatter which implements the Formatter.
package text

import (
	"regexp"
	"strings"
	"sync"

	"github.com/gxlog/gxlog/iface"
)

var headerRegexp = regexp.MustCompile("{{([^:%]*?)(?::([^%]*?))?(%.*?)?}}")

// A Formatter implements the interface iface.Formatter.
//
// All methods of a Formatter are concurrency safe.
// A Formatter MUST be created with New.
type Formatter struct {
	header      string
	minBufSize  int
	enableColor bool

	colorMgr  *colorMgr
	appenders []*headerAppender
	suffix    string

	lock sync.Mutex
}

// New creates a new Formatter with the config.
func New(config Config) *Formatter {
	config.setDefaults()
	formatter := &Formatter{
		minBufSize:  config.MinBufSize,
		enableColor: config.EnableColor,
		colorMgr:    newColorMgr(),
	}
	formatter.SetHeader(config.Header)
	formatter.MapColors(config.ColorMap)
	return formatter
}

// Header returns the header of the Formatter.
func (formatter *Formatter) Header() string {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	return formatter.header
}

// SetHeader sets the header of the Formatter.
// For details of all supported fields in a header, see the comment of Config.
func (formatter *Formatter) SetHeader(header string) {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	formatter.header = header
	formatter.appenders = formatter.appenders[:0]
	var staticText string
	for header != "" {
		indexes := headerRegexp.FindStringSubmatchIndex(header)
		if indexes == nil {
			break
		}
		begin, end := indexes[0], indexes[1]
		staticText += header[:begin]
		element, property, fmtspec := extractElement(indexes, header)
		if formatter.addAppender(element, property, fmtspec, staticText) {
			staticText = ""
		}
		header = header[end:]
	}
	formatter.suffix = staticText + header
}

// MinBufSize returns the min buf size of the Formatter.
func (formatter *Formatter) MinBufSize() int {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	return formatter.minBufSize
}

// SetMinBufSize sets the min buf size of the Formatter.
// The size must NOT be negative. If the size is 0, 256 is used.
func (formatter *Formatter) SetMinBufSize(size int) {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	if size == 0 {
		formatter.minBufSize = 256
	} else {
		formatter.minBufSize = size
	}
}

// ColorEnabled returns whether colorization is enabled in the Formatter.
func (formatter *Formatter) ColorEnabled() bool {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	return formatter.enableColor
}

// EnableColor enables colorization in the Formatter.
func (formatter *Formatter) EnableColor() {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	formatter.enableColor = true
}

// DisableColor disables colorization in the Formatter.
func (formatter *Formatter) DisableColor() {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	formatter.enableColor = false
}

// Color returns the color of the level in the Formatter.
func (formatter *Formatter) Color(level iface.Level) Color {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	return formatter.colorMgr.Color(level)
}

// SetColor sets the color of the level in the Formatter.
func (formatter *Formatter) SetColor(level iface.Level, color Color) {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	formatter.colorMgr.SetColor(level, color)
}

// MapColors maps the color of levels in the Formatter according to the colorMap.
// The color of a level is left to be unchanged if it is not in the map.
func (formatter *Formatter) MapColors(colorMap map[iface.Level]Color) {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	formatter.colorMgr.MapColors(colorMap)
}

// MarkedColor returns the color of a log that is marked.
func (formatter *Formatter) MarkedColor() Color {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	return formatter.colorMgr.MarkedColor()
}

// SetMarkedColor sets the color of a log that is marked.
func (formatter *Formatter) SetMarkedColor(color Color) {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	formatter.colorMgr.SetMarkedColor(color)
}

// Format implements the interface Formatter. It formats a Record.
func (formatter *Formatter) Format(record *iface.Record) []byte {
	formatter.lock.Lock()
	defer formatter.lock.Unlock()

	var left, right []byte
	if formatter.enableColor {
		if record.Aux.Marked {
			left, right = formatter.colorMgr.MarkedColorEars()
		} else {
			left, right = formatter.colorMgr.ColorEars(record.Level)
		}
	}

	buf := make([]byte, 0, formatter.minBufSize)
	buf = append(buf, left...)
	for _, appender := range formatter.appenders {
		buf = appender.AppendHeader(buf, record)
	}
	buf = append(buf, formatter.suffix...)
	buf = append(buf, right...)

	return buf
}

func (formatter *Formatter) addAppender(element, property, fmtspec,
	staticText string) bool {
	appender := newHeaderAppender(element, property, fmtspec, staticText)
	if appender == nil {
		return false
	}
	formatter.appenders = append(formatter.appenders, appender)
	return true
}

func extractElement(indexes []int, header string) (element, property,
	fmtspec string) {
	element = strings.ToLower(getField(header, indexes[2], indexes[3]))
	property = getField(header, indexes[4], indexes[5])
	fmtspec = getField(header, indexes[6], indexes[7])
	if fmtspec == "%" {
		fmtspec = ""
	}
	return element, property, fmtspec
}

func getField(header string, begin, end int) string {
	if begin < end {
		return strings.TrimSpace(header[begin:end])
	}
	return ""
}
