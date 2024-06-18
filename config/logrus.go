package config

//
//var once sync.Once
//var instance *logrus.Logger
//
//// GetLogger trả về một instance duy nhất của logrus.Logger
//func GetLogger() *logrus.Logger {
//	once.Do(func() {
//		instance = logrus.New()
//		instance.SetFormatter(&logrus.JSONFormatter{})
//		instance.SetOutput(os.Stdout)       // Hoặc thiết lập output mong muốn, ví dụ file
//		instance.SetLevel(logrus.InfoLevel) // Thiết lập level log mặc định
//	})
//	return instance
//}
//
//// Các hàm helper cho log
//func Info(args ...interface{}) {
//	GetLogger().Info(args...)
//}
//
//func Infof(format string, args ...interface{}) {
//	GetLogger().Infof(format, args...)
//}
//
//func Error(args ...interface{}) {
//	GetLogger().Error(args...)
//}
//
//func Errorf(format string, args ...interface{}) {
//	GetLogger().Errorf(format, args...)
//}
