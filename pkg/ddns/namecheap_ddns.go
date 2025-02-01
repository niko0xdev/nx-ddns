package ddns

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/niko0xdev/nx-ddns/internal/utils"
)

type NameCheapError struct {
	Err1 string `xml:"Err1"`
}

type NameCheapResponse struct {
	ErrorCount int            `xml:"ErrCount"`
	Errors     NameCheapError `xml:"errors"`
}

// api key should be: <username>:<password>
func UpdateNameCheapDNSRecord(apiKey, domain, hostname, ip string) error {
	var response NameCheapResponse

	// Link from Namecheap knowledge article.
	// https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-to-dynamically-update-the-hosts-ip-with-an-http-request/
	ncURL := NAMECHEAP_DNS_API_HOST + "?host=" + hostname + "&domain=" + domain + "&password=" + apiKey + "&ip=" + ip

	client := &http.Client{Timeout: httpTimeout}

	req, err := http.NewRequest("GET", ncURL, nil)
	if err != nil {
		// fmt.Println(1, err.Error())
		return err
	}

	// req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*")
	// req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	// req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		// fmt.Println(2, err.Error())
		return err
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Below function removes first line (below line) from response body because golang xml encoder does not support utf-16
	// <?xml version="1.0" encoding="utf-16"?>
	modifyBodyBytes := func(bodyBytes []byte) []byte {

		bodyString := string(bodyBytes)

		read_lines := strings.Split(bodyString, "\n")

		var updatedString string

		for i, line := range read_lines {
			if i != 0 {
				updatedString = fmt.Sprintf("%s%s\n", updatedString, line)
			}
		}

		return []byte(updatedString)
	}

	err = xml.Unmarshal(modifyBodyBytes(bodyBytes), &response)
	if err != nil {
		return err
	}

	if response.ErrorCount != 0 {
		return &utils.CustomError{ErrorCode: -1, Err: errors.New(response.Errors.Err1)}
	}

	return nil
}
