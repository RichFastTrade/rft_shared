package utils

var IntervalMap = map[string]Interval{
	"1m": {
		1,
		"1m",
		"1 minute",
		"* * * * *",
		"6h",
		func(m int64, h int64, d int64, dd int64) bool { return true },
	},
	"5m": {
		5,
		"5m",
		"5 minutes",
		"*/5 * * * *",
		"1d6h",
		func(m int64, h int64, d int64, dd int64) bool { return m%5 == 0 },
	},
	"15m": {
		15,
		"15m",
		"15 minutes",
		"*/15 * * * *",
		"3d18h",
		func(m int64, h int64, d int64, dd int64) bool { return m%15 == 0 },
	},
	"30m": {
		30,
		"30m",
		"30 minutes",
		"*/30 * * * *",
		"7d12h",
		func(m int64, h int64, d int64, dd int64) bool { return m%30 == 0 },
	},
	"1h": {
		60 * 1,
		"1h",
		"1 hour",
		"0 * * * *",
		"15d",
		func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 },
	},
	"2h": {
		60 * 2,
		"2h",
		"2 hours",
		"0 */2 * * *",
		"1M",
		func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 && h%2 == 0 },
	},
	//"3h": {
	//	60 * 3,
	//	"3h",
	//	"3 hours",
	//	"0 */3 * * *",
	//	"45d",
	//	func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 && h%3 == 0 },
	//},
	"4h": {
		60 * 4,
		"4h",
		"4 hours",
		"0 */4 * * *",
		"2M",
		func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 && h%4 == 0 },
	},
	"6h": {
		60 * 6,
		"6h",
		"6 hours",
		"0 */6 * * *",
		"3M",
		func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 && h%6 == 0 },
	},
	"8h": {
		60 * 8,
		"8h",
		"8 hours",
		"0 */8 * * *",
		"4M",
		func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 && h%8 == 0 },
	},
	"12h": {
		60 * 12,
		"12h",
		"12 hours",
		"0 */12 * * *",
		"6M",
		func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 && h%12 == 0 },
	},
	"1d": {
		60 * 24,
		"1d",
		"1 day",
		"0 0 * * *",
		"1y",
		func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 && h%24 == 0 },
	},
	"2d": {
		60 * 24 * 2,
		"2d",
		"2 days",
		"0 0 * * *",
		"2y",
		func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 && h%24 == 0 && d%2 == 0 },
	},
	"3d": {
		60 * 24 * 3,
		"3d",
		"3 days",
		"0 0 * * *",
		"3y",
		func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 && h%24 == 0 && d%3 == 0 },
	},
	"5d": {
		60 * 24 * 5,
		"5d",
		"5 days",
		"0 0 * * *",
		"5y",
		func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 && h%24 == 0 && d%5 == 0 },
	},
	"1w": {
		60 * 24 * 7,
		"1w",
		"1 week",
		"0 0 * * *",
		"7y",
		func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 && h%24 == 0 && d%7 == 0 },
	},
	"1M": {
		60 * 24 * 28,
		"1M",
		"1 month",
		"0 0 * * *",
		"30y",
		func(m int64, h int64, d int64, dd int64) bool { return m%60 == 0 && h%24 == 0 && dd == 0 },
	},
}

type Interval struct {
	Minutes           int64
	String            string
	TimescaleInterval string
	CronExpression    string
	ValidateInterval  string
	Task              func(int64, int64, int64, int64) bool
}
