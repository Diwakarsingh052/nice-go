package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
)

const tokenStr = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcGkgcHJvamVjdCIsInN1YiI6IjEwMSIsImV4cCI6MTY5NDE3NDM3NiwiaWF0IjoxNjk0MTcxMzc2LCJyb2xlcyI6WyJVU0VSIl19.K_Tnlg2inFM8MB64q7TH7f-s45ejRPBaIxpRlu7EIAKjI1leWkhZmv_u3uzm_qMUaGNtgTxmsNfI6T4CADdRK03rsWe9jD3fuNC0Ay9cPGJ1neUzC97ZTH5tWvJSaXN82f2oyQ36-qdkExJodntgL5rYDdebok8CvfukCMSaum1C8SR4S_vTS3cxWDrLwPdAfC9LI1tqwsRxrx7AqO0TNwzwq3AVWGpLKUbVuqdQ-fOhWa09d_ojBF_-rUrN35Hm2HDPf4eazhWXlAXMnJXIoDTLlnleTeMVvYxLn8DuIrCxBQT-hNvX8sOF0mr6Npyntcv7TzFUpSF414Nz6Eqebg`

type claims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

func main() {
	PublicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		log.Fatalln("not able to read pem file")
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(PublicPEM)
	_ = publicKey
	if err != nil {
		log.Fatalln(err)
	}

	var c claims

	token, err := jwt.ParseWithClaims(tokenStr, &c, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		fmt.Println("parsing token", err)
		return

	}
	if !token.Valid {
		fmt.Println("invalid token")
		return
	}
	fmt.Printf("%+v", c)

}
