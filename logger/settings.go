package logger

import "github.com/gxlog/gxlog/iface"

// Config returns a copy of Config of the Logger.
func (log *Logger) Config() *Config {
	log.lock.Lock()
	defer log.lock.Unlock()

	copyConfig := *log.config
	return &copyConfig
}

// SetConfig sets the copy of config to the Logger. The config must NOT be nil.
// If the config is invalid, it returns an error and the config of the Logger
// is left to be unchanged.
func (log *Logger) SetConfig(config *Config) error {
	log.lock.Lock()
	defer log.lock.Unlock()

	copyConfig := *config
	log.config = &copyConfig
	return nil
}

// UpdateConfig will call fn with copy of the config of the Logger, and then
// sets copy of the returned config to the Logger. The fn must NOT be nil.
// If the returned config is invalid, it returns an error and the config of
// the Logger is left to be unchanged.
//
// Do NOT call methods of the Logger within fn, or it will deadlock.
func (log *Logger) UpdateConfig(fn func(Config) Config) error {
	log.lock.Lock()
	defer log.lock.Unlock()

	config := fn(*log.config)
	log.config = &config
	return nil
}

// Level returns the level of the Logger.
func (log *Logger) Level() iface.Level {
	log.lock.Lock()
	defer log.lock.Unlock()

	return log.config.Level
}

// SetLevel sets the level of the Logger.
func (log *Logger) SetLevel(level iface.Level) {
	log.lock.Lock()
	defer log.lock.Unlock()

	log.config.Level = level
}

// TimeLevel returns the time level of the Logger.
func (log *Logger) TimeLevel() iface.Level {
	log.lock.Lock()
	defer log.lock.Unlock()

	return log.config.TimeLevel
}

// SetTimeLevel sets the time level of the Logger.
func (log *Logger) SetTimeLevel(level iface.Level) {
	log.lock.Lock()
	defer log.lock.Unlock()

	log.config.TimeLevel = level
}

// PanicLevel returns the panic level of the Logger.
func (log *Logger) PanicLevel() iface.Level {
	log.lock.Lock()
	defer log.lock.Unlock()

	return log.config.PanicLevel
}

// SetPanicLevel sets the panic level of the Logger.
func (log *Logger) SetPanicLevel(level iface.Level) {
	log.lock.Lock()
	defer log.lock.Unlock()

	log.config.PanicLevel = level
}

// TrackLevel returns the track level of the Logger.
func (log *Logger) TrackLevel() iface.Level {
	log.lock.Lock()
	defer log.lock.Unlock()

	return log.config.TrackLevel
}

// SetTrackLevel sets the track level of the Logger.
func (log *Logger) SetTrackLevel(level iface.Level) {
	log.lock.Lock()
	defer log.lock.Unlock()

	log.config.TrackLevel = level
}

// ExitLevel returns the exit level of the Logger.
func (log *Logger) ExitLevel() iface.Level {
	log.lock.Lock()
	defer log.lock.Unlock()

	return log.config.ExitLevel
}

// SetExitLevel sets the exit level of the Logger.
func (log *Logger) SetExitLevel(level iface.Level) {
	log.lock.Lock()
	defer log.lock.Unlock()

	log.config.ExitLevel = level
}

// Filter returns the filter of the Logger.
func (log *Logger) Filter() Filter {
	log.lock.Lock()
	defer log.lock.Unlock()

	return log.config.Filter
}

// SetFilter sets the filter of the Logger.
func (log *Logger) SetFilter(filter Filter) {
	log.lock.Lock()
	defer log.lock.Unlock()

	log.config.Filter = filter
}

// Flags returns the flags of the Logger.
func (log *Logger) Flags() Flag {
	log.lock.Lock()
	defer log.lock.Unlock()

	return log.config.Flags
}

// SetFlags sets the flags of the Logger.
func (log *Logger) SetFlags(flags Flag) {
	log.lock.Lock()
	defer log.lock.Unlock()

	log.config.Flags = flags
}

// Enable enables the flags of the Logger.
func (log *Logger) Enable(flags Flag) {
	log.lock.Lock()
	defer log.lock.Unlock()

	log.config.Flags |= flags
}

// Disable disables the flags of the Logger.
func (log *Logger) Disable(flags Flag) {
	log.lock.Lock()
	defer log.lock.Unlock()

	log.config.Flags &^= flags
}
