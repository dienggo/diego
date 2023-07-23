// Package logger
package logger

import (
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

const (
	// FunctionKey holds the function field
	FunctionKey = "function"

	// PackageKey holds the package field
	PackageKey = "package"

	// LineKey holds the line field
	LineKey = "line"

	// FileKey holds the file field
	FileKey = "file"

	// MessageKey holds the message field
	MessageKey = "message"
)

const (
	maximumCallerDepth int = 25
	knownLogrusFrames  int = 4
)

var (
	// qualified package name, cached at first use
	logrusPackage string

	// Positions in the call stack when tracing to report the calling method
	minimumCallerDepth int

	// Used for caller information initialisation
	callerInitOnce sync.Once
)

// Formatter decorates log entries with function name and package name (optional) and line number (optional)
type Formatter struct {
	ChildFormatter logrus.Formatter
	// When true, line number will be tagged to fields as well
	Line bool
	// When true, package name will be tagged to fields as well
	Package bool
	// When true, file name will be tagged to fields as well
	File bool
	// When true, only base name of the file will be tagged to fields
	BaseNameOnly bool
}

// Format the current log entry by adding the function name and line number of the caller.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	function, file, line := f.getCurrentPosition(entry)

	packageEnd := strings.LastIndex(function, ".")
	functionName := function[packageEnd+1:]

	data := logrus.Fields{FunctionKey: functionName}
	if f.Line {
		data[LineKey] = line
	}
	if f.Package {
		packageName := function[:packageEnd]
		data[PackageKey] = packageName
	}
	if f.File {
		if f.BaseNameOnly {
			data[FileKey] = filepath.Base(file)
		} else {
			data[FileKey] = file
		}
	}
	for k, v := range entry.Data {
		data[k] = v
	}
	entry.Data = data

	return f.ChildFormatter.Format(entry)
}

func (f *Formatter) getCurrentPosition(entry *logrus.Entry) (string, string, string) {
	fr := getCaller()

	if nil == fr {
		return "", "", ""
	}

	return fr.Function, fr.File, strconv.Itoa(fr.Line)
}

func getCaller() *runtime.Frame {
	//cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, maximumCallerDepth)
		cl := runtime.Callers(5, pcs)

		// dynamic get the package name and the minimum caller depth
		for i := 0; i < cl; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "sirupsen/logrus") {
				logrusPackage = getPackageName(funcName)
				break
			}
		}

		minimumCallerDepth = knownLogrusFrames
	})

	lpkg := `github.com/oni-kit/oni-rest-engine-go/src/pkg/logger`

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if !(pkg == logrusPackage || pkg == lpkg) {
			return &f //nolint:scopelint
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}

// getPackageName reduces a fully qualified function name to the package name
// There really ought to be to be a better way...
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}
