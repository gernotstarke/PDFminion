Feature: Concatenate several existing PDFs into one

  Scenario: A user wants to concatenate two PDFs into one
  Given A PDF file with 2 pages
    And a second PDF file with 2 pages
  When PDF files are concatenated
  Then the resulting file has four pages