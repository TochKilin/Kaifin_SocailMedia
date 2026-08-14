package logs

import (
	"fmt"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type CustomLog struct {
	MessageID string
	LogReason string
	Function  string
	File      string
	Line      int
}

func (e *CustomLog) LogToString() string {
	return fmt.Sprintf("MessageID: %s,LogReason:%s, Funtion: %s, File: %s, Line: %d",
		e.MessageID, e.LogReason, e.Function, e.File, e.Line)
}

func NewCustomLog(messageID string, logDesc string, logType ...string) *CustomLog {
	pc, file, line, ok := runtime.Caller(1)
	function := "unknown"
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			function = fn.Name()
		}
	}
	msg := &CustomLog{
		MessageID: messageID,
		LogReason: logDesc,
		Function:  function,
		File:      file,
		Line:      line,
	}

	switch {
	//  when Log and Application Stop soon
	case len(logType) > 0 && logType[0] == "fatal":
		log.Fatal().Str("MessageID", msg.MessageID).
			Str("LogReason", "logged : `"+msg.Function+"()`").
			Str("Function", msg.Function).
			Str("File", msg.File).
			Int("Line", msg.Line).
			Msg(msg.LogReason)

	// log.Error() // system stay run ing
	case len(logType) > 0 && logType[0] == "error":
		log.Error().Str("MessageID", msg.MessageID).
			Str("LogReason", "Logged : `"+msg.Function+"()`").
			Str("Function", msg.Function).
			Str("File", msg.File).
			Int("Line", msg.Line).
			Msg(msg.LogReason)
	// log.Warn() //
	case len(logType) > 0 && logType[0] == "warn":
		log.Warn().Str("MessageID", msg.MessageID).
			Str("LogReason", "Logged : `"+msg.Function+"()`").
			Str("Function", msg.Function).
			Str("File", msg.File).
			Int("Line", msg.Line).
			Msg(msg.LogReason)

		// 	logs.NewCustomLog(
		// 	"UserNotFound",
		// 	"user not found",
		// 	"warn",
		// )

	// log.Info() // simple log INFO INFO
	case len(logType) > 0 && logType[0] == "info":
		log.Info().Str("MessageID", msg.MessageID).
			Str("LogReason", "Logged : `"+msg.Function+"()`").
			Str("Function", msg.Function).
			Str("File", msg.File).
			Int("Line", msg.Line).
			Msg(msg.LogReason)

		// 	logs.NewCustomLog(
		// 	"LoginSuccess",
		// 	"user login success",
		// 	"info",
		// )

	// log.Debug() // show tracking process
	case len(logType) > 0 && logType[0] == "debug":
		log.Debug().Str("MessageID", msg.MessageID).
			Str("LogReason", "Logged : `"+msg.Function+"()`").
			Str("Function", msg.Function).
			Str("File", msg.File).
			Int("Line", msg.Line).
			Msg(msg.LogReason)

		// logs.NewCustomLog(
		// 	"DBFatal",
		// 	"database unavailable",
		// 	"fatal",
		// )
	default:
		if zerolog.GlobalLevel() >= zerolog.InfoLevel {
			log.Info().Str("MessageID", msg.MessageID).
				Str("LogReason", "Logged : `"+msg.Function+"()`").
				Str("Function", msg.Function).
				Str("File", msg.File).
				Int("Line", msg.Line).
				Msg(msg.LogReason)
		}
	}

	// 	logs.NewCustomLog(
	// 	"UserCreated",
	// 	"user created",
	// )
	// send log json to other server
	return msg
}

//this customlog for show error at where  function , file  , lin
