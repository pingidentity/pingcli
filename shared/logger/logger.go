package logger

import (
	"github.com/pingidentity/pingcli/internal/output"
)

type SharedLogger struct{}

func (gl SharedLogger) convertFields(fields map[string]string) map[string]interface{} {
	convertedFields := make(map[string]interface{}, len(fields))
	for k, v := range fields {
		convertedFields[k] = v
	}

	return convertedFields
}

func (gl SharedLogger) Message(message string, fields map[string]string) error {
	output.Message(message, gl.convertFields(fields))

	return nil
}

func (gl SharedLogger) Success(message string, fields map[string]string) error {
	output.Success(message, gl.convertFields(fields))

	return nil
}

func (gl SharedLogger) Warn(message string, fields map[string]string) error {
	output.Warn(message, gl.convertFields(fields))

	return nil
}

func (gl SharedLogger) UserError(message string, fields map[string]string) error {
	output.UserError(message, gl.convertFields(fields))

	return nil
}

func (gl SharedLogger) UserFatal(message string, fields map[string]string) error {
	output.UserFatal(message, gl.convertFields(fields))

	return nil
}

func (gl SharedLogger) PluginError(message string, fields map[string]string) error {
	output.PluginError(message, gl.convertFields(fields))

	return nil
}
