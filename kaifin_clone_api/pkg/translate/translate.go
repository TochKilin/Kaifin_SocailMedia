package translate

import (
	// standard lirbary
	"fmt"
	"log"
	"path/filepath"
	//custom internal

	"github.com/gofiber/fiber/v3"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"

	"kaifin_clone_api/pkg/logs"
	error_response "kaifin_clone_api/pkg/responses"
)

// Global Variable store all ms
var bundle *goi18n.Bundle

// this fn load file translates
func Init() *error_response.ErrorResponse {

	//set defualt for ms is english
	bundle = goi18n.NewBundle(language.English)
	// tell i18n is file yulm
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// list of file of translate
	localeFiles := []string{
		"pkg/i18n/localize/en.yaml",
		"pkg/i18n/localize/km.yaml",
		"pkg/i18n/localize/zh.yaml",
	}

	// loop all file translate
	for _, file := range localeFiles {
		// load file to bundle
		_, err := bundle.LoadMessageFile(filepath.Join(file))
		if err != nil {
			log.Printf("Error loading local file %s: %v", file, err)
			logs.NewCustomLog("translate_error", err.Error(), "error")
			return &error_response.ErrorResponse{
				MessageID: "ErrorLoadMessage",
				Err:       err,
			}
		}
	}

	// load success show tru translate
	return nil
}

// this func is translate message and err from client reqest
func TranslateWithError(c fiber.Ctx, key string, templateData ...map[string]any) (string, *error_response.ErrorResponse) {
	if bundle == nil {
		initErr := Init()
		if initErr != nil {
			logs.NewCustomLog(
				"I18nNotInit",
				initErr.ErrorString(),
				"error",
			)
		}
		return "", &error_response.ErrorResponse{
			MessageID: key,
			Err:       fmt.Errorf("translation service is unavailable"),
		}
	}

	// bind lang from header
	lang := c.Get("Accept-Language", "en")
	//Localizer use for translate Client to wnat it
	localizer := goi18n.NewLocalizer(bundle, lang)
	data := map[string]any{}
	if len(templateData) > 0 && templateData[0] != nil {
		data = templateData[0]
	}

	msg, err := localizer.Localize(&goi18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: data,
	})
	if err != nil {
		log.Printf("Error localizing message ID %s: %v", key, err)
		logs.NewCustomLog("TranslationNotFound", err.Error(), "error")
		return "", &error_response.ErrorResponse{
			MessageID: key,
			Err:       fmt.Errorf("Translation not found"),
		}
	}
	return msg, nil
}

// call func Translatewitherror and return string
func Translate(c fiber.Ctx, key string) string {
	msg, err := TranslateWithError(c, key)
	// check error
	if err != nil {
		return key
	}
	return msg
}

// TranslateWithError() → Return (string, error) for Handle Error។
//Translate() → Return តែ string want Message very fast
