# 5. Commands and config options

Date: 2024-12-01

## Status

Accepted

## Context

We need to decide which commands and config options to support.
"General" in this sense means commands that relate to the overall processing.
Commands can determine the behavior of the application, such as the language of text stamped onto the page (e.g. English, German, Spanish etc.), merging of PDFs, debug-mode, help screen.


## Decision

>Please note: The first releases of PDFminion might support only a limited set of commands and configuration options.
We will support the following commands and options:

| **Command/Config Name**  | **Long Command**            | **Short Command**       | **Description**                                                                                                                    |
|---------------------------|-----------------------------|--------------------------|------------------------------------------------------------------------------------------------------------------------------------|
| **Help**                 | `--help`                   | `-h`, `-?`, `?`         | Displays a list of supported commands and their usage. Example: `pdfminion --help`                                                |
| **Version**              | `--version`                | `-v`                    | Displays the current version of PDFminion. Example: `pdfminion --version`                                                        |
| **Source Directory**     | `--source <directory>`     | `-s <directory>`        | Specifies the input directory for PDF files. Example: `pdfminion --source ./input`                                               |
| **Target Directory**     | `--target <directory>`     | `-t <directory>`        | Specifies the output directory for processed files. Creates the directory if it doesn’t exist. Example: `pdfminion --target ./out`|
| **Force Overwrite**      | `--force`                  | `-f`                    | Allows overwriting existing files in the target directory. Example: `pdfminion --force`                                          |
| **Language**             | `--language <code>`        | `-l <code>`             | Sets the language for stamped text. Supports `EN` (English) and `DE` (German). Default: `EN`. Example: `pdfminion --language DE` |
| **Evenify**              | `--evenify {on\|off}`       | `-e {on\|off}`           | Enables or disables adding blank pages for even page counts. Example: `pdfminion --evenify=off`                                  |
| **Blank Page Text**       | `--blankpagetext <text>`   | `-b <text>`             | Specifies text printed on blank pages added during evenification. Example: `pdfminion --blankpagetext "deliberately left blank"`              |
| **Defaults**             | `--defaults`               |                          | Prints all current default settings. Example: `pdfminion --defaults`.                                                             |
| **Debug Mode**           | `--debug`                  |                          | Enables debug mode for detailed logs. Example: `pdfminion --debug`                                                               |
| **Merge**                | `--merge <filename>`       | `-m <filename>`         | Merges input files into a single PDF. Uses default name if `<filename>` not provided. Example: `pdfminion --merge combined.pdf`   |
| **Table of Contents**    | `--toc`                    |                          | Generates a table-of-contents PDF. Supported from v1.5.0. Example: `pdfminion --toc`                                             |
| **Config File**          | `--config <filename>`      | `-c <filename>`         | Loads configuration from a file. Overrides conflicting command-line options. Example: `pdfminion --config settings.json`         |
| **List Languages**       | `--list-languages`         | `-ll`                   | Lists all available languages for the `--language` option. Example: `pdfminion --list-languages`                                 |
| **Running Head**         | `--running-head <text>`    |                          | Sets text for the running head at the top of each page. Example: `pdfminion --running-head "Document Title"`                     |
| **Chapter Prefix**       | `--chapter-prefix <text>`  |                          | Specifies prefix for chapter numbers. Default: "Chapter". Example: `pdfminion --chapter-prefix "Ch."`                           |
| **Page Prefix**          | `--page-prefix <text>`     |                          | Sets prefix for page numbers. Default: "Page". Example: `pdfminion --page-prefix "Page"`                                         |
| **Separator**            | `--separator <symbol>`     |                          | Defines the separator between chapter, page number, and total count. Default: `-`. Example: `pdfminion --separator " | "`        |
| **Page Count Prefix**    | `--page-count-prefix <text>`|                          | Sets prefix for total page count. Default: "of". Example: `pdfminion --page-count-prefix "out of"`                               |

## Consequences

