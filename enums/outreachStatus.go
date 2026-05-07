package enums

type OutreachStatus int

const (
	NotSent OutreachStatus = 0
	Success OutreachStatus = 1
	Failed  OutreachStatus = 2
)

func (o OutreachStatus) String() string {
	switch o {
	case NotSent:
		return "NotSent"
	case Success:
		return "Success"
	case Failed:
		return "Failed"
	default:
		return "Unknown"
	}
}
