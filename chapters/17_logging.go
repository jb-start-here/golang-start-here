//go:build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
)

func printLine() {
	fmt.Println("--------------------------------------------------------")
}

func beforeSlog() {
	defaultLogger := log.Default() // Returns the standards logger from go log package with
	// opinionated default settings from the devs

	defaultLogger.Println("This is a log message")  // 2024/08/26 22:46:46 This is a log message
	// as opposed to fmt.Println, log.Println adds a timestamp to the log message which is a minimum requirement for a log message

	// you can customize the timestamp formatting a little bit
	// this is flag bitwise OR operation using the pipe operator
	defaultLogger.SetFlags(log.Ltime | log.Lshortfile)
	defaultLogger.Println("This is a log message")
	// 22:53:27 17_logging.go:22: This is a log message

	// There are shorthand methods you can use to direclty log methods without getting the default logger
	// This is just syntactic sugar. All methods on the default logger can be invoked directly from log namespace
	// and it will be delegated to the default logger obtained by log.Default() method.
	// since we changed the default logger setting above, those will still carry over here as well
	log.Println("This is a log message") // 23:02:14 17_logging.go:30: This is a log message
	// this simply calls the Default() method and then calls Println on the returned logger under the hood

	// you can also set a default prefix for all logs
	log.SetPrefix("[1234]") // this will also delegate the Setter method to the underlying default logger. 
	// same as log.Default().SetPrefix("[1234]")

	log.Println("This is a log message") // [1234]23:07:34 17_logging.go:39: This is a log message

	// you can also set flags for the default logger
	log.SetFlags(log.Lmsgprefix)
	log.Println("This is a log message") // [1234]This is a log message

	// you can also set the output for the default logger - it takes in any implementation of io.Writer
	// by default, it writes to os.Stderr
	// you can set it to a file, or a buffer or any other writer

	// Lets set it to a null writer from io package
	log.SetOutput(io.Discard)
	log.Println("This is a discarded log message") // wont print anything

	// you can also create many custom loggers with different set of settings using log.New()
	logger1 := log.New(os.Stderr, "logger1: ", log.Ldate | log.Lshortfile)
	logger2 := log.New(os.Stderr, "logger2: ", log.Ltime | log.Lmicroseconds)

	logger1.Println("operation successful") // logger1: 2024/08/26 17_logging.go:59: operation successful
	logger2.Println("operation failed") // logger2: 23:19:05.519638 operation failed

	// calls to log.Println still go to the default logger however
	log.Println("This is a default log message") // wont print anything since we set the output to io.Discard

	// unless we change the default logger output back to os.Stderr
	log.SetOutput(os.Stderr)
	// We cannot set the default logger to a custom logger created using log.New()
	log.Println("This is a the final log message") // [1234]This is a the final log message
}

// Why log is not good enough?
//     Unstructured logging
// 	   Limited customization
// 	   No built-in log levels
// 	   Slow performance
// 	   No custom formatting/ JSON/Text etc
// 	 These five points encapsulate the main reasons why developers often choose third-party 
//   logging libraries or the newer log/slog package for production environments.


// Lets explore slog package
func basicSlogging() {
	// order : DEBUG, INFO, WARN, ERROR
	// default log level is INFO
	// Whatever level you set, that and all the severity levels above that will be logged
	slog.SetLogLoggerLevel(slog.LevelWarn) // only WARN and ERROR will be logged
	slog.Debug("This is a debug log message") // will not be logged
	slog.Info("This is a info log message") // will not be logged
	slog.Warn("This is a warn log message")
	slog.Error("This is a error log message")

	printLine()

	slog.SetLogLoggerLevel(slog.LevelDebug)  // DEBUG, INFO, WARN, ERROR will be logged
	slog.Debug("This is a debug log message")
	slog.Info("This is a info log message")
	slog.Warn("This is a warn log message")
	slog.Error("This is a error log message")

	printLine()

	// Like with log slog.Debug is syntactic sugar for slog.Default().Debug
	// It calls this under the hood
	slog.Default().Debug("This is a debug log message without using syntactic sugar")

	// slog package works in conjunction to log package. It is a wrapper around log package.
	// it inherits logs default logger's flags and settings for its default logger 
	log.SetFlags(log.Lshortfile)
	slog.Default().Info("This is a info log message")

	// Infact slog.Info() calls slog.Default().Info() which in turns calls log.Default().Println()
	// slog is responsible for formatting and filtering levels before finally utilizing log package to log the message
}

func slogfields() {
	// one featiure of slog is that you can pass in multiple key-value pairs as arguments to log additionally
	// as attributes of your log message

	slog.Info("This is a info log message", slog.Attr{Key: "key1", Value: slog.StringValue("value1")})
	// or alternatively as a short hand you can do
	slog.Info("This is a info log message", slog.String("key1", "value1"))
	// or even shorter hand you cana do
	slog.Info("This is a info log message", "key1", "value1",)

	// All three will print
	// 	2024/08/27 10:49:13 INFO This is a info log message key1=value1

	// You can also pass in arbitrary number of key-value pairs
	slog.Info("This is a info log message", "key1", "value1", "key2", "value2")
	// 2024/08/27 10:50:33 INFO This is a info log message key1=value1 key2=value2

	// you can also group attributes together
	slog.Info(
		"grouped log msg",
		slog.String("key1", "value1"),
		slog.Group("group1", slog.String("nestedKey1", "value3"), slog.String("nestedKey2", "value4")),
		slog.Int("key2", 56),
	)
	// 2024/08/27 10:58:02 INFO grouped log msg key1=value1 group1.nestedKey1=value3 group1.nestedKey2=value4 key2=56

	// if you want all logs for a certain code block to have the same key value pairs you dont have to repeat yourself
	// create custom child loggers and use them instead

	child1 := slog.With("key1", "value1", "key2", "value2")
	child2 := slog.With("key3", "value3", "key4", "value4")
	child3 := slog.With(slog.Group("grouped", "count", 52, "version", "1.0.0"))

	child1.Info("This is a child 1 info log message")
	// 2024/08/27 11:01:07 INFO This is a child 1 info log message key1=value1 key2=value2
	child2.Info("This is a child 2 info log message")
	// 2024/08/27 11:01:07 INFO This is a child 2 info log message key3=value3 key4=value4
	child3.Info("This is a child 3 info log message")
	// 2024/08/27 11:02:17 INFO This is a child 3 info log message grouped.count=52 grouped.version=1.0.0
}

type User struct {
	Name string
	Age  int
}
func (u User) LogValue() slog.Value {
	return slog.StringValue(fmt.Sprintf("%s is %d years old", u.Name, u.Age))
}

func logValuerInterfaceExample() {
	u := User{Name: "John", Age: 25}
	slog.Info("This is a info log message", "user", u,)
	// [1234]17_logging.go:166: INFO This is a info log message user="John is 25 years old"
}

func slogHandlers() {
	// each slog logger is associated with a handler that does the formatting for the log message 
	// and all the accomapnying key-value pairs before sending it to the log package for actual printing
	// you can create custom handlers to format the log message as you see fit
	// By default slog uses a text handler instance with dev opinionated sane defaults 
	// which formats the log message as a text message 
	// and it also provides a JSON handler which formats the log message as a JSON message

	// Text or json is the most common formats for logging, so they are provided out of the box, 
	// but you can create your own custom handler if you want to log in a different format or if
	// you want to log to a different output like a file or a buffer

	// Lets say we want to use text handler with different settings, just use a new instance of text handler
	// with different constructor args

	customTextHandler := slog.NewTextHandler(
		os.Stderr, // lets write to stderr for now
		&slog.HandlerOptions{ // lets use all available customizers
			Level: slog.LevelWarn, // only log warn and above
			AddSource: true, // add source file and line number
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr { // replacer func, can be used to redact passwords
				if a.Key == "password" {
					return slog.String("password", "********")
				}
				return a
			},
		},
	)
	customLogger := slog.New(customTextHandler)

	customLogger.Warn("This is a warn log message", "key1", "value1", "password", "123456")
	// time=2024-08-27T11:23:34.061-07:00 level=WARN source=17_logging.go:186 msg="This is a iwarnnfo log message" key1=value1 password=********

	// The default logger is left untouched however
	slog.Warn("This is a warn log message", "key1", "value1", "password", "123456")
	// 2024/08/27 11:24:18 WARN This is a warn log message key1=value1 password=123456

	// unless you set the default logger to the custom logger - this also changes the logs default  logger as well.

	originalDefault := slog.Default()
	slog.SetDefault(customLogger)
	slog.Error("This is a err log message", "key1", "value1", "password", "123456")
	// time=2024-08-27T11:41:17.309-07:00 level=ERROR source=17_logging.go:197 msg="This is a err log message" key1=value1 password=********

	// lets revert back to the original default logger
	slog.SetDefault(originalDefault)

	// All the custom attribute handling methods and child loggers we saw earlier can be used with custom handler powered logger as well

	// You can also specify custom attributes for your custom logger that you want displayed for everybody
	// and force any and all attributed added later to be included in that group

	// You can use with attrs to add default attributes to the logger

	handler := slog.NewJSONHandler(os.Stdout, nil) // lets try a JSON handler
	customGroupedLogger := slog.New(handler).WithGroup("program_info").With("main", "default")
	
	child := customGroupedLogger.With(
		slog.Int("pid", os.Getpid()),
		slog.String("go_version", "vx.xx"),
	)
	
	child.Warn(
		"storage is 90% full",
		slog.String("available_space", "900.1 MB"),
	)
	// {
	// 	"time":"2024-08-27T11:49:57.238631-07:00",
	// 	"level":"WARN",
	// 	"msg":"storage is 90% full",
	// 	"program_info":{
	//    "main":"default",
	// 		"pid":46038,
	// 		"go_version":"vx.xx",
	// 		"available_space":"900.1 MB"
	// 	}
	// }
}



func main() {
	beforeSlog()
	basicSlogging()
	
	slogfields()
	logValuerInterfaceExample()
	slogHandlers()
}
