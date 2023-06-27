package tenkiGetter

type Config struct {
	RunMode    Mode
	OfficeCode string
}

type Mode int

const (
	Day Mode = iota + 1
	Week
)

func NewConfig(mode Mode, str string) *Config {
	code := replaceOfficeName(str)
	return &Config{RunMode: mode, OfficeCode: code}
}

func (m Mode) String() string {
	switch m {
	case Day:
		return "day"
	case Week:
		return "week"
	default:
		return "unknown"
	}
}

func (m Mode) GetMode() string {
	switch m {
	case Day:
		return "overview_forecast"
	case Week:
		return "overview_week"
	default:
		return "unknown"
	}
}

func replaceOfficeName(name string) string {
	/* うまく動かなかった
	offices, err := GetOffice()
	if err != nil {
		return "", err
	}

	for i := 010000; i <= 480000; i++ {
		//fmt.Println(i)
		if data, ok := offices.CheckGet(string(fmt.Sprint("%06d", i))); ok {
			fmt.Println(data.MustString("name"))
			if name == data.MustString("name") {
				return string(rune(i)), nil
			}
		}
	}
	return "", nil
	*/
	switch name {
	case "宗谷地方":
		return "011000"
	case "上川・留萌地方":
		return "012000"
	case "網走・北見・紋別地方":
		return "013000"
	case "十勝地方":
		return "014030"
	case "釧路・根室地方":
		return "014100"
	case "胆振・日高地方":
		return "015000"
	case "石狩・空知・後志地方":
		return "016000"
	case "渡島・檜山地方":
		return "017000"
	case "青森県":
		return "020000"
	case "岩手県":
		return "030000"
	case "宮城県":
		return "040000"
	case "秋田県":
		return "050000"
	case "山形県":
		return "060000"
	case "福島県":
		return "070000"
	case "茨城県":
		return "080000"
	case "栃木県":
		return "090000"
	case "群馬県":
		return "100000"
	case "埼玉県":
		return "110000"
	case "千葉県":
		return "120000"
	case "東京都":
		return "130000"
	case "神奈川県":
		return "140000"
	case "新潟県":
		return "150000"
	case "富山県":
		return "160000"
	case "石川県":
		return "170000"
	case "福井県":
		return "180000"
	case "山梨県":
		return "190000"
	case "長野県":
		return "200000"
	case "岐阜県":
		return "210000"
	case "静岡県":
		return "220000"
	case "愛知県":
		return "230000"
	case "三重県":
		return "240000"
	case "滋賀県":
		return "250000"
	case "京都府":
		return "260000"
	case "大阪府":
		return "270000"
	case "兵庫県":
		return "280000"
	case "奈良県":
		return "290000"
	case "和歌山県":
		return "300000"
	case "鳥取県":
		return "310000"
	case "島根県":
		return "320000"
	case "岡山県":
		return "330000"
	case "広島県":
		return "340000"
	case "山口県":
		return "350000"
	case "徳島県":
		return "360000"
	case "香川県":
		return "370000"
	case "愛媛県":
		return "380000"
	case "高知県":
		return "390000"
	case "福岡県":
		return "400000"
	case "佐賀県":
		return "410000"
	case "長崎県":
		return "420000"
	case "熊本県":
		return "430000"
	case "大分県":
		return "440000"
	case "宮崎県":
		return "450000"
	case "奄美地方":
		return "460040"
	case "鹿児島県":
		return "460100"
	case "沖縄本島地方":
		return "471000"
	case "大東島地方":
		return "472000"
	case "宮古島地方":
		return "473000"
	case "八重山地方":
		return "474000"

	default:
		return ""
	}
}
