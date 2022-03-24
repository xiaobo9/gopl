package funcp

func doActionArgs(action func(int, int) int) int {
	return action(1, 2)
}

func doAction(action func()) {
	action()
}
