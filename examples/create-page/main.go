package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cryptowizard0/go-notion"
	"github.com/sanity-io/litter"
)

type httpTransport struct {
	w io.Writer
}

// RoundTrip implements http.RoundTripper. It multiplexes the read HTTP response
// data to an io.Writer for debugging.
func (t *httpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	res.Body = io.NopCloser(io.TeeReader(res.Body, t.w))

	return res, nil
}

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("NOTION_API_KEY")
	buf := &bytes.Buffer{}
	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &httpTransport{w: buf},
	}
	client := notion.NewClient(apiKey, notion.WithHTTPClient(httpClient))

	var parentPageID string
	flag.StringVar(&parentPageID, "parentPageId", "", "Parent page ID.")
	flag.Parse()

	params := notion.CreatePageParams{
		ParentType: notion.ParentTypePage,
		ParentID:   parentPageID,
		Title: []notion.RichText{
			{
				Text: &notion.Text{
					Content: "Create Page Example",
				},
			},
		},
		Children: []notion.Block{

			notion.BlockDTO{
				Image: &notion.ImageBlock{
					Type: notion.FileTypeExternal,
					External: &notion.FileExternal{
						URL: "https://picsum.photos/600/200.jpg",
					},
				},
			},
		},
	}
	page, err := client.CreatePage(ctx, params)
	if err != nil {
		log.Fatalf("Failed to create page: %v", err)
	}

	decoded := map[string]interface{}{}
	if err := json.NewDecoder(buf).Decode(&decoded); err != nil {
		log.Fatal(err)
	}

	// Pretty print JSON reponse.
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	if err := enc.Encode(decoded); err != nil {
		log.Fatal(err)
	}

	// Pretty print parsed `notion.Page` value.
	litter.Dump(page)
}
