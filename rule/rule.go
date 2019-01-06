package rule

// Rule 规则
type Rule struct {
	// pathPrefix / pathStrip / etc.
	Type    string
	Content string
	Host    string
	Port    int
}

// Match 规则匹配
func (rule *Rule) Match(url string) (bool, string) {
	if rule.Type == "pathPrefix" {
		// 前缀匹配
		if url[:len(rule.Content)] == rule.Content {
			return true, url
		}
	} else if rule.Type == "path" {
		// 完全匹配
		if rule.Content == url {
			return true, url
		}
	} else if rule.Type == "pathPrefixStrip" {
		// 前缀匹配，并去掉前缀
		if url[:len(rule.Content)] == rule.Content {
			return true, url[len(rule.Content):]
		}
	}
	return false, ""
}
