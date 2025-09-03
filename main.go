package main

import (
    "archive/zip"
    "bytes"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"

    "github.com/go-git/go-git/v5"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
    // 1. Clone o repositório
    repoURL := "https://github.com/joaobcandido/posto-de-gasolina.git"
    clonePath := "./repo-clone"

    // Remove a pasta se já existir
    os.RemoveAll(clonePath)

    _, err := git.PlainClone(clonePath, false, &git.CloneOptions{
        URL:      repoURL,
        Progress: os.Stdout,
    })
    if err != nil {
        log.Fatalf("Erro ao clonar repositório: %v", err)
    }
    fmt.Println("Repositório clonado!")

    // 2. Compactar o diretório em um ZIP (em memória)
    var buf bytes.Buffer
    zipWriter := zip.NewWriter(&buf)

    err = filepath.Walk(clonePath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        relPath, _ := filepath.Rel(clonePath, path)
        file, err := os.Open(path)
        if err != nil {
            return err
        }
        defer file.Close()

        zipFile, err := zipWriter.Create(relPath)
        if err != nil {
            return err
        }
        _, err = io.Copy(zipFile, file)
        if err != nil {
            return err
        }
        return nil
    })
    if err != nil {
        log.Fatalf("Erro ao criar ZIP: %v", err)
    }
    zipWriter.Close()
    fmt.Println("ZIP criado em memória!")

    // 3. Configurar sessão S3 (LocalStack)
    sess, err := session.NewSession(&aws.Config{
        Region:           aws.String("us-east-1"),
        Endpoint:         aws.String("http://localhost:4566"),
        S3ForcePathStyle: aws.Bool(true),
        Credentials:      credentials.NewStaticCredentials("test", "test", ""),
    })
    if err != nil {
        log.Fatalf("Erro ao criar sessão AWS: %v", err)
    }
    s3Client := s3.New(sess)

    // 4. Fazer upload do ZIP para o S3
    bucket := "meu-bucket"
    zipKey := "repo-clone.zip"
    _, err = s3Client.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(zipKey),
        Body:   bytes.NewReader(buf.Bytes()),
    })
    if err != nil {
        log.Fatalf("Erro ao enviar ZIP para o S3: %v", err)
    }
    fmt.Printf("Arquivo %s enviado para o S3!\n", zipKey)
}