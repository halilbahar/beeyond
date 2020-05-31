Feature: Test for getting all custom application

  Scenario: Get list with all custom applications
    Given url 'http://localhost:8080/custom-application'
    And header Accept = 'application/json'
    When method get
    Then status 200
