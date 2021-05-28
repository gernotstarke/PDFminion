Feature: Consent to usage analytics data collection

  @manual
  Scenario: A user has to consent to usage analytics
  Given a user who has never used the application before
  When the user starts the application
  Then the user is given the choice to consent to usage statstics collection


  Scenario: A user consents to usage analytics
    Given a user who consents to usage analytics
    When the user uses the application
    Then usage data is collected

