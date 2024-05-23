package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUIExample(t *testing.T) {
    
    page := GetPage()

    _, err := page.Goto("https://example.com")
    assert.NoError(t, err, "should navigate to example.com")

    title, err := page.Title()
    assert.NoError(t, err, "should get the page title")
    assert.Equal(t, "Example Domain", title, "title should be 'Example Domain'")

    //can be safely removed, just to show you how to click on elements
    time.Sleep(2 * time.Second)

    page.Locator("body > div > p:nth-child(3) > a").Click()

    time.Sleep(2 * time.Second)

}
