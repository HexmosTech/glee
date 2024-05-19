---
title:
   default: new default title
   hn: title from glee to HN
   reddit: title from glee to reddit

# title: hello
authors:
- sample@gmail.com
tags: ["draft"]
featured: true
status: draft

excerpt: This is a  excerpt
feature_image: ./test_images/smiley.png
sidebar_toc: true
# code_hilite_theme: solarized-dark
slug: testing-glee-1
---


# My Simple Markdown File

 

This is basic Markdown file with some common formatting elements.

## Headers

You can create headers using the `#` symbol. There are six levels of headers:

# Header 1
## Header 2
### Header 3
#### Header 4
##### Header 5
###### Header 6

## Text Formatting

You can make text **bold** using double asterisks or double underscores, and you can make it *italic* using single asterisks or single underscores.

## Lists

### Ordered List

1. Item 1
2. Item 2
3. Item 3

### Unordered List

- Item A
- Item B
- Item C

## Links

You can attach links like this: [glee](https://github.com/HexmosTech/glee).

## Images

You can embed images like this:
### jpg

![sticky](./test_images/sticky.jpg)

### png

![Markdown Logo](https://markdown-here.com/img/icon256.png)
![smiley](./test_images/smiley.png)

### jpeg
![jpeg](./test_images/img.jpeg)


### gif
![gif](./test_images/Animhorse.gif)


### svg

![svg](./test_images/glee_banner.svg)

### ico

![ico](./test_images/icon.ico)

### heic
<!-- ![heic](./test_images/apple.heic) -->
## Code

You can include inline code using backticks (`) like `code`.

For code blocks, use triple backticks (```):

```python
def hello_world():
    print("Hello, world!")
```

```golang
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
```

## Table 

| Name      | Age | Occupation |
| --------- | --- | ---------- |
| John      | 30  | Engineer   |
| Mary      | 25  | Designer   |
| Richard   | 35  | Teacher    |


## support html

<ul>
  <li>Coffee</li>
  <li>Tea
    <ul>
      <li>Black</li>
      <li>Green tea</li>
    </ul>
  </li>
  <li>Milk</li>
</ul>
