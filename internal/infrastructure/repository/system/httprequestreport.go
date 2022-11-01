package system

import (
	"errors"
	"io"
	"net/http"
	"pachico/snitch/internal/domain"
	"strings"
)

type HTTPRequestReportRepository struct {
}

func (r *HTTPRequestReportRepository) GetHTTPRequestReport(url string) (domain.HTTPRequestReport, error) {
	report := domain.HTTPRequestReport{
		HTTPRequest: domain.HTTPRequest{
			URL: url,
		},
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return report, errors.New("URL must start with 'http://' or 'https://'")
	}

	response, err := http.Get(url)
	if err != nil {
		return report, errors.New("error resolving or requesting " + url + ": " + err.Error())
	}
	defer response.Body.Close()

	maxSize := int64(1 << 20) // 1 megabyte
	limitedReader := io.LimitedReader{R: response.Body, N: maxSize}

	responseBody, err := io.ReadAll(&limitedReader)
	if err != nil {
		return report, errors.New("error reading the response body of " + url + ": " + err.Error())
	}

	if limitedReader.N <= 0 {
		return report, errors.New("response body exceeds maximum size limit of " + url)
	}

	report.HTTPResponse = domain.HTTPResponse{
		StatusCode: response.StatusCode,
		Size:       response.ContentLength,
		Header:     response.Header,
		Body:       string(responseBody),
	}

	return report, nil
}
