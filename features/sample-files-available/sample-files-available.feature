Feature: Sample files are available

  Scenario: A directory containing sample files is available
    When Existence of "sample-files-for-testing/" is checked
    Then The directory containing sample files is available

