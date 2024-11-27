package utils

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type LogMessage struct {
	Env      string
	Host     string
	Service  string
	Port     string
	Userid   string
	Uuid     string
	Start    time.Time
	Elapsed  time.Duration
	Request  string
	Method   string
	Function string
	Desc     string
	Type     string
}

var AppLog *LoggerClass

type LoggerClass struct {
	Log  *logrus.Logger
	File *os.File
	Sync sync.Mutex
}

func LogSetUp(service, port, level, suffix string) *LoggerClass {
	var logLevel logrus.Level

	/* verifico si existe el directorio syslogs */
	if _, err := os.Stat("./syslogs"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err = os.Mkdir("syslogs", 0755)
		if err != nil {
			log.Panic("no se pudo crear la carpeta de logs...")
			return nil
		}
	}

	hoy := GetDate()
	nombreArchivo := fmt.Sprintf("./syslogs/%s_%s_%d_%d_%d%s.log", service, port, hoy.Year(), hoy.Month(), hoy.Day(), suffix)
	fmt.Printf("%s-Abriendo archivo de logs en:%s", service, nombreArchivo)
	AppLog = &LoggerClass{}
	AppLog.Log = logrus.New()

	file, err := os.OpenFile(
		nombreArchivo,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		AppLog.Log.Fatal(err)
	}

	switch level {
	case "0":
		logLevel = logrus.PanicLevel
	case "1":
		logLevel = logrus.FatalLevel
	case "2":
		logLevel = logrus.ErrorLevel
	case "3":
		logLevel = logrus.WarnLevel
	case "4":
		logLevel = logrus.InfoLevel
	case "5":
		logLevel = logrus.DebugLevel
	case "6":
		logLevel = logrus.TraceLevel
	default:
		logLevel = logrus.ErrorLevel
	}

	AppLog.Log.SetOutput(file)
	AppLog.Log.SetFormatter(&logrus.JSONFormatter{})
	AppLog.Log.SetLevel(logLevel)
	AppLog.Log.SetNoLock()
	AppLog.File = file

	return AppLog
}

func LogFormatMessages(msg LogMessage) map[string]interface{} {

	atributos := make(map[string]interface{})

	v := reflect.ValueOf(msg)

	// we only accept structs
	if v.Kind() != reflect.Struct {
		log.Printf("Me enviaron ago que no estruct")
	} else {
		// Lo proceso
		typeOfS := v.Type()

		for i := 0; i < v.NumField(); i++ {
			// fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
			atributos[typeOfS.Field(i).Name] = v.Field(i).Interface()
		}

	}
	return atributos
}

func addElapsed(fields map[string]interface{}) map[string]interface{} {
	if _, ok := fields["Start"]; ok {
		start := fields["Start"].(time.Time)
		fields["Elapsed"] = time.Since(start)
	}
	return fields
}

func (logger *LoggerClass) WriteLogsTrace(msg string, fields map[string]interface{}) {
	// logger.Sync.Lock()
	// defer logger.Sync.Unlock()
	// debo chequer si esiste el Start para agregar el Elapsed....

	logger.Log.WithFields(addElapsed(fields)).Trace(msg)
}

func (logger *LoggerClass) WriteLogsWarn(msg string, fields map[string]interface{}) {
	// logger.Sync.Lock()
	// defer logger.Sync.Unlock()
	logger.Log.WithFields(addElapsed(fields)).Warn(msg)
}

func (logger *LoggerClass) WriteLogsWError(msg string, fields map[string]interface{}) {
	// logger.Sync.Lock()
	// defer logger.Sync.Unlock()
	logger.Log.WithFields(addElapsed(fields)).Error(msg)
}

func (logger *LoggerClass) WriteLogsInfo(msg string, fields map[string]interface{}) {
	// logger.Sync.Lock()
	// defer logger.Sync.Unlock()
	logger.Log.WithFields(addElapsed(fields)).Info(msg)
}

func (logger *LoggerClass) WriteLogsDebug(msg string, fields map[string]interface{}) {
	// logger.Sync.Lock()
	// defer logger.Sync.Unlock()
	logger.Log.WithFields(addElapsed(fields)).Debug(msg)
}
