package instapaper

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

type Api interface {
	GetUnreadBookmarks() ([]Bookmark, error)
}

type api struct {
	logger         *logrus.Logger
	consumerKey    string
	consumerSecret string
	accessToken    string
	accessSecret   string
}

func init() {
	logger = logrus.New()

	logger.Formatter = &logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02T15:04:05",
		FullTimestamp:   true,
	}

	logger.Out = os.Stdout
	logger.SetReportCaller(true)
}

func New() *api {

	return &api{
		logger: logger,
	}
}
