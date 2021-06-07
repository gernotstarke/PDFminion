@user @directory
Feature: Target directory - required properties
  User can select a target directory with a file/directory chooser.

  Rules:
    - Target and source directory must be different
    - If user selects a target directory containing PDF files, status line shall contain an error message


Scenario: Empty directory is valid target
  When user selects "EmptyDir" as target
  Then validity of target dir is "true"


Scenario: Target directory containing PDF is invalid
  When user selects "OnePDF" as target
  Then validity of target dir is "false"
  And application status contains "error: target must not contain PDFs"