package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"

	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ghodss/yaml"
	"github.com/jessevdk/go-flags"
	"github.com/pelletier/go-toml"

	// "github.com/russross/blackfriday/v2"
	img "github.com/OhYee/goldmark-image"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"

	goldmarkHTML "github.com/yuin/goldmark/renderer/html"
	"golang.org/x/net/html"
)

var config *toml.Tree

type Post struct {
	ID           string `json:"id"`
	UpdatedAt    string `json:"updated_at"`
	Mobiledoc    string `json:"mobiledoc"`
	FeatureImage string `json:"feature_image"`
}

type ResponseData struct {
	Posts []Post `json:"posts"`
}

var opts struct {
	ShowConfig bool `short:"c" long:"config" description:"Show glee configuration global file"`
	Debug      bool   `short:"d" long:"debug" description:"debug mode"`
	File       string `description:"Markdown file to process"`
	Help 	  	bool 	`short:"h" long:"help" description:"Show this help message"`
	Version 	bool	`short:"v" long:"version" description:"Show version"`
}

var postsApiBase string

var buf bytes.Buffer
var theme string
var md = goldmark.New(

	goldmark.WithExtensions(
		extension.GFM, // Includes table, fenced code, and code highlight extensions
		extension.Table,
		img.NewImg("image", nil),
		// highlighting.NewHighlighting(
		// 	highlighting.WithStyle(theme), //theme list https://github.com/yuin/goldmark-highlighting/issues/37

		// ),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(), // Enable automatic heading IDs for TOC
	),
	goldmark.WithRendererOptions(
		goldmarkHTML.WithHardWraps(),
		goldmarkHTML.WithUnsafe(), // This also enables unsafe rendering
	),
)

func getPostId(slug string, headers http.Header) (*Post, error) {
	ghostVersion := config.GetDefault("ghost-configuration.GHOST_VERSION", "").(string)
	ghostUrl := config.GetDefault("ghost-configuration.GHOST_URL", "").(string)

	if ghostVersion == "v5" {
		postsApiBase = ghostUrl + "/api/admin/posts/"
	} else {
		postsApiBase = ghostUrl + "/api/" + ghostVersion + "/admin/posts/"
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", postsApiBase+"slug/"+slug+"/", nil)
	if err != nil {
		return nil, err
	}
	req.Header = headers

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to communicate with the Ghost Admin API: %s", body)
	}

	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if len(data.Posts) > 0 {
		return &data.Posts[0], nil
	}

	return nil, errors.New("no posts found for the given slug")
}

func toHTML(markdown string) string {
	// fmt.Println("Converting markdown to HTML...", markdown)
	start := "<!--kg-card-begin: html-->"
	end := "<!--kg-card-end: html-->"
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		panic(err)
	}
	// fmt.Printf("HTML: %s\n", buf.String())
	return start + buf.String() + end // Return the generated HTML string
}

func initLogging(debug bool) {
	log.SetOutput(os.Stdout)

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// You can customize the log formatter if needed
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

func addBlogConfigurations(meta map[string]interface{}) map[string]interface{} {
	if config == nil {
		log.Fatal("Configuration is not initialized. Call loadGlobalConfig to initialize it.")
	}

	globalSidebarTOC := config.GetDefault("blog-configuration.SIDEBAR_TOC", "").(bool)
	// globalFeatured := config.GetDefault("blog-configuration.FEATURED", "").(bool)
	// globalStatus := config.GetDefault("blog-configuration.STATUS", "").(string)

	defaultStyle := `pre { line-height: 125%; }
   td.linenos .normal { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   span.linenos { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   td.linenos .special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   span.linenos.special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   .article-image {
   max-width: 600px;
   margin: 0 auto !important;
   float: none !important;
   }`

	sidebarTocHead := `<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/tocbot/4.12.3/tocbot.css">
<style>
   .gh-content {
   position: relative;
   }
   .gh-toc > .toc-list {
   position: relative;
   }
   .toc-list {
   overflow: hidden;
   list-style: none;
   }
   @media (min-width: 1300px) {
   .gh-sidebar {
   position: absolute; 
   top: 0;
   bottom: 0;
   margin-top: 4vmin;
   margin-left: 20px;
   grid-column: wide-end / main-end; /* Place the TOC to the right of the content */
   width: inline-block;
   white-space: nowrap;
   }
   .gh-toc-container {
   position: sticky; /* On larger screens, TOC will stay in the same spot on the page */
   top: 4vmin;
   }
   }
   .gh-toc .is-active-link::before {
   background-color: var(--ghost-accent-color); /* Defines TOC accent color based on Accent color set in Ghost Admin */
   } 
</style>`
	sidebarTocFooter := `<script src="https://cdnjs.cloudflare.com/ajax/libs/tocbot/4.12.3/tocbot.min.js"></script>
<script>
   const parent = document.querySelector(".gh-content.gh-canvas");
   // Create the <aside> element
   const asideElement = document.createElement("aside");
   asideElement.setAttribute("class", "gh-sidebar");
   //asideElement.style.zIndex = 0; // sent to back so it doesn't show on top of images
   
   // Create the container div for title and TOC
   const containerElement = document.createElement("div");
   containerElement.setAttribute("class", "gh-toc-container");
   
   // Create the title element
   const titleElement = document.createElement("div");
   titleElement.textContent = "Table of Contents";
   titleElement.style.fontWeight = "bold";
   containerElement.appendChild(titleElement);
   
   // Create the <div> element for TOC
   const divElement = document.createElement("div");
   divElement.setAttribute("class", "gh-toc");
   containerElement.appendChild(divElement);
   
   // Append the <div> element to the <aside> element
   asideElement.appendChild(containerElement);
   parent.insertBefore(asideElement, parent.firstChild);
   
   tocbot.init({
       // Where to render the table of contents.
       tocSelector: '.gh-toc',
       // Where to grab the headings to build the table of contents.
       contentSelector: '.gh-content',
       // Which headings to grab inside of the contentSelector element.
       headingSelector: 'h1, h2, h3, h4',
       // Ensure correct positioning
       hasInnerContainers: true,
   });
   
   // Get the table of contents element
   const toc = document.querySelector(".gh-toc");
   const sidebar = document.querySelector(".gh-sidebar");
   
   // Check the number of items in the table of contents
   const tocItems = toc.querySelectorAll('li').length;
   
   // Only show the table of contents if it has more than 5 items
   if (tocItems > 2) {
     sidebar.style.display = 'block';
   } else {
     sidebar.style.display = 'none';
   }
</script>`

	if existingHead, ok := meta["codeinjection_head"].(string); ok {
		meta["codeinjection_head"] = existingHead + "<style>" + defaultStyle + "</style>"
	} else {
		meta["codeinjection_head"] = "<style> " + defaultStyle + "</style>"
	}
	
	if globalSidebarTOC {
		if existingHead, ok := meta["codeinjection_head"].(string); ok {
			meta["codeinjection_head"] = existingHead + sidebarTocHead
			log.Debug("Done Sidebar TOC Code injection")
		} else {
			meta["codeinjection_head"] = sidebarTocHead
		}
		meta["codeinjection_foot"] = sidebarTocFooter
	}

	return meta
}

func loadGlobalConfig() {
	configPath := filepath.Join(os.Getenv("HOME"), ".glee.toml")
	var err error
	config, err = toml.LoadFile(configPath)
	if err != nil {
		getTOMLFile(configPath)
		os.Exit(0)
	}
}

func getTOMLFile(configPath string) {
	log.Error("The configuration file at %s was not found.\n", configPath)

	var configResponse string
	log.Info("Would you like me to create the configuration file? (yes/no): ")
	fmt.Scanln(&configResponse)

	if configResponse == "yes" || configResponse == "y" {
		// Your existing code to create the configuration file
		// ...

	} else {
		os.Exit(0)
	}
}

func makeRequest(headers http.Header, body map[string]interface{}, pid string, updated_at string) {
	var method, apiEndpoint string
	ghostVersion := config.GetDefault("ghost-configuration.GHOST_VERSION", "").(string)
	ghostUrl := config.GetDefault("ghost-configuration.GHOST_URL", "").(string)

	if ghostVersion == "v5" {
		postsApiBase = ghostUrl + "/api/admin/posts/"
	} else {
		postsApiBase = ghostUrl + "/api/" + ghostVersion + "/admin/posts/"
	}
	if pid == "" {
		method = http.MethodPost
		apiEndpoint = postsApiBase + "?source=html"
	} else {
		method = http.MethodPut
		body["posts"].([]map[string]interface{})[0]["updated_at"] = updated_at
		apiEndpoint = postsApiBase + pid + "?source=html"
	}

	client := &http.Client{}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		log.Error("Error marshaling body to JSON:", err)
		return
	}

	req, err := http.NewRequest(method, apiEndpoint, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Error("Error creating request:", err)
		return
	}
	req.Header = headers
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error executing request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		var errorMessage string
		errorBytes, _ := ioutil.ReadAll(resp.Body) 
		if err := json.Unmarshal(errorBytes, &errorMessage); err != nil {	
			errorMessage = string(errorBytes)
		}
		log.Error("Request failed with status: %s. Error message: %s\n", resp.Status, errorMessage)
		return
	}

	// Unmarshal the response
	var responseData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		log.Error("Error unmarshaling response:", err)
		return
	}

	// Log the result
	if pid == "" {
		log.Info("Created new post")
	} else {
		log.Info("Updated existing post based on slug")
	}
	fmt.Printf("Blog preview link: %s\n", responseData["posts"].([]interface{})[0].(map[string]interface{})["url"])
}

func injectMultiTitles(meta map[string]interface{}) error {
	meta["codeinjection_head"] = ""

	_, ok := meta["title"].(string)
	if !ok {
		titleDataMap, ok := meta["title"].(map[string]interface{})
		if !ok {
			log.Error("missing default title")
			return nil
		}

		defaultTitle, ok := titleDataMap["default"].(string)
		if !ok {
			log.Error("missing default title in multi-title")
			return nil
		}

		meta["title"] = defaultTitle

		titleDataBytes, err := json.Marshal(titleDataMap)
		if err != nil {
			log.Error("error marshaling title data: %v", err)
			return err
		}

		titleDataStr := string(titleDataBytes)
		log.Debug("multi title: ", titleDataStr)

		meta["codeinjection_head"] = fmt.Sprintf(`<script>
			changetitle(%s);
			</script>`, titleDataStr)
	}

	return nil
}

func postToGhost(metadata map[string]interface{}, content string) {
	injectMultiTitles(metadata)
	if val, ok := metadata["code_hilite_theme"]; ok {
		theme = val.(string)
		log.Debug("Theme from markdown:", theme)
	} else {
		codeTheme := config.GetDefault("blog-configuration.CODE_HILITE_THEME", "").(string)
		if codeTheme != "" {
			theme = codeTheme
		} else {
			theme = "monokai" // Default theme if none is specified
		}
		log.Debug("Default/Global Configured theme:", theme)
	}
	md = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			img.NewImg("image", nil),
			highlighting.NewHighlighting(
				highlighting.WithStyle(theme),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			goldmarkHTML.WithHardWraps(),
			goldmarkHTML.WithUnsafe(),
		),
	)

	metadata = addBlogConfigurations(metadata)
	metadata["html"] = toHTML(content)
	token, err := getJWToken()
	if err != nil {
		log.Error("Failed generate jwt token: %v", err)
	} else{
		log.Debug("Generated jwt token")
	}
	headers := http.Header{}
	headers.Set("Authorization", "Ghost "+token)
	post, err := getPostId(metadata["slug"].(string), headers)
	if err != nil {
		log.Error("Error: Unable to communicate with the Ghost Admin API. Please verify your Ghost configurations: %v\n", err)
		return
	}
	var pid, updated_at, htmlData, featureImage string
	if post == nil {
		log.Info("No post found for the given slug.")
	} else {
		pid, updated_at, htmlData, featureImage = post.ID, post.UpdatedAt, post.Mobiledoc, post.FeatureImage
	}
	if _, ok := metadata["feature_image"]; ok {
		uploadFeatureImage(metadata, token, featureImage)
	} else {
		log.Fatal("Feature image not provided")
		metadata["feature_image"] = ""
	}

	uploadedImages, err := uploadImages(token, htmlData)
	if err != nil{
		log.Error("Error uploading images:", err)
	} else{
		log.Info("Uploaded Blog Images")
	}
	result, err := replaceImageLinks(metadata, uploadedImages)
	if err != nil {
		log.Error("Error replacing image links:", err)
	}
	metadata["html"] = result

	postObj := metadata
	body := map[string]interface{}{
		"posts": []map[string]interface{}{postObj},
	}
	makeRequest(headers, body, pid, updated_at)

}

func replaceImageLinks(metadata map[string]interface{}, imgMap map[string]string) (string, error) {
	
	htmlStr, ok := metadata["html"].(string)
	if !ok {
		return "", errors.New("metadata does not contain a string under the 'html' key")
	}

	
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return "", err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for i, attr := range n.Attr {
				switch {
				case n.Data == "img" && attr.Key == "src":
					fallthrough
				case n.Data == "a" && attr.Key == "href":
					if newSrc, ok := imgMap[attr.Val]; ok {
						n.Attr[i].Val = newSrc
					}
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			f(child)
		}
	}
	f(doc)

	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func uploadImages(token, htmlData string) (map[string]string, error) {
	uploadedImages := make(map[string]string)
	imageBackend := config.GetDefault("image-configuration.IMAGE_BACKEND", "").(string)
	blogImageList := []string{}

	pattern := `<img[^>]+src="([^"]+)"[^>]*>`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(buf.String(), -1)

	if imageBackend == "ghost" {
		var err error
		blogImageList, err = getImageFromPost(htmlData)
		if err != nil {
			log.Error("Error getting images from post:", err)
			
		}
	}
	var mdlibImages []string
	for _, m := range matches {
		if len(m) > 1 {
			mdlibImages = append(mdlibImages, m[1])
		}
	}


	for _, image := range mdlibImages {
		hashValue, imageData, err := imageToHash(image) // Implement imageToHash function
		if err != nil {
			log.Error("Error calculating hash:", err)
			return nil, err
		}
		
		var imageLink string
		if strings.HasPrefix(image, "http://") || strings.HasPrefix(image, "https://") {
			if imageBackend == "s3" {
				imageLink, err = uploadToS3(imageData, hashValue)
			} else if imageBackend == "ghost" {
				imageLink, err = uploadToGhost(token, imageData, hashValue, blogImageList)
			}
		} else {
			if imageBackend == "s3" {
				log.Debug("Uploading image to AWS S3")
				imageLink, err = uploadToS3(image, hashValue)
			} else if imageBackend == "ghost" {
				imageLink, err = uploadToGhost(token, image, hashValue, blogImageList)
				log.Debug("Uploading image to Ghost Database")
			}
		}

		if err != nil {
			return nil, err
		}

		uploadedImages[image] = imageLink

	}
	return uploadedImages, nil
}

func uploadToGhost(token, imageData, hashName string, blogImageList []string) (string, error) {
	ghostVersion := config.GetDefault("ghost-configuration.GHOST_VERSION", "").(string)
	ghostUrl := config.GetDefault("ghost-configuration.GHOST_URL", "").(string)
	var postsApiBase string
	if ghostVersion == "v5" {
		postsApiBase = ghostUrl + "/api/admin/images/upload/"
	} else {
		postsApiBase = ghostUrl + "/api/" + ghostVersion + "/admin/images/upload/"
	}

	

	for _, name := range blogImageList {
		hashValue := strings.Split(hashName, ".")[0]
		filename := filepath.Base(name)
		filenameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename))
		if hashValue == filenameWithoutExt {
			log.Debug("The image already exists and is being reused: ", name)
			return name, nil
		}
	}

	// Open the file for reading
	file, err := os.Open(imageData)
	
	if err != nil {
		return "", errors.New("Failed to open image file: " + err.Error())
	}
	defer file.Close()
	// Create a buffer to hold the multipart encoded form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	contentType := mime.TypeByExtension(filepath.Ext(imageData))
	
	if contentType == "" {
		contentType = "application/octet-stream" 
	}

	part, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": []string{`form-data; name="file"; filename="` + hashName + `"`},
		"Content-Type":        []string{contentType},
	})

	
	if err != nil {
		return "", errors.New("Failed to create form file part: " + err.Error())
	}
	_, err = io.Copy(part, file)
	
	if err != nil {
		return "", errors.New("Failed to copy image data to form part: " + err.Error())
	}

	err = writer.WriteField("ref", hashName)
	if err != nil {
		return "", errors.New("Failed to write ref field: " + err.Error())
	}

	
	err = writer.Close()

	if err != nil {
		return "", errors.New("Failed to close multipart writer: " + err.Error())
	}

	req, err := http.NewRequest("POST", postsApiBase, body)

	if err != nil {
		return "", errors.New("Failed to create HTTP request: " + err.Error())
	}
	req.Header.Set("Authorization", "Ghost "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.New("Failed to send HTTP request: " + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return "", errors.New("Failed to upload image: HTTP status " + resp.Status)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", errors.New("Failed to decode JSON response: " + err.Error())
	}

	imageLink := result["images"].([]interface{})[0].(map[string]interface{})["url"].(string)

	return imageLink, nil
}

func getImageFromPost(postJSON string) ([]string, error) {
	var post map[string]interface{}
	if err := json.Unmarshal([]byte(postJSON), &post); err != nil {
		log.Error("Error during JSON unmarshaling:", err)
	}
	htmlContent := ""
	if cards, ok := post["cards"].([]interface{}); ok {
		if len(cards) > 0 {
			if cardData, ok := cards[0].([]interface{}); ok && len(cardData) > 1 {
				if cardHtml, ok := cardData[1].(map[string]interface{}); ok {
					if htmlStr, ok := cardHtml["html"].(string); ok {
						htmlContent = htmlStr
					}
				}
			}
		}
	}

	imageList := []string{}
	z := html.NewTokenizer(strings.NewReader(htmlContent))
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "img" {
				for _, attr := range t.Attr {
					if attr.Key == "src" {
						imageList = append(imageList, attr.Val)
						break
					}
				}
			}
		}
	}

	return imageList, nil

}

func sha256Sum(filename string) (string, error) {
	h := sha256.New()

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	hashInBytes := h.Sum(nil)
	hash := hex.EncodeToString(hashInBytes)

	return hash, nil
}

func imageToHash(image string) (string, string, error) {
	var tp string // Temporary file path
	var err error
	var fileExtension string
	var hashValue string

	if strings.HasPrefix(image, "http://") || strings.HasPrefix(image, "https://") {

		iext := filepath.Ext(image)
		// fmt.Println("iext", iext)
		tempDir := getTempDir()

		tp = filepath.Join(tempDir, "img"+iext)
		resp, err := http.Get(image)
		if err != nil {
			return "", "", err
		}
		defer resp.Body.Close()

		outFile, err := os.Create(tp)
		if err != nil {
			return "", "", err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, resp.Body)
		if err != nil {
			return "", "", err
		}
		fileExtension = iext
		hashValue, err = sha256Sum(tp)
	} else {
		// fmt.Println("Downloading image:", image)
		hashValue, err = sha256Sum(image)
		fileExtension = filepath.Ext(image)
	}

	if err != nil {
		return "", "", err
	}

	imageName := hashValue + fileExtension
	return imageName, tp, nil
}
func getTempDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("TEMP")
	}
	return "/tmp/"
}

func uploadFeatureImage(meta map[string]interface{}, token string, featureImage string) {
	if featureImagePath, ok := meta["feature_image"].(string); ok {
		hashValue, err := sha256Sum(featureImagePath)
		if err != nil {
			log.Error("Error calculating SHA-256 sum:", err)
			return
		}

		fileExtension := filepath.Ext(featureImagePath)
		imageName := hashValue + fileExtension
		imageBackend := config.GetDefault("image-configuration.IMAGE_BACKEND", "").(string)

		if imageBackend == "s3" {
			log.Debug("Uploading feature image to AWS S3")
			meta["feature_image"], err = uploadToS3(featureImagePath, imageName)
			if err != nil {
				log.Error("Error uploading file to S3:", err)
			} else {
				log.Info("File uploaded to S3 successfully.")
			}
		} else {
			log.Debug("Uploading feature image to Ghost Database")
			var featureImgList []string
			if featureImage != "" {
				featureImgList = []string{featureImage}
			} else {
				featureImgList = []string{}
			}
			meta["feature_image"], err = uploadToGhost(token, featureImagePath, imageName, featureImgList)
			if err != nil {
				log.Error("Error uploading feature image to Ghost:", err)
			}

		}

		log.Info("Uploaded feature image")
	} else {
		log.Info("Feature image not provided")
	}
}

func getMimeType(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	return mime.TypeByExtension(ext), nil
}

func uploadToS3(localFilePath, s3FilePath string) (string, error) {
	accessKeyID := config.GetDefault("aws-s3-configuration.ACCESS_KEY_ID", "").(string)
	secretAccessKey := config.GetDefault("aws-s3-configuration.SECRET_ACCESS_KEY", "").(string)
	bucketName := config.GetDefault("aws-s3-configuration.BUCKET_NAME", "").(string)
	s3BaseUrl := config.GetDefault("aws-s3-configuration.S3_BASE_URL", "").(string)
	region := config.GetDefault("aws-s3-configuration.REGION", "").(string)

	creds := credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region), Credentials: creds},
	)
	if err != nil {
		log.Error("Error creating session:", err)
		return "", err
	}

	
	svc := s3.New(sess)

	
	mimeType, err := getMimeType(localFilePath)
	if err != nil {
		log.Error("Error determining MIME type:", err)
		return "", err
	}

	
	file, err := os.Open(localFilePath)
	if err != nil {
		log.Error("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	
	stat, err := file.Stat()
	if err != nil {
		log.Error("Error getting file stats:", err)
		return "", err
	}
	size := stat.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(s3FilePath),
		Body:        bytes.NewReader(buffer),
		ContentType: aws.String(mimeType),
	}

	_, err = svc.PutObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fallthrough
			default:
				log.Error(aerr.Error())
			}
		} else {
			log.Error(err.Error())
		}
		return "", err
	}

	s3URL := s3BaseUrl + s3FilePath
	log.Debug("Image s3 url", s3URL)
	return s3URL, nil
}

func getJWToken() (string, error) {

	ghostAdminKey := config.GetDefault("ghost-configuration.ADMIN_API_KEY", "").(string)
	ghostVersion := config.GetDefault("ghost-configuration.GHOST_VERSION", "").(string)

	parts := strings.Split(ghostAdminKey, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid admin key format")
	}
	id, secret := parts[0], parts[1]

	var audValue string
	if ghostVersion == "v5" {
		audValue = "/admin/"
	} else {
		audValue = "/" + ghostVersion + "/admin/"
	}

	iat := time.Now().Unix()
	exp := iat + 5*60 // expires in  5 minutes

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": iat,
		"exp": exp,
		"aud": audValue,
	})

	token.Header["kid"] = id

	secretBytes, err := hex.DecodeString(secret)
	if err != nil {
		return "", fmt.Errorf("failed to decode secret: %v", err)
	}

	signedToken, err := token.SignedString(secretBytes)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return signedToken, nil
}
var version string

func main() {
	loadGlobalConfig()
	if len(version) == 0 {
		version = "vUnset"
	}
	if opts.Version{
		fmt.Println(version)
	}

	parser := flags.NewParser(&opts, flags.Default)
	parser.Usage = "Usage: glee <markdown_file_path>"
	args, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
	if opts.ShowConfig {
			printConfiguration()
			return
		}

	if opts.Help {
		parser.WriteHelp(os.Stdout)
		return
	
	}

	if opts.Debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
		
	

	if len(args) == 1 {
		filePath := args[0]
		content, err := ioutil.ReadFile(filePath)

		if err != nil {
			log.Fatalf("Failed to read file %s: %v", filePath, err)
		}

		contentStr := string(content)

	index := strings.Index(contentStr, "\n---\n")
	if index == -1 {
		log.Fatal("Could not find YAML delimiter")
	}

	yamlFrontMatter := contentStr[:index]
	markdownContent := contentStr[index+len("\n---\n"):]
	
	var metadata map[string]interface{}
	err = yaml.Unmarshal([]byte(yamlFrontMatter), &metadata)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	postToGhost(metadata, markdownContent)
	} else {
		log.Fatal("Please provide a markdown file.")
	}
}

func printConfiguration() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}

	configPath := filepath.Join(usr.HomeDir, ".glee.toml")
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	fmt.Printf("Configuration path: %s\n", configPath)
	fmt.Println("-------------------")
	fmt.Println(string(content))
}
