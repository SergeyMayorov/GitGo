package main

import (
	v10 "v10"
	v11 "v11"
	v20 "v20"
	v21 "v21"
)

func main() {
	v10.Do("patients.txt", "result_v10.txt")
	v11.Do("patients.txt", "result_v11.txt")
	v20.Do("patients.txt", "result_v20.txt")
	v21.Do("patients.txt", "result_v21.txt")
}
