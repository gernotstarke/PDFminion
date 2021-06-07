// Package sample_pdfs contains constant declarations for several sample PDF files,
// plus their (relative) path information, so that other test cases can be DRY.
package sample_pdfs

// private constant

// public constants

// DirPrefix is the relative path to the directory containing samples
const DirPrefix = "../sample-files-for-testing/"

// OnePageFile has exactly one page
const OnePageFile = DirPrefix + "OnePDF/sample-A4-portrait-1pg.pdf"

// ThreePageFile has exactly three pages
const ThreePageFile = DirPrefix + "TwoPDFs/make"

// DisguisedMarkdownFile is a file with PDF extension, but markdown content
const DisguisedMarkdownFile = DirPrefix + "md-disguised-as-pdf.pdf"
