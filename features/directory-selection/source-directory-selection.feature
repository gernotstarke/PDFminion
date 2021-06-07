@user @directory
Feature: Source directory - required properties
  User can select a source directory with a file/directory chooser.

  Rules:
    - If user selects a directory without PDF files, an appropriate warning is given
    - Even directories without PDF files are valid

Scenario: User selects empty directory
  When user selects "EmptyDir" as source
  Then validity of source dir is true


Scenario: User selects dir with one PDF
  When user selects "OnePDF" as source
  Then validity of source dir is true