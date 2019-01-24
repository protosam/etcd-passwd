package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/protosam/etcd-passwd"

	"github.com/kless/osutil/user/crypt/sha512_crypt"
	"crypto/rand"
	"encoding/hex"

)

var (
	flagEtcdServer = flag.String("etcd-server", "http://localhost:2379", "etcd endpoint")
	flagEtcdPrefix = flag.String("etcd-prefix", "/etcd-passwd", "etcd prefix")

	flagName  = flag.String("name", "", "Name")
	flagPassword  = flag.String("password", "!!", "Password")
	flagUID   = flag.Int("uid", -1, "UID")
	flagGID   = flag.Int("gid", -1, "GID")
	flagGecos = flag.String("gecos", "", "Gecos")
	flagDir   = flag.String("home", "", "Home")
	flagShell = flag.String("shell", "/bin/sh", "Login shell")
)

func run() error {
	home := filepath.Join("/home", *flagName)
	if *flagName == "" {
		return errors.New("invalid name")
	}
	if *flagUID < 0 {
		return errors.New("invalid UID")
	}
	if *flagGID < 0 {
		return errors.New("invalid GID")
	}
	if *flagDir != "" {
		home = *flagDir
	}

	var password string
	var err error
	if *flagPassword == "!!" || *flagPassword == "" {
		password = "!!"
	}else{
		password, err = shadow_word(*flagPassword)
		if err != nil {
			panic(err)
		}
	}
	
	return etcdpasswd.AddUser(&etcdpasswd.Passwd{*flagName, password, etcdpasswd.UID(*flagUID), etcdpasswd.GID(*flagGID), *flagGecos, home, *flagShell})
}

func main() {
	flag.Parse()

	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}



func shadow_word(password string) (string, error) {
	b := make([]byte, 12)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	shadow_word := hex.EncodeToString(b)

	c := sha512_crypt.New()
	hash, err := c.Generate([]byte(password), []byte("$6$" + shadow_word))

	return hash, err
}
