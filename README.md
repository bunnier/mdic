# mdic - Markdown Images Cleaner

The tool will help you maintain the image relative paths of markdown files and cleanup no reference images.

Compile:

```bash
cd ./mdic
go build
```

Usage:

```bash
mdic [-dfh] [-i image folder] [-m markdown doc folder] 
```

Options:

- `-d` Set the option to delete no reference images.
- `-f` Set the option to fix image relative paths of markdown documents.
- `-h` Show this help.
- `-i string` The folder images save in.
- `-m string` The folder markdown documents save in.
