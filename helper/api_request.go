package helper

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gordan0410/common/dto"
	"github.com/gordan0410/common/enum/wh_error"
)

const defaultApiTimeout time.Duration = 60

func SendApiRequest(req *dto.SendApiRequest) (dto.SendApiResponse, error) {
	if req.Timeout == 0 {
		req.Timeout = defaultApiTimeout * time.Second
	}

	bodyReader := bytes.NewReader(req.Body)

	httpRequest, err := http.NewRequest(req.Method, req.Url, bodyReader)
	if err != nil {
		return dto.SendApiResponse{}, err
	}

	for k, v := range req.Header {
		httpRequest.Header.Set(k, v)
	}

	tr := http.Transport{
		DisableKeepAlives: true,
	}

	client := http.Client{
		Timeout:   req.Timeout,
		Transport: &tr,
	}

	startTime := time.Now()
	resp, err := client.Do(httpRequest)
	endTime := time.Now()
	respData := dto.SendApiResponse{
		HttpResponse: resp,
		CostTime:     endTime.Sub(startTime),
	}
	if err != nil {
		return respData, err
	}

	if resp == nil {
		return respData, wh_error.RecordNotExistsError
	}

	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return respData, err
	}
	respData.Data = respBytes

	return respData, nil
}
