/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"log"
	"os"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "115.190.57.118:9000"
		accessKeyID := "minioadmin"
		secretAccessKey := "minioadmin"
		useSSL := false

		// Initialize minio client object.
		minioClient, err := minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("%#v\n", minioClient) // minioClient is now setup
		filepath := args[0]
		file, err := os.Open(filepath)
		if err != nil {
			panic("文件打开")
		}
		defer file.Close()

		fileStat, err := file.Stat()
		if err != nil {
			fmt.Println(err)
			return
		}
		buckname := "yuhang"
		typer := "image/png"
		objectName := time.Now().Format("2006-01-02-15-04-05") + "图片"
		uploadInfo, err := minioClient.PutObject(context.Background(), buckname, objectName, file, fileStat.Size(), minio.PutObjectOptions{ContentType: typer})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Successfully uploaded bytes: ", uploadInfo)
		fmt.Println("upload called")
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
