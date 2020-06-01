Feature: Test for creating a custom application

  Background:
    Given url 'http://localhost:8080/authentication/login'
    And request { username: 'it123456', password: 'passme'}
    When method post
    * def accessToken = response.access_token

  @valid
  Scenario: Create a custom application
    Given url 'http://localhost:8080/custom-application'
    And header Authorization = 'Bearer ' + accessToken
    And header Accept = 'application/json'
    And request { content:'Some yaml', note: 'Some note', status: 'PENDING'}
    When method post
    Then status 204

  @invalidContent
  Scenario: Create an invalid custom application
    Given url 'http://localhost:8080/custom-application'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: 'No content' }, { content:'', note: 'Some note', status: 'PENDING'}
    When method post
    Then status 422