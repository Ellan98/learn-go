/*
 * @Date: 2024-08-05 08:11:51
 * @LastEditTime: 2024-08-05 08:19:45
 * @FilePath: \gemini\main.go
 * @description: 注释
 */

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	// option.WithAPIKey(os.Getenv("")
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyAniWlvp8fsdAuBGmCG3TMWj2ASW7nkcoI"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a AI and magic"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("响应：", resp)
}
