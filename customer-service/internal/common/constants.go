package common

const (
	Production  = "PRODUCTION"
	Development = "DEVELOPMENT"
)

const (
	CustomerStatusWaiting      = "WAITING"       // Waiting for the request (food)
	CustomerStatusEating       = "EATING"        // Eating, can ask for more items
	CustomerStatusNeedsPayment = "NEEDS_PAYMENT" // Finished eating, payment needed
	CustomerStatusPaid         = "PAID"          // Paid, order done
)

const (
	TableStatusAvailable   = "AVAILABLE"
	TableStatusUnavailable = "UNAVAILABLE"
)
