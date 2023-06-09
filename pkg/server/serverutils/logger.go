/*
 * Copyright (c) 2022-2023 Sienna Lloyd
 *
 * Licensed under the PolyForm Internal Use License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

package serverutils

import (
	dlog "github.com/lni/dragonboat/v3/logger"
	zlog "github.com/rs/zerolog"
)

var (
	rootLogger zlog.Logger
	rootLoggerSet = false
)

func SetRootLogger(logger zlog.Logger) {
	if rootLoggerSet {
		return
	}
	rootLogger = logger
	rootLoggerSet = true
}

func DragonboatLoggerFactory(pkgName string) dlog.ILogger {
	logz := rootLogger.With().Str("pkg", pkgName).Logger()
	return DragonboatLoggerAdapter{
		logger: logz,
	}
}

var _ dlog.ILogger = DragonboatLoggerAdapter{}

type DragonboatLoggerAdapter struct {
	logger zlog.Logger
}

var internalSeverityMap = map[dlog.LogLevel]zlog.Level{
	dlog.CRITICAL: zlog.FatalLevel,
	dlog.ERROR:    zlog.ErrorLevel,
	dlog.WARNING:  zlog.WarnLevel,
	dlog.INFO:     zlog.InfoLevel,
	dlog.DEBUG:    zlog.DebugLevel,
}

var reverseInternalSeverityMap = map[zlog.Level]dlog.LogLevel{
	zlog.FatalLevel: dlog.CRITICAL,
	zlog.ErrorLevel: dlog.ERROR,
	zlog.WarnLevel:  dlog.WARNING,
	zlog.InfoLevel:  dlog.INFO,
	zlog.DebugLevel: dlog.DEBUG,
}

func (l DragonboatLoggerAdapter) SetLevel(logLevel dlog.LogLevel) {
	l.logger = l.logger.Level(internalSeverityMap[logLevel])
}

func (l DragonboatLoggerAdapter) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}

func (l DragonboatLoggerAdapter) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}

func (l DragonboatLoggerAdapter) Warningf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}

func (l DragonboatLoggerAdapter) Errorf(format string, args ...interface{}) {
	l.logger.Error().Stack().Msgf(format, args...)
}

func (l DragonboatLoggerAdapter) Panicf(format string, args ...interface{}) {
	l.logger.Panic().Msgf(format, args...)
}

