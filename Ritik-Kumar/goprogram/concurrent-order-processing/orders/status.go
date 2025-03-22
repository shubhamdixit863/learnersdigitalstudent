package orders

type Status string

const (
	Processing Status = "Processing"
	Shipped    Status = "Shipped"
	Failed     Status = "Failed"
)
