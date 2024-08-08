package map_test

import "fmt"

func Map() {
	//website:=[]string{"https://google.com","https://aws.com"} -> Khong the nho noi url can co 1 index
	website := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}
	fmt.Println(website)
	fmt.Println(website["Amazon"])
	website["LinkedIn"] = "https://linkedin.com"
	fmt.Println(website)
	delete(website, "Google")
	fmt.Println(website)

	courseRatings := map[string]float64{}
	courseRatings["go"] = 4.7
	fmt.Println(courseRatings)

	var testcourseRating = map[string]float64{}
	fmt.Println(testcourseRating)
	testcourseRating["go"] = 4.7
	testcourseRating["react"] = 4.8
	testcourseRating["angular"] = 5.2
	fmt.Println(testcourseRating)

	userNames := make([]string, 2, 5)
	userNames[0] = "Julie"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)
	for index, value := range userNames {
		fmt.Println("Index: ", index)
		fmt.Println("Value", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)

	}
}
