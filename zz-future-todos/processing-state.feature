Feature: Internal processing state is accessible
  As a developer I want the internal processing state to be accessible
  when the application runs:

  Rules:
  * The list of source PDFs shall be available
  * The list of (already processed) target PDFs shall be available
  * The current (active) page number shall be available

Scenario: Processing state at beginning of work
  Given "OnePDF" is a source directory
  And it contains "sample-A4-portrait-1pg.pdf"
  When processing is started
  Then processing.sourceDir is "OnePDF"
  And processing.sourcePDFs is make "sample-A4-portrait.1pg.pdf"