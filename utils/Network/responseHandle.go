package Network

import (
	"bytes"
	"github.com/wgpsec/EndpointSearch/utils/Error"
	"io"
	"net/http"
)

func HandleResponse(resp *http.Response) (bodyString string) {
	bodyBuf := new(bytes.Buffer)
	_, err := io.Copy(bodyBuf, resp.Body)
	Error.HandleError(err)
	bodyString = bodyBuf.String()
	return bodyString
}
