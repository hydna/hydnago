package hydna

import "errors"
import "io"
import "io/ioutil"
import "net/http"
import "strings"
import "bytes"


const payload_max_size = 65530


func Send(url string, body io.Reader) error {

    resp, err := doRequest(url, body, false)

    if err != nil {
        return err
    }

    return checkResponse(resp)
}

func Emit(url string, body io.Reader) error {

    resp, err := doRequest(url, body, true)

    if err != nil {
        return err
    }

    return checkResponse(resp)
}

func prefixUrl(url string) string {
    if strings.HasPrefix(url, "http://") == false ||
        strings.HasPrefix(url, "https://") == false {
        return "http://" + url
    }

    return url
}

func doRequest(url string, body io.Reader, emit bool) (*http.Response, error) {

    url = prefixUrl(url)

    if body == nil {
        return nil, errors.New("Body cannot be nil")
    }

    all_bytes, err := ioutil.ReadAll(body)
    
    if err != nil || len(all_bytes) > payload_max_size {
        return nil, errors.New("Payload max size reached");
    }

    data := bytes.NewBuffer(all_bytes)

    client := &http.Client{}

    req, err := http.NewRequest("POST", url, data)

    if err != nil {
        return nil, err
    }

    if emit {
        req.Header.Add("X-Emit", "yes")
    }

    return client.Do(req)
}

func checkResponse(resp *http.Response) error {

    if resp.StatusCode == 400 {
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return err
        }
        return errors.New(string(body[:]))
    }

    resp.Body.Close()

    return nil
}