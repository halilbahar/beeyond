Feature: Update a custom application status

  Background:
    Given url 'http://localhost:8080/authentication/login'
    And request { username: 'it123456', password: 'passme'}
    When method post
    * def accessToken = response.access_token

  @validApprove
  Scenario: Approve a custom application
    Given url 'http://localhost:8080/custom-application/approve/2'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request {}
    When method put
    Then status 204

  @invalidApprove
  Scenario: Test approve of non existent custom application
    Given url 'http://localhost:8080/custom-application/approve/9223372036854775807'
    And header Content-Type = 'application/json'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request {}
    When method put
    Then status 404

  @validDeny
  Scenario: Deny a custom application
    Given url 'http://localhost:8080/custom-application/deny/1'
    And header Authorization = 'Bearer ' + accessToken
    And request ''
    When method put
    Then status 204

  @invalidDeny
  Scenario: Test deny of non existent custom application
    Given url 'http://localhost:8080/custom-application/deny/9223372036854775807'
    And header Authorization = 'Bearer ' + accessToken
    And request ''
    When method put
    Then status 404