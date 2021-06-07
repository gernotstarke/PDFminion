@user @directory
Feature: Target directory - required properties
  User can select a target directory with a file/directory chooser.

  Rules:
    - Target and source directory must be different
    - If user selects a target directory containing PDF files, an appropriate warning is given
    -

Scenario: User selects empty directory
  When user selects "EmptyDir" as source
  Then validity of source dir is true


Scenario: User selects dir with one PDF
  When user selects "OnePDF" as source
  Then validity of source dir is true