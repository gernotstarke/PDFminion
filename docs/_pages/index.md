---
title: "Welcome to PDFminion"

layout: page

permalink: /

header:
    overlay_image: /assets/images/header.webp
    overlay_filter: rgba(0, 0, 0, 0.6)

    actions:
        - label: "GitHub Repo <i class='fab fa-github'></i>"
          url: https://github.com/arc42/pdfminion
          blank: true

excerpt: "**For all those who like handouts with page numbers and running headers!**"
---



<section>

     <p>
        PDFminion adds page numbers and running-headers on pdf documents, helping to produce useful handouts.

        It's open-source, runs on all major platforms and is free to use.
        <br>    
        <div class="logo">
            <img src="assets/images/arc42-logo.png" alt="arc42 logo">
        </div>

    </p>
 <p>
   Brought to you by <strong>arc42</strong>
</p>

</section>




<br>
<hr class="section-sep">
<br>



<section id="features">

    <h1>Features</h1>

        <div class="box-container">

        <div class="box box--primary box-third">
            <img src="assets/images/functions/page-number.png" alt="page-numbering" class="img-half">
            <h5>Page Numbers</h5>
            <p>Adds consecutive page numbers to all documents.</p>
        </div>

        <div class="box box--primary box-third">
            <img src="assets/images/functions/running-header.png" alt="running-head" class="img-half">
            <h5>Running Header</h5>
            <p>Adds a header to the top of all pages in all documents. </p>
        </div>

        <div class="box box--primary box-third">
            <img src="assets/images/functions/mascot.png" alt="we have a sweet mascot" class="img-half">
            <h5>Add Personal Touch</h5>
            <p>Who said PDF documents couldn't have a personal touch? Thanx to Lisa, @teapot418.</p>
        </div>
    
        <div class="box box--primary box-third">
            <img src="assets/images/functions/chapter-number.png" alt="chapter numbers" class="img-half">
            <h5>Chapter Number</h5>
            <p>Every document makes a chapter within the output. This feature allows to add the chapter number to all pages.</p>
        </div>
        <div class="box box--primary box-third">
            <img src="assets/images/functions/toc.png" alt="table of contents" class="img-half">
            <h5>Table of Contents</h5>
            <p>Create table-of-contents with document names, their chapter number and starting page number.</p>
        </div>
        <div class="box box--primary box-third">
            <img src="assets/images/functions/merge.png" alt="merge documents" class="img-half">
            <h5>Merge</h5>
            <p>Create a single output document from all inputs. The output file name is configurable.</p>
        </div>
    </div>

</section>



<br>
<hr class="section-sep">
<br>


<section id="installation" markdown="1">

<h1>Installation</h1>
    
PDFminion runs on all major (desktop) operating systems. 
You can download an appropriate version and install it yourself, or you can use one of our installation options.

If you're nerdy enough, you can [fork the repo](https://github.com/arc42/pdfminion), and build your own executable version, either with `make` or `go build`.

<div class="box-container">

<div class="box box--primary box-third" markdown="1">
<h4><i class="fab fa-apple"></i> MacOS</h4>
<p>How to install on MacOS</p>
PDFminion will be installable with [Homebrew](https://brew.sh), the package manager for MacOS.
But currently, this installation method is not available...

</div>

<div class="box box--primary box-third" markdown="1">
<h4><i class="fab fa-windows"></i> Windows</h4>
<p>How to install on Windows</p>
PDFminion will be installable with [Chocolatey](https://chocolatey.org/), the package manager for Windows.
But currently, this installation method is not available...
</div>

<div class="box box--primary box-third" markdown="1">
<h4><i class="fab fa-linux"></i> Linux</h4>
<p>How to install on Linux</p>
PDFminion will be installable with [Snapcraft](https://snapcraft.io/snapcraft), the package manager for Linux, that works for most distributions.
But currently, this installation method is not available...
</div>

</div>



</section>


<br>
<hr class="section-sep">
<br>


<section id="terminology"></section>


<h1>Terminology</h1>

<img src="assets/images/page-terminology.png" alt="page terminology"/>


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
| **List Languages**       | `--list-languages` | `-ll`                   | Lists all available languages for the `--language` option. Example: `pdfminion --list-languages` |
| **Running Head**         | `--running-head <text>`    | | Sets text for the running head at the top of each page. Example: `pdfminion --running-head "Document Title"`|
| **Chapter Prefix**       | `--chapter-prefix <text>`  | | Specifies prefix for chapter numbers. Default: "Chapter". Example: `pdfminion --chapter-prefix "Ch."`|
| **Page Prefix**          | `--page-prefix <text>`     | | Sets prefix for page numbers. Default: "Page". Example: `pdfminion --page-prefix "Page"` |
| **Separator**            | `--separator <symbol>`     |  | Defines the separator between chapter, page number, and total count. Default: `-`. Example: `pdfminion --separator " | "`        |
| **Page Count Prefix**    | `--page-count-prefix <text>`| | Sets prefix for total page count. Default: "of". Example: `pdfminion --page-count-prefix "out of"` |



<br>
<hr class="section-sep">
<br>


<section id="examples">

    <h1>Examples</h1>

<div markdown="1" >

</div>

</section>

