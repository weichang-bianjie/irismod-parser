package integration

import (
	"encoding/hex"
	"fmt"
	. "github.com/kaifei-bianjie/common-parser/codec"
	"github.com/kaifei-bianjie/common-parser/utils"
)

func (s IntegrationTestSuite) TestEvm() {
	cases := []SubTest{
		{
			"EthereumTx",
			EthereumTx,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func EthereumTx(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)

	txBytes, err := hex.DecodeString("0AAF040AFB030A1F2F65746865726D696E742E65766D2E76312E4D7367457468657265756D547812D7030A87030A1A2F65746865726D696E742E65766D2E76312E4C6567616379547812E80208B59D0412013118F0D71C222A3078304537363246344431313433394231313330443430323939353332386236333463423963393937332A013032E401F242432A0000000000000000000000006BAD6587869C76F83C37A66C89296A84BB9320A50000000000000000000000005C66430EA2CF2AAD776CE10F78D14FDB1E3467C8000000000000000000000000000000000000000000000000000000000003309C000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000A000000000000000000000000000000000000000000000000000000000000000030A0A0A00000000000000000000000000000000000000000000000000000000003A011C422011CE6864FE8CA1646B72672AA4B9EACC10A5E397CE5B672E430F24E4E069AD674A203BF7539451B8961CEDC8B6045CD3B03410E4A4132A69DBA861117658CA5C9C71110000000000B074401A42307861396136613036366430323730613565336164333830386532306636373239383135373931316135393835303634316431626561386664323662353033333238FA3F2E0A2C2F65746865726D696E742E65766D2E76312E457874656E73696F6E4F7074696F6E73457468657265756D5478121612140A0E0A0475676173120634373030303010F0D71C")
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	for _, msg := range authTx.GetMsgs() {
		if mtDoc, ok := s.Evm.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(mtDoc))
		}
	}
}