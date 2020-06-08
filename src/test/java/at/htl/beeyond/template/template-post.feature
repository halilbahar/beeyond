Feature: Test for creating a template

  Background:
    Given url baseUrl
    Given path '/authentication/login'
    And request { username: 'testteacher', password: 'teacher'}
    When method post
    * def accessToken = response.access_token

  Scenario: Create a new template
    Given path '/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: 'Test Template', description: 'This is a test template', content: 'Some yaml' }
    When method post
    Then status 204