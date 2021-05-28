Feature: Add footer
  This feature allows users to add chapter number or title as footer to PDFs.

  Background:
    These are the assumptions we make for the following scenarios:

    * The horizontal size of the PDF files allows for at least 20 characters footer
    * The font size for footers is fixed (11pt)

  Scenario: Add footer to PDF
  Given A non-empty text for the footer
  When PDF files are processed
  Then this text appears as footer of all pages