package rule

import (
	"github.com/wgpsec/EndpointSearch/internal/config"
	"github.com/wgpsec/EndpointSearch/pkg"
	"strings"
)

func JudgeEndpoint(respList ...pkg.ResponseData) (resultList []string) {
	for _, resp := range respList {
		for _, rule := range config.R.RuleText {
			if HeaderRuleHit(resp.Header, rule.Header) && BodyRuleHit(resp.Body, rule.Body) {
				resultList = append(resultList, resp.Url)
				break // 一旦找到匹配的规则，就不再检查其他规则
			}
		}
	}
	return resultList
}

func HeaderRuleHit(headers []string, ruleHeaders []string) bool {
	for _, header := range headers {
		for _, ruleHeader := range ruleHeaders {
			if strings.Contains(header, ruleHeader) {
				return true
			}
		}
	}
	return false
}

func BodyRuleHit(body string, ruleBodies []string) bool {
	for _, ruleBody := range ruleBodies {
		if strings.Contains(body, ruleBody) {
			return true
		}
	}
	return false
}
