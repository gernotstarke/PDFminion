Feature: Usage analytics

  Background: User consented to usage analytics

  Scenario: what data is collected for usage analytics
  Given user uses the application to process several PDFs
  When user has chosen configuration options
  And has started processing
  Then all configuration settings
    And the number of PDFs
    And the dimensions and pagecount of PDFs
    And operating system are collected
    And stored  in local storage


