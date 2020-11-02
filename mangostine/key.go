package mangostine

import (
	"bytes"
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip39"
)

var (
	DefaultHome = os.ExpandEnv("$HOME/.mangostine/")
)

type Keypair struct {
	Mnemonic   string `json:"mnemonic"`
	Privatekey []byte `json:"privatekey"`
	Publickkey []byte `json:"publickkey"`
}

// keysCmd represents the keys command
func CreateKeyPairCommand() *cobra.Command {

	keysCmd := &cobra.Command{
		Use:   "keypair",
		Short: "create the account with mnemonic, private key, public key and address",
		Long:  `create the account with mnemonic, private key, public key and address`,
		RunE: func(_ *cobra.Command, _ []string) error {

			node, err := createNode()
			kp := []Keypair{}
			for j := 0; j < node; j++ {
				if err != nil {
					fmt.Println("keypair creating", err)
				}
				_ = strconv.Itoa(j)
				entropy, _ := bip39.NewEntropy(256)
				mnemonic, _ := bip39.NewMnemonic(entropy)
				seed := bip39.NewSeed(mnemonic, "123456")
				first := seed[:32]
				privatekey := ed25519.NewKeyFromSeed(first)
				publicKey := ed25519.PublicKey(privatekey)
				keys := Keypair{Mnemonic: mnemonic, Privatekey: privatekey, Publickkey: publicKey}
				kp = append(kp, keys)

			}

			kpjson, _ := json.Marshal(kp)
			outputDocumet, err := makeOutputFilepath(DefaultHome)
			err = ioutil.WriteFile(outputDocumet, kpjson, 0644)
			if err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}

	return keysCmd
}

func makeOutputFilepath(rootDir string) (string, error) {
	writePath := filepath.Join(rootDir)
	if err := EnsureDir(writePath, 0700); err != nil {
		return "", err
	}
	return filepath.Join(writePath, fmt.Sprintf("key.json")), nil
}

//Take the input from the terminal for creating the node
func createNode() (int, error) {
	var n int
	fmt.Println("Enter an integer value for keypair creation  : ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
	}
	if n > 20000 {
		fmt.Println("please enter below 15")
		os.Exit(1)
	}

	fmt.Println("You have entered : ", n)
	return n, nil
}

func WriteDataToFileAsJSON(data interface{}, filedir string) (int, error) {
	//write data as buffer to json encoder
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return 0, err
	}
	file, err := os.OpenFile(filedir, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return 0, err
	}
	n, err := file.Write(buffer.Bytes())
	if err != nil {
		return 0, err
	}
	return n, nil
}
