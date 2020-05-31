Feature: Test for creating a custom application

  @valid
  Scenario: Create a custom application
    Given url 'http://localhost:8080/custom-application'
    And header Content-Type = 'application/json'
    And request { name: "it123456" }, { content:'Some yaml', note: 'Some note', status: 'PENDING'}
    When method post
    Then status 204

  @invalid
  Scenario: Create an invalid custom application
    Given url 'http://localhost:8080/custom-application'
    And header Content-Type = 'application/json'
    And request { name: "does not exist" }, { content:'', note: 'Some note', status: 'PENDING'}
    When method post
    Then status 422