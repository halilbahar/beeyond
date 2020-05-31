Feature: Test for getting all templates

  Scenario: Get all templates
    Given url 'http://localhost:8080/template'
    And header Accept = 'application/json'
    When method get
    Then status 200