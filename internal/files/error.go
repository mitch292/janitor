package files

import (
	"log"
	"os"
	"path"

	"github.com/mitch292/janitor/internal/constants"
)

type internalErrorLog struct {
	location string
}

// NewErrorLog will create the internal error log file.
func NewErrorLog() *internalErrorLog {
	return &internalErrorLog{
		location: path.Join(constants.INTERNAL_DIR_NAME, constants.INTERNAL_ERROR_LOG_FILE),
	}
}

// WriteErrorToLog will write an error to the internal janitor's closet log (.janitors_closet/.janitors_error_log)
func (e *internalErrorLog) WriteErrorToLog(message string) (err error) {
	errorLogFile, err := os.OpenFile(e.location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer errorLogFile.Close()

	logger := log.New(errorLogFile, "janitor: ", log.LstdFlags)
	logger.Println(message)

	return
}
