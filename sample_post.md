---
title:  'testing sample markdown file'
authors:
- linz07m@gmail.com
# tags: []
# featured: false
# status: draft
# excerpt: null,
feature_image: ./test_images/Animhorse.gif
# sidebar_toc: false
code_hilite_theme: vim
slug: testing-glee
---

<!-- [TOC] -->

```python
import shutil
import ansible.constants as C
from ansible.executor.task_queue_manager import TaskQueueManager
from ansible.module_utils.common.collections import ImmutableDict
from ansible.inventory.manager import InventoryManager
from ansible.parsing.dataloader import DataLoader
from ansible.vars.manager import VariableManager
from ansible import context
from ansible.executor.playbook_executor import PlaybookExecutor
import os
from ansible.utils.display import Display
import getpass

loader = DataLoader()
Configure the CLI arguments
context.CLIARGS = ImmutableDict(
    listtasks=False,
    listhosts=False,
    syntax=False,
    connection="smart",
    forks=10,
    become=False,
    verbosity=4,
    check=False,
    start_at_task=None,
    become_method= 'sudo',
      become_user= None, 
      become_ask_pass= True,
)
```

```bash
python3 -m nuitka --onefile   \
--include-package-data=ansible:'*.py' \
--include-package-data=ansible:'*.yml' \
--include-data-files=one_installer.yml=one_installer.yml \
 executor.py
```
# My Simple Markdown File

This is  basic Markdown file with some common formatting elements.

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
![heic](./test_images/apple.heic)
## Code

You can include inline code using backticks (`) like `code`.

For code blocks, use triple backticks (```):

```python
def hello_world():
    print("Hello, world!")
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


