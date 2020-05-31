Feature: Test for deleting a template

  @valid
  Scenario: Delete a template
    Given url 'http://localhost:8080/template/1'
    And header Accept = 'application/json'
    When method delete
    Then status 204

  @invalid
  Scenario: Delete a custom application
    Given url 'http://localhost:8080/template/9999999999'
    And header Accept = 'application/json'
    When method delete
    Then status 404