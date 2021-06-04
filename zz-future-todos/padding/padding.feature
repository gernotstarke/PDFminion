Feature: Padding to a certain number of pages

  Padding (another term: filling up) a PDF up to a given number of pages.
  Often used in conjunction with concatenation.

  Scenario: Make document eight pages long
  Given A PDF with five pages
  When I choose padding to eight pages
    And start processing
  Then the resulting PDF contains eight pages
