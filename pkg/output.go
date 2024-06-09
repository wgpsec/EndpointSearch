package pkg

import (
	"bufio"
	"fmt"
	"github.com/wgpsec/EndpointSearch/utils/Error"
	"os"
	"sort"
)

func WriteToFile(writeResultList []string, output string) {
	file, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	Error.HandlePanic(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	sort.Strings(writeResultList)
	if len(writeResultList) != 0 {
		for _, i := range writeResultList {
			fmt.Println("[+] service endpoint:", i)
			fmt.Fprintln(writer, i)
		}
	}
}
