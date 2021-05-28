# mdim - Markdown Images Maintainer

[![Go](https://github.com/bunnier/mdim/actions/workflows/go.yml/badge.svg)](https://github.com/bunnier/mdim/actions/workflows/go.yml)
[![CodeQL](https://github.com/bunnier/mdim/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/bunnier/mdim/actions/workflows/codeql-analysis.yml)

## Function

Now:

- Fixed wrong image relative path after you move document.
- Download the web images in your docs to local folder.
- Clean up no reference image.

Next:

- Read the paths from enviroment variable.

## Usage

Install:

```bash
cd ./mdim
go install
```

Usage:

```bash
mdim [-h] [-d] [-f] [-i imageFolder] [-m markdownFolder] 
```

Options:

- `-s` Set the option to save markdown document changes, otherwise print scan result only.
- `-w` Set the option to download web images to imageFolder. This option often be set with the `-s` option, otherwise although images have been download to imageFolder, the path in document still be web paths.
- `-d` Set the option to delete no reference images, otherwise print the paths only.
- `-i imageFolder` Must be not empty. The folder images save in.
- `-m markdownFolder` Must be not empty. The folder markdown documents save in.
- `-h` Show this help.
