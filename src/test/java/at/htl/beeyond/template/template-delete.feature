Feature: Test for deleting a template

  Background:
    Given url baseUrl
    Given path '/authentication/login'
    And request { username: 'testteacher', password: 'teacher'}
    When method post
    * def accessToken = response.access_token

  @deleteExistingTemplate
  Scenario: Delete a template
    Given path '/template/1'
    And header Authorization = 'Bearer ' + accessToken
    And header Accept = 'application/json'
    When method delete
    Then status 204