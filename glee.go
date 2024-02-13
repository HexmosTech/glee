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
	"net/http"
	"os"
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

var postsApiBase string

var buf bytes.Buffer
var md = goldmark.New(

		goldmark.WithExtensions(
			extension.GFM, // Includes table, fenced code, and code highlight extensions
			extension.Table,
			img.NewImg("image", nil),
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
				
			),
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
		return nil, fmt.Errorf("Unable to communicate with the Ghost Admin API: %s", body)
	}

	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if len(data.Posts) > 0 {
		return &data.Posts[0], nil
	}

	return nil, errors.New("No posts found for the given slug")
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
	// Check if config is nil
	if config == nil {
		log.Fatal("Configuration is not initialized. Call viewTOMLFile to initialize it.")
	}

	globalSidebarTOC := config.GetDefault("blog-configuration.SIDEBAR_TOC", "").(bool)
	globalFeatured := config.GetDefault("blog-configuration.FEATURED", "").(bool)
	globalStatus := config.GetDefault("blog-configuration.STATUS", "").(string)

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

	theme := ` .codehilite .hll { background-color: #49483e }
   .codehilite  { background: #303030; color: #f8f8f2 }
   .codehilite .c { color: #75715e } /* Comment */
   .codehilite .err { color: #960050; background-color: #1e0010 } /* Error */
   .codehilite .k { color: #66d9ef } /* Keyword */
   .codehilite .l { color: #ae81ff } /* Literal */
   .codehilite .n { color: #f8f8f2 } /* Name */
   .codehilite .o { color: #f92672 } /* Operator */
   .codehilite .p { color: #f8f8f2 } /* Punctuation */
   .codehilite .ch { color: #75715e } /* Comment.Hashbang */
   .codehilite .cm { color: #75715e } /* Comment.Multiline */
   .codehilite .cp { color: #75715e } /* Comment.Preproc */
   .codehilite .cpf { color: #75715e } /* Comment.PreprocFile */
   .codehilite .c1 { color: #75715e } /* Comment.Single */
   .codehilite .cs { color: #75715e } /* Comment.Special */
   .codehilite .gd { color: #f92672 } /* Generic.Deleted */
   .codehilite .ge { font-style: italic } /* Generic.Emph */
   .codehilite .gi { color: #a6e22e } /* Generic.Inserted */
   .codehilite .gs { font-weight: bold } /* Generic.Strong */
   .codehilite .gu { color: #75715e } /* Generic.Subheading */
   .codehilite .kc { color: #66d9ef } /* Keyword.Constant */
   .codehilite .kd { color: #66d9ef } /* Keyword.Declaration */
   .codehilite .kn { color: #f92672 } /* Keyword.Namespace */
   .codehilite .kp { color: #66d9ef } /* Keyword.Pseudo */
   .codehilite .kr { color: #66d9ef } /* Keyword.Reserved */
   .codehilite .kt { color: #66d9ef } /* Keyword.Type */
   .codehilite .ld { color: #e6db74 } /* Literal.Date */
   .codehilite .m { color: #ae81ff } /* Literal.Number */
   .codehilite .s { color: #e6db74 } /* Literal.String */
   .codehilite .na { color: #a6e22e } /* Name.Attribute */
   .codehilite .nb { color: #f8f8f2 } /* Name.Builtin */
   .codehilite .nc { color: #a6e22e } /* Name.Class */
   .codehilite .no { color: #66d9ef } /* Name.Constant */
   .codehilite .nd { color: #a6e22e } /* Name.Decorator */
   .codehilite .ni { color: #f8f8f2 } /* Name.Entity */
   .codehilite .ne { color: #a6e22e } /* Name.Exception */
   .codehilite .nf { color: #a6e22e } /* Name.Function */
   .codehilite .nl { color: #f8f8f2 } /* Name.Label */
   .codehilite .nn { color: #f8f8f2 } /* Name.Namespace */
   .codehilite .nx { color: #a6e22e } /* Name.Other */
   .codehilite .py { color: #f8f8f2 } /* Name.Property */
   .codehilite .nt { color: #f92672 } /* Name.Tag */
   .codehilite .nv { color: #f8f8f2 } /* Name.Variable */
   .codehilite .ow { color: #f92672 } /* Operator.Word */
   .codehilite .w { color: #f8f8f2 } /* Text.Whitespace */
   .codehilite .mb { color: #ae81ff } /* Literal.Number.Bin */
   .codehilite .mf { color: #ae81ff } /* Literal.Number.Float */
   .codehilite .mh { color: #ae81ff } /* Literal.Number.Hex */
   .codehilite .mi { color: #ae81ff } /* Literal.Number.Integer */
   .codehilite .mo { color: #ae81ff } /* Literal.Number.Oct */
   .codehilite .sa { color: #e6db74 } /* Literal.String.Affix */
   .codehilite .sb { color: #e6db74 } /* Literal.String.Backtick */
   .codehilite .sc { color: #e6db74 } /* Literal.String.Char */
   .codehilite .dl { color: #e6db74 } /* Literal.String.Delimiter */
   .codehilite .sd { color: #e6db74 } /* Literal.String.Doc */
   .codehilite .s2 { color: #e6db74 } /* Literal.String.Double */
   .codehilite .se { color: #ae81ff } /* Literal.String.Escape */
   .codehilite .sh { color: #e6db74 } /* Literal.String.Heredoc */
   .codehilite .si { color: #e6db74 } /* Literal.String.Interpol */
   .codehilite .sx { color: #e6db74 } /* Literal.String.Other */
   .codehilite .sr { color: #e6db74 } /* Literal.String.Regex */
   .codehilite .s1 { color: #e6db74 } /* Literal.String.Single */
   .codehilite .ss { color: #e6db74 } /* Literal.String.Symbol */
   .codehilite .bp { color: #f8f8f2 } /* Name.Builtin.Pseudo */
   .codehilite .fm { color: #a6e22e } /* Name.Function.Magic */
   .codehilite .vc { color: #f8f8f2 } /* Name.Variable.Class */
   .codehilite .vg { color: #f8f8f2 } /* Name.Variable.Global */
   .codehilite .vi { color: #f8f8f2 } /* Name.Variable.Instance */
   .codehilite .vm { color: #f8f8f2 } /* Name.Variable.Magic */
   .codehilite .il { color: #ae81ff } /* Literal.Number.Integer.Long */`

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

	// Append the default styles to the head
		// Ensure meta["codeinjection_head"] is a string before appending
	if existingHead, ok := meta["codeinjection_head"].(string); ok {
		meta["codeinjection_head"] = existingHead +"</style>"+ defaultStyle + theme+"</style>"
	} else {
		meta["codeinjection_head"] ="<style> "+defaultStyle + theme+"</style>"
	}

	// Conditionally append sidebar TOC head and footer
	if globalSidebarTOC {
		if existingHead, ok := meta["codeinjection_head"].(string); ok {
			meta["codeinjection_head"] = existingHead + sidebarTocHead
		} else {
			meta["codeinjection_head"] = sidebarTocHead
		}
		meta["codeinjection_foot"] = sidebarTocFooter
	}

	// Assuming meta is a map[string]interface{}, add or update relevant fields
	fmt.Printf("SIDEBAR_TOC: %v\n", globalSidebarTOC)
	fmt.Printf("FEATURED: %v\n", globalFeatured)
	fmt.Printf("STATUS: %v\n", globalStatus)

	return meta
}

func viewTOMLFile() {
	configPath := filepath.Join(os.Getenv("HOME"), ".glee.toml")
	var err error
	config, err = toml.LoadFile(configPath)
	if err != nil {
		getTOMLFile(configPath)
		os.Exit(0)
	}
}

func getTOMLFile(configPath string) {
	fmt.Printf("The configuration file at %s was not found.\n", configPath)

	var configResponse string
	fmt.Print("Would you like me to create the configuration file? (yes/no): ")
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
		// fmt.Printf("%#v\n", body)
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
		fmt.Println("Error marshaling body to JSON:", err)
		return
	}

	req, err := http.NewRequest(method, apiEndpoint, bytes.NewBuffer(bodyJson))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header = headers
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		var errorMessage string
		errorBytes, _ := ioutil.ReadAll(resp.Body) // Read the response body
		if err := json.Unmarshal(errorBytes, &errorMessage); err != nil {
			// If the error message is not JSON, just use the raw response body
			errorMessage = string(errorBytes)
		}
		fmt.Printf("Request failed with status: %s. Error message: %s\n", resp.Status, errorMessage)
		return
	}

	// Unmarshal the response
	var responseData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return
	}

	// Log the result
	if pid == "" {
		fmt.Println("Created new post")
	} else {
		fmt.Println("Updated existing post based on slug")
	}
	fmt.Printf("Blog preview link: %s\n", responseData["posts"].([]interface{})[0].(map[string]interface{})["url"])
}

func postToGhost(metadata map[string]interface{}, content string) {
	// Add configurations to metadata
	metadata=addBlogConfigurations(metadata)
	metadata["html"] = toHTML(content)
	token, err := getJWToken()
	if err != nil {
		log.Fatalf("Failed generate jwt token: %v", err)
	}
	headers := http.Header{}
	headers.Set("Authorization", "Ghost "+token)
	post, err := getPostId(metadata["slug"].(string), headers)
	if err != nil {
		fmt.Printf("Error: Unable to communicate with the Ghost Admin API. Please verify your Ghost configurations: %v\n", err)
		return
	}

	// fmt.Printf("Post ID: %s\nUpdated At: %s\nHTML Data: %s\nFeature Image: %s\n",pid, updated_at, html_data, feature_image)
	var pid, updated_at, htmlData, featureImage string
	if post == nil {
		fmt.Println("No post found for the given slug.")
	} else {
		pid, updated_at, htmlData, featureImage = post.ID, post.UpdatedAt, post.Mobiledoc, post.FeatureImage
	}
	if _, ok := metadata["feature_image"]; ok {
		uploadFeatureImage(metadata, token, featureImage)
	} else {
		metadata["feature_image"] = ""
	}
	

	uploadedImages, err := uploadImages(token, htmlData)
	result, err := replaceImageLinks(metadata, uploadedImages)
	metadata["html"] = result
	
	
	
	// fmt.Println("Uploaded images:", uploadedImages)
	postObj := metadata
	body := map[string]interface{}{
		"posts": []map[string]interface{}{postObj},
	}
	makeRequest(headers, body, pid, updated_at)

}

func replaceImageLinks(metadata map[string]interface{}, imgMap map[string]string) (string, error) {
	// Extract the HTML string from the metadata map
	htmlStr, ok := metadata["html"].(string)
	if !ok {
		return "", errors.New("metadata does not contain a string under the 'html' key")
	}

	// Parse the HTML string into a DOM tree
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return "", err
	}

	// Traverse the DOM tree to replace img src attributes and a href attributes
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for i, attr := range n.Attr {
				switch {
				case n.Data == "img" && attr.Key == "src":
					fallthrough
				case n.Data == "a" && attr.Key == "href":
					if newSrc, ok := imgMap[attr.Val]; ok {
						// Replace the old src or href with the new one
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

	// Render the modified DOM tree back into an HTML string
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

	// Regular expression to find img tags and extract src attributes
	pattern := `<img[^>]+src="([^"]+)"[^>]*>`
	// htmlData = `<html><body><img src="http://example.com/image1.jpg"></body></html>`
	re := regexp.MustCompile(pattern)

	// Use the regex to find all matches in the htmlData
	matches := re.FindAllStringSubmatch(buf.String(), -1)
	// fmt.Println("Matches:", matches)

	// Extract the src attribute from each match
	var mdlibImages []string
	for _, m := range matches {
		if len(m) >   1 {
			mdlibImages = append(mdlibImages, m[1])
		}
	}
	// fmt.Println("mdlibImages:", mdlibImages)

	for _, image := range mdlibImages {
		hashValue, imageData, err := imageToHash(image) // Implement imageToHash function
		if err != nil {
			fmt.Println("Error calculating hash:", err)
			return nil, err
		}
		// fmt.Println("image", image)
		var imageLink string
		if strings.HasPrefix(image, "http://") || strings.HasPrefix(image, "https://") {
			if imageBackend == "s3" {
				imageLink, err = uploadToS3(imageData, hashValue, log.StandardLogger())
			} else if imageBackend == "ghost" {
				imageLink, err = uploadToGhost(token, imageData, hashValue, blogImageList)
			}
		} else {
			if imageBackend == "s3" {
				imageLink, err = uploadToS3(image, hashValue, log.StandardLogger())
			} else if imageBackend == "ghost" {
				imageLink, err = uploadToGhost(token, image, hashValue, blogImageList)
			}
		}

		if err != nil {
			return nil, err
		}

		uploadedImages[image] = imageLink
	}

	return uploadedImages, nil
}


func uploadToGhost(token, imageData, hashValue string, blogImageList []string) (string, error) {
	// Implementation goes here
	return "", nil
}

func getImageFromPost(postJSON string) ([]string, error) {
	var post map[string]interface{}
	if err := json.Unmarshal([]byte(postJSON), &post); err != nil {
		return nil, err
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

	// Parse the HTML and find all image tags
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

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Copy the file's content to the hash
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	// Calculate the hash
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
		hashValue,err =sha256Sum(image)
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
			fmt.Println("Error calculating SHA-256 sum:", err)
			return
		}

		fileExtension := filepath.Ext(featureImagePath)
		imageName := hashValue + fileExtension
		imageBackend := config.GetDefault("image-configuration.IMAGE_BACKEND", "").(string)

		if imageBackend == "s3" {
			// Upload to S3
			// You would need to implement the s3Upload function
			// s3Upload(S3BucketName, featureImagePath, imageName)
			// fmt.Println("Image uploaded to S3:", imageName)

			meta["feature_image"], err = uploadToS3(featureImagePath, imageName, log.StandardLogger())
			// fmt.Println("Feature image uploaded to S3", meta["feature_image"])
			if err != nil {
				fmt.Println("Error uploading file to S3:", err)
			} else {
				fmt.Println("File uploaded to S3 successfully.")
			}
		} else {
			meta["feature_image"] = imageName
			// Upload to Ghost's admin API
			// uploadToGhost(GhostAdminAPI, token, featureImagePath, imageName)
		}

		// Update the feature_image in the metadata

		fmt.Println("Uploaded feature image")
	} else {
		fmt.Println("Feature image not provided")
	}
}

func getMimeType(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	return mime.TypeByExtension(ext), nil
}

func uploadToS3(localFilePath, s3FilePath string, logger *log.Logger) (string, error) {
	accessKeyID := config.GetDefault("aws-s3-configuration.ACCESS_KEY_ID", "").(string)
	secretAccessKey := config.GetDefault("aws-s3-configuration.SECRET_ACCESS_KEY", "").(string)
	bucketName := config.GetDefault("aws-s3-configuration.BUCKET_NAME", "").(string)
	s3BaseUrl := config.GetDefault("aws-s3-configuration.S3_BASE_URL", "").(string)
	region := config.GetDefault("aws-s3-configuration.REGION", "").(string)

	// Create a custom credentials provider
	creds := credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region), Credentials: creds},
	)
	if err != nil {
		logger.Println("Error creating session:", err)
		return "", err
	}

	// Get the S3 service client
	svc := s3.New(sess)

	// Determine the MIME type of the file
	mimeType, err := getMimeType(localFilePath)
	if err != nil {
		logger.Println("Error determining MIME type:", err)
		return "", err
	}

	// Read the local file
	file, err := os.Open(localFilePath)
	if err != nil {
		logger.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	// Get the file size and read the file into a buffer
	stat, err := file.Stat()
	if err != nil {
		logger.Println("Error getting file stats:", err)
		return "", err
	}
	size := stat.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Upload the file to S3

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
				logger.Println(aerr.Error())
			}
		} else {
			logger.Println(err.Error())
		}
		return "", err
	}

	// Construct the full URL of the uploaded file
	// fmt.Println("s3FilePath", s3FilePath)
	s3URL := s3BaseUrl + s3FilePath
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

func main() {
	// newMarkdown()
	initLogging(true)
	viewTOMLFile()
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run glee.go <markdown_file>")
		os.Exit(1)
	}

	// Read the markdown file from the command line argument
	filePath := os.Args[1]
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filePath, err)
	}

	// Convert to string
	contentStr := string(content)

	// Find the index of the second delimiter to isolate the YAML front matter
	index := strings.Index(contentStr, "\n\n---\n\n")
	if index == -1 {
		log.Fatal("Could not find YAML front matter delimiter")
	}

	// Extract the YAML front matter and the rest of the content
	yamlFrontMatter := contentStr[:index]
	markdownContent := contentStr[index+len("\n\n---\n\n"):]

	// Parse the YAML front matter into a Go struct or map
	var metadata map[string]interface{}
	err = yaml.Unmarshal([]byte(yamlFrontMatter), &metadata)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	// Call the postToGhost function with metadata and content
	postToGhost(metadata, markdownContent)
}
