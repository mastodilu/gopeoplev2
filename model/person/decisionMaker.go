package person

// DecisionMaker represents the brain needed to take a decision
type DecisionMaker interface {
	IsRightPartner(intelligent1, intelligent2, pretty1, pretty2 int, sex1, sex2 byte) bool
}
