@developer
Feature: Building PDFminion with make and Makefile
  This is to demonstrate if or how purely technical scenarios might look like.

  @skip
  Scenario: We have a Makefile
    When The repository is checked out
    Then "Makefile" is present
    And  "README.md" is present



  Scenario Outline: Working with scenario outlines
    Given The repository is checked out
    Then <file> is present

    Examples:
      | file           |
      | "Makefile"     |
      | ".gitignore"   |
      | "LICENSE"      |
      | "package.json" |
      | "assets/simple-cucumber-report.js" |