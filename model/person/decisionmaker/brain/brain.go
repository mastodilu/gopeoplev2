package brain

// Brain implements DecisionMaker interface
type Brain struct{}

// IsRightPartner returns true if b is the right partner for a
func (b *Brain) IsRightPartner(intelligent1, intelligent2, pretty1, pretty2 int, sex1, sex2 byte) bool {
	return sex2 != sex1 &&
		pretty2 >= pretty1 &&
		intelligent2 >= intelligent1
}
