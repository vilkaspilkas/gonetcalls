package gonetcalls

import (
    "time"
    "bytes"
    "net/http"
    "io/ioutil"
)

func HttpPost(url string, headers map[string]string, data *bytes.Buffer) (code int, clen, rtt int64, body string, rheaders http.Header, err_ret string) {
    startTime := time.Now()
    client := &http.Client{}

    //Prepare request
    req, err := http.NewRequest("POST", url, data)
    if err != nil {
        rtt = int64(time.Since(startTime).Nanoseconds()/1000)
        return 0, 0, rtt, "", nil, "Failed to create http request." + err.Error()
    }

    for hname,hvalue := range(headers) {
        req.Header.Add(hname, hvalue)
    }

    //Send the request
    resp,err := client.Do(req)
    if err != nil {
        rtt = int64(time.Since(startTime).Nanoseconds()/1000)
        return 0, 0, rtt, "", nil, "Failed to get response from server." + err.Error()
    }

    //Process the response
    code = resp.StatusCode
    clen = resp.ContentLength
    defer resp.Body.Close()
    rbody, errb := ioutil.ReadAll(resp.Body)
    if errb != nil {
        rtt = int64(time.Since(startTime).Nanoseconds()/1000)
        return code, clen, rtt, "", nil, "Failed reading response body." + errb.Error()
    }
    body = string(rbody)
    if clen == -1 {
        clen = int64(len(body))
    }
    rtt = int64(time.Since(startTime).Nanoseconds()/1000)
    return code, clen, rtt, body, resp.Header, ""
}

func HttpGet(url string, headers []string) (code int, clen, rtt int64, body string, rheaders http.Header, err_ret string) {
    startTime := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        rtt = int64(time.Since(startTime).Nanoseconds()/1000)
        return 0, 0, rtt, "", nil, "Failed to get response from server." + err.Error()
    }

    code = resp.StatusCode
    clen = resp.ContentLength
    defer resp.Body.Close()
    rbody, errb := ioutil.ReadAll(resp.Body)
    if errb != nil {
        rtt = int64(time.Since(startTime).Nanoseconds()/1000)
        return code, clen, rtt, "", nil, "Failed reading response body." + errb.Error()
    }
    body = string(rbody)
    rtt = int64(time.Since(startTime).Nanoseconds()/1000)
    return code, clen, rtt, body, resp.Header, ""
}