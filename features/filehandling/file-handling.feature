@developer
Feature: Convenience functions for PDF files
  In order to handle PDF files convenience functions for various file-related operations are needed.

  Background:
    Given A directory named "sample-files-for-testing" containing sample PDFs

  Scenario Outline: Count number of PDF files in directory

    When PDF files in <directory> are counted
    Then the number of PDF files should be <count>

    Examples:
      | directory    | count |
      | "OnePDF"     | 1     |
      | "TwoPDFs"    | 2     |
      | "TwelvePDFs" | 12    |
      | "NoPDF"      | 0     |
      | "EmptyDir"   | 0     |

  Scenario Outline: Collect PDF files in directory

    When PDF files in <directory> are collected
    Then the list of PDF files shall be <fileList>

    Examples:
      | directory          | fileList                                                                                                                                                                                                                                                                                                                                                                               |
      | "OnePDF"           | "sample-A4-portrait-1pg.pdf"                                                                                                                                                                                                                                                                                                                                                           |
      | "TwoPDFs"          | "sample-A4-portrait-1pg.pdf sample-A4-portrait-3pgs.pdf"                                                                                                                                                                                                                                                                                                                               |
      | "TwelvePDFs"       | "01_sample-A4-portrait-1pg.pdf 02_sample-A4-portrait-3pgs.pdf 03_sample-A4-portrait-11pgs.pdf 04_sample-A4-portrait-101pgs.pdf 05_sample-color-A4-portrait-3pg.pdf 06_sample-A4-portrait-4pgs.pdf 07_sample-A4-portrait-1pg.pdf 08_sample-A4-portrait-1pg.pdf 09_sample-A4-portrait-1pg.pdf 10_sample-A4-portrait-1pg.pdf 11_sample-A4-portrait-1pg.pdf 12_sample-A4-portrait-1pg.pdf" |
      | "FourFilesTwoPdfs" | "sample-A4-portrait-1pg.pdf sample-A4-portrait-3pgs.pdf"                                                                                                                                                                                                                                                                                                                               |

