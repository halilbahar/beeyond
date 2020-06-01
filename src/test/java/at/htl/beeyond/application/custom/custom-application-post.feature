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

  @invalidNote
  Scenario: Create an invalid custom application
    Given url 'http://localhost:8080/custom-application'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: 'Note too long' }, { content:'Some content', note: 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.', status: 'PENDING'}
    When method post
    Then status 422