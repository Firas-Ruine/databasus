package intervals

type IntervalType string

const (
	IntervalEvery30Minutes IntervalType = "EVERY_30_MINUTES"
	IntervalHourly         IntervalType = "HOURLY"
	IntervalDaily          IntervalType = "DAILY"
	IntervalWeekly         IntervalType = "WEEKLY"
	IntervalMonthly        IntervalType = "MONTHLY"
	IntervalCron           IntervalType = "CRON"
)
