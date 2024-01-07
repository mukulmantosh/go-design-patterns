package main

func main() {
	// Construct two DataListener observers and give each one a name

	listener1 := DataListener{
		Name: "Listener 1",
	}
	listener2 := DataListener{
		Name: "Listener 2",
	}

	subj := &DataSubject{}

	subj.registerObserver(listener1)
	subj.registerObserver(listener2)

	subj.ChangeItem("Monday!")
	subj.ChangeItem("Sunday!")
	subj.ChangeItem("Friday!")

	//unregister
	subj.unregisterObserver(listener2)
	subj.ChangeItem("Tuesday!")

}
