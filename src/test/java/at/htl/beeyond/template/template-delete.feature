Feature: Test for deleting a template

  Background:
    Given url 'http://localhost:8080/authentication/login'
    And request { username: 'it123456', password: 'passme'}
    When method post
    * def accessToken = response.access_token

  @valid
  Scenario: Delete a template
    Given url 'http://localhost:8080/template/1'
    And header Authorization = 'Bearer ' + accessToken
    And header Accept = 'application/json'
    When method delete
    Then status 204

  @invalid
  Scenario: Delete a custom application
    Given url 'http://localhost:8080/template/9999999999'
    And header Authorization = 'Bearer ' + accessToken
    And header Accept = 'application/json'
    When method delete
    Then status 404