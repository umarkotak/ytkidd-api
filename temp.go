package main

import (
	"context"
	"fmt"
	"os"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"google.golang.org/api/option"
)

// TextToSpeech generates speech from text and saves it to a file.
// text: The text to be converted to speech.
// filePath: The path where the output audio file (MP3) will be saved.
func TextToSpeech(text, filePath string) error {
	ctx := context.Background()

	// --- SETUP & AUTHENTICATION ---
	// The client will attempt to authenticate in the following order:
	// 1. From an environment variable `GCP_CREDENTIALS_JSON` containing the service account key.
	// 2. From the file path specified in `GOOGLE_APPLICATION_CREDENTIALS`.
	// 3. From Application Default Credentials (ADC) if on a GCP environment or configured locally.
	//
	// You will also need to enable the Text-to-Speech API in your Google Cloud project
	// and install the necessary Go packages:
	// go get cloud.google.com/go/texttospeech/apiv1
	// go get google.golang.org/api/option

	var client *texttospeech.Client
	var err error

	// Check for credentials in the environment variable.
	credsJSON := os.Getenv("GOOGLE_SERVICE_ACCOUNT")
	if credsJSON != "" {
		// If the env var is set, use the JSON credentials directly.
		client, err = texttospeech.NewClient(ctx, option.WithCredentialsJSON([]byte(credsJSON)))
	} else {
		// Otherwise, fall back to default authentication (ADC, file path, etc.).
		client, err = texttospeech.NewClient(ctx)
	}

	if err != nil {
		// If client creation fails, return the error.
		return fmt.Errorf("failed to create texttospeech client: %w", err)
	}
	// It's good practice to close the client when the function finishes.
	defer client.Close()

	// --- BUILD THE REQUEST ---
	// The SynthesizeSpeechRequest contains the text, voice, and audio configuration.
	req := &texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("NEUTRAL"). You can change these to other voices.
		// For a full list of available voices, see:
		// https://cloud.google.com/text-to-speech/docs/voices
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "id-ID",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		// Select the type of audio file you want to create.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	// --- PERFORM THE API CALL ---
	// Call the SynthesizeSpeech method to generate the audio.
	resp, err := client.SynthesizeSpeech(ctx, req)
	if err != nil {
		// If the API call fails, return the error.
		return fmt.Errorf("failed to synthesize speech: %w", err)
	}

	// --- SAVE THE AUDIO TO A FILE ---
	// The response contains the audio content as a byte slice.
	// We use os.WriteFile to save this data to the specified file path.
	// The file permissions 0644 are standard for a file that is readable by everyone
	// but only writable by the owner.
	err = os.WriteFile(filePath, resp.AudioContent, 0644)
	if err != nil {
		return fmt.Errorf("failed to write audio file: %w", err)
	}

	fmt.Printf("Successfully generated speech and saved to %s\n", filePath)
	return nil
}

// main function to demonstrate the TextToSpeech function.
// func main() {
// 	// Define the text you want to convert and the output file name.
// 	textToConvert := "Nabi Adam adalah manusia pertama yang diciptakan oleh Allah dari tanah liat. Allah membentuk tubuhnya dengan sempurna, lalu meniupkan ruh ke dalamnya sehingga ia hidup. Allah mengajarkan kepada Adam nama-nama segala sesuatu, menjadikannya khalifah di bumi, dan memberinya pasangan bernama Hawa. Mereka berdua tinggal di surga, hidup dengan segala kenikmatan, dan diperbolehkan memakan apa saja kecuali buah dari satu pohon yang terlarang. Namun, iblis menggoda keduanya hingga mereka memakan buah tersebut."
// 	outputFilePath := "output.mp3"

// 	// Call the function.
// 	err := TextToSpeech(textToConvert, outputFilePath)
// 	if err != nil {
// 		// If an error occurs, log it and exit.
// 		log.Fatalf("An error occurred: %v", err)
// 	}
// }
