Feature: Test for getting all templates

  Background:
    Given url baseUrl

  Scenario: Get all templates
    Given path '/authentication/login'
    And request { username: 'teststudent', password: 'student'}
    When method post
    Then status 200
    And print response
    * def accessToken = response.access_token

    Given path '/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Accept = 'application/json'
    When method get
    Then status 200