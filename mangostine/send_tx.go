package mangostine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

type Transaction struct {
	Amount    uint64
	Address   []byte
	PublicKey []byte
	Signature []byte
	Sequence  uint64
	Nonce     uint64
}

// keysCmd represents the keys command
func SendTxCommand() *cobra.Command {

	sendCmd := &cobra.Command{
		Use:   "sendTx",
		Short: "send transcation network",
		Long:  `send transcation`,
		RunE: func(_ *cobra.Command, _ []string) error {
			kp := []Keypair{}
			file, _ := ioutil.ReadFile(DefaultHome + "key.json")
			err := json.Unmarshal([]byte(file), &kp)
			if err != nil {
				fmt.Println("reading json key file ", err)
			}
			fmt.Println("reading key file %v", kp)
			// for i := 0; i < len(kp); i++ {

			// }
			return nil
		},
	}

	return sendCmd
}

func make_transaction() {

}
