@developer
Feature: Convenience functions for PDF files
  In order to handle PDF files convenience functions for various file-related operations are needed.

  @wip
  Scenario Outline: Count number of PDF files in directory

    Given Sample PDF files under "sample-files-for-testing"
    When Number of PDF files in <directory> is counted
    Then the number of PDF files should be <count>

    Examples:
      | directory    | count |
      | "OnePDF"     | 1     |
      | "TwoPDFs"    | 1     |
      | "TwelvePDFs" | 12    |
      | "NoPDF"      | 0     |
      | "EmptyDir"   | 0     |

  @wip
  Scenario: Can create temporary directory for writing
    When a temporary directory is created
    Then a file can be created there