package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	controllers "github.com/dgonzalez/pid/pkg/controllers"
	pid "github.com/dgonzalez/pid/pkg/controllers/pid"
)

func main() {
	var controller controllers.Controller = pid.New(0.01, 0.01, 0.5, 1)
	controller.SetTarget(500)

	file, err := os.OpenFile("output.csv", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	spew.Dump(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{"Iteration", "Value"})

	lastValue := 0.0
	for i := 0; i < 1000; i++ {
		delta := controller.Update(lastValue)
		lastValue += delta
		spew.Dump(lastValue)
		valueString := fmt.Sprintf("%f", lastValue)
		spew.Dump(valueString)
		iteration := fmt.Sprintf("%d", i)

		writer.Write([]string{iteration, valueString})
	}
	writer.Flush()
}
