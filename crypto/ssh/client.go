package service

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"time"

	log "github.com/ckeyer/logrus"
	"golang.org/x/crypto/ssh"
)

func GenerateSSHKey() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Errorf("generate rsa key failed, %s", err)
		return "", "", err
	}

	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	buf := &bytes.Buffer{}
	err = pem.Encode(buf, block)
	if err != nil {
		return "", "", err
	}

	priKey := buf.String()
	sg, err := ssh.ParsePrivateKey(buf.Bytes())
	if err != nil {
		return "", "", err
	}

	buf.Reset()
	buf.WriteString(sg.PublicKey().Type())
	buf.WriteString(" ")
	buf.WriteString(base64.StdEncoding.EncodeToString(sg.PublicKey().Marshal()))
	buf.WriteString(" ")
	buf.WriteString("robot@nevis.io")
	pubKey := buf.String()

	return priKey, pubKey, nil
}

func RSAAuthMethod(priKey string) (ssh.AuthMethod, error) {
	sg, err := ssh.ParsePrivateKey([]byte(priKey))
	if err != nil {
		log.Errorf("parse private key failed, %s", err)
		return nil, err
	}
	return ssh.PublicKeys(sg), nil
}

type SSHClient struct {
	*ssh.Client
	*ssh.Session
	WorkDir string
}

func NewSSHClient(user, endporint, priKey string) (*SSHClient, error) {
	sg, err := RSAAuthMethod(priKey)
	if err != nil {
		log.Errorf("get rsa auth method failed, %s", err)
		return nil, err
	}

	cfg := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{sg},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 5,
	}

	cli, err := ssh.Dial("tcp", endporint, cfg)
	if err != nil {
		log.Errorf("dial ssh failed, %s", err)
		return nil, err
	}

	ss, err := cli.NewSession()
	if err != nil {
		log.Errorf("get ssh session failed, %s", err)
		return nil, err
	}

	return &SSHClient{cli, ss, ""}, nil
}

func (s *SSHClient) Close() {
	if s.Session != nil {
		s.Session.Close()
	}
	if s.Client != nil {
		s.Client.Close()
	}
}
