package common

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"reflect"
	"strings"
	"time"
)

var Log *zap.SugaredLogger
var lumberJackLogger *lumberjack.Logger

func SetupLogger() {
	lumberJackLogger = &lumberjack.Logger{
		Filename:  LoggerConfig.LogDir,   //日志文件的位置
		MaxSize:   LoggerConfig.MaxSize,  //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxAge:    LoggerConfig.MaxAge,   //保留旧文件的最大天数
		Compress:  LoggerConfig.Compress, //是否压缩/归档旧文件
		LocalTime: true,
	}
	ws := zapcore.AddSync(lumberJackLogger)
	if LoggerConfig.Stdout {
		ws = zap.CombineWriteSyncers(ws, os.Stdout)
	}
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.FunctionKey = "F"
	cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(LocateTimeFormat))
	}
	encoder := zapcore.NewConsoleEncoder(cfg)

	l, err := zapcore.ParseLevel(LoggerConfig.LogLevel)
	if err != nil {
		panic(err)
	}

	core := zapcore.NewCore(encoder, ws, l)
	Log = zap.New(core, zap.AddCaller()).Sugar()
	Log.Infof("setup logger ok, level: %s", l)
}

func ShutdownLogger() {
	lumberJackLogger.Rotate()
}

func J(value interface{}) string {
	res, err := json.Marshal(value)
	if err != nil {
		Log.Warnf("J err: %v", err)
	}
	return string(res)
}

func P(value interface{}) string {
	return print(reflect.ValueOf(value), 1)
}

func petty(level int) string {

	//res := "\n"
	//for i := 0; i <= level; i++ {
	//	res += "  "
	//}
	//return res

	return "  "

}

func print(value reflect.Value, level int) string {
	v := reflect.Indirect(value)
	switch v.Kind() {
	case reflect.Invalid:
		return "<nil>"
	case reflect.Bool:
		return cast.ToString(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return cast.ToString(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return cast.ToString(v.Uint())
	case reflect.Float32, reflect.Float64:
		return cast.ToString(v.Float())
	case reflect.String:
		str := v.String()
		if len(str) > 100 {
			return str[:100] + "..."
		}
		return str
	case reflect.Array, reflect.Slice:
		var res []string
		for i := 0; i < v.Len(); i++ {
			res = append(res, print(v.Index(i), level+1))
		}
		return fmt.Sprintf("[%s%s%s]", petty(level), strings.Join(res, petty(level)), petty(level-1))
	case reflect.Map:
		var res []string
		iter := v.MapRange()
		for iter.Next() {
			res = append(res, iter.Key().String()+": "+print(iter.Value(), level+1))
		}
		return fmt.Sprintf("{%s%s%s}", petty(level), strings.Join(res, petty(level)), petty(level-1))
	case reflect.Struct:
		var res []string
		for i := 0; i < v.NumField(); i++ {
			res = append(res, v.Type().Field(i).Name+": "+print(v.Field(i), level+1))
		}
		return fmt.Sprintf("%s({%s%s%s})", v.Type().String(), petty(level), strings.Join(res, petty(level)), petty(level-1))
	case reflect.Interface:
		return print(v.Elem(), level)
	}
	return "?" + v.Type().String() + "?"
}
