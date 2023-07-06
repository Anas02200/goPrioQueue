package SDiffApiClient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var Wg sync.WaitGroup

func ExecuteRequest() {
	Wg.Add(1)

	callSdClient()

}

func callSdClient() {
	log.Println("Calling SD client")
	//defer Wg.Done()
	img := StableDiffusionProcessingTxt2Img{
		Prompt: "galaxy imploding",
	}
	marshal, _ := json.Marshal(img)

	post, err := http.Post("http://127.0.0.1:7860/sdapi/v1/txt2img", "application/json", bytes.NewBuffer(marshal))
	defer post.Body.Close()

	log.Println(post.StatusCode)

	if post.StatusCode == http.StatusOK {

		//bodyBytes, err := io.ReadAll(post.Body)

		decoder := json.NewDecoder(post.Body)

		var response TextToImageResponse

		//choose one of these unmarshall or decode

		//json.Unmarshal(bodyBytes, response)
		decoder.Decode(response)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(response.Info)
		log.Println(response.Images)

	} else {
		log.Println(err, post.StatusCode)
	}
}
