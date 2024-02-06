# glee: Dev-friendly Blogging Setup

<div align="center">

[![Binary Build And Release](https://github.com/HexmosTech/glee/actions/workflows/build-and-release.yml/badge.svg)](https://github.com/HexmosTech/glee/actions/workflows/build-and-release.yml)

</div>

<img src="assets/glee-readme-banner.png" width="90%" />

<div align="center">
<a href="https://www.producthunt.com/posts/glee-2?utm_source=badge-top-post-badge&utm_medium=badge&utm_souce=badge-glee&#0045;2" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/top-post-badge.svg?post_id=418974&theme=light&period=daily" alt="glee - Dev&#0045;friendly&#0032;blogging&#0032;setup | Product Hunt" style="margin-top:20px;width: 250px; height: 54px;" width="250" height="54" /></a>
</div>

## Overview

Most technical teams struggle to put together a modern and dev-friendly blogging setup.
Our free product `glee` helps your devs compose, backup & collaborate on blog posts using markdown files
so that they can ship awesome content without frustration.

With glee, you can create and update [Ghost](https://ghost.org/) blogs. Since glee operates on plain markdown files, your blog posts are now git-friendly, allowing for easy collaboration.

`glee` command will read **metadata** from the YAML preface of your Markdown post ([sample_post.md](https://github.com/HexmosTech/glee/blob/main/sample_post.md?plain=1)), convert the post content into HTML, Store the content images either in your Ghost database or on AWS S3, and then publish them to your Ghost platform. Set up the `glee` CLI tool with a single command.

## Benefits

- Publish markdown files into Ghost blog post
- Install and configure with minimal effort
- Multiple image backends (AWS S3, ghost)
- Create and update posts with a single command
- Support for syntax highlighting and Table of Contents
- Collaborate with content writers in your team
- 100% Free and Open Source Software

## Watch Demo

[![IMAGE ALT TEXT](http://img.youtube.com/vi/nlM4b65_GSU/0.jpg)](http://www.youtube.com/watch?v=nlM4b65_GSU "glee: Dev-friendly Blogging Setup")

## Installation/Update

Run the following command to either install or update `glee`:

### For Linux/MacOS systems or Linux via [WSL](https://ubuntu.com/wsl):

```bash
wget -O - https://raw.githubusercontent.com/HexmosTech/glee/main/install.sh | bash
```

### For Windows:

Open the Command Prompt (cmd) as an administrator and execute the following command:

```cmd
powershell -Command "(New-Object Net.WebClient).DownloadFile('https://raw.githubusercontent.com/HexmosTech/glee/main/install.bat', 'install.bat'); Start-Process 'install.bat';"
```

Alternatively, you can download the [executable (exe) file](https://github.com/HexmosTech/glee/releases/latest/download/glee_windows.exe) and then move it into the `system32` folder using the command:

```cmd
Move-Item -Path "C:\Path\to\Downloads\glee_windows.exe" -Destination "C:\Windows\system32\glee.exe"
```

Note: If you encounter any security issues on Windows, turn off the [Real-time protection](https://support.microsoft.com/en-us/windows/turn-off-defender-antivirus-protection-in-windows-security-99e6004f-c54c-8509-773c-a4d776b77960) in the virus and threat protection settings.

## Configuration

After the installation, `glee` will create a configuration file ([.glee.toml](https://github.com/HexmosTech/glee/blob/main/.glee.toml)) in your home directory.

Open the configuration file `$HOME/.glee.toml` and modify the ghost, image backend and AWS S3 credential (optional).

### Ghost Configuration

#### Ghost Admin API Key

Admin API keys are used to generate short-lived single-use JSON Web Tokens (JWTs), which are then used to authenticate a request (GET,POST,PUT) using Ghost Admin API.

- Admin API keys can be obtained by creating a new Custom Integration under the Integrations screen in Ghost Admin.
 <p align="left">
  <a href="">
  <img alt="img-name" src="assets/glee-custom-integration.png" width="450"> 
    <br/>
   </a>
</p>

- Save the Custom Integration and Copy the Admin API Key to [.glee.toml](https://github.com/HexmosTech/glee/blob/main/.glee.toml) file.

 <p align="left">
  <a href="">
  <img alt="img-name" src="assets/glee-admin-api-edited.jpg" width="450"> 
    <br/>
   </a>
</p>

#### Ghost Version

Include the Ghost platform version in the TOML file.
You can find the version in the Ghost admin settings.
The version notation is as follows: 'v4' represents version 4, 'v5' represents version 5, and so forth.

 <p align="left">
  <a href="">
  <img alt="img-name" src="assets/ghost -version.png" width="450"> 
    <br/>
   </a>
</p>

#### Ghost URL

The `GHOST_URL` represents the domain where your Ghost blog is hosted.

### Blog Configuration

The `blog-configuration` section in the `.glee.toml` file serves as the global configuration for all blog posts published using `glee`. For instance, if `sidebar_toc` is set to `true` in the `blog-configuration`, then all blog posts published through glee will have the sidebar table of contents enabled. However, you have the flexibility to customize the configuration for individual blogs by utilizing the local configuration defined within the `YAML` structure of your markdown file.

### Image Storage Backend Configuration

All images in the markdown file can be stored either in your `ghost database` or an `AWS S3` bucket. We calculate the hash for each image and use that as the filename in `s3`. This ensures that each unique image is stored only once in the server and that there are no naming conflicts.

1. **Your Ghost Database (default)**

You can store the image in the same db where your content resides. To use Ghost as an image backend provide

```toml
[image-configuration]

IMAGE_BACKEND = "ghost"
```

in the [.glee.toml](https://github.com/HexmosTech/glee/blob/main/.glee.toml#L13) file.

2. **AWS S3**

Or, you can store the images in your `AWS S3` bucket as well. To use `S3` as an image backend provide

```toml
[image-configuration]

IMAGE_BACKEND = "s3"
```

in the [.glee.toml](https://github.com/HexmosTech/glee/blob/main/.glee.toml#L13) file.

Also, Configure the `S3` Credentials in the [.glee.toml](https://github.com/HexmosTech/glee/blob/main/.glee.toml) file.

Find further [information](https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html) and [tutorial](https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html) to learn more about `AWS S3`.

## Usage

After installation and configuration, you can convert Markdown file into a Ghost blog post using the following command:

```bash
glee your-post.md
```

## Markdown File Structure

The Markdown file used by `glee` consists mainly of two parts:

- A YAML Interface for metadata
- Content

```markdown
---
yaml
---

[TOC]
your content
```

### Example markdown file

See [sample_post.md](https://github.com/HexmosTech/glee/blob/main/sample_post.md?plain=1) for learning how to structure an example post.
Find additional field reference in [official docs](https://ghost.org/docs/admin-api/#posts).

## Features


### Platform-Specific Titles
With Glee, you can customize the title of your article for users coming from various platforms such as Reddit, HN, Medium, etc. The `yaml` syntax for handling multiple titles is as follows:



```yaml
title:
   default: new default title
   hn: title from glee for HN
   reddit: title from glee for Reddit
```
If you only need a single title, use the following syntax:

```yaml
title: your default title
```

Additionally, you can enhance the user experience by adding the following CSS style to create a transition effect when switching titles:

```css
<style>
    .article-title {
      opacity: 0;
       filter: blur(3px);
      animation: fadeIn 1s ease 1s forwards;
    }

    @keyframes fadeIn {
      to {
        opacity: 1;
         filter: blur(0px);
      }
    }
</style>
```

Include the above code snippet in the `Admin Dashboard -> Settings -> Code Injection -> Site Header`.

Remember to specify the src query parameter when sharing your article on platforms. For example: https://journal.hexmos.com/spam-detection-ml/?src=reddit

### Specifying author

The `authors` field in the markdown frontmatter can specify multiple
**staff emails**. Note that this is different from _member emails_.

### Draft vs Publishing

The YAML field `status` determines status of the post. Pick `status: draft` or `status: published`
as required.

### Slug field to support updating

If your post doesn't contain a `slug` field (post name in the URL), then `glee` will not publish.
This is to help with future updates/edits from the markdown file. If you see this error, give a url
friendly fragment as `slug` in your markdown:

> ERROR: Include a URL friendly slug field in your markdown file and retry! This is required to support updates

### Automatically generate Table of Contents (TOC)

`glee` support two kinds of TOC.

1. TOC in Content

For Adding TOC include the string `[TOC]` in your content area:

```markdown
---
yaml
---

[TOC]
your content
```

2. TOC as Sidebar

The YAML field `sidebar_toc` determines including sidebar table of content. Pick `sidebar_toc:true` or `sidebar_toc:false` as required.

### Syntax Highlighting

glee supports five themes: Monokai, Native, Pastie, Vim, and Fruity, for highlighting your code block. You can configure the theme globally in the glee configuration file inside `blog-post-configuration` or for a specific blog post using the `code_hilite_theme` option in the YAML structure. The default theme is Monokai.

Languages supported: https://pygments.org/languages/

Fenced code blocks docs: https://python-markdown.github.io/extensions/fenced_code_blocks/

### Collaboration

When multiple team members are working simultaneously on the same Ghost blog, they can collaborate seamlessly using any version control system. `glee` will update the blog content with each `glee` command.

## Debugging

Utilize the `--debug` option with your glee command to uncover underlying issues.

```bash
 glee sample_post.md --debug
```

## View Configuration

Utilize the `--config` option with your glee command to view the glee configurations.

```bash
glee sample_post.md --config
```

## Local Testing

Clone the repository and test the `glee` tool locally.

### Option 1: Build into a binary

Create a local standalone executable using nuitka. run the command:

```bash
./installbin.sh
```

After it's done, you can simply do:

```bash
glee your-post.md
```

### Option 2: Poetry standard method

```bash
poetry shell
python glee.py your-post.md
```

### Blog Post about glee

[glee for Ghost: Why we abandoned the Web Editor and Adopted Markdown Files](https://journal.hexmos.com/glee/)

## Acknowledgement

- The `glee` standalone single binary is created using [Nuitka](https://nuitka.net/doc/user-manual.html).
- `glee` utilizes the [Ghost Admin API](https://ghost.org/docs/admin-api/) for interaction with the Ghost blog platform.
