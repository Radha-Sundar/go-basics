package crypto

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	plaintext := `A paragraph is a collection of sentences all related to a central topic, idea, or theme. Paragraphs act as structural tools for writers to organize their thoughts into an ideal progression, and they also help readers process those thoughts effortlessly. Imagine how much harder reading and writing would be if everything was just one long block of text. `
	ct := EncryptAes128Cbc("12349876", plaintext)
	pr := fmt.Sprintf("ENCRYPTED=%v", ct)
	fmt.Println(pr)
}

func TestDecryptAes128Cbc(t *testing.T) {
	ct := "7e4e431f4f4a8e583557f75141e1ac41f949a1171ad0248c81ced40ae7d1727c97a5929efd8af22c78d8a32dcf08b5dd73e530fc82de8e3fe6bf3f846d1c0eb6be5fabf915189ed97c4ad662f5fbc3008df0bee410cba8fb0b1a4992d8fb83e0cf99546638cc58cc0dece6d14c0a95dc4d224e5cae6e6c0450f1f1e21babfe8471449733575570a86ee79f477bdb7cef4720153e41e962e00fb41d0ea38254bf66b84e2eef6325dc918443d540dd2d17cb0a2df20b1b95dd2ef03779ad12feff57fb09553691b544a964b4c6f38de8a59ce14adbaa35003dfb878793f91b07d659c996cb24a8997f5352140787c7178b284eaed153e29527278a9050836da807db669e614382468eb8caeca758c0f8dbfb7b5bb93ac79a206fc015db6098e64cb6a1fbf584998fb19a2d33dd45e8e033b7cff7bcb3979fa938920fd64da62c1a07b6cfa6424c494d58d36345ead26b7f8566cbad0a77a0e39eeadf6aa0397d0a49792bdef6eb2d26265413f115c77e44\n"
	pt := DecryptAes128Cbc("12349876", ct)
	pr := fmt.Sprintf("DECRYPTED=%v", pt)
	fmt.Println(pr)
}

func TestSha256(t *testing.T) {
	type User struct {
		Name   string
		Age    int
		Active bool
	}
	req := &User{
		Name:   "Bob",
		Age:    10,
		Active: true,
	}
	b, _ := json.Marshal(req)
	pt := Sha256(b, []byte("12349876"))
	pr := fmt.Sprintf("Sha256=%v", pt)
	fmt.Println(pr)
}
