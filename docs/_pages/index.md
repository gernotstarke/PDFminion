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
        <br>
        It's open-source, runs on all major platforms and is free to use.
        <br>    
        <p>Brought to you by</p>
        <div class="logo">
            <img src="assets/images/arc42-logo.png" alt="arc42">
        </div>

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
            <p>Who said PDF documents couldn't have a personal touch? Thanx to <a href="https://www.sketchnotes.tech/">Lisa, @teapot418</a>.</p>
        </div>
    
        <div class="box box--primary box-third">
            <img src="assets/images/functions/chapter-number.png" alt="chapter numbers" class="img-half">
            <h5>Chapter Number</h5>
            <p>Every document makes a chapter within the output. This feature allows to add the chapter number to all pages.</p>
        </div>


       <div class="box box--primary box-third">
            <img src="assets/images/functions/evenify.png" alt="evenify" class="img-half">
             <h5>Evenify</h5>
             <p>Make all documents have an even number of pages by adding a blank page to files with odd page count.</p>
       </div>
 <div class="box box--primary box-third">
            <img src="assets/images/functions/privacy.png" alt="privacy" class="img-half">
             <h5>Privacy</h5>
             <p>We value your privacy: Your documents are yours, forever. We do neither collect data nor retain any information about your documents.</p>
       </div>

 <div class="box box--primary box-third">
            <img src="assets/images/functions/multi-language.png" alt="multi-language" class="img-half">
             <h5>Multi-Language</h5>
             <p>Handle a number of natural languages, starting with English (EN)m German (DE), French (FR) and Spanish (ES).</p>
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

<h2>Evenify</h2>

To evenify a file means adding a blank page at the end of the file if the page-count is odd (1, 3, 5 or such).
That means that the first page of every file in a group will always start on the front-page of paper,
even in case of double-sided printing.

Chapters in technical or scientific books traditionally start on odd (right-hand) pages to ensure consistency, 
readability, and prominence, aligning with classic book design practices.


<img src="assets/images/page-terminology-evenify.png" alt="evenify"/>


<br>
<hr class="section-sep">
<br>


<section id="usage"></section>


<h1>Usage</h1>


<h2>Basic Commands</h2>

| **Name**  | **Long Command**  | **Short Command** | **Description**    |
|-----------|-------------------|-------------------|--------------------|
| **Source Directory**     | `--source <directory>`     | `-s <directory>`        | Specifies the input directory for PDF files. Example: `pdfminion --source ./input`|
| **Target Directory**     | `--target <directory>`     | `-t <directory>`        | Specifies the output directory for processed files. Creates the directory if it doesn’t exist. Example: `pdfminion --target ./out`|
| **Force Overwrite**      | `--force`                  | `-f`                    | Allows overwriting existing files in the target directory. Example: `pdfminion --force` |


<h2>Processing Commands</h2>

Set the running head, the page- and chapter prefix etc.

| **Name**  | **Long Command**  | **Short Command** | **Description**    |
|-----------|-------------------|-------------------|--------------------|
| **Running Head**         | `--running-head <text>`    | | Sets text for the running head at the top of each page. Example: `pdfminion --running-head "Document Title"`|
| **Chapter Prefix**       | `--chapter-prefix <text>`  | | Specifies prefix for chapter numbers. Default: "Chapter". Example: `pdfminion --chapter-prefix "Ch."`|
| **Page Prefix**          | `--page-prefix <text>`     | | Sets prefix for page numbers. Default: "Page". Example: `pdfminion --page-prefix "Page"` |
| **Separator**            | `--separator <symbol>`     |  | Defines the separator between chapter, page number, and total count. Default: `-`. Example: `pdfminion --separator " | "`        |
| **Page Count Prefix**    | `--page-count-prefix <text>`| | Sets prefix for total page count. Default: "of". Example: `pdfminion --page-count-prefix "out of"` |
| **Evenify**  | `--evenify {on\|off}`  | `-e {on\|off}`  | Enables or disables adding blank pages for even page counts. Example: `pdfminion --evenify=off |


<h2>Information</h2>

| **Name**  | **Long Command**  | **Short Command** | **Description** |
|-----------|-------------------|-------------------|-----------------|
| **Help**                 | `--help`                   | `-h`, `-?`, `?`| Displays a list of supported commands and their usage. Example: `pdfminion --help`|
| **Version**              | `--version`                | `-v`           | Displays the current version of PDFminion. Example: `pdfminion --version` |
| **Defaults**             | `--defaults`               |                | Prints all current default settings. Example: `pdfminion --defaults`.  |
| **Debug Mode**           | `--debug`                  |                | Enables debug mode for detailed logs. Example: `pdfminion --debug`    |


<h2>Multi-Language Support</h2>

PDFMinion provides defaults for page processing for several languages.
With these commands you can change these defaults and provide your own values.


| **Name**| **Long Command**  | **Short Command** | **Description** |
|-----------|-------------------|-------------------|-----------------|
| **List Languages**| `--list-languages` | `-ll` | Lists all available languages for the `--language` option. Example: `pdfminion --list-languages` |
| **Language**      | `--language <code>`        | `-l <code>`     | Sets the language for stamped text. Supports `EN` (English) and `DE` (German). Default: `EN`. Example: `pdfminion --language DE` |
| **Blank Page Text** | `--blankpagetext <text>`   | `-b <text>`     | Specifies text printed on blank pages added during evenification. Example: `pdfminion --blankpagetext "deliberately left blank"`|




<h2>File-Related Commands</h2>
After all files have been processed, you may merge them or create a table-of-contents.

| **Name**  | **Long Command**  | **Short Command** | **Description** |
|-----------|-------------------|-------------------|-----------------|
| **Merge** | `--merge <filename>`       | `-m <filename>` | Merges input files into a single PDF. Uses default name if `<filename>` not provided. Example: `pdfminion --merge combined.pdf`   |
| **Table of Contents**  | `--toc`   |  | Generates a table-of-contents PDF. Supported from v1.5.0. Example: `pdfminion --toc`|
| **Config File**  | `--config <filename>`  | `-c <filename>` | Loads configuration from a file. Overrides conflicting command-line options. Example: `pdfminion --config settings.json`  |



<br>
<hr class="section-sep">
<br>


<section id="examples">

    <h1>Examples</h1>

<div markdown="1" >

</div>

</section>

