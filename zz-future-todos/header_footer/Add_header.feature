Feature: Add header
  This feature allows users to add textual headers to PDFs.

  Background:
    These are the assumptions we make for the following scenarios:

    * The horizontal size of the PDF files allows for at least 20 characters header
    * The font size for headers is fixed (11pt)

  Scenario: Add header to PDF
  Given A non-empty text for the header
  When PDF files are processed
  Then this text appears as header of all pages