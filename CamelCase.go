package hackerranksubmission

func camelcase(s string) int32 {
	var counter int32 = 1

	for _, value := range s {

		if value >= 65 && value <= 90 {
			counter += 1
		}
	}
	return counter
}
