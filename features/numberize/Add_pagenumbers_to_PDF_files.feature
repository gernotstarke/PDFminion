Feature:  Add page numbers to existing PDF files

  Background:
    Given A source directory
    And a different target directory


  @numbering
  Scenario: As a user I want to add page numbers to a single PDF file

    When  the source directory contains a single PDF file
    Then sequential page numbers should have been added to this PDF file starting from 1


  @numbering
  Scenario: As a user I want to add page numbers to a number of PDF files

    When the source directory contains several PDF file
    Then sequential page numbers should have been added to the PDF files

  @numbering @sorting
  Scenario: Several files shall be numbered in alphabetical order
  When the source directory contains several PDF files
    Then numbering is done in alphabetical order of the filenames