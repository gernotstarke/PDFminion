Feature: (Schematic) preview of the output documents

  As a user I want to get a visual preview
  how the resulting PDFs will look like,
  when the current configuration settings are applied

 @manual
  Scenario: View effect of numbering a single PDF
  Given A single PDF document with two pages

  When I enable the configuration setting add-page-numbers
  Then A schematic preview of the output is shown
   And it contains a symbolic page number