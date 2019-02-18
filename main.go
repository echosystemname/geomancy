package main

import "fmt"
import "math/rand"
import "time"
import "strconv"

type Figure struct {
	contents [4]int
}

//https://www.princeton.edu/~ezb/geomancy/geostep.html
func main() {
	mothers := cast()
	daughters := daughters(mothers)
	nieces := nieces(mothers, daughters)
	witnesses := witnesses(nieces)
	judge := judge(witnesses)

	prettyPrint(mothers, daughters, nieces, witnesses, judge)
}

func cast() [4]Figure {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var mothers [4]Figure

	for i, v := range mothers {
		v.contents[0] = (r.Intn(100) % 2) + 1
		v.contents[1] = (r.Intn(100) % 2) + 1
		v.contents[2] = (r.Intn(100) % 2) + 1
		v.contents[3] = (r.Intn(100) % 2) + 1
		mothers[i] = v
	}
	return mothers
}

func daughters(mothers [4]Figure) [4]Figure {
	var daughters [4]Figure

	daughter1 := Figure{}
	daughter1.contents[0] = mothers[0].contents[0]
	daughter1.contents[1] = mothers[1].contents[0]
	daughter1.contents[2] = mothers[2].contents[0]
	daughter1.contents[3] = mothers[3].contents[0]

	daughter2 := Figure{}
	daughter2.contents[0] = mothers[0].contents[1]
	daughter2.contents[1] = mothers[1].contents[1]
	daughter2.contents[2] = mothers[2].contents[1]
	daughter2.contents[3] = mothers[3].contents[1]

	daughter3 := Figure{}
	daughter3.contents[0] = mothers[0].contents[2]
	daughter3.contents[1] = mothers[1].contents[2]
	daughter3.contents[2] = mothers[2].contents[2]
	daughter3.contents[3] = mothers[3].contents[2]

	daughter4 := Figure{}

	daughter4.contents[0] = mothers[0].contents[3]
	daughter4.contents[1] = mothers[1].contents[3]
	daughter4.contents[2] = mothers[2].contents[3]
	daughter4.contents[3] = mothers[3].contents[3]

	daughters[0] = daughter1
	daughters[1] = daughter2
	daughters[2] = daughter3
	daughters[3] = daughter4

	return daughters
}

func nieces(mothers [4]Figure, daughters [4]Figure) [4]Figure {
	var nieces [4]Figure

	nieces[0] = add(mothers[0], mothers[1])
	nieces[1] = add(mothers[2], mothers[3])
	nieces[2] = add(daughters[0], daughters[1])
	nieces[3] = add(daughters[2], daughters[3])

	return nieces
}

func witnesses(neices [4]Figure) [2]Figure {
	return [2]Figure{add(neices[0], neices[1]), add(neices[2], neices[3])}
}

func judge(witnesses [2]Figure) Figure {
	return add(witnesses[0], witnesses[1])
}

func add(a Figure, b Figure) Figure {
	result := Figure{}

	for i, _ := range result.contents {
		result.contents[i] = ((a.contents[i] + b.contents[i]) % 2) + 1
	}
	return result
}

func prettyPrint(mothers [4]Figure, daughters [4]Figure, nieces [4]Figure, witnesses [2]Figure, judge Figure) {
	fmt.Println()
	fmt.Println()
	fmt.Print(strconv.Itoa(daughters[3].contents[0]) + " ")
	fmt.Print(strconv.Itoa(daughters[2].contents[0]) + " ")
	fmt.Print(strconv.Itoa(daughters[1].contents[0]) + " ")
	fmt.Print(strconv.Itoa(daughters[0].contents[0]) + " ")
	fmt.Print(strconv.Itoa(mothers[3].contents[0]) + " ")
	fmt.Print(strconv.Itoa(mothers[2].contents[0]) + " ")
	fmt.Print(strconv.Itoa(mothers[1].contents[0]) + " ")
	fmt.Print(strconv.Itoa(mothers[0].contents[0]))
	fmt.Println()
	fmt.Print(strconv.Itoa(daughters[3].contents[1]) + " ")
	fmt.Print(strconv.Itoa(daughters[2].contents[1]) + " ")
	fmt.Print(strconv.Itoa(daughters[1].contents[1]) + " ")
	fmt.Print(strconv.Itoa(daughters[0].contents[1]) + " ")
	fmt.Print(strconv.Itoa(mothers[3].contents[1]) + " ")
	fmt.Print(strconv.Itoa(mothers[2].contents[1]) + " ")
	fmt.Print(strconv.Itoa(mothers[1].contents[1]) + " ")
	fmt.Print(strconv.Itoa(mothers[0].contents[1]))
	fmt.Println()
	fmt.Print(strconv.Itoa(daughters[3].contents[2]) + " ")
	fmt.Print(strconv.Itoa(daughters[2].contents[2]) + " ")
	fmt.Print(strconv.Itoa(daughters[1].contents[2]) + " ")
	fmt.Print(strconv.Itoa(daughters[0].contents[2]) + " ")
	fmt.Print(strconv.Itoa(mothers[3].contents[2]) + " ")
	fmt.Print(strconv.Itoa(mothers[2].contents[2]) + " ")
	fmt.Print(strconv.Itoa(mothers[1].contents[2]) + " ")
	fmt.Print(strconv.Itoa(mothers[0].contents[2]))
	fmt.Println()
	fmt.Print(strconv.Itoa(daughters[3].contents[3]) + " ")
	fmt.Print(strconv.Itoa(daughters[2].contents[3]) + " ")
	fmt.Print(strconv.Itoa(daughters[1].contents[3]) + " ")
	fmt.Print(strconv.Itoa(daughters[0].contents[3]) + " ")
	fmt.Print(strconv.Itoa(mothers[3].contents[3]) + " ")
	fmt.Print(strconv.Itoa(mothers[2].contents[3]) + " ")
	fmt.Print(strconv.Itoa(mothers[1].contents[3]) + " ")
	fmt.Print(strconv.Itoa(mothers[0].contents[3]))
	fmt.Println()
	fmt.Println(" a")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
