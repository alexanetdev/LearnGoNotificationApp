package enums

type OutreachType int

const (
	Reminder OutreachType = 1
	Blast    OutreachType = 2
)

func (o OutreachType) String() string {
	switch o {
	case Reminder:
		return "Reminder"
	case Blast:
		return "Blast"
	default:
		return "Unknown"
	}
}
