Feature:  Visual support

  In order to avoid making mistakes when using the application
  As an application user
  I want to receive visual support


  Scenario: Tooltips on UI areas where choices are required
    Given The users moves the mouse cursor over the apolication window
    When the cursor hovers over one of the following fields

    | field |
    | source_dir   |
    | target_dir   |
    | numberize    |
    | page prefix  |
    | chapter prefix |
    | evenify      |
    | evenify_text |
    | add_header   |
    | header       |
    | concatenate  |

    Then an appropriate tooltip becomes visible
