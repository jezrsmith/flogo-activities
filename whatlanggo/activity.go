package whatlanggo

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/abadojack/whatlanggo"
)

// log is the default package logger
var log = logger.GetLogger("activity-whatlango")

const (
	ivPhrase   = "phrase"
	ivLanguage = "language"
	ivScript   = "script"

	ovValue = "value"
)

// WhatlanggoActivity is a whatlanggo activity implementation
type WhatlanggoActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new WhatlanggoActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &WhatlanggoActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *WhatlanggoActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *WhatlanggoActivity) Eval(context activity.Context) (done bool, err error) {

	phrase := context.GetInput(ivPhrase).(string)

	if phrase != "" {
		log.Debugf("Checking language for phrase: ", phrase)
	} else {
		log.Debugf("Nil string?: ", phrase)
	}

	info := whatlanggo.Detect(phrase)
	log.Debugf("Language:", whatlanggo.LangToString(info.Lang), "Script:", whatlanggo.Scripts[info.Script])

	context.SetOutput(ivLanguage, whatlanggo.LangToString(info.Lang))
	context.SetOutput(ivScript, whatlanggo.Scripts[info.Script])

	return true, nil
}
