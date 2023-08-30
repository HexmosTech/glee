# glee - A CLI for Ghost blog

[[_TOC_]]

`glee` can publish markdown posts into the Hexmos Journal. The
command will read **metadata** from the post's YAML preface, 
convert the post content into HTML and finally push to journal.

## Install

### Experimental Option 1: Build into a binary

```
./installbin.sh
```

After it's done, you can simply do:

```
glee your-post.md
```

### Option 2: Poetry standard method

```
poetry shell
python glee.py your-post.md
```



Get repo and then run:

```sh
poetry install
```

## Usage

```py
poetry shell
glee.py sample_post.md
```

### Example markdown file

See `sample_post.md` for learning how to structure an example post. 
Find additional field reference in [official docs](https://ghost.org/docs/admin-api/#posts).

### Specifying author

The `authors` field in the markdown frontmatter can specify multiple
**staff emails**. Note that this is different from *member emails*.

### Draft vs Publishing

The YAML field `status` determines status of the post. Pick `status: draft` or `status: published`
as required.

### Slug field to support updating

If your post doesn't contain a `slug` field (post name in the URL), then `glee` will not publish.
This is to help with future updates/edits from the markdown file. If you see this error, give a url
friendly fragment as `slug` in your markdown:

> ERROR: Include a URL friendly slug field in your markdown file and retry! This is required to support updates

### Automatically generate Table of Contents (TOC)

Include the string `[TOC]` in your content area:

```
---
yaml
---
[TOC]
your content
```

### Image handling

Presently all images in the input post are uploaded to an s3 bucket.
We calculate the hash for each image, and use that as the filename in s3.
This ensures that each unique image is stored only once in the server and
that there are no naming conflicts.

TODO: Check whether the hash exists in s3 before we try to upload.

### Syntax Highlighting

Languages supported: https://pygments.org/languages/

Fenced code blocks docs: https://python-markdown.github.io/extensions/fenced_code_blocks/

## Google doc to markdown conversion

The following tool sort of works:

1. Source: https://github.com/Mr0grog/google-docs-to-markdown
2. Usable demo: https://mr0grog.github.io/google-docs-to-markdown/

## TODO

1. Report link on the terminal if published
1. Test image inclusion, say through base64
1. List options for staff emails 
1. Better argument handling for the CLI
1. More advanced/powerful md library