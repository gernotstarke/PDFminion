---
title: "Welcome to PDFminion"

layout: page

permalink: /

header:
    overlay_image: /assets/images/header.webp
    overlay_filter: rgba(0, 0, 0, 0.6)

    actions:
        - label: "Get Started <i class='fab fa-github'></i>"
          url: https://github.com/gernotstarke/pdfminion
          blank: true

excerpt: "**For all those who like handouts - and page numbers even more so!**"
---



<section>

    <p>
        PDFminion adds page numbers and running-headers on pdf documents, helping to produce nicely formated

        It runs on all platforms and is free to use.
    </p>

    <div class="container">
        <div class="logo">
            <img src="assets/images/arc42-logo.png" alt="arc42 logo">
        </div>
        <div class="text">
            Brought to you by <strong>arc42</strong> + <strong>Gernot Starke</strong>
        </div>
    </div>
</section>




<br>
<hr class="section-sep">
<br>



<section id="features">

    <h1>Features</h1>

        <div class="box-container">

        <div class="box box--primary box-third">
            <img src="assets/images/functions/page-number.png" alt="page-numbering">
        </div>

        <div class="box box--primary box-third">
            <img src="assets/images/functions/running-header.png" alt="running-head">
        </div>

        <div class="box box--primary box-third">
            <img src="assets/images/functions/mascot.png" alt="we have a sweet mascot">
        </div>

        <div class="box box--primary box-third">
            <img src="assets/images/functions/chapter-number.png" alt="chapter numbers">
        </div>
        <div class="box box--primary box-third">
            <img src="assets/images/functions/toc.png" alt="table of contents">
        </div>
        <div class="box box--primary box-third">
            <img src="assets/images/functions/merge.png" alt="merge documents">
        </div>
    </div>

</section>



<br>
<hr class="section-sep">
<br>


<section id="installation">

<h1>Installation</h1>
    
PDFminion will be installable with [Homebrew](https://brew.sh), the package manager for MacOS.
But currently, this simple installation method is not available...

    <div class="box-container">

        <div class="box box--primary box-third">
            <h4><i class="fab fa-apple"></i> MacOS</h4>
            <p>How to install on MacOS</p>

           </div>

        <div class="box box--primary box-third">
            <h4><i class="fab fa-windows"></i> Windows</h4>
            <p>How to install on Windows</p>
        </div>

        <div class="box box--primary box-third">
            <h4><i class="fab fa-linux"></i> Linux</h4>
            <p>How to install on Linux</p>
        </div>

    </div>



</section>


<br>
<hr class="section-sep">
<br>


<section id="usage"></section>


<h1>Usage</h1>

| **Name**  | **Long Command**  | **Short Command** | **Description**    |
|-----------|-------------------|-------------------|--------------------|
| **Help**                 | `--help`                   | `-h`, `-?`, `?`         | Displays a list of supported commands and their usage. Example: `pdfminion --help`                                                |
| **Version**              | `--version`                | `-v`                    | Displays the current version of PDFminion. Example: `pdfminion --version`                                                        |
| **Source Directory**     | `--source <directory>`     | `-s <directory>`        | Specifies the input directory for PDF files. Example: `pdfminion --source ./input`                                               |
| **Target Directory**     | `--target <directory>`     | `-t <directory>`        | Specifies the output directory for processed files. Creates the directory if it doesn’t exist. Example: `pdfminion --target ./out`|
| **Force Overwrite**      | `--force`                  | `-f`                    | Allows overwriting existing files in the target directory. Example: `pdfminion --force`                                          |
| **Language**             | `--language <code>`        | `-l <code>`             | Sets the language for stamped text. Supports `EN` (English) and `DE` (German). Default: `EN`. Example: `pdfminion --language DE` | 
| **Evenify**              | `--evenify {on\|off}`       | `-e {on\|off}`           | Enables or disables adding blank pages for even page counts. Example: `pdfminion --evenify=off`                                  |
| **Blank Page Text**      | `--blankpagetext <text>`   | `-b <text>`             | Specifies text printed on blank pages added during evenification. Example: `pdfminion --blankpagetext "deliberately left blank"`              |
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



<br>
<hr class="section-sep">
<br>


<section id="examples">

    <h1>Examples</h1>

<div markdown="1" >

</div>

</section>

