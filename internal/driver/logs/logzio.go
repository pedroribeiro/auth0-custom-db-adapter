package logs

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/bshuster-repo/logruzio"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func Init() {
	InitLogrus()

	ctx := logrus.Fields{
		"env": os.Getenv("APP_ENV"),
	}
	hook, err := logruzio.New(os.Getenv("LOGZ_IO_TOKEN"), "users", ctx)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.AddHook(hook)
}

func InitLogrus() {
	var formater = &zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	l := logrus.WithFields(logrus.Fields{})
	l.Logger.SetReportCaller(true)
	l.Logger.SetFormatter(formater)
	l.Logger.SetLevel(logrus.InfoLevel)
}
