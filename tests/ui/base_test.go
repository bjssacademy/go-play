package tests

import (
	"os"
	"testing"

	"github.com/bjssacademy/go-play/helpers"
	pw "github.com/playwright-community/playwright-go"
)

var (
    pwHelper *helpers.PlaywrightHelper
)

func init() {
    helpers.InstallPlaywright()
}

//Why use TestMain? https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc
func TestMain(m *testing.M) {
    
    //launches a chromium browser for playwright
    pwHelper = helpers.NewPlaywrightHelper("chromium", pw.BrowserTypeLaunchOptions{
        Headless: pw.Bool(false), //set to true for headless mode
    })

    code := m.Run()

    pwHelper.Close()

    os.Exit(code)
}

func GetPage() pw.Page {
    return pwHelper.NewPage()
}
