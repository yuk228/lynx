package handler

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/google/uuid"

	"github.com/yuk228/lynx/config"
)

func GetBinary(message string, speaker string) ([]byte, error) {
	urlParts := []string{config.VoiceBoxURL, "/audio_query?text=", url.QueryEscape(message), "&speaker=", speaker}
	url := strings.Join(urlParts, "")
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Set("accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error getting query: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	synthUrlParts := []string{config.VoiceBoxURL, "/synthesis?speaker=", speaker, "&enable_interrogative_upspeak=true"}
	synthUrl := strings.Join(synthUrlParts, "")
	req_s, _ := http.NewRequest(http.MethodPost, synthUrl, resp.Body)
	req_s.Header.Set("accept", "audio/wav")
	req_s.Header.Set("Content-Type", "application/json")
	resp_s, err := client.Do(req_s)
	if err != nil {
		log.Printf("error getting binary: %v", err)
		return nil, err
	}

	defer resp_s.Body.Close()

	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, resp_s.Body)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func ToWav(b []byte, guildId string) (string, error) {
	uuid := uuid.New()
	path := fmt.Sprintf("%s_%s.wav", guildId, uuid)
	file, err := os.Create(path)
	if err != nil {
		log.Printf("error creating wav file: %v", err)
		return "", err
	}
	defer file.Close()
	file.Write(b)
	return path, nil
}
