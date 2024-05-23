package helpers

import (
	"log"

	pw "github.com/playwright-community/playwright-go"
)

type PlaywrightHelper struct {
    Playwright *pw.Playwright
    Browser    pw.Browser
    Page       pw.Page
}

func NewPlaywrightHelper(browserName string, launchOptions pw.BrowserTypeLaunchOptions) *PlaywrightHelper {
    playwright, err := pw.Run()
    if err != nil {
        log.Fatalf("could not start Playwright: %v", err)
    }

	browserType := playwright.Chromium

	if browserName == "chromium" || browserName == "" {
		browserType = playwright.Chromium
	} else if browserName == "firefox" {
		browserType = playwright.Firefox
	} else if browserName == "webkit" {
		browserType = playwright.WebKit
	}

    browser, err := browserType.Launch(launchOptions)
    if err != nil {
        log.Fatalf("could not launch browser: %v", err)
    }
    return &PlaywrightHelper{
        Playwright: playwright,
        Browser:    browser,
    }
}

/*func NewPlaywrightHelper(browserType pw.BrowserType, launchOptions pw.BrowserTypeLaunchOptions) *PlaywrightHelper {
    playwright, err := pw.Run()
    if err != nil {
        log.Fatalf("could not start Playwright: %v", err)
    }
    browser, err := browserType.Launch(launchOptions)
    if err != nil {
        log.Fatalf("could not launch browser: %v", err)
    }
    return &PlaywrightHelper{
        Playwright: playwright,
        Browser:    browser,
    }
}*/

/*func NewPlaywrightHelper(headless bool) *PlaywrightHelper {
    playwright, err := pw.Run()
    if err != nil {
        log.Fatalf("could not start Playwright: %v", err)
    }
    browser, err := playwright.Chromium.Launch(pw.BrowserTypeLaunchOptions{
        Headless: pw.Bool(headless),
    })
    if err != nil {
        log.Fatalf("could not launch browser: %v", err)
    }
    return &PlaywrightHelper{
        Playwright: playwright,
        Browser:    browser,
    }
}*/

func (ph *PlaywrightHelper) NewPage() pw.Page {
    page, err := ph.Browser.NewPage()
    if err != nil {
        log.Fatalf("could not create new page: %v", err)
    }
    ph.Page = page
    return page
}

func (ph *PlaywrightHelper) Close() {
    if err := ph.Browser.Close(); err != nil {
        log.Fatalf("could not close browser: %v", err)
    }
    if err := ph.Playwright.Stop(); err != nil {
        log.Fatalf("could not stop Playwright: %v", err)
    }
}

func InstallPlaywright() {
    if err := pw.Install(); err != nil {
        log.Fatalf("could not install Playwright: %v", err)
    }
}
