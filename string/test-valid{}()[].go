// LC20. 有效的括号
package main

import "container/list"

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var l list.List
	for i := range s {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			l.PushFront(s[i])
		} else {
			if l.Front() == nil {
				return false
			}
			v := l.Front().Value.(byte)
			l.Remove(l.Front())
			switch s[i] {
			case '}':
				if v != '{' {
					return false
				}
			case ']':
				if v != '[' {
					return false
				}
			case ')':
				if v != '(' {
					return false
				}
			}

		}
	}
	if l.Front() != nil {
		return false
	}

	return true

}
