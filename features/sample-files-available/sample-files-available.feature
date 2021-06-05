Feature: Sample files are available

  Background: The directory "./resources" containing sample files is available


  Scenario: Can detect that directory is empty
    Given an empty directory
    When Number of files is checked
    Then 0 is returned

  Scenario Outline: As developer I require several sample PDFs for testing

    When <directory> <file> is checked
    Then <file> should be <valid>

    Examples:
      | directory | file                            | valid    |
      | OnePDF  | sample-A4-portrait-1pg.pdf    | true     |
      | TwoPDFs | sample-A4-portrait-3pgs.pdf   | true     |
      | resources | md-disguised-as-pdf.pdf    | false    |