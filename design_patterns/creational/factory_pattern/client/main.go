package main

import (
	"fmt"
	"github.com/ntnghiatn/leetcode-golang/design_patterns/creational/factory_pattern/abstract_factory"
	"log"
)

func main() {

	log.Println(testFactory())
	log.Println(doTask())

}

func testFactory() string {
	// using factory to init product
	newShoe := abstract_factory.NewShoe()
	r := fmt.Sprintf("Logoname: %v and size: %d", newShoe.GetLogo(), newShoe.GetSize())
	return r
}

func doTask() string {
	newCombo, err := abstract_factory.GetWinterCombo("summeer")
	if err != nil {
		return err.Error()
	}

	//
	newCombo.GetShoe()
	newCombo.GetShort()
	return "Done!"
}
