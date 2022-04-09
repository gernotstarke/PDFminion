@developer
Feature: Evenify PDF files
  New chapters or parts shall always start on the right (odd) side of a book or handout.
  Therefore the preceeding file needs to have an even number of pages.

  @wip
  Scenario Outline: We can add a blank page at the end of a PDF file.

    Given A single <PDFfile> located in <directory>
    And has <pagecount> number of pages.
    When A blank page is added to the file
    Then the resulting output file has <pagecount+1> pages.

    Examples:
      | directory    | PDFfile                               | pagecount |
      | "TwoPDFs"    | "sample-A4-portrait-1pg.pdf"          | 1         |
      | "TwoPDFs"    | "sample-A4-portrait-3pgs.pdf"         | 3         |
      | "TwelvePDFs" | "03_sample-A4-portrait-11pgs.pdf"     | 11        |
      | "TwelvePDFs" | "05_sample-color-A4-portrait-3pg.pdf" | 3         |
      | "TwelvePDFs" | "04_sample-A4-portrait-101pgs.pdf"    | 101       |
