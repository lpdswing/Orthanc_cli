package core

import "fmt"

func ListStudy() {
	fmt.Println("this is list study function.")
}

func DeleteAllOverview() {
	fmt.Println("delete all overview function.")
}

func DeleteOverviewByStudy(study_uid string) {
	fmt.Printf("delete overview by study :[%v]\n", study_uid)
}
