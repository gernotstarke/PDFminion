@developer
Feature: Sample files are available

  Scenario: The directory containing sample files is available
    When Existence of "sample-files-for-testing/" is checked
    Then The directory containing sample files is available


  Scenario: A directory without PDF files exists
    Given Directory "sample-files-for-testing/EmptyFolder"
    When Number of PDF files is counted
    Then 0 is returned

  Scenario: A directory with 1 PDF file exists
    Given Directory "sample-files-for-testing/OnePDF"
    When Number of PDF files is counted
    Then 1 is returned

  Scenario: A directory with 2 PDF files exists
    Given Directory "sample-files-for-testing/TwoPDFs"
    When Number of PDF files is counted
    Then 2 is returned

