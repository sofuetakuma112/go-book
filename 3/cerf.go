package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{ // 識別名を作成
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	// SSL証明書はX509をベースにしているのでcrypto/x509ライブラリを利用して証明書を作る
	template := x509.Certificate{ // 証明書の構成を設定するための構造体Certificate
		SerialNumber: serialNumber, // 認証局によって発行される一意の番号(ローカルでHTTPSを試すだけなら適当な大きい数字でOK?)
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour), // 有効期限を証明書が作成された日から1年間とする
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // X509証明書がサーバー認証に使用されることを示す
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, // X509証明書がサーバー認証に使用されることを示す
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")}, // 作成される証明書を127.0.0.1でだけ効力を持つように設定する
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048) // RSAの秘密鍵を生成(公開鍵も含まれている) CAが本来発行するもの

	// ここで作成された証明書は1つの秘密鍵と公開鍵のペアを、証明書の作成用途と、共通鍵の作成用途に使いまわしているため、サーバーのなりすましにクライアントが気付けないのが問題
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk) // CAが発行した秘密鍵を用いて証明書(CAが発行したものとは別の公開鍵を含んでいる)を作成する(DER形式のバイトデータのスライスを生成する)
	certOut, _ := os.Create("cert.pem") // ファイルが存在しない場合、モード0666（umaskの前）で作成され、返されたFileのメソッドは、I/Oに使用することができる。
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)}) // 証明書を発行するときに使用した秘密鍵をPEM符号化
	keyOut.Close()
}
