@developer
Feature: Building PDFminion with make and Makefile
  This is to demonstrate if or how purely technical scenarios might look like.


  Scenario: We have a Makefile
    When The repository is checked out
    Then "Makefile" is present
    And  "README.adoc" is present

  @wip
  Scenario: We have a License
    When The repository is checked out
    Then we have a "License" file

  @wip
  Scenario Outline: Working with scenario outlines
    Given The repository is checked out
    Then <file> is present

    Examples:
      | file          |
      | "Makefile"    |
      | ".gitignore"  |
      | "LICENSE"     |
      | "assets/simple-cucumber-report.js" |