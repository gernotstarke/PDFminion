Feature: Evenify - pad PDF to even page count
  The user wants every section of a printed handout to start on a right (uneven) page

  # "evenify" means adding a single (usually blank) page so the resulting file has an even number of pages

  Scenario: A PDF files with an uneven number of pages shall be evenified
  Given a PDF file with one page
  When the user has turned evenify to on
  And processing is started
  Then the resulting output file has two pages


  Scenario: A PDF files with an even number of pages shall be evenified
    Given a PDF file with two pages
    When the user has turned evenify to on
    And processing is started
    Then the resulting output file has two pages

