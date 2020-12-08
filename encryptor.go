package main

import (
    "errors"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strings"

    "github.com/Funny-Systems-OSS/CloudSQL-Proxy-Hardening-Common"
)

var (
    version = flag.Bool("v", false, "Print the version of the proxy and exit.")
    instanceID = flag.Int("i", -1, "The instance ID which the cloud_sql_proxy will be set.")
    credentialFilePath = flag.String("f", "", "The json file be used to retrieve Service Account credential in cloud_sql_proxy.")
    outputFilePath = flag.String("o", "", "If provided, it is treated as the store path of encrypted file. Default to be the same place as the input with filename '<FILENAME>.encrypted'.")
)

const (
    versionString = "1.0.0"

    usage = `
Usage:
  encrypt_funny -f [credential file] -i [instance ID]

Options:
`
)

func init(){
    fmt.Println(funny.Funny)
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, usage)
        flag.VisitAll(func(f *flag.Flag) {
            usage := strings.Replace(f.Usage, "\n", "\n    ", -1)
            fmt.Fprintf(os.Stderr, "  -%s\n    %s\n\n", f.Name, usage)
        })
    }
}

func checkFlags() error {
    if *instanceID == -1 {
        return errors.New("Must specify ID of credential Instance.")
    }
    
    if *credentialFilePath == "" {
        return errors.New("Must specify path of credential file.")
    }

    if *outputFilePath == "" {
        *outputFilePath = *credentialFilePath + ".encrypted"
    }
    return nil
}

func readDataFromFile(filepath string) ([]byte, error) {
    return ioutil.ReadFile(filepath)
}

func writeDataToFile(filepath string, data []byte) error {
    return ioutil.WriteFile(filepath, data, 666)
}

func main() {
    flag.Parse()

    if *version {
        fmt.Println("Encrypt Funny:", versionString)
        return
    }

    if err := checkFlags(); err != nil {
        log.Fatal(err)
    }

    log.Printf("Reading file from \"%s\".\n", *credentialFilePath)
    plaintext, err := readDataFromFile(*credentialFilePath)
    if err != nil {
        log.Fatal("File not found.")
    }

    key := funny.KeyGenerator(*instanceID + 69)
    nonce := funny.NonceGenerator(*instanceID + 6969)

    log.Println("Encrypting file...")
    ciphertext := funny.Encrypt(plaintext, []byte(key), []byte(nonce))
    if err = writeDataToFile(*outputFilePath, []byte(ciphertext)); err != nil {
        log.Fatal(err)
    }
    log.Println("Done.")
    log.Printf("\"%s\" saved.\n", *outputFilePath)
    
    log.Println("Validating...")
    byteCiphertext, err := readDataFromFile(*outputFilePath)
    if err != nil {
        log.Fatal("Output file not found.")
    }
    log.Println("OK.")

    if !funny.Validate(funny.Decrypt(byteCiphertext, []byte(key), []byte(nonce)), plaintext){
        log.Println("Some shit happened. The enrypted file might not work.")
    } else {
        log.Println("Task complete.")
    }
}
