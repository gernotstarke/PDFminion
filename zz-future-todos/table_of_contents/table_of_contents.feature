Feature: Add table of contents
  As a user I want generate a table of contents from existing PDF files
  with only level-1 headings
  taken from the PDF metadata

  @develop @toc
  Scenario: Read title from PDF
    Given the PDF single.pdf with title "single-title"
    When the system extracts the title from PDF metadata
    Then the result is "single-title"


  @manual @toc
  Scenario: Generate ToC for single PDF
  Given a single PDF file named single.pdf
    And the title of that file is "Single"
  When the user generates the ToC
  Then the a PDF named ToC.pdf is created
    And that file contains one entry names "Single"
