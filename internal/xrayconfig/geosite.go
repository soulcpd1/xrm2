package xrayconfig

import (
	"bufio"
	"os"
	"strings"
)

// LoadDomainList загружает список доменов в форматах:
//   - domain.com
//   - 0.0.0.0 domain.com
//   - 127.0.0.1 domain.com
func LoadDomainList(path string) ([]Rule, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var rules []Rule
	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "!") {
			continue
		}
		// Убираем inline-комментарии
		if idx := strings.Index(line, "#"); idx >= 0 {
			line = strings.TrimSpace(line[:idx])
		}
		if line == "" {
			continue
		}

		// Обрабатываем hosts-формат: "0.0.0.0 domain.com" или "127.0.0.1 domain.com"
		fields := strings.Fields(line)
		var domain string
		if len(fields) >= 2 && (fields[0] == "0.0.0.0" || fields[0] == "127.0.0.1") {
			domain = fields[1]
		} else if len(fields) == 1 {
			domain = fields[0]
		} else {
			continue
		}

		// Пропускаем localhost и служебные
		if domain == "localhost" || domain == "0.0.0.0" || domain == "127.0.0.1" || domain == "::1" {
			continue
		}

		rules = append(rules, Rule{
			Type:   DomainRule,
			Value:  domain,
			Action: "block",
		})
	}

	if err := scanner.Err(); err != nil {
		return rules, err
	}
	return rules, nil
}
