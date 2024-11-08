package ifsc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBankName_BankName(t *testing.T) {
	assert := assert.New(t)
	actual, ierr := GetBankName("WBSC0DJCB01")
	assert.Nil(ierr)
	assert.Equal("Darjeeling District Central Co-operative Bank", actual)
}

func TestGetBankName_Sublet(t *testing.T) {
	assert := assert.New(t)
	fixtureData := getSubletFixture()
	for input, _ := range fixtureData {
		ownerBankCode := input[0:4]
		actual, err := GetBankName(input)
		assert.Nil(err)
		expected, err := GetBankName(ownerBankCode)
		assert.Nil(err)
		assert.Equal(expected, actual)
	}
}
func TestGetBankName_CustomSublet_Success(t *testing.T) {
	assert := assert.New(t)
	fixtureData := getCustomSubletFixture()
	for input, expected := range fixtureData {
		actual, err := GetBankName(input)
		assert.Nil(err)
		assert.Equal(expected, actual)
	}

}

func getCustomSubletFixture() map[string]string {
	return map[string]string{
		"KSCB0006001": "Tumkur District Central Bank",
		"WBSC0KPCB01": "Kolkata Police Co-operative Bank",
		"YESB0ADB002": "Amravati District Central Co-operative Bank",
	}
}

func getSubletFixture() map[string]string {
	return map[string]string{
		"SKUX": "IBKL0116SBK",
		"SPTX": "IBKL0116SSB",
		"VCOX": "IBKL0116VMC",
		"AURX": "IBKL01192AC",
		"NMCX": "IBKL0123NMC",
		"MSSX": "IBKL01241MB",
		"TNCX": "IBKL01248NC",
		"URDX": "IBKL01263UC",
	}
}

func TestValidateBankCode(t *testing.T) {
	type args struct {
		bankCodeInput string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"success",
			args{"ABCX"},
			true,
		},
		{
			"failure",
			args{"Aaaa"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateBankCode(tt.args.bankCodeInput); got != tt.want {
				t.Errorf("ValidateBankCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	var result interface{}
	err := LoadFile("validator_asserts.json", &result, "../../tests")
	if err != nil {
		t.Fatalf("error: %+v", err)
	}
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{}
	for key, value := range result.(map[string]interface{}) {
		for inp_key, expected_val := range value.(map[string]interface{}) {
			tests = append(tests, struct {
				name string
				args args
				want bool
			}{key + ":" + inp_key, args{inp_key}, expected_val.(bool)})
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.code); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBankCodeFromIfsc(t *testing.T) {
	tests := []struct {
		name     string
		ifscCode string
		code     string
	}{
		{"Andhra Pragathi Grameena Bank", "APGB0000001", "APGB"},
		{"AU Small Finance Bank", "AUBL0002122", "AUBL"},
		{"Axis Bank", "UTIB0000175", "UTIB"},
		{"Bandhan Bank", "BDBL0001934", "BDBL"},
		{"Bank of Baroda", "BARB0DIGVIJ", "BARB"},
		{"Bank of India", "BKID0003250", "BKID"},
		{"Bank of Maharashtra", "MAHB0000014", "MAHB"},
		{"Canara Bank", "CNRB0000386", "CNRB"},
		{"Central Bank of India", "CBIN0280580", "CBIN"},
		{"CITI Bank", "CITI0000007", "CITI"},
		{"City Union Bank", "CIUB0000498", "CIUB"},
		{"Development Bank of Singapore", "DBSS0IN0827", "DBSS"},
		{"DCB Bank", "DCBL0000368", "DCBL"},
		{"Deutsche Bank", "DEUT0797TRS", "DEUT"},
		{"Dhanlaxmi Bank", "DLXB0000143", "DLXB"},
		{"Equitas Small Finance Bank", "ESFB0007017", "ESFB"},
		{"Esaf Small Finance Bank", "ESMF0001148", "ESMF"},
		{"Federal Bank", "FDRL0002127", "FDRL"},
		{"HDFC Bank", "HDFC0000569", "HDFC"},
		{"ICICI Bank", "ICIC0003486", "ICIC"},
		{"IDBI", "IBKL0000427", "IBKL"},
		{"IDFC FIRST Bank", "IDFB0042642", "IDFB"},
		{"Indian Bank", "IDIB000J005", "IDIB"},
		{"Indian Overseas Bank", "IOBA0000417", "IOBA"},
		{"Indusind Bank", "INDB0000065", "INDB"},
		{"Jana Small Finance Bank", "JSFB0003074", "JSFB"},
		{"Karnataka Bank", "KARB0000372", "KARB"},
		{"Karnataka Vikas Grameena Bank", "KVGB0006301", "KVGB"},
		{"Karur Vysya Bank", "KVBL0002203", "KVBL"},
		{"Kotak Mahindra Bank", "KKBK0002936", "KKBK"},
		{"Paytm Payments Bank", "PYTM0123456", "PYTM"},
		{"Punjab & Sind Bank", "PSIB0000411", "PSIB"},
		{"Punjab National Bank", "PUNB0022400", "PUNB"},
		{"RBL Bank", "RATN0000243", "RATN"},
		{"Shivalik Small Finance Bank", "SMCB0001017", "SMCB"},
		{"South Indian Bank", "SIBL0000361", "SIBL"},
		{"Standard Chartered Bank", "SCBL0036064", "SCBL"},
		{"State Bank of India", "SBIN0012211", "SBIN"},
		{"Suryoday Small Finance Bank", "SURY0000018", "SURY"},
		{"Tamilnad Mercantile Bank", "TMBL0000113", "TMBL"},
		{"Catholic Syrian Bank", "CSBK0000262", "CSBK"},
		{"Chembur Nagarik Sahakari Bank", "ICIC00CNSBL", "CNSX"},
		{"Cosmos Co-operative Bank", "COSB0000056", "COSB"},
		{"Hongkong & Shanghai Banking Corporation", "HSBC0380002", "HSBC"},
		{"Jammu and Kashmir Bank", "JAKA0AHAMAD", "JAKA"},
		{"Kalupur Commercial Co-operative Bank", "KCCB0RJT057", "KCCB"},
		{"Varachha Co-operative Bank", "VARA0289008", "VARA"},
		{"UCO Bank", "UCBA0000081", "UCBA"},
		{"Ujjivan Small Finance Bank", "UJVN0004516", "USFB"},
		{"Union Bank of India", "UBIN0531511", "UBIN"},
		{"Yes Bank", "YESB0000533", "YESB"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bankDetails, err := GetBankDetailsFromIfscCode(test.ifscCode)
			assert := assert.New(t)
			assert.Nil(err)
			assert.Equal(bankDetails.Code, test.code)
			assert.Equal(bankDetails.Name, test.name)
		})
	}
}
