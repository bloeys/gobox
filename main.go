package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	genUuidv1Count = flag.Uint("uuidv1", 0, "Generates N UUID V1s. Use -uuidv1 N for N UUIDs")
	genUuidv4Count = flag.Uint("uuidv4", 0, "Generates N UUID V4s. Use -uuidv4 N for N UUIDs")
	genUuidv6Count = flag.Uint("uuidv6", 0, "Generates N UUID V6s. Use -uuidv6 N for N UUIDs")
	genUuidv7Count = flag.Uint("uuidv7", 0, "Generates N UUID V7s. Use -uuidv7 N for N UUIDs")
	decodeJwtToken = flag.String("decode-jwt", "", "Decodes the provided JWT token and prints its JSON without verifying")
)

func main() {

	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	if *genUuidv1Count != 0 {
		genUuids(1, *genUuidv1Count)
	}

	if *genUuidv4Count != 0 {
		genUuids(4, *genUuidv4Count)
	}

	if *genUuidv6Count != 0 {
		genUuids(6, *genUuidv6Count)
	}

	if *genUuidv7Count != 0 {
		genUuids(7, *genUuidv7Count)
	}

	if *decodeJwtToken != "" {
		decodeJwt(*decodeJwtToken)
	}
}

func genUuids(version uint, count uint) {

	var genFunc func() (uuid.UUID, error)

	switch version {
	case 1:
		genFunc = uuid.NewUUID
	case 4:
		genFunc = uuid.NewRandom
	case 6:
		genFunc = uuid.NewV6
	case 7:
		genFunc = uuid.NewV7
	default:
		panic(fmt.Errorf("unknown uuid version=%d. Supported values are: 1, 4, 6, 7", version))
	}

	for i := 0; i < int(count); i++ {
		id := uuid.Must(genFunc())
		fmt.Println(id.String())
	}
}

func decodeJwt(token string) {

	type PrintingDecodedToken struct {
		Header jwt.MapClaims
		Body   jwt.Claims
	}

	jwtParser := jwt.NewParser()

	mapClaims := jwt.MapClaims{}
	decodedToken, _, err := jwtParser.ParseUnverified(token, &mapClaims)
	panicIfErr(err)

	printingDecodedToken := PrintingDecodedToken{Header: decodedToken.Header, Body: decodedToken.Claims}

	headerPrettyJsonBytes, err := json.MarshalIndent(&printingDecodedToken, "", "  ")
	panicIfErr(err)

	fmt.Println(string(headerPrettyJsonBytes))
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
