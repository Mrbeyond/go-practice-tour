package apicall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
)

// A Multpart/form-data request
func TestMedilleryProjectsCount() {

	url := "https://medilleryserver.medillery.com/api/projects_range_count"
	var wbuf bytes.Buffer
	formdata := multipart.NewWriter(&wbuf)
	//TODO: public key needed but removed
	formdata.WriteField("from", "2019-01-01")
	formdata.WriteField("to", "2023-03-23")

	formdata.Close()
	req, err := http.NewRequest(http.MethodPost, url, &wbuf)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("X-external-source", "binary")
	req.Header.Set("Content-Type", formdata.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err, ": is err")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("There is an error, non 200 status")
		return
	}

	var payload map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&payload)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(payload)
}
