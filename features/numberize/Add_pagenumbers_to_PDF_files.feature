@numbering @user
Feature:  Add page numbers to existing PDF files

  Background:
    Given "sample-files-for-testing/" as source directory
    And a temporary target directory


  Scenario: Can add page number to PDF containing single page
    Given A PDF with a single page
    When  minion processing is started
    Then PDF pagecount is 1


  Scenario: As a user I want to add page numbers to a number of PDF files
    When the source directory contains several PDF files
    Then sequential page numbers should have been added to the PDF files


  Scenario: Several files shall be numbered in alphabetical order
    When the source directory contains several PDF files
    Then numbering is done in alphabetical order of the filenames


  Scenario: Number Single File
    Given I have a sourceDirectory with 1 file and 10 pages
    When I start numbering
    Then the pagecount should be 10