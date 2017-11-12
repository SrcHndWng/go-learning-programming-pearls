package main

func searchName(pushed string, names personNames, min int, max int) personNames {
	i := min + ((max - min) / 2)
	if pushed < names[i].pushNumber {
		return searchName(pushed, names, min, i)
	} else if pushed > names[i].pushNumber {
		return searchName(pushed, names, i, max)
	}

	var result personNames
	result = append(result, names[i])
	for j := i - 1; j >= min; j-- {
		if pushed == names[j].pushNumber {
			result = append(result, names[j])
		}
	}
	for j := i + 1; j <= max; j++ {
		if pushed == names[j].pushNumber {
			result = append(result, names[j])
		}
	}
	return result
}
