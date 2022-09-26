package integration

import (
	"encoding/hex"
	"fmt"
	. "github.com/kaifei-bianjie/common-parser/codec"
	"github.com/kaifei-bianjie/common-parser/utils"
)

func (s IntegrationTestSuite) TestService() {
	cases := []SubTest{
		{
			"DefineService",
			DefineService,
		},
		{
			"BindService",
			BindService,
		},
		{
			"UpdateServiceBinding",
			UpdateServiceBinding,
		},
		{
			"ServiceSetWithdrawAddress",
			ServiceSetWithdrawAddress,
		},
		{
			"DisableServiceBinding",
			DisableServiceBinding,
		},
		{
			"EnableServiceBinding",
			EnableServiceBinding,
		},
		{
			"RefundServiceDeposit",
			RefundServiceDeposit,
		},
		{
			"CallService",
			CallService,
		},
		{
			"PauseRequestContext",
			PauseRequestContext,
		},
		{
			"StartRequestContext",
			StartRequestContext,
		},
		{
			"KillRequestContext",
			KillRequestContext,
		},
		{
			"UpdateRequestContext",
			UpdateRequestContext,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func DefineService(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A85020A82020A212F697269736D6F642E736572766963652E4D7367446566696E655365727669636512DC010A0A77777465737430363135120F746573746465736372697074696F6E1A0474657374222A696161316571766B667468747272393367347039717370703534773664746A74726E32376172377270772A04746573743284017B22696E707574223A7B2274797065223A226F626A656374222C2270726F70657274696573223A7B226964223A7B2274797065223A22737472696E67227D7D7D2C226F7574707574223A7B2274797065223A226F626A656374222C2270726F70657274696573223A7B226E616D65223A7B2274797065223A22737472696E67227D7D7D7D12650A510A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A21038BE4539785F0621A19066EB2DA45C11DCC5FCCFF5D58E89C670BB80D251CC1B712040A02080118C61012100A0A0A05756972697312013210C09A0C1A40C429BF0C4C9705B2F508D669DF7831911D89BCC7FC0CA5CDF004B09B5AAEA172362D34AB026415128ABEDC940D91B6C17AF2898009C408D6957DB301200CBFF4")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Service.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func BindService(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AD5010AD2010A1F2F697269736D6F642E736572766963652E4D736742696E645365727669636512AE010A0E736572766963652D77616E676A75122A696161317A616565767173326C756A7265306B796C367563667072366C7868716370736D7632733566741A120A057569726973120932303030303030303022147B227072696365223A223230307569726973227D285032187B22696E707574223A7B7D2C226F7574707574223A7B7D7D3A2A696161317A616565767173326C756A7265306B796C367563667072366C7868716370736D76327335667412690A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A21035FBC38095298439E20335EE4C72A74DCFB1E5BCAAED6973416A180A9933B33DE12040A020801180D12150A0F0A057569726973120631303030303010C09A0C1A4018969A0237A1D4D7BD2E0FB3A0CC057490604D274DFBE18B71CAE7968409E68B6A62EB9A5038CDFA5355183B653F0C5F5C5ADA7928025BDC251D37BAA1CDBF49")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Service.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func UpdateServiceBinding(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AD9010AD6010A282F697269736D6F642E736572766963652E4D73675570646174655365727669636542696E64696E6712A9010A0A736572766963652D776A122A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E1A110A0575697269731208313030303030303022147B227072696365223A223130307569726973227D283232187B22696E707574223A7B7D2C226F7574707574223A7B7D7D3A2A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E12690A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103F5330F7602F5AA3994F462C71380E2FC0288FAA228890987B2F069DD9DCC7BBB12040A020801180312150A0F0A057569726973120631303030303010C09A0C1A404EC2322E9152A57B8633D9825409B186FA3AE988D6778AA10718F471C5F4B8DE74D6C342E6512D4D1E5152F879B87E372BF0D3AA14763F22340AB5AC64CA033A")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Service.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func ServiceSetWithdrawAddress(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A85010A82010A262F697269736D6F642E736572766963652E4D736753657457697468647261774164647265737312580A2A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E122A696161317A616565767173326C756A7265306B796C367563667072366C7868716370736D76327335667412690A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103F5330F7602F5AA3994F462C71380E2FC0288FAA228890987B2F069DD9DCC7BBB12040A020801180412150A0F0A057569726973120631303030303010C09A0C1A400E71CA95B9282C15E7773E5C7712A3FD109F3B4829A176A458AECA402F12CC6C14468C6FE3213BE7D597AF46CAFA1B60F9394B993BB09C5A1A251AC91F912652")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Service.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func DisableServiceBinding(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A94010A91010A292F697269736D6F642E736572766963652E4D736744697361626C655365727669636542696E64696E6712640A0A736572766963652D776A122A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E1A2A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E12580A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103F5330F7602F5AA3994F462C71380E2FC0288FAA228890987B2F069DD9DCC7BBB12040A0208011809120410C09A0C1A401B86FFA705EE787B013AF7C0D50904B0EFBD1F821FEC9BC0A611A372F7E1602C39C080B886E9526C5591A9C46E231F76AE780F12F9DECD164F03D3646AD91E6D")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Service.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func EnableServiceBinding(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AA6010AA3010A282F697269736D6F642E736572766963652E4D7367456E61626C655365727669636542696E64696E6712770A0A736572766963652D776A122A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E1A110A05756972697312083130303030303030222A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E12580A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103F5330F7602F5AA3994F462C71380E2FC0288FAA228890987B2F069DD9DCC7BBB12040A0208011807120410C09A0C1A40D205CC400941FD783FA4527035298DAF61357D62A31D68C38E45831C633E529939F186DB33D5BAE68F8281D25797CF839F88FEBDFAECFF900B560D2D807AC229")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Service.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func RefundServiceDeposit(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A93010A90010A282F697269736D6F642E736572766963652E4D7367526566756E64536572766963654465706F73697412640A0A736572766963652D776A122A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E1A2A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E12580A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103F5330F7602F5AA3994F462C71380E2FC0288FAA228890987B2F069DD9DCC7BBB12040A020801180A120410C09A0C1A40D04C5596D48ED1A46CE83040BA265BD5528D520AC2E627D34718D747006DEBDD71665E400375B3CD82ABC6528E7645B9C806A33820AA6468B549919480284646")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Service.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func CallService(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AD9010AD6010A1F2F697269736D6F642E736572766963652E4D736743616C6C5365727669636512B2010A0A736572766963652D796D122A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E1A2A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E22387B22686561646572223A7B7D2C226964223A226373303031222C226E616D65223A22737A6332222C2264617461223A22646174616373227D2A100A057569726973120731303030303030306412580A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103F5330F7602F5AA3994F462C71380E2FC0288FAA228890987B2F069DD9DCC7BBB12040A0208011811120410C09A0C1A40F6AA651B64438E9ECB588B12DFFAEE1D727E016A25F2439D50C0E3D07778EE5D3909D29633DBAD35C5CEC758A558ACE9501A9742040E7F9049B4EE9C9FA7B172")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Service.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func PauseRequestContext(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AAC010AA9010A272F697269736D6F642E736572766963652E4D7367506175736552657175657374436F6E74657874127E0A503145334245453439454436443245443137443533344631344339313244443931313736464633314434304537463438353535314430423044323731303235353030303030303030303030303030303030122A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E12580A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103F5330F7602F5AA3994F462C71380E2FC0288FAA228890987B2F069DD9DCC7BBB12040A0208011814120410C09A0C1A400D28196D289999903FCF00C368DC5E897AFE9C1CADF25506950E00819E69A8F62E93C1AD51DF7D27A92424C267A34E05A12040C8A1F32ADD0CD1C1154EDA15A0")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Service.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func StartRequestContext(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AAC010AA9010A272F697269736D6F642E736572766963652E4D7367537461727452657175657374436F6E74657874127E0A503145334245453439454436443245443137443533344631344339313244443931313736464633314434304537463438353535314430423044323731303235353030303030303030303030303030303030122A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E12580A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103F5330F7602F5AA3994F462C71380E2FC0288FAA228890987B2F069DD9DCC7BBB12040A0208011815120410C09A0C1A408CFE8CD7CEDC9FC902549FE50651750675917335D58630EFFA30F2D081F6B29B1C4440520B95F82B97E8E0C060AA2CF48AF7585F010A4CA31241615CC3FD149A")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Service.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func KillRequestContext(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AAB010AA8010A262F697269736D6F642E736572766963652E4D73674B696C6C52657175657374436F6E74657874127E0A503145334245453439454436443245443137443533344631344339313244443931313736464633314434304537463438353535314430423044323731303235353030303030303030303030303030303030122A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E12580A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103F5330F7602F5AA3994F462C71380E2FC0288FAA228890987B2F069DD9DCC7BBB12040A0208011816120410C09A0C1A401A82D464019B44DA658CA9B17600415AB3E07F03D8A53CA96F16A797A803C86B462F09A2ED8D6C0D8D9FF0DED2CCA0217F7BFCAAA6AB3015846554F954879C88")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Service.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func UpdateRequestContext(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AF2010AEF010A282F697269736D6F642E736572766963652E4D736755706461746552657175657374436F6E7465787412C2010A503145334245453439454436443245443137443533344631344339313244443931313736464633314434304537463438353535314430423044323731303235353030303030303030303030303030303030122A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E1A2A696161316878797874796779386476643674337672367A6D6B30617072347032756C756E76363036726E22100A05756972697312073230303030303030C80138C80112580A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103F5330F7602F5AA3994F462C71380E2FC0288FAA228890987B2F069DD9DCC7BBB12040A0208011813120410C09A0C1A406C37CDD590F5CF3ADC002CC49B8D4F98AD9170EF8D1CA70D92C9FEB666D41AC235D7E29395E2FE956A51689FD002EDB7591CAAE77A492A22BC4AB04701967707")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Service.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
