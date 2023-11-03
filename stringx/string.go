package stringx

import "bytes"

// Camel2Snake 驼峰转蛇形
func Camel2Snake(s string) string {
	var b bytes.Buffer
	for i, v := range s {
		if v >= 'A' && v <= 'Z' {
			if i > 0 {
				b.WriteString("_")
			}
			b.WriteString(string(v + 32))
		} else {
			b.WriteString(string(v))
		}
	}
	return b.String()
}

// Snake2Camel 蛇形转驼峰
func Snake2Camel(s string) string {
	var b bytes.Buffer
	upper := true
	for _, v := range s {
		if v == '_' {
			upper = true
		} else {
			if upper {
				b.WriteString(string(v - 32))
				upper = false
			} else {
				b.WriteString(string(v))
			}
		}
	}
	return b.String()
}
