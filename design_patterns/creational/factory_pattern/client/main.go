package main

import (
	"fmt"
	"github.com/ntnghiatn/leetcode-golang/design_patterns/creational/factory_pattern/abstract_factory"
	"log"
)

func main() {

	//log.Println(testFactory())
	log.Println(doTask())

}

func doTask() string {
	newCombo, err := abstract_factory.GetWinterCombo("winter")
	if err != nil {
		return err.Error()
	}

	//
	fmt.Printf(
		"newCombo.GetShoe.GetLogo \"%v\" and Short:: %v",
		newCombo.GetShoe().GetLogo(),
		newCombo.GetShort(),
	)

	//newCombo.GetShort()
	return "Done!"
}

func testFactory() string {
	// using factory to init product
	newShoe := abstract_factory.NewConcreteShoe()
	r := fmt.Sprintf("Logoname: %v and size: %d", newShoe.GetLogo(), newShoe.GetSize())
	return r
}
