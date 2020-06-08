Feature: Test for getting all applications

  Scenario: Get list with all applications
    Given url baseUrl + '/application'
    And header Content-Type = 'application/json'
    When method get
    Then status 200