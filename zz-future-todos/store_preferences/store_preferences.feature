Feature: Store preferences
  As a user I want to persistently store preferences between invocations of the application

  @usage_analytics @preferences
  Scenario: Store preference for collecting usage analytics
  Given A user is asked for permission to anonymous usage analytics collection
  When the user consent to usage analytics
  Then the configuration of usage analytics is persisted in preference store