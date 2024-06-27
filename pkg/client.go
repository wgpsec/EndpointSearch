package pkg

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/wgpsec/EndpointSearch/utils/Error"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

func GenerateHTTPClient(timeOut int, proxyURL string) *http.Client {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 5 {
				return fmt.Errorf("stopped after 5 redirects")
			}
			return nil
		},
	}

	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if timeOut > 0 {
		timeoutDuration := time.Duration(timeOut) * time.Second
		client.Timeout = timeoutDuration
	}

	if proxyURL != "" {
		proxyURLParsed, err := url.Parse(proxyURL)
		if err != nil {
			Error.HandleError(fmt.Errorf("[-] Invalid proxy URL: %v\n", err))
			os.Exit(1)
		}

		switch proxyURLParsed.Scheme {
		case "http", "https":
			client.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxyURLParsed),
			}

		case "socks5", "socks5h":
			username, password, proxyAddress := ParseSockURL(proxyURLParsed.String())
			var dialer proxy.Dialer
			var err error

			if username != "" {
				proxyAuth := &proxy.Auth{
					User:     username,
					Password: password,
				}
				dialer, err = proxy.SOCKS5("tcp", proxyAddress, proxyAuth, nil)
			} else {
				dialer, err = proxy.SOCKS5("tcp", proxyAddress, nil, nil)
			}

			if err != nil {
				Error.HandleError(fmt.Errorf("[-] Failed to connect to the proxy server: %v\n", err))
				os.Exit(1)
			}

			client.Transport = &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return dialer.Dial(network, addr)
				},
			}

		default:
			Error.HandleError(fmt.Errorf("[-] Unsupported protocol: %v\n", proxyURLParsed.Scheme))
			os.Exit(1)
		}
	}

	return client
}

func ParseSockURL(urlStr string) (username, password, proxyAddress string) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", "", ""
	}

	if parsedURL.User != nil {
		username = parsedURL.User.Username()
		password, _ = parsedURL.User.Password()
	}

	proxyAddress = parsedURL.Host

	return username, password, proxyAddress
}
