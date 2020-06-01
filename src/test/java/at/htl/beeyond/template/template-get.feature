Feature: Test for getting all templates

  Scenario: Get all templates
    Given url 'http://localhost:8080/authentication/login'
    And request { username: 'it123456', password: 'passme'}
    When method post
    * def accessToken = response.access_token

    Given url 'http://localhost:8080/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Accept = 'application/json'
    When method get
    Then status 200