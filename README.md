pid is a PID controller (as you might have guessed by the name). If you need more information on what a PID contorller is, [Wikipedia is your friend](https://en.wikipedia.org/wiki/PID_controller).

# Quickstart

Super straight-forward:
```go
  var controller controllers.Controller = pid.New(0.01, 0.01, 0.5, 1)
	controller.SetTarget(500)

	file, err := os.OpenFile("output.csv", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{"Iteration", "Value"})

	lastValue := 0.0
	for i := 0; i < 1000; i++ {
		delta := controller.Update(lastValue)
		lastValue += delta
		valueString := fmt.Sprintf("%f", lastValue)
		iteration := fmt.Sprintf("%d", i)

		writer.Write([]string{iteration, valueString})
	}
	writer.Flush()
```
The above snippet will create a PID controller and iterate for 1000 times until the target
number converges creating a CSV with all the iterations.
