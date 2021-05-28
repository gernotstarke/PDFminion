Feature:  Visual feedback

  In order to avoid making mistakes when using the application
  As an application user
  I want to receive visual feedback on my selections

@manual
Scenario: The start button shall be enabled only if valid directories have been selected
    Given No directories have been selected
    Then stamping is not possible
    When I select a source directory
    And it contains one or more PDF files
    And I select a target directory
    And that is different from the source directory
    And the target directory does not contain any files that conflict with PDF files from the source directory
    Then stamping becomes possible

 @@manual
Scenario: A pie chart of the selected files together with their pagecount is displayed once a valid source has been chosen
   Given No source directory has been selected
   When The user selects a source directory
   And the source directory contains three PDF files with 10, 20 and 30 pages
   Then a pie chart is shown, with the names and pagecount of these files