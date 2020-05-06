package go_interface_logger

import "testing"

func TestLoggerImpl_ErrorF(t *testing.T) {
	DefaultLogger.ErrorF("%s", "yetrhsfh")
	DefaultLogger.ErrorWithStackF("%s", "yetrhsfh")
}
