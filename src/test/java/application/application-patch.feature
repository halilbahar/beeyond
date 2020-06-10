Feature: Tests for patching applications

  Background:
    Given url baseUrl
    Given path '/authentication/login'
    And request { username: 'testteacher', password: 'teacher'}
    When method post
    * def accessToken = response.access_token

  @approve
  Scenario: Approve an application
    Given path '/approve/1'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request ''
    When method patch
    Then status 204

  @deny
  Scenario: Deny a application
    Given path '/deny/1'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request ''
    When method patch
    Then status 204