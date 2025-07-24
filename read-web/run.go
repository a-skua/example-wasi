package read

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"go.bytecodealliance.org/cm"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	whttp "github.com/a-skua/go-wasi/http"
)

var URL string

func Run() cm.BoolResult {
	body, err := fetchURL(URL)
	if err != nil {
		slog.Error("Failed to read response body", "error", err)
		os.Exit(1)
	}
	defer body.Close()

	html, err := io.ReadAll(body)
	if err != nil {
		slog.Error("Failed to read response body", "error", err)
		os.Exit(1)
	}

	md, err := htmlToMD(string(html))
	if err != nil {
		slog.Error("Failed to convert HTML to Markdown", "error", err)
		os.Exit(1)
	}

	fmt.Println(string(md))

	return cm.ResultOK
}

func fetchURL(url string) (io.ReadCloser, error) {
	var c whttp.Client
	r, err := c.Get(url)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", r.StatusCode)
	}

	return r.Body, nil
}

func htmlToMD(html string) (string, error) {
	return htmltomarkdown.ConvertString(html)
}
