package logrus

import (
	"errors"
	"testing"
)

func TestMixFormat(t *testing.T) {
	SetFormatter(&MixFormatter{TimestampFormat: "2006-01-02 15:04:05"})

	WithField("hello", "world").WithFields(Fields{"foo": "bar", "A": "B", "123": "321", "err": errors.New("This is an error")}).Info("Info信息")
}
