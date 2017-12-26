package main

type instruction struct {
	register  string
	action    string
	value     int
	condition conditional
}

type conditional struct {
	register string
	operator string
	value    int
}

func (i instruction) apply(registers *map[string]int) {
	switch i.action {
	case "dec":
		if i.condition.isTrue(*registers) {
			(*registers)[i.register] -= i.value
		}
		break

	case "inc":
		if i.condition.isTrue(*registers) {
			(*registers)[i.register] += i.value
		}
		break
	}
}

func (c conditional) isTrue(registers map[string]int) bool {
	switch c.operator {
	case ">":
		return registers[c.register] > c.value
	case "<":
		return registers[c.register] < c.value
	case "==":
		return registers[c.register] == c.value
	case ">=":
		return registers[c.register] >= c.value
	case "!=":
		return registers[c.register] != c.value
	case "<=":
		return registers[c.register] <= c.value
	default:
		return false
	}
}
