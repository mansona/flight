package figma

import (
	"errors"
	"flight/internal/core"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/antchfx/jsonquery"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

var (
	documentId            = "TLnoT5AYQfy3tZ0H68BgOr" // TODO: Make configurable by environment
	header     req.Header = nil
)

// Finds all nodes that represent icons in a given Figma document
func FindIconsInDoc() ([]core.Icon, error) {
	log.Println("Querying Figma file")

	var icons []core.Icon

	// Setup our Figma API headers
	header = req.Header{
		"Accept":        "application/json",
		"X-FIGMA-TOKEN": os.Getenv("FIGMA_TOKEN"),
	}

	// Create the URL to call
	url := fmt.Sprintf("https://api.figma.com/v1/files/%s", documentId)

	// We need a reasonable depth returned from our Figma document
	param := req.Param{
		"depth":    "10",
		"geometry": "paths",
	}

	// Call the API
	r, err := req.Get(url, header, param)

	if err != nil {
		log.Fatal(err)
	}

	if r.Response().StatusCode != 200 {
		return nil, errors.New("could not find icons in figma document")
	}

	// Parse the json response into jsonquery which allows xpath-like node searching
	doc, err := jsonquery.Parse(strings.NewReader(r.String()))

	if err != nil {
		log.Fatal(err)
	}

	// Look for our icons...
	for _, n := range jsonquery.Find(doc, "//*[type='COMPONENT_SET']/children/*") {
		// Find the size of the size
		iconSize := strings.Split(n.SelectElement("name").InnerText(), "=")

		// Find the name from the parent element
		iconName := n.Parent.Parent.SelectElement("name").InnerText()

		// Comebine size and name to get our icon name
		iconFilename := iconName + "-" + iconSize[1]

		// Get the fill geomtery (path) so we can diff it for changes

		// TODO: Use fill geometry to actually create SVGs??!!

		// Create new icon object
		i := core.Icon{
			Name:       iconFilename,
			FigmaID:    n.InnerText(),
			Fingeprint: core.MD5("1"),
		}

		icons = append(icons, i)
	}

	return icons, nil
}

// Given an Icon, returns a URL to view it online
func ExportIconToURL(icon core.Icon) (string, error) {
	url := fmt.Sprintf("https://api.figma.com/v1/images/%s", documentId)

	param := req.Param{
		"ids":    icon.FigmaID,
		"format": "svg",
		"scale":  "1",
	}

	r, err := req.Get(url, header, param)

	if err != nil {
		log.Fatal(err)
		return "", errors.New("could not get export icon")
	}

	log.Println(fmt.Sprintf("Exporting %s", icon.Name))

	value := gjson.Get(r.String(), fmt.Sprintf("images.%s", icon.FigmaID))

	return value.String(), nil
}
