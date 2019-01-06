package logger

import "github.com/gxlog/gxlog/iface"

// The Flag defines the flag type of Logger.
type Flag int

// All available flags of Logger here. If a flag is set, its corresponding
// feature is enabled.
const (
	Prefix Flag = 0x01 << iota
	Contexts
	DynamicContexts // to be effective, the flag Contexts must also be set
	Mark
	Limit // corresponding to limiting by count and limiting by time
)

// The Filter type defines a function type which is as the log filter.
// Do NOT call methods of the Logger within a filter, or it will deadlock.
type Filter func(*iface.Record) bool

// A Config is used to configure a Logger.
// A Config should be created with NewConfig.
type Config struct {
	// Level is the level of Logger, logs with lower level will be omitted.
	Level iface.Level
	// TimeLevel is the level of a log when Time or Timef is called.
	TimeLevel iface.Level
	// PanicLevel is the level of a log when Panic or Panicf is called.
	PanicLevel iface.Level
	// TrackLevel is the auto backtracking level of Logger.
	// If the level of a log is NOT lower than the TrackLevel, the stack of
	// the current goroutine will also be output.
	TrackLevel iface.Level
	// ExitLevel is the auto exiting level of Logger.
	// If the level of a log is NOT lower than the ExitLevel, the Logger
	// will call os.Exit after outputs log.
	ExitLevel iface.Level
	// Filter is the log filter of Logger. If it is not nil and it returns
	// false, the log will be omitted.
	Filter Filter
	Flags  Flag
}

// NewConfig creates a new Config. By default, the Level is iface.Trace, the
// TimeLevel is iface.Trace, the PanicLevel is iface.Fatal, the TrackLevel is
// iface.Fatal, the ExitLevel is iface.Off, the Filter is nil and all flags are
// set.
func NewConfig() *Config {
	return &Config{
		Level:      iface.Trace,
		TimeLevel:  iface.Trace,
		PanicLevel: iface.Fatal,
		TrackLevel: iface.Fatal,
		ExitLevel:  iface.Off,
		Flags:      Prefix | Contexts | DynamicContexts | Mark | Limit,
	}
}

// WithLevel sets the Level of the Config and returns the Config.
func (cfg *Config) WithLevel(level iface.Level) *Config {
	cfg.Level = level
	return cfg
}

// WithTimeLevel sets the TimeLevel of the Config and returns the Config.
func (cfg *Config) WithTimeLevel(level iface.Level) *Config {
	cfg.TimeLevel = level
	return cfg
}

// WithPanicLevel sets the PanicLevel of the Config and returns the Config.
func (cfg *Config) WithPanicLevel(level iface.Level) *Config {
	cfg.PanicLevel = level
	return cfg
}

// WithTrackLevel sets the TrackLevel of the Config and returns the Config.
func (cfg *Config) WithTrackLevel(level iface.Level) *Config {
	cfg.TrackLevel = level
	return cfg
}

// WithExitLevel sets the ExitLevel of the Config and returns the Config.
func (cfg *Config) WithExitLevel(level iface.Level) *Config {
	cfg.ExitLevel = level
	return cfg
}

// WithFilter sets the Filter of the Config and returns the Config.
func (cfg *Config) WithFilter(filter Filter) *Config {
	cfg.Filter = filter
	return cfg
}

// WithFlags sets the Flags of the Config and returns the Config.
func (cfg *Config) WithFlags(flags Flag) *Config {
	cfg.Flags = flags
	return cfg
}

// WithEnabled enables flags of the Config and returns the Config.
func (cfg *Config) WithEnabled(flags Flag) *Config {
	cfg.Flags |= flags
	return cfg
}

// WithDisabled disables flags of the Config and returns the Config.
func (cfg *Config) WithDisabled(flags Flag) *Config {
	cfg.Flags &^= flags
	return cfg
}
