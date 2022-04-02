@developer
Feature: Count pages in PDF files
  In order to stamp PDF files, correct page counts are needed.

  @wip
  Scenario Outline: Count number of pages in PDF files

    Given A <directory> under "sample-files-for-testing" containing <PDFfile>
    When Number of pages is counted
    Then the number of pages should be <pagecount>

    Examples:
      | directory    | PDFfile                               | pagecount |
      | "TwoPDFs"    | "sample-A4-portrait-1pg.pdf"          | 1         |
      | "TwoPDFs"    | "sample-A4-portrait-3pgs.pdf"         | 3         |
      | "TwelvePDFs" | "03_sample-A4-portrait-11pgs.pdf"     | 11        |
      | "TwelvePDFs" | "05_sample-color-A4-portrait-3pg.pdf" | 3         |
      | "TwelvePDFs" | "04_sample-A4-portrait-101pgs.pdf"    | 101       |

